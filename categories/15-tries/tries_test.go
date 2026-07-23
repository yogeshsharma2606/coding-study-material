package tries

import (
	"reflect"
	"testing"
)

func TestTrie(t *testing.T) {
	tr := NewTrie()
	tr.Insert("apple")
	if !tr.Search("apple") {
		t.Error("search apple")
	}
	if tr.Search("app") {
		t.Error("app is not a word")
	}
	if !tr.StartsWith("app") {
		t.Error("prefix app")
	}
	tr.Insert("app")
	if !tr.Search("app") {
		t.Error("app now a word")
	}
}

func TestWordDictionary(t *testing.T) {
	w := NewWordDictionary()
	w.AddWord("bad")
	w.AddWord("dad")
	w.AddWord("mad")
	if w.Search("pad") {
		t.Error("pad absent")
	}
	if !w.Search("bad") {
		t.Error("bad present")
	}
	if !w.Search(".ad") {
		t.Error(".ad should match")
	}
	if !w.Search("b..") {
		t.Error("b.. should match")
	}
}

func TestReplaceWords(t *testing.T) {
	got := ReplaceWords([]string{"cat", "bat", "rat"},
		[]string{"the", "cattle", "was", "rattled", "by", "the", "battery"})
	want := []string{"the", "cat", "was", "rat", "by", "the", "bat"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v", got)
	}
}

func TestLongestWord(t *testing.T) {
	if got := LongestWord([]string{"w", "wo", "wor", "worl", "world"}); got != "world" {
		t.Errorf("got %q", got)
	}
	if got := LongestWord([]string{"a", "banana", "app", "appl", "ap", "apply", "apple"}); got != "apple" {
		t.Errorf("got %q", got)
	}
}

func TestSuggestedProducts(t *testing.T) {
	got := SuggestedProducts([]string{"mobile", "mouse", "moneypot", "monitor", "mousepad"}, "mouse")
	want := [][]string{
		{"mobile", "moneypot", "monitor"},
		{"mobile", "moneypot", "monitor"},
		{"mouse", "mousepad"},
		{"mouse", "mousepad"},
		{"mouse", "mousepad"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v", got)
	}
}

func TestMapSum(t *testing.T) {
	m := NewMapSum()
	m.Insert("apple", 3)
	if got := m.Sum("ap"); got != 3 {
		t.Errorf("got %d", got)
	}
	m.Insert("app", 2)
	if got := m.Sum("ap"); got != 5 {
		t.Errorf("got %d", got)
	}
	m.Insert("apple", 5) // overwrite
	if got := m.Sum("ap"); got != 7 {
		t.Errorf("got %d", got)
	}
}

func TestFindMaximumXOR(t *testing.T) {
	if got := FindMaximumXOR([]int{3, 10, 5, 25, 2, 8}); got != 28 {
		t.Errorf("got %d", got)
	}
	if got := FindMaximumXOR([]int{0}); got != 0 {
		t.Errorf("got %d", got)
	}
}
