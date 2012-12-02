package mytesting

import "testing"

type T struct {
	*testing.T
}

func (t *T) expectEq(a, b interface{}) {
	if a != b {
		t.Errorf("expectEq(%v, %v)", a, b)
	}
}

func (t *T) expectTrue(a bool) {
	if !a {
		t.Errorf("expectTrue(%v)", a)
	}
}

func (t *T) expectFalse(a bool) {
	if a {
		t.Errorf("expectFalse(%v)", a)
	}
}

func (t *T) assertEq(a, b interface{}) {
	if a != b {
		t.Fatalf("assertEq(%v, %v)", a, b)
	}
}

func (t *T) assertTrue(a bool) {
	if !a {
		t.Fatalf("assertTrue(%v)", a)
	}
}

func (t *T) assertFalse(a bool) {
	if a {
		t.Fatalf("assertFalse(%v)", a)
	}
}
