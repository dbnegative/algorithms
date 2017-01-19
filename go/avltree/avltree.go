package avltree

import "fmt"

//AVLTree basic AVLTree data structure
type AVLTree struct {
	LeftChild  *AVLTree
	RightChild *AVLTree
	Height     int
	Key        int
}

//New returns a new intialised AVLTree node
func New(lc *AVLTree, rc *AVLTree, height int, key int) *AVLTree {
	root := &AVLTree{
		LeftChild:  lc,
		RightChild: rc,
		Height:     height,
		Key:        key,
	}
	return root
}

//max returns the maximun between two ints
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//returns the lowest minimum value node
func minValue(root *AVLTree) *AVLTree {
	current := root
	for current.LeftChild != nil {
		current = current.LeftChild
	}
	return current
}

//Height returns the nodes height unless nil then it returns -1
func Height(root *AVLTree) int {
	if root == nil {
		return -1
	}
	return root.Height
}

//Child return the child node of the unbalanced node
func Child(root *AVLTree) *AVLTree {
	child := root.LeftChild
	if child == nil {
		child = root.RightChild
	}
	return child
}

//GetKey helper function return the key of the node, if nil return nil
func (root *AVLTree) GetKey() interface{} {
	if root != nil {
		return root.Key
	}
	return nil
}

//Insert a new node into the AVLTree (assumes no duplicates) recursively and return the modified
//avl tree
func (root *AVLTree) Insert(key int) *AVLTree {

	if root == nil {
		root = New(nil, nil, 0, key)
		fmt.Printf("Creating New Node %d \n", key)
		return root
	}
	//check whether to place left
	if key < root.Key {
		root.LeftChild = root.LeftChild.Insert(key)
	}
	//check whether to place right
	if key > root.Key {
		root.RightChild = root.RightChild.Insert(key)
	}

	//recalculate the height of this node after insertion
	root.Height = max(Height(root.LeftChild), Height(root.RightChild)) + 1

	//check if AVLTree
	return balance(root)
}

//Search recursively through AVL tree for Key (Log n time)
func (root *AVLTree) Search(key int) *AVLTree {

	if root.Key == key {
		return root
	}
	if root.LeftChild != nil {
		if key < root.Key {
			return root.LeftChild.Search(key)
		}
	}
	if root.RightChild != nil {
		if key > root.Key {
			return root.RightChild.Search(key)
		}
	}
	return nil
}

//Delete removes a node matching the key recursively, returns the
//reblanced tree
/* Ascii diagram to show deletion
--------------------------------------------------------------------------
       /            /
      4            4
    /   \         /  \
   2     5  ->   3    5
    \
     3
-----------------------------------------------------------------------*/
func (root *AVLTree) Delete(key int) *AVLTree {

	if root == nil {
		return root
	}

	if key < root.Key {
		root.LeftChild = root.LeftChild.Delete(key)
	} else if key > root.Key {
		root.RightChild = root.RightChild.Delete(key)
	} else {
		//does it have children? if only one set root to that child
		if root.LeftChild == nil || root.RightChild == nil {
			node := Child(root)
			// No child case
			if node == nil {
				//unlink node as it has no children
				root = nil
			} else {
				//overwrite deleted node with child
				root = node
			}
		} else {
			//root has 2 children so find the lowest inorder successor
			node := minValue(root.RightChild)
			//delete the inorder successor
			root.Key = node.Key
			root.RightChild = root.RightChild.Delete(node.Key)
		}
	}

	if root == nil {
		return root
	}
	//rebalance after delete
	return balance(root)
}

//Balances a subtree according to it's height
func balance(root *AVLTree) *AVLTree {

	//Balance factor between left and right subtree
	factor := Height(root.LeftChild) - Height(root.RightChild)

	//Get the node below
	childnode := Child(root)

	//unbalanced in left subtree
	if factor > 1 {
		//Check the height difference of the node below to work out rotations
		//greater > than 0 do right rotation as node is unbalanced on the left
		//less < than 0 do Left Right rotation as node is unbalanced on the right subtree
		//the inverse of the above is true if factor is less than -1

		if Height(childnode.LeftChild)-Height(childnode.RightChild) > 0 {
			root = rotateRight(root)
		} else {
			root = rotateLeftRight(root)
		}
	}

	//unbalanced in right subtree
	if factor < -1 {
		if Height(childnode.LeftChild)-Height(childnode.RightChild) < 0 {
			root = rotateLeft(root)
		} else {
			root = rotateRightLeft(root)
		}
	}
	//return the node
	return root
}

//rotate a left subtree right
func rotateRight(root *AVLTree) *AVLTree {
	node := root.LeftChild
	root.LeftChild = node.RightChild
	node.RightChild = root
	root.Height = max(Height(root.LeftChild), Height(root.RightChild)) + 1
	node.Height = max(Height(node.LeftChild), Height(node.RightChild)) + 1
	return node
}

/* Ascii Diagram to show Right Left rotation
-------------------------------------------------------------------------------------------
 *                                 *                                *
  \                                 \                                \
   6    (0,2) = -2                   6      (0,2)  -> Left            7
    \                                 \                              /  \
     10 (1,0) =  1  -> Right           7    (0,1)                   6    10
    /                                   \
   7                                     10
-------------------------------------------------------------------------------------------
**/

//rotate a left subtree right then rotate the new root left
func rotateRightLeft(root *AVLTree) *AVLTree {
	root.RightChild = rotateRight(root.RightChild)
	root = rotateLeft(root)
	return root
}

//rotate a right subtree left
func rotateLeft(root *AVLTree) *AVLTree {
	node := root.RightChild
	root.RightChild = node.LeftChild
	node.LeftChild = root
	root.Height = max(Height(root.LeftChild), Height(root.RightChild)) + 1
	node.Height = max(Height(node.LeftChild), Height(node.RightChild)) + 1
	return node
}

/* Ascii Diagram to show Left Right rotation
-------------------------------------------------------------------------------------------
     *                                 *                              *
    /                                 /                              /
   6   (2, 0) =  2                   6  (2,0)   -> Right      (1,1) 5
  /                                 /                              /  \
 4     (0, 1) = -1  -> Left        5    (1,0)               (0,0) 4    10 (0,0)
  \                               /
   5   (0, 0) = 0                4
-------------------------------------------------------------------------------------------
**/

//rotate a right subtree left then rotate the new root right
func rotateLeftRight(root *AVLTree) *AVLTree {
	root.LeftChild = rotateLeft(root.LeftChild)
	root = rotateRight(root)
	return root
}
