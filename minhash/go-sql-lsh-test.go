package main

import (
	"database/sql"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"

	sqllsh "github.com/ekzhu/go-sql-lsh"

	_ "github.com/mattn/go-sqlite3"
)

func creatTempFileBench(t *testing.B) *os.File {
	tmpfile, err := ioutil.TempFile("", "_test")
	if err != nil {
		t.Fatal(err)
	}
	return tmpfile
}

func removeTempFileBench(t *testing.B, tempfile *os.File) {
	if err := tempfile.Close(); err != nil {
		t.Fatal(err)
	}
}

func randomSigs(n, size int) []sqllsh.Signature {
	random := rand.New(rand.NewSource(1))
	sigs := make([]sqllsh.Signature, n)
	for i := 0; i < n; i++ {
		sigs[i] = make(sqllsh.Signature, size)
		for d := 0; d < size; d++ {
			sigs[i][d] = uint(random.Int63())
		}
	}
	return sigs
}

func runSqlite(k, l, n, nq int, b *testing.B) {
	// Inialize database
	f := creatTempFileBench(b)
	db, err := sql.Open("sqlite3", f.Name())
	if err != nil {
		b.Fatal(err)
	}

	// Initalize data
	lsh, err := sqllsh.NewSqliteLsh(k, l, "lshtable", db)
	if err != nil {
		b.Fatal(err)
	}
	sigs := randomSigs(n, k*l)
	ids := make([]int, len(sigs))
	for i := range sigs {
		ids[i] = i
	}
	qids := rand.Perm(len(ids))[:nq]

	// Inserting
	start := time.Now()
	err = lsh.BatchInsert(ids, sigs)
	if err != nil {
		b.Fatal(err)
	}
	dur := float64(time.Now().Sub(start)) / float64(time.Second)
	log.Printf("Batch inserting %d signatures takes %.4f seconds", len(sigs), dur)

	// Indexing
	start = time.Now()
	lsh.Index()
	if err != nil {
		b.Fatal(err)
	}
	dur = float64(time.Now().Sub(start)) / float64(time.Second)
	log.Printf("Building index takes %.4f seconds", dur)

	// Query
	start = time.Now()
	for _, i := range qids {
		out := make(chan int)
		go func() {
			err := lsh.Query(sigs[i], out)
			if err != nil {
				b.Error(err)
			}
			close(out)
		}()
		for _ = range out {
		}
	}
	dur = float64(time.Now().Sub(start)) / float64(time.Millisecond)
	log.Printf("%d queries, average %.4f ms / query",
		len(qids), dur/float64(nq))

	removeTempFileBench(b, f)
}

func BenchmarkSqliteLsh128(b *testing.B) {
	runSqlite(2, 64, 10000, 100, b)
}

func BenchmarkSqliteLsh256(b *testing.B) {
	runSqlite(4, 64, 10000, 100, b)
}

func BenchmarkSqliteLsh512(b *testing.B) {
	runSqlite(8, 64, 10000, 100, b)
}

func main() {
	b := &testing.B{}
	BenchmarkSqliteLsh128(b)
	BenchmarkSqliteLsh256(b)
	BenchmarkSqliteLsh512(b)
}

/*

$ go run -v go-sql-lsh-test.go
command-line-arguments
2016/08/06 12:24:14 Batch inserting 10000 signatures takes 3.8678 seconds
2016/08/06 12:24:17 Building index takes 2.4401 seconds
2016/08/06 12:24:17 100 queries, average 1.5561 ms / query
2016/08/06 12:24:28 Batch inserting 10000 signatures takes 10.7476 seconds
2016/08/06 12:24:38 Building index takes 9.9105 seconds
2016/08/06 12:24:38 100 queries, average 1.8124 ms / query
2016/08/06 12:25:00 Batch inserting 10000 signatures takes 21.2290 seconds
2016/08/06 12:25:08 Building index takes 7.5356 seconds
2016/08/06 12:25:08 100 queries, average 2.2338 ms / query

*/
