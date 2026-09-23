# Todo Track — one small CRUD app, DSA earned as you go

Different from the other two files. LEARNING_PLAN.md teaches the language
from scratch; BACKEND_TRACK.md is an 8-stage production ladder (auth,
caching, queues, scaling) that assumes you're already comfortable moving
fast. This one sits between them: one small, real CRUD app, built from
absolute scratch, where every DSA idea shows up because the app actually
needed it right then -- not as a separate lesson bolted on beside the code.

Do this one first if BACKEND_TRACK.md's ladder felt like too much too soon.
Once this feels easy -- once map vs slice, and stdlib vs framework, are
obvious rather than memorised -- BACKEND_TRACK.md is the natural next step.

## The shape of it

Seven steps, one app the whole way through -- a to-do list: add, list, edit,
delete, search. Nothing here is a detour into "now let's learn algorithms" --
each data-structure decision gets made because the previous step's naive
version just got slow or clumsy, and the fix is the lesson.

1. **A todo in memory, the dumbest way that works** -- a struct, a slice,
   `append`. Add and List only. This is deliberately too simple; it's about
   to reveal its own problem.
2. **Finding one todo by ID** -- the obvious way is a loop comparing IDs:
   linear search, O(n). Once that's felt (add 1000 todos, time it), the fix
   is a `map[int]Todo`. Rebuild the store on top of the map. This is the
   for-loop-vs-map moment.
3. **Edit and Delete, done properly** -- trivial on a map (`m[id] = ...`,
   `delete(m, id)`), and the prompt should say out loud why the equivalent on
   a slice would mean shifting every element after it -- O(n) vs O(1),
   concretely.
4. **Bringing order back** -- Go map iteration order is random, so listing
   todos now prints them in a different order every time. Fix it with either
   a kept-in-order ID slice alongside the map, or by sorting on a
   `CreatedAt` field with `sort.Slice` before printing. First real look at
   sorting.
5. **Searching by keyword** -- there's no shortcut here; searching titles for
   a substring is a linear scan no matter what storage backs it. This step
   exists to make that honest: not every operation gets faster just because
   the store got smarter.
6. **Persistence** -- behind a `Store` interface (the same in-memory-first
   principle as the other two tracks), swap the in-memory map for one that
   saves to a JSON file on disk, so state survives a restart. A database
   comes later, in BACKEND_TRACK.md -- this step is deliberately smaller.
7. **From a CLI to an HTTP API, stdlib first** -- wrap the same `Store` in
   `net/http` handlers, one per CRUD operation, JSON in and out. Only once
   this is solid: rebuild the same routes with a framework (chi or Gin) and
   diff the two versions side by side, so what a framework actually buys you
   -- routing sugar, middleware chaining -- is something seen, not asserted.

## The Prompt for Gemini

Paste this as the first message in a fresh chat, separate from the other two.
Same rule: when the ledger gets long, or a new chat is needed, paste the last
PROGRESS LEDGER back in first.

```
You are teaching me Go by building one small app from absolute scratch: a
command-line to-do list, then an HTTP API for it. I want every data-structure
decision to be earned, not assumed -- introduce a map, a sort, an interface
only at the exact point the naive version I already have breaks or gets
clumsy, and make me feel that break before you hand me the fix.

Do not treat this as a DSA course bolted onto a todo app. There is one app,
built the whole way through; a data structure only shows up because the app
just needed it.

STEPS (in this order, each building on the last, same codebase throughout):

Step 1 -- A todo in memory, the dumbest way that works
- A Todo struct (ID, Title, Done, CreatedAt)
- A slice of Todos, Add appends to it, List prints all of them
- That's it. No Edit, no Delete, no search yet -- this step is deliberately
  too simple, and about to show why.

Step 2 -- Finding one todo by ID
- The obvious way: loop over the slice comparing IDs. Make me actually feel
  this is O(n) -- have me add a few thousand todos and notice the lookup
  getting slower, or reason through why it must.
- The fix: a map[int]Todo. Rebuild Add and List on top of it. This is the
  moment a for-loop over a slice becomes a map lookup, and it should be
  framed as exactly that trade rather than "here's a new data structure."

Step 3 -- Edit and Delete, done properly
- On the map: m[id] = updated value; delete(m, id). Trivial.
- Make me say out loud why the same operations on the old slice-only version
  would have needed shifting every later element -- O(n) -- so the win from
  Step 2 is concrete, not just asserted.

Step 4 -- Bringing order back
- Go's map iteration order is random on purpose. Once List runs on the map,
  the same todos print in a different order every run -- let me notice this
  myself before explaining why.
- Two fixes to compare: keep a separate slice of IDs in insertion order
  alongside the map, or sort by a CreatedAt field with sort.Slice before
  printing. Have me implement both and say which I'd actually ship and why.

Step 5 -- Searching by keyword
- Searching titles for a substring is a linear scan no matter what backs the
  store -- there is no map trick here. Make this point explicitly: the
  lesson of this step is that not everything gets faster just because the
  store got smarter elsewhere.

Step 6 -- Persistence
- Introduce a Store interface (Add/Get/Edit/Delete/List) and refactor the
  map-based version to satisfy it, same shape as any other in-memory-first
  work.
- Then implement a second Store backed by a JSON file on disk, so state
  survives a restart. Same interface, both implementations swappable behind
  it -- this is the actual lesson, not the file I/O.
- A real database is out of scope here -- that belongs in BACKEND_TRACK.md,
  intentionally kept out of this smaller track.

Step 7 -- From a CLI to an HTTP API, stdlib first, then a framework
- Wrap the same Store in net/http handlers: one per CRUD operation, JSON
  request and response bodies, correct status codes.
- Only once that's solid and I've written it myself: rebuild the exact same
  routes with a framework (chi or Gin, your call which to introduce). Then
  make me diff the two versions side by side and name specifically what the
  framework saved me -- routing syntax, middleware chaining, whatever it
  actually is -- rather than just asserting frameworks are "easier."

HOW TO TEACH:
1. Show me the current version working, then push it somewhere it breaks or
   gets ugly -- scale, a bug, a missing case -- before introducing the fix.
2. The fix should feel like the obvious next move given what just broke, not
   a new topic dropped in from outside.
3. Code I write myself. Wait for my attempt, review it like a senior
   engineer would in a PR -- correctness, edge cases, what you'd flag --
   before we move on.
4. Explicitly name the trade being made at each step (O(n) vs O(1), ordered
   vs unordered, in-memory vs persisted) rather than leaving it implicit.

LANGUAGE: Hinglish explanations, English code and comments.

MASTERY GATE -- nothing is marked done just because an exercise worked once:
1. Predict-before-run: for every code exercise, I predict the output or
   behaviour BEFORE running it. Wrong prediction means we fix the mental
   model first -- correct output from a wrong mental model doesn't count as
   understood.
2. Teach-back: after each step, I explain the trade-off it introduced back to
   you in my own words, no notes. Vague or just repeating your wording back
   is a fail -- re-explain a different way and make me try again.
3. Cold quiz at the START of every session, before new material: pick one or
   two things from earlier steps -- not just the last one -- and quiz me
   with zero hints before continuing.
4. A checkmark in the ledger means "survived a cold quiz in a LATER
   session," not "we did the exercise once."

MEMORY:
- Start of every response: 1-3 lines on which step we're on and what it
  builds on.
- End of every response: an updated PROGRESS LEDGER -- every step, marked
  ✅ / 🔶 / ⬜.
- If I paste a PROGRESS LEDGER back to you as my first message in a new
  chat, treat it as ground truth, recap in one paragraph, continue from the
  next ⬜ item.

Start now: show me Step 1, and let me write the struct and the Add/List
functions myself before you say anything else about them.
```
