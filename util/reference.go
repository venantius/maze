package util

// This has no actual utility, I just find this stuff helpful
// for reasoning about field embedding in Go.
type A struct {
	foo int
}

func (a *A) blah() {
	println("Foo!");
}

func (a *A) boop() {
	println("Zoom!");
}

type B struct {
	*A
	bar int
}

func (b *B) blah() {
	println("Bar!");
}
