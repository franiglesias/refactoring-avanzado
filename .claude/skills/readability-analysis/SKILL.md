---
name: readability-analysis
description: Analyze the readability of a specific code unit (class, function, method, or block) using Cognitive Complexity (Campbell/SonarSource) as the primary metric, plus text-on-code proxies (identifier length, comment density, magic literal count, comment/identifier consistency). Produces three refactoring proposals targeted at the readability drivers. Use when the user asks "is this readable", "score readability", "cognitive complexity", "is this hard to read", or wants refactoring proposals oriented around readability rather than control-flow complexity.
---

# Readability Analysis

Quantify how hard a code unit is to *read*, not just to execute. Cognitive Complexity is the primary signal because it's the only readability metric that is hand-computable, interpretable, and explicitly designed around what makes code hard to follow (nesting, broken linear flow, mixed boolean operators). Text-on-code metrics support it as secondary signals.

This skill is complementary to `complexity-analysis`: that one measures *computational* complexity (CC, Halstead, MI, Big-O); this one measures *cognitive* complexity. A unit can be cyclomatic-simple and cognitively hard, or vice versa.

## Scope of this skill

- **One unit per invocation** — class, function, method, or block named by the user.
- **Read-only by default.** Do not edit unless the user picks a refactor and asks for the diff.
- **Hand-computed numbers are approximate.** Cognitive Complexity has well-defined rules and matches tooling closely; the text-on-code proxies are softer signals — treat them as smoke alarms, not measurements.
- **Cognitive Complexity is not Cyclomatic Complexity.** They share some triggers but differ on `else`, `switch`, boolean sequences, and nesting. Do not paper over the differences.

## Primary metric: Cognitive Complexity

Defined by G. Ann Campbell (SonarSource, 2018) as a more readability-oriented alternative to McCabe's CC. It penalizes nesting and breaks in linear flow.

Start at 0. The unit's score is the sum of increments below.

### B1 — Increments for breaks in linear flow (each adds 1)

- `if`
- `else if`
- `else`
- Ternary `? :`
- `switch` (the whole `switch` adds 1, regardless of how many `case` labels)
- `for`, `for...of`, `for...in`, `while`, `do-while`
- `catch`
- `goto LABEL`, `break LABEL`, `continue LABEL` (labelled jumps only)
- Each method in a recursion cycle (direct or mutual recursion)
- **Sequences of mixed binary logical operators**: count one for each *switch* between `&&` and `||` in a chain. `a && b && c` adds 1. `a && b || c` adds 2 (one for `&&`, one for switching to `||`). `a && b || c && d` adds 3.

### B2 — Nesting increment (added on top of B1)

For each of these constructs nested inside another nesting structure, add the current nesting depth:

- `if` / `else if` / `else` / ternary
- `switch`
- `for` / `while` / `do-while`
- `catch`

**Nesting depth** is incremented by entering: the same list above, plus inner functions / lambdas / closures.

So a single `if` at the top of a method adds 1. An `if` inside a `for` adds 1 (B1) + 1 (depth=1) = 2. An `if` inside a `for` inside an `if` adds 1 + 2 = 3.

### What does *not* increment

- Method/function declarations.
- Sequential statements.
- Each `case` of a `switch` (only the `switch` itself counts).
- The first `else if` is *not* special — it counts like any other.
- Null coalescing (`??`), optional chain (`?.`), and short-circuit `&&`/`||` *as a single chain* of one operator add nothing on their own — only switches between operators count.
- `try` and `finally` (only `catch` counts).

### Risk bands (per method, SonarSource defaults)

- 0–5: easy
- 6–10: moderate
- 11–15: difficult
- 16+: very difficult — refactor

### Differences vs cyclomatic complexity

Important so the two skills don't disagree without explanation:

| Construct | Cyclomatic | Cognitive |
|---|---|---|
| `else` | 0 | 1 |
| `switch` with N cases | N | 1 |
| Each `case` | 1 | 0 |
| `&&` / `||` (each occurrence) | 1 each | only on operator switch |
| Nesting | flat | weighted by depth |
| Recursion | 0 | 1 per method in cycle |
| Ternary | 1 | 1 (+ nesting) |

A 30-case `switch` is CC=30 but Cognitive=1. A triple-nested `if` is CC=3 but Cognitive=1+2+3=6.

## Secondary signals: text-on-code proxies

These are *supporting* observations, not primary metrics. Report them when they reinforce or contradict the Cognitive Complexity reading.

- **Average identifier length** — sweet spot 8–16 chars. Below 6 suggests cryptic single-letter or abbreviated names; above 20 suggests responsibilities packed into one name. Report mean and outliers.
- **Identifier consistency** — does the same concept always use the same word? (`user` vs `account` vs `customer` for the same thing is a readability tax.)
- **Comment density** — comments per logical LOC. Below ~5% on non-trivial logic *or* above ~30% are both smells (under-documented vs. comment-as-band-aid).
- **Comment/identifier coherence** — do the comments use vocabulary that matches the identifiers, or do they translate? Mismatched vocabularies signal that names don't carry their weight.
- **Magic literals** — count untyped/unnamed numeric and string literals in domain logic (excluding obvious `0`, `1`, `-1`, `''`, and array indices). High count means readers must context-switch to interpret values.
- **Boolean parameter count** — boolean parameters in calls force the reader to look up which `true` means what. Count them at call sites within the unit.

## How to apply this skill

1. **Locate the unit.** Read the file. Confirm exact line range. If the user named something ambiguous, pick the most likely target and say which.
2. **Compute Cognitive Complexity walking the unit top to bottom.** Maintain a running nesting depth. For each construct, write the increment in the form `+1 (if) + 2 (depth)` so the user can audit the count. Show the running total.
3. **Compute the secondary signals.** Hand-sample if the unit is large; say what you sampled.
4. **Identify the readability drivers.** Name *why* the unit reads as it does — deep nesting, long boolean chains, cryptic names, magic literals, etc. The drivers are what the refactors target.
5. **Propose exactly three refactorings**, ordered by expected readability impact. Each must include:
   - One-line summary
   - Concrete change (extract X, invert Y, name Z)
   - Which signals it moves and roughly by how much
   - Trade-off (does it harm anything else — perf, line count, learnability)
6. **Stop.** Do not refactor unless the user picks one.

## Output shape

When invoked, produce sections in this order:

1. **Unit** — `path:start-end`, signature, logical and physical LOC.
2. **Cognitive Complexity** — total score, band, plus a short walk-through showing the increments line-by-line or block-by-block. This is the load-bearing part of the analysis; do not skip the walk-through.
3. **Secondary signals** — small bullet list:
   - Avg identifier length, with notable outliers
   - Comment density, with quick verdict
   - Magic literals count, with examples
   - Other signals if relevant (boolean params, vocabulary drift)
4. **Drivers** — 2–4 sentences naming what makes this unit hard to read. Tie back to specific lines.
5. **Refactoring proposals** — three, numbered, ordered by impact. For each: summary, change, signals moved, trade-off.
6. **Caveats** — what was sampled vs counted, comparisons to a tool's exact output, anything ambiguous.

## Anti-patterns to avoid

- Reporting a Cognitive Complexity number without the walk-through. The number alone teaches nothing; the *path* the count walked is the actual feedback.
- Treating Cognitive Complexity as interchangeable with cyclomatic complexity. They diverge on `else`, `switch`, boolean chains, and nesting. State the divergence when both metrics are present.
- Letting text-on-code proxies dominate. They are soft signals; if they say "bad" but Cognitive says "fine", trust Cognitive.
- Recommending renames as the primary fix. A rename moves an identifier-length signal but rarely moves Cognitive Complexity, which is usually the bigger lever.
- Producing more or fewer than three refactor proposals — the constraint is what forces prioritization.
- Quoting precise percentages on subjective text-on-code signals ("identifier readability index = 73.2%"). Round, qualify, or omit.
- Confusing nesting depth with indentation level. Nesting depth counts only the constructs in the B2 list, not braces or destructuring.
