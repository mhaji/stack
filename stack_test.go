package stack

import "testing"

func TestPop(t *testing.T) {
	s := NewStack()
	s.Push("a")
	s.Push("b")
	s.Push("c")


	for i, elem := range []string{"c", "b", "a"} {
		v, err := s.Pop();
		if err != nil {
			t.Fatal("Expected err to be nil")
		}
		if v != elem {
			t.Fatalf("Expected v to be '%s'", elem)
		}
		if s.Size() != (2-i) {
			t.Fatalf("Expected size to be %d", (2-i))
		}
	}
}

func TestPopEmptyStackErr(t *testing.T) {
	s := NewStack()

	_, err := s.Pop();
	if err != EmptyStackErr {
		t.Fatal("Expected err to be EmptyStackErr but was %v", err)
	}
}


func TestSize(t *testing.T) {
	s := NewStack()
	s.Push(1)
	s.Push(2)

	if s.Size() != 2 {
		t.Fatal("Unexpected size, expected 2")
	}

	s.Pop()

	if s.Size() != 1 {
		t.Fatal("Unexpected size, expected 1")
	}
}