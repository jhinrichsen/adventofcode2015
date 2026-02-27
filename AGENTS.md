# Claude Code Guidelines for Advent of Code 2022

## Startup
- **MUST** run `beans prime` before starting work in this repository.

## 🚨 CRITICAL: Solution Confidentiality

**Solutions MUST NOT appear anywhere except in unit test `want` values.**

Prohibited locations:
- ❌ Git commit messages
- ❌ Pull request titles/descriptions
- ❌ Code comments
- ❌ Documentation files (README.adoc, AGENTS.md)
- ❌ Console output or logs

Permitted location (solutions **MUST** appear only here):
- ✅ Unit test files: `TestDayXXPart1` and `TestDayXXPart2` (the `want` parameter)

**Example commit messages:**
```
✅ Good: perf(day17): optimize water retention logic by 30%
❌ Bad:  feat(day17): solve part 2 - answer is 30410
```

---

## Critical Rules

### Requirement Keywords
The key words **MUST**, **MUST NOT**, **SHOULD**, and **MAY** in this document are to be interpreted as described in RFC 2119 and RFC 8174.

### Absolute Prohibition
- **MUST NOT** use recursion anywhere in this repository (puzzle code, helpers, or tests)

### Function Signatures (PRIMARY RULE)
- **MUST** implement: `func DayXX(<input>) uint`
- **SHOULD** use: `func DayXX(<input>, part1 bool) uint` unless alternatives are more elegant
- **IF** parser required (input cannot be directly processed by `input_test.go` functions):
  - Parser: `func NewDayXX(<input>) DayXXPuzzle` (return by value)
  - Combined: `func DayXX(puzzle DayXXPuzzle) uint`
- Alternative/custom implementations used for comparison or optimization (for example `Day01Branchless`) **MAY** use any practical signature.

### Export Policy (MANDATORY)
- **MUST** export only the day entrypoints and parser/puzzle types: `DayXX*`, `NewDayXX`, `DayXXPuzzle`
- Keep internal helpers unexported (e.g. parser internals, evaluators, temporary structs)
- Do **NOT** add exported convenience methods on puzzle structs unless explicitly required

### File Access Prohibition
- Puzzles must not perform I/O
- **MUST** use `input_test.go` helpers for all file reads in tests; puzzle code **MUST NOT** read files

### uint Pattern (MANDATORY)
- **ALL** puzzle return types that are counts/sums/totals/amounts must be `uint`
- Push `uint` contract up the entire call chain
- Area, perimeter, distances, prices are inherently non-negative
- Example: `func exploreRegion(...) (area, perimeter uint)`

### Coordinate System
- **MUST** use `x`/`y` throughout; **MUST NOT** use row/col naming
- `dimX` (width), `dimY` (height) for dimensions
- `grid[y][x]` indexing pattern
- `startY, startX int` parameter order

### Data Types
- **MUST** use `byte` for ASCII characters (A-Z, 0-9, symbols)
- **MUST NOT** use `rune` - unnecessary UTF-8 overhead for AoC
- Use `[]byte(string)` for conversion, not manual loops

### Algorithm Requirements
- **MUST NOT** use recursion
- **MUST** use iterative implementations with explicit stacks: `[]image.Point`
- Use `image.Point{X: x, Y: y}` for coordinates

### Modern Go Patterns (MANDATORY)
- **MUST** use the latest Go language and standard library features available in the Go version declared in `go.mod`
- Use `for range N` instead of `for i := 0; i < N; i++` (range over integers)
- Use `slices` package: `slices.Equal`, `slices.Contains`, `slices.Sort`
- Use `maps` package: `maps.Equal`, `maps.Clone` when needed
- Use `clear(map)` and `clear(slice)` for efficient clearing
- Use `min()` and `max()` built-in functions

### Error Handling (MANDATORY)
- **MUST NOT** silently ignore errors with blank identifier `_`
- **MUST NOT** panic - AoC problems should not panic
- **MUST** handle errors gracefully (continue, skip, use zero value)
- AoC input is always valid, so errors won't occur in practice
- Bad: `n, _ := strconv.Atoi(line)`
- Bad: `n, err := strconv.Atoi(line); if err != nil { panic(err) }`
- Good: `n, err := strconv.Atoi(line); if err != nil { continue }` (skip invalid)
- Good: `n, err := strconv.Atoi(line); if err != nil { n = 0 }` (use default)

### Test Structure
- Table-driven tests with external files
- `testdata/dayXX_example1.txt` not inline strings
- **MUST NOT** use multiline string literals in tests; **MUST** use external testdata files
- `lines := linesFromFilename(t, filename)` in tests only
- Multiple examples: use `example1Filename(day)`, `example2Filename(day)`, etc.
- Available filename functions: `exampleFilename()`, `exampleNFilename()`, `example1Filename()`, `example2Filename()`, `example3Filename()`, `filename()`

### Input Parsing (Flexible)
- Parser is optional; you **MAY** use a parser when beneficial for complexity
- `func DayXX(input []byte)` - fine if puzzle can parse bytes directly
- `func DayXX(lines []string)` - fine if puzzle needs line-based input
- `func NewDayXX()` + `DayXX(puzzle)` - use for complex data structures
- Choose the most appropriate input format for each puzzle's needs
- Use appropriate `input_test.go` helper functions

### Day Migration Pattern
When migrating legacy days to the new style, apply this pattern:
- **MUST** make `DayXX(input, part1 bool)` (or equivalent canonical solver) the single source of truth
- Alternative/custom implementations (for comparison or optimization) **MAY** exist separately (for example `DayXXBranchless`)
- **MUST NOT** keep redundant wrapper entrypoints that duplicate part1/part2 logic
- **MUST** use zero-padded two-digit day naming everywhere: `DayXX`, `NewDayXX`, `TestDayXX...`, `BenchmarkDayXX...`, and regex patterns like `^TestDayXX` / `^BenchmarkDayXX`
- **MUST** route parse/validation errors through a parser step (`NewDayXX(...)`) when parsing is non-trivial; solvers **MUST NOT** silently skip malformed records (for example `if err != nil { continue }`)
- **MUST** keep benchmarks as one-line calls via shared benchmark helpers (for example `bench(...)`, `benchLines(...)`, `benchWithParser(...)`)
- **MUST** keep file-based tests as one-line calls via shared test helpers (for example `testSolver(...)`, `testLines(...)`, `testWithParser(...)`)
- **SHOULD** keep example tables local to the test function that uses them (inline/local scope)
- **SHOULD** reduce nonessential file-level globals/constants when straightforward to inline without harming readability

## Commit Message Convention

Use conventional commits with day number as scope:

**Format:** `<type>(<scope>): <description>`

**Examples:**
- `feat(day14): add part 1 solution`
- `feat(day14): add part 2 with binary search`
- `test(day14): add example tests for part 1`
- `refactor(day14): convert to table-driven tests`
- `fix(day14): correct ORE calculation for surplus`
- `docs(day14): add algorithm explanation`

**Types:**
- `feat`: New feature/solution
- `fix`: Bug fix
- `perf`: Performance optimization
- `refactor`: Code refactoring
- `test`: Test additions/changes
- `docs`: Documentation
- `chore`: Build/tooling changes

**Scope:**
- Use `dayXX` (lowercase, zero-padded) for day-specific commits
- Omit scope for repository-wide changes (e.g., `chore: update AGENTS.md`)

## Benchmark Optimization Workflow

When optimizing puzzle solutions, follow this workflow to measure and document performance improvements:

### Performance Targets

**Goal:** "make total" should run in **under 1 second** with **minimal memory usage**

Focus on:
- Low `B/op` (bytes per operation)
- Low `allocs/op` (allocations per operation)
- Minimize allocations in hot loops
- Prefer arrays/slices over maps for bounded data ranges
- Reuse buffers and reset slice lengths rather than reallocating

### Process

**IMPORTANT:** Make ONE change at a time to ensure clear attribution of improvements.

1. **Create baseline benchmark (b0)**
   ```bash
   go test -run=^$ -bench=DayXXPart.$ -count=8 -benchmem > b0
   ```
   Remove the last two lines (PASS and ok lines) from b0

2. **Optimize the code**
   Apply ONE targeted performance improvement following the guidelines above

3. **Run benchmark again (b1)**
   ```bash
   go test -run=^$ -bench=DayXXPart.$ -count=8 -benchmem > b1
   ```
   Remove the last two lines from b1

4. **Compare with benchstat**
   ```bash
   benchstat b0 b1
   ```

5. **Document if worthwhile (>5% improvement)**
   - Add a "Day XX" section to README.adoc if not present
   - Include the benchstat output
   - Add a short, concise explanation of the optimization
   - **MUST NOT** use bold (**) in README.adoc - use proper AsciiDoc formatting
   - **MUST NOT** include solution values in documentation

6. **Commit changes**
   ```bash
   git add .
   git commit -m "perf(dayXX): [optimization description]"
   ```

7. **Repeat** until performance targets are met or no further significant improvements

### Example README.adoc Entry

```asciidoc
== Day 01: The Tyranny of the Rocket Equation

Optimized fuel calculation by eliminating redundant allocations.

----
name       old time/op    new time/op    delta
Day01-16     1.23µs ± 2%    0.95µs ± 1%  -22.76%

name       old alloc/op   new alloc/op   delta
Day01-16      512B ± 0%       0B ± 0%  -100.00%
----
```
