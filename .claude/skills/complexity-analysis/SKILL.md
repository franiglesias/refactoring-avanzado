---
name: complexity-analysis
description: Analyze the complexity of a specific code unit (class, function, method, or block) using cyclomatic complexity, Halstead metrics, maintainability index, and Big-O algorithmic complexity, then propose three refactorings ordered by impact. Use when the user asks to "analyze complexity", "measure complexity", "compute metrics", "is this too complex", "Halstead", "cyclomatic", "maintainability index", or asks for refactoring proposals grounded in metrics for a specific unit.
---

# Complexity Analysis

Quantify the complexity of one named code unit and propose three concrete refactorings, each tied to the metric it would move. The point is to give the user a defensible reading — not a vibes-based "this looks gnarly" — and to convert that reading into actionable next steps.

## Scope of this skill

- **One unit per invocation.** The user names a class, function, method, or block. Do not analyze siblings, callers, or the whole file unless asked.
- **Read-only by default.** Produce numbers and proposals. Do not edit the file unless the user explicitly asks for a refactor to be applied.
- **Manual computation is approximate.** Halstead and maintainability numbers from hand-counting reveal trends and order of magnitude, not exact tooling output. Say so when reporting.
- **Big-O depends on assumptions.** External data structures (e.g. `Map` vs `Array`) drive complexity. Make assumptions explicit in the analysis.

## Metrics to compute

### Cyclomatic complexity (McCabe)

Start at 1, then add 1 for each independent decision point:

- `if`, `else if`, `case` (each `case` label, not the `switch` itself)
- `for`, `while`, `do-while`, `for...of`, `for...in`
- `catch` clause
- Each `&&`, `||`, `??` in a boolean expression
- Ternary `? :`
- Optional chain `?.` when it short-circuits a branch (count once per chain, not per `?`)

Do **not** count `else` separately (it is the implicit branch of the matching `if`). Do not count `try` or `finally`.

**Risk bands:**
- 1–4: simple, low risk
- 5–7: moderate
- 8–10: complex, review carefully
- 11+: high risk, refactor likely warranted

### Halstead metrics

Count from the unit's source:

- `n1` = number of **distinct** operators (keywords, operators, function calls, brackets-as-grouping)
- `n2` = number of **distinct** operands (identifiers, literals)
- `N1` = total operator occurrences
- `N2` = total operand occurrences

Then:

- Vocabulary: `η = n1 + n2`
- Length: `N = N1 + N2`
- Volume: `V = N × log₂(η)`
- Difficulty: `D = (n1 / 2) × (N2 / n2)`
- Effort: `E = D × V`
- Estimated time (seconds): `T = E / 18`
- Estimated delivered bugs: `B = V / 3000`

For hand counting, treat as operators: control keywords, assignment, arithmetic, comparison, logical, member access (`.`), call `()`, indexing `[]`, `new`, `await`, `yield`, `typeof`, `?.`, `??`, ternary parts, brackets used to group. Treat as operands: identifiers and literal values. Be consistent within one analysis.

### Maintainability Index (Microsoft variant, 0–100 scale)

```
MI = max(0, (171 − 5.2 × ln(V) − 0.23 × CC − 16.2 × ln(LOC)) × 100 / 171)
```

Where `V` is Halstead Volume, `CC` is cyclomatic complexity, `LOC` is logical lines of code (exclude blank lines and pure-comment lines).

**Bands:**
- ≥ 85: highly maintainable
- 65–84: moderately maintainable
- < 65: difficult to maintain

### Algorithmic complexity (Big-O)

Report time and space complexity in terms of the relevant input sizes (name them — `n = items.length`, `m = workshops per item`, etc.). Identify:

- Nested iteration → multiply
- Sequential phases → take the max
- Recursion → solve the recurrence (or state it: T(n) = aT(n/b) + f(n))
- Hidden costs of standard ops (array `includes` is O(n), `Map.get` is O(1), `JSON.parse` is O(input size))

State assumptions when external types are involved.

### Lines of code

Report **logical LOC** (statements, excluding blank lines and pure comments) and **physical LOC** (raw line span) — they tell different stories.

## How to apply this skill

1. **Locate the unit.** Read the file. Confirm the exact byte/line range you are scoring. If the user named something ambiguous ("the loop"), pick the most likely target and say which.
2. **Compute each metric** following the definitions above. Show enough of the count that the user can sanity-check it (e.g. "CC = 1 + 3 ifs + 1 ternary + 2 `&&` = 7"). Don't bury the math.
3. **Interpret the numbers together.** A high CC with low Halstead volume is different from a low CC with high volume. Call out which dimension dominates.
4. **Identify the structural drivers.** Before proposing refactors, name *why* the unit scores as it does — nested branches, repeated state, deep call chains, accidental coupling, etc.
5. **Propose exactly three refactorings**, ordered by expected impact on the dominant metric. Each proposal must include:
   - A one-line summary
   - The concrete change (extract X, replace Y with Z, invert Z)
   - Which metrics it moves and roughly by how much
   - Any risk or trade-off (behavior change, perf, readability cost)
6. **Stop there.** Do not start refactoring unless the user picks one and asks for the diff.

## Output shape

When invoked, produce sections in this order:

1. **Unit** — `path:start-end`, signature/name, logical and physical LOC.
2. **Metrics** — a small table:

   | Metric | Value | Band |
   |---|---|---|
   | Cyclomatic | … | … |
   | Halstead V / D / E | … | … |
   | Maintainability Index | … | … |
   | Big-O time / space | … | — |

3. **Computation notes** — short bullet list showing how CC and the Halstead counts were tallied, plus stated assumptions for Big-O.
4. **Interpretation** — 2–4 sentences naming the dominant complexity drivers.
5. **Refactoring proposals** — three, numbered, ordered by impact. For each: summary, change, metric movement, trade-off.
6. **Caveats** — manual-count approximation note; any ambiguity in unit boundaries or external types.

## Anti-patterns to avoid

- Reporting metrics without showing the count — the user can't sanity-check or learn from a black-box number.
- Quoting Halstead values to four decimal places. They are estimates; one decimal at most.
- Producing more or fewer than three proposals. The constraint forces prioritization.
- Proposing refactors that contradict the surrounding codebase's conventions (e.g. introducing a class in a code style that prefers free functions).
- Conflating physical and logical LOC.
- Treating cyclomatic complexity as the only signal — a unit can be CC=3 and still unmaintainable due to high Halstead difficulty or deep call chains.
- Recommending tooling instead of doing the analysis ("run eslint-complexity"). The user asked you, not the tool.
