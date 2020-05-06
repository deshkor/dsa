package main

import (
	"log"
	"os"
	"testing"
)

type testValues struct {
	input  int
	expect bool
	size   int
}

var (
	addTests    []testValues
	removeTests []testValues
	list        List
)

func TestMain(m *testing.M) {
	addTests = []testValues{
		{input: 3, expect: false, size: 1},
		{input: 10, expect: false, size: 2},
		{input: 5, expect: false, size: 3},
		{input: 5, expect: true, size: 3},
		{input: 20, expect: false, size: 4},
	}

	removeTests = []testValues{
		{input: 3, expect: false, size: 3},
		{input: 123, expect: true, size: 3},
		{input: 20, expect: false, size: 2},
		{input: 5, expect: false, size: 1},
	}

	os.Exit(m.Run())
}

func TestAdd(t *testing.T) {
	log.Println("Running add tests.")

	for i, test := range addTests {
		err := list.Add(test.input)

		if (err != nil) != test.expect {
			t.Errorf("[run: %d] Unexpected response. Got: %s. Expect: %v", i,
				err, test.expect)
		}

		if list.Size() != test.size {
			t.Errorf("[run: %d] Unexpected list size. Got: %d. Expect: %d", i,
				list.Size(), test.size)
		}
	}
}

func TestRemove(t *testing.T) {
	log.Println("Running remove tests.")

	for i, test := range removeTests {
		err := list.Remove(test.input)

		if (err != nil) != test.expect {
			t.Errorf("[run: %d] Unexpected response. Got: %s. Expect: %v", i,
				err, test.expect)
		}

		if list.Size() != test.size {
			t.Errorf("[run: %d] Unexpected list size. Got: %d. Expect: %d", i,
				list.Size(), test.size)
		}
	}
}
