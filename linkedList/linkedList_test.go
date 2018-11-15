package linkedList

import "testing"

func TestLinkedList(t *testing.T) {

	l := &List{}
	l.Append("s")
	l.Append("s")
}

func TestSearchList(t *testing.T) {
	names := []string{
		"Bob",
		"Burgers",
		"Barbra",
		"Jimmy",
		"George",
		"Immigrant",
	}
	l := &List{}
	for _, name := range names {
		l.Append(name)
	}

	results := l.Search("im")
	exp := map[string]bool{
		names[3]: true,
		names[5]: true,
	}
	if len(results) != len(exp) {
		t.Fatal("not all results returned from search:", results, "Expected: ", exp)
	}
}
