package adventofcode2015

import "testing"

//nolint:unused // shared cross-repo test helper pattern
func testWithParser[P any, R comparable](
	t *testing.T,
	day uint8,
	filenameFunc func(uint8) string,
	part1 bool,
	parser func([]string) (P, error),
	solver func(P, bool) R,
	want R,
) {
	t.Helper()
	lines := linesFromFilename(t, filenameFunc(day))
	puzzle, err := parser(lines)
	if err != nil {
		t.Fatal(err)
	}
	got := solver(puzzle, part1)
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

//nolint:unused // shared cross-repo test helper pattern
func testSolver[R comparable](
	t *testing.T,
	day uint8,
	filenameFunc func(uint8) string,
	part1 bool,
	solver func([]byte, bool) (R, error),
	want R,
) {
	t.Helper()
	buf := fileFromFilename(t, filenameFunc, day)
	got, err := solver(buf, part1)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

//nolint:unused // shared cross-repo test helper pattern
func testLines[R comparable](
	t *testing.T,
	day uint8,
	filenameFunc func(uint8) string,
	part1 bool,
	solver func([]string, bool) R,
	want R,
) {
	t.Helper()
	lines := linesFromFilename(t, filenameFunc(day))
	got := solver(lines, part1)
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

//nolint:unused // shared cross-repo test helper pattern
func bench[T any](
	b *testing.B,
	day uint8,
	part1 bool,
	solver func([]byte, bool) (T, error),
) {
	b.Helper()
	puzzle := file(b, day)
	for b.Loop() {
		_, _ = solver(puzzle, part1)
	}
}

//nolint:unused // shared cross-repo test helper pattern
func benchWithParser[P any, R any](
	b *testing.B,
	day uint8,
	part1 bool,
	parser func([]string) (P, error),
	solver func(P, bool) R,
) {
	b.Helper()
	lines := linesFromFilename(b, filename(day))
	for b.Loop() {
		puzzle, _ := parser(lines)
		_ = solver(puzzle, part1)
	}
}

func benchLines[R any](
	b *testing.B,
	day uint8,
	part1 bool,
	solver func([]string, bool) R,
) {
	b.Helper()
	lines := linesFromFilename(b, filename(day))
	for b.Loop() {
		_ = solver(lines, part1)
	}
}
