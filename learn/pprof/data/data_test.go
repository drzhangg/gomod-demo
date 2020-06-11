package data

import "testing"

const url = "https://github.com/drzhangg"

func TestAdd(t *testing.T) {
	s := Add(url)
	if s == ""{
		t.Errorf("Test.add error")
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(url)
	}
}
