package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var defaultListInsertSeq = []int{34, 44, 3, 12, -1, 33, -23, 100}

func getDefaultPopulatedTree() (*Tree, error) {
	tree := Tree{}
	var err error

	for _, val := range defaultListInsertSeq {
		err = tree.Insert(val)
		if err != nil {
			return nil, err
		}
	}

	return &tree, nil
}

func TestInsertNode(t *testing.T) {
	tree := Tree{}

	testCases := []struct {
		Value       int
		ListSize    int
		ExpectError bool
	}{
		{Value: 25, ListSize: 1, ExpectError: false},
		{Value: 10, ListSize: 2, ExpectError: false},
		{Value: 5, ListSize: 3, ExpectError: false},
		{Value: 30, ListSize: 4, ExpectError: false},
		{Value: 35, ListSize: 5, ExpectError: false},
		{Value: 30, ListSize: 5, ExpectError: true},
	}

	for _, tc := range testCases {
		err := tree.Insert(tc.Value)
		if tc.ExpectError {
			assert.NotNil(t, err, "Should have got an error")
		} else {
			assert.Nil(t, err, "Should have got an error")
		}

		assert.Equal(t, tc.ListSize, tree.Size, "Wrong list size")
	}
}

func TestSearchNode(t *testing.T) {
	nilList := Tree{}
	populatedList, err := getDefaultPopulatedTree()

	assert.Nil(t, err, "Shouldn't have failed while initializing default tree values")

	testCases := []struct {
		Value       int
		List        *Tree
		ExpectNode  bool
		ExpectError bool
	}{
		{Value: 10, List: &nilList, ExpectNode: false, ExpectError: true},
		{Value: 10, List: populatedList, ExpectNode: false, ExpectError: false},
		{Value: 33, List: populatedList, ExpectNode: true, ExpectError: false},
		{Value: -23, List: populatedList, ExpectNode: true, ExpectError: false},
		{Value: 34, List: populatedList, ExpectNode: true, ExpectError: false},
	}

	for _, tc := range testCases {
		node, err := tc.List.Search(tc.Value)

		if tc.ExpectError {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
		}

		if tc.ExpectNode {
			assert.NotNil(t, node)
			assert.Equal(t, node.Value, tc.Value)
		} else {
			assert.Nil(t, node)
		}
	}
}

/*
						10
				5				15
			3		7		13		17
		  1   4   6	  8	  11  14  16  19

*/


func TestRemoveNode(t *testing.T) {
	populatedList, err := getDefaultPopulatedTree()
	assert.Nil(t, err, "Shouldn't have failed while initializing default tree values")

	tcs := []struct {
		val int
		err bool
	}{
		{val: 34},
		{val: 3},
		{val: -1},
		{val: -23},
		{val: 100},
		{val: 1000, err: true},
	}

	for i, tc := range tcs {
		fmt.Println("Starting test:",i)

		err := populatedList.Remove(tc.val)

		if tc.err {
			assert.NotNil(t, err, fmt.Sprintf("Expected an error for tc: %d", i))
		} else {
			assert.Nil(t, err, fmt.Sprintf("Unexpected error for tc: %d", i))
		}
	}
}

func TestGetLeastLeafNode(t *testing.T) {
	fmt.Println("TODO: TestGetLeastLeafNode")
}

func TestGetParent(t *testing.T) {
	fmt.Println("TODO: TestGetParent")
}
