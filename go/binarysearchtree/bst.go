package bst

import "fmt"

//Bintree struct
type Bintree struct {
	id         int
	data       int
	root       int
	leftChild  *Bintree
	rightChild *Bintree
}

var nodes = 0

//New instantiate a new Bintree struct
func New(data int, root int, lc *Bintree, rc *Bintree) *Bintree {
	bt := &Bintree{
		id:         nodes,
		data:       data,
		root:       root,
		leftChild:  lc,
		rightChild: rc,
	}
	nodes++
	return bt
}

//Insert 's new data
func (bt *Bintree) Insert(data int) bool {

	if data == bt.data {
		return true
	}

	if data < bt.data {
		if bt.leftChild == nil {
			bt.leftChild = New(data, bt.id, nil, nil)
			return true
		}
		bt.leftChild.Insert(data)

	} else {
		if bt.rightChild == nil {
			bt.rightChild = New(data, bt.id, nil, nil)
			return true
		}
		bt.rightChild.Insert(data)
	}
	return false
}

//Search for data in tree
func (bt *Bintree) Search(data int) bool {

	if bt.data == data {
		return true
	}

	if data < bt.data {
		if bt.leftChild != nil {
			return bt.leftChild.Search(data)
		}
	}

	if data > bt.data {
		if bt.rightChild != nil {
			return bt.rightChild.Search(data)
		}
	}

	return false
}

//Print out node data
func (bt *Bintree) Print() {

	fmt.Printf("root node: %v data: %d \n", bt.id, bt.data)

	if bt.leftChild != nil {
		bt.leftChild.Print()
	}
	if bt.rightChild != nil {
		bt.rightChild.Print()
	}
}

//Depth - returns depth of tree
func (bt *Bintree) Depth() int {
	depth := 0
	for nodes > 0 {
		nodes = nodes / 2
		depth++
	}
	return depth
}
