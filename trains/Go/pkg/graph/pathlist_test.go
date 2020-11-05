package graph

import (
	"testing"
)

func TestPathList(t *testing.T) {
	testPathList := NewPathList()

	if testPathList.Last() != nil {
		t.Fatal("empty PathList.Last() returned non-nil")
	}

	if testPathList.Len() != 0 {
		t.Fatal("empty PathList.Last() returned non-zero value")
	}

	testPathList.Append(generateSimplePath())

	if testPathList.Len() != 1 {
		t.Fatal("single entry PathList.Len() returned non-single value")
	}
}

func TestPathListDedup(t *testing.T) {
	testPathList := NewPathList()

	testPathOne := generateSimplePath()
	testPathTwo := generateSimplePath()

	for i := 0; i < 4; i++ {
		testPathList.Append(testPathOne, testPathTwo)
	}

	testPathList.Dedup()

	if testPathList.Len() != 2 {
		t.Fatal("PathList.Dudup() failed to dedup")
	}
}
