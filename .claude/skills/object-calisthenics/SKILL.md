---
name: object-calisthenics
description: Apply Jeff Bay's nine Object Calisthenics rules to new or recently changed code. Use when the user asks to review, refactor, or write new code under Object Calisthenics constraints, or mentions "calisthenics", "OO discipline", or "tighten this class". Targets new code by default — does not refactor unrelated existing code unless explicitly requested.
---

# Object Calisthenics

Nine constraints from Jeff Bay (*The ThoughtWorks Anthology*, 2008) that, taken together, push code toward small, expressive, well-encapsulated objects. They are deliberately strict — treat them as rails for *new* code, not a stick to beat legacy code with.

## Scope of this skill

- **Apply to new or just-changed code** — files in the current diff, code about to be written, or a class the user explicitly named.
- **Do not** sweep across unrelated existing code "to comply." Touching working code that wasn't part of the request is out of scope.
- When a rule cannot be honored without distorting the design, say so explicitly and explain the trade-off rather than forcing compliance.
- These rules are heuristics. If the user pushes back on one, defer to their judgement and record the decision.

## The nine rules

### 1. One level of indentation per method
Each method has at most one level of nesting. Extract inner blocks into well-named private methods. Loops inside conditionals (or vice versa) are the most common smell.

### 2. Don't use the `else` keyword
Prefer early returns, guard clauses, polymorphism, or the Null Object pattern. `else` branches usually hide a missing abstraction or an unstated default.

### 3. Wrap all primitives and strings (in domain code)
A `string email` or `int ageInYears` is a primitive obsession. Introduce a value object (`Email`, `Age`) so invariants live with the type. Exception: primitives at the edge of the system (DTOs, serialization, framework adapters) and inside the value object itself.

### 4. First-class collections
A class that holds a collection holds *only* that collection. No other instance fields. This forces collection-related behavior (filtering, summing, grouping) into the wrapper instead of leaking across the codebase.

### 5. One dot per line (Law of Demeter)
`a.getB().getC().doSomething()` reaches through other objects' internals. Tell, don't ask: move the behavior closer to the data. Fluent builders/DSLs and immutable chains (e.g. `Stream`, `Optional`) are exceptions where each call returns the same conceptual object.

### 6. Don't abbreviate
If a name is so long it tempts abbreviation, the responsibility is probably too big. Split the class. `usrMgr` is never an improvement over `UserManager`, but `UserManager` is itself a hint that a more specific name is waiting.

### 7. Keep all entities small
Soft ceilings: classes ≤ 50 lines, methods ≤ 10 lines, packages/modules ≤ 10 classes. Treat these as smoke alarms, not hard limits — when one trips, look for a missing concept.

### 8. No classes with more than two instance fields
Forces you to find the cohesive cluster of state. A class with five fields is usually two or three classes wearing a trench coat. The two-field limit is the most aggressive rule — relax to "few" if the alternative is meaningless wrappers.

### 9. No getters/setters/public properties
Encapsulate behavior, not state. Instead of `account.getBalance() < amount`, write `account.canAfford(amount)`. Getters are acceptable at serialization boundaries (e.g. mapping to a DTO) and for value objects whose whole purpose is to expose a value.

## How to apply this skill

1. **Identify the target surface.** Read the diff or the file the user named. Do not expand scope.
2. **Walk the rules in order.** For each rule, list concrete violations in the target code with file:line references.
3. **Propose changes, don't make them silently.** For each violation, suggest the smallest refactor that honors the rule, and call out cases where complying would harm the design.
4. **Stop at the diff boundary.** If a violation lives in code you didn't change, mention it as an observation but do not refactor it.
5. **Summarize trade-offs.** End with a short note on which rules were applied, which were skipped and why.

## Anti-patterns to avoid when applying

- Refactoring far beyond the user's request to chase compliance.
- Wrapping primitives at API/persistence boundaries where they shouldn't be wrapped.
- Splitting a class purely to hit the two-field rule when the resulting pieces have no independent meaning.
- Replacing every getter with a behavioral method, including in true value objects where exposing the value is the point.
- Citing the rules as authority ("rule 8 says…") instead of explaining the design benefit.

## Output shape

When invoked, produce:

1. A short bullet list of violations grouped by rule, each with `path:line` and a one-line description.
2. Suggested refactors, ordered by impact.
3. A trade-offs section listing rules deliberately not applied and why.