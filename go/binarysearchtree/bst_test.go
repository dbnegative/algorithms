package bst

import "testing"

func TestNew(t *testing.T) {
	t.Log("Testing creation of root node")
	data := 5
	//create root node
	bst := New(data, 0, nil, nil)
	//fmt.Println(bst)
	if bst.data != data && bst != nil {
		t.Errorf("%s", "Root node not created, Data was not inserted")
	}
}

func TestMultiInsert(t *testing.T) {
	t.Log("Testing multi inserts to BST")
	data := []int{10, 20, 3, 44, 56, 88, 4, 3, 5555, 43, 33, 65}
	//create root node
	bst := New(data[0], 0, nil, nil)
	// add all elements
	for i := range data[0:] {
		bst.Insert(data[i])
	}
	if bst.leftChild == nil {
		t.Errorf("%s", "Data was not found in leftchild")
	}

	if bst.rightChild == nil {
		t.Errorf("%s", "Data was not found in rightchild")
	}
}

func TestSearch(t *testing.T) {
	t.Log("Testing search BST")
	data := []int{10, 20, 3, 88, 7, 67, 100, 2, 5, 8}
	//create root node
	bst := New(data[0], 0, nil, nil)
	for i := range data[0:] {
		bst.Insert(data[i])
	}

	for i := range data {
		if !bst.Search(data[i]) {
			t.Errorf("%s", "Could not find data in tree")
		}
	}
}

func TestPrint(t *testing.T) {
	t.Log("Testing printout of BST")
	data := []int{10, 20, 3, 88, 7, 67, 100, 2, 5, 8}
	nodes = 0
	//create root node
	bst := New(data[0], 0, nil, nil)
	for i := range data[0:] {
		bst.Insert(data[i])
	}

	bst.Print()
	//OUTPUT:
	//root node: 0 data: 10
	//root node: 2 data: 3
	//root node: 7 data: 2
	//root node: 4 data: 7
	//root node: 8 data: 5
	//root node: 9 data: 8
	//root node: 1 data: 20
	//root node: 3 data: 88
	//root node: 5 data: 67
	//root node: 6 data: 100
}
