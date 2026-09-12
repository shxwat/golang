# My Concurrent Link Checker Deadlocked Before Printing a Single Result

> How a small mistake in channel flow blocked the producer, every worker, and the result collector in my Go program.

I recently built a concurrent URL health checker in Go.

The idea was simple: keep a list of URLs, distribute them among a fixed number of workers, apply a timeout to every HTTP request, and collect the results through a channel.

Instead of starting a separate goroutine for every URL, I created a small worker pool with three workers. This gave me control over how many HTTP requests could run at the same time.

The program compiled successfully. Every request also had a three-second timeout. From the code, it looked like the workers should process URLs and send their results back to the main goroutine.

<!-- SCREENSHOT 1: Attach the ORIGINAL worker code screenshot here. It should show the worker function, the 3-second context timeout, HTTP request, and result sends. -->

![Original URL-checker worker implementation](ATTACH-SCREENSHOT-1-HERE)

*The original worker handled HTTP requests and timeouts, but its lifecycle was still incomplete.*

But when I ran the program, the terminal remained completely silent.

No URL status appeared. No timeout message appeared. Even the first “URL Health Check Started” message was never printed.

The program had not crashed—it was still running—but it could not make any progress.

<!-- SCREENSHOT 2: Attach the FROZEN TERMINAL screenshot here. It should show the command used to run the program and the blank terminal below it. -->

![The link checker running without producing output](ATTACH-SCREENSHOT-2-HERE)

*The program started successfully but remained stuck without printing a single result.*

## Finding the blocking point

Both `jobs` and `results` were unbuffered channels.

With an unbuffered channel, a send can complete only when another goroutine is ready to receive the value. This creates a direct handoff between the sender and receiver.

My main goroutine started three workers and then began sending every URL into the `jobs` channel. The result collection loop appeared only after the complete job-submission loop.

<!-- SCREENSHOT 3: Attach the ORIGINAL main-flow screenshot here. It should show both unbuffered channels, worker creation, synchronous jobs loop, wg.Wait goroutine, and result loop. Important old lines: 104-128. -->

![Original job submission and result collection flow](ATTACH-SCREENSHOT-3-HERE)

*The main goroutine attempted to submit every job before it started receiving worker results.*

The first three URLs could be received by the three workers. The main goroutine then tried to send another URL.

Meanwhile, the workers completed their first HTTP requests and tried to send their results into the unbuffered `results` channel. However, the main goroutine was not receiving results yet—it was still blocked while trying to submit the next job.

The program reached this state:

```text
Main goroutine
    └── blocked while sending the next job

Worker 1
    └── blocked while sending a result

Worker 2
    └── blocked while sending a result

Worker 3
    └── blocked while sending a result

Result receiver
    └── not reached yet
```

The workers could not publish their results, so they could not return to the `jobs` channel. The main goroutine could not submit another job because no worker was available to receive it.

Every important goroutine was waiting for another blocked goroutine.

## Confirming it with a goroutine dump

Reading the code suggested a deadlock, but I wanted evidence from the running program.

I generated a goroutine dump using `SIGQUIT`. The dump showed the main goroutine in the `[chan send]` state at the job-submission line. It also showed the workers in the same state at the line where they attempted to publish results.

<!-- SCREENSHOT 4: Attach the GOROUTINE DUMP screenshot here. It must visibly contain: goroutine 1 [chan send], main.main(), Link_Checker/main.go:115, main.worker(), and Link_Checker/main.go:39. -->

![Goroutine dump showing the blocked producer and workers](ATTACH-SCREENSHOT-4-HERE)

*The dump confirmed the deadlock: main was blocked while submitting a job at line 115, while workers were blocked while publishing results at line 39.*

This was more useful than simply saying “the program is hanging.” The dump showed exactly which goroutines were blocked and the source lines responsible for the block.

It also confirmed that the HTTP timeout was not the solution to this problem. The timeout protected the HTTP request, but it could not stop a worker from blocking indefinitely while sending to a channel.

## Fixing the pipeline

The main goroutine was responsible for two jobs:

1. Sending URLs to the workers.
2. Receiving completed results from the workers.

The deadlock happened because it tried to finish the first responsibility before starting the second one.

I moved job submission into a separate producer goroutine. That allowed the main goroutine to begin consuming results immediately while URLs were still being submitted.

The producer also became responsible for closing the `jobs` channel after submitting the final URL.

The pipeline now looked like this:

```text
Job producer → jobs channel → workers → results channel → main
```

If workers temporarily blocked while sending results, the main goroutine could receive those results. The workers could then continue and accept more jobs.

<!-- SCREENSHOT 5: Attach the FIXED main-flow screenshot here. Current fixed lines: 106-133. It must show worker creation, job producer goroutine, wg.Wait goroutine, and result loop. -->

![Corrected concurrent pipeline](ATTACH-SCREENSHOT-5-HERE)

*Moving job submission into its own goroutine allowed job production and result collection to progress independently.*

## The second bug: workers never reported completion

While fixing the channel flow, I found another lifecycle problem.

The program called `wg.Add(1)` for every worker, but the original worker function never called `wg.Done()`.

That meant `wg.Wait()` could never finish. Even if every URL was processed, the goroutine responsible for closing `results` would wait forever. The main result loop would then also wait forever because the channel was never closed.

I added `defer wg.Done()` at the beginning of the worker function. Now each worker decrements the `WaitGroup` exactly once when it exits after the `jobs` channel is closed.

<!-- SCREENSHOT 6: Attach the FIXED worker screenshot here. Current fixed lines: 17-44. The function signature and defer wg.Done() must be clearly visible. -->

![Worker reporting completion through the WaitGroup](ATTACH-SCREENSHOT-6-HERE)

*Each worker now marks itself complete before exiting, allowing the results channel to close safely.*

This made the shutdown sequence clear:

1. The producer sends every URL and closes `jobs`.
2. Workers finish their remaining jobs and exit.
3. Every worker calls `wg.Done()`.
4. `wg.Wait()` returns and `results` is closed.
5. The main result loop finishes naturally.

## The result after the fix

After making both changes, the checker began printing results as soon as workers completed their requests.

The results did not always appear in the same order as the input URLs. That was expected: faster requests finished before slower ones. The important part was that all workers could continue processing jobs and the program eventually exited normally.

<!-- SCREENSHOT 7: Attach the FINAL SUCCESSFUL TERMINAL screenshot here. It should show multiple URL results and end with “All checks completed successfully!!”. -->

![Successful output from the corrected link checker](ATTACH-SCREENSHOT-7-HERE)

*After fixing the pipeline and worker lifecycle, the checker processed every URL and exited normally.*

## What I learned

This project changed how I think about channels.

Using a channel does not automatically make a program correctly concurrent. I still need to design the complete communication flow and make sure every send has a receiver that can actually run.

The main lessons I took from this bug were:

- An unbuffered channel send is a synchronization point, not a background queue.
- Producers and consumers must be able to make progress at the same time.
- Channel ownership should be clear, especially when deciding who closes a channel.
- Every `WaitGroup.Add` needs a matching completion signal.
- An HTTP timeout protects the request, not the entire worker lifecycle.
- A goroutine dump can turn a silent hang into a precise, evidence-backed diagnosis.

The final fix was small, but understanding why it worked required me to reason about the entire pipeline—not just the line where the program appeared to stop.

The program compiled successfully before the fix. It simply had no possible path forward.

That was the real lesson: concurrency bugs are often not about syntax. They are about whether the system can continue making progress.

---

**Tags:** `Go` `Golang` `Concurrency` `Backend Development` `Debugging` `Worker Pool`
