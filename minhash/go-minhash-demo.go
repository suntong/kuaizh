package main

import (
	"fmt"

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

		fmt.Printf("m1: %v\n", m1.Signature())
		fmt.Printf("m2: %v\n", m2.Signature())
		fmt.Printf("Similarity: %v\n", (m1.Similarity(m2)))
	}
}

/*

$ go run minhash-test.go
m1: [2080620544968867365 8563714006720870342 2321955019530081269 748419707729262257 3830358558201948858 1245459476431551274 719322203384083329 12825625163505565688 168963933333021279 1984971849485248784]
m2: [875043171984725454 2288412544560513636 1717007218637525576 748419707729262257 5270773809508254720 1245459476431551274 719322203384083329 10601423695814348436 168963933333021279 1984971849485248784]
Similarity: 0.5

*/
