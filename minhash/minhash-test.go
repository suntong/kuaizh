package main

import (
	"github.com/dgryski/go-metro"
	"github.com/dgryski/go-minhash"
	"github.com/dgryski/go-spooky"
)

func mhash(b []byte) uint64 { return metro.Hash64(b, 0) }

func main() {

	tests := []struct {
		s1 []string
		s2 []string
	}{
		{
			[]string{"hello", "world", "foo", "baz", "bar", "zomg"},
			[]string{"goodbye", "world", "foo", "qux", "bar", "zomg"},
		},
	}

	for _, tt := range tests {
		m1 := minhash.NewMinWise(spooky.Hash64, mhash, 10)
		m2 := minhash.NewMinWise(spooky.Hash64, mhash, 10)

		for _, s := range tt.s1 {
			m1.Push([]byte(s))
		}

		for _, s := range tt.s2 {
			m2.Push([]byte(s))
		}

		print(m1.Similarity(m2))
	}
}

/*

$ go run minhash-test.go
+5.000000e-001

*/
