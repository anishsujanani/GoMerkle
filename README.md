# GoMerkle
A merkle-tree implementation in Go.

## Motivation:
- Learning Go
- Merkle trees are great

## API
`go doc -all gomerkle`
```
package gomerkle // import "github.com/anishsujanani/gomerkle"

Package gomerkle provides functions to create Merkle-trees and perform
common operations on the data structures involved. Anish Sujanani, January
2022.

TYPES

type MerkleNode struct {
	// Has unexported fields.
}

func MerkleTree(Content string, LeafSize int) MerkleNode
    MerkleTree creates the Merkle tree and returns the root node of type
    'MerkleNode'.

func (m MerkleNode) BreadthFirstSearch() []MerkleNode
    BreadthFirstSearch returns a slice of MerkleNode(s) gathered from a
    level-wise ordering of the tree starting from the node it is invoked by.

func (m MerkleNode) DepthFirstSearch(order string) []MerkleNode
    DepthFirstSearch returns a slice containing MerkleNode(s) gathered from a
    depth-first-search on the tree starting from the node it is invoked by.
    Ordering is decided based on the input parameter:
    (preorder|inorder|postorder).

func (m MerkleNode) EqualTo(t MerkleNode) bool
    EqualTo returns the result of node hash equality. Calling this function with
    nodes in corresponding positions in two trees will return the equivalence of
    those nodes, and therefore those sub-trees (if they are not leaves).

func (m *MerkleNode) GetHash() string
    GetHash returns the SHA-256 hash of a MerkleNode's raw-text.

func (m MerkleNode) GetHeight() int
    GetHeight returns the height of the Merkle tree.

func (m MerkleNode) GetInconsistentLeaves(t MerkleNode) []MerkleNode
    GetInconsistentLeaves returns a list of MerkleNode(s) from the 't' tree that
    differ from the 'm' tree. It is advised to first check if the two trees are
    the same height. If heights differ, there is no point in calling this
    function as the 't' tree has obviously changed.

func (m MerkleNode) GetLeaves() []MerkleNode
    GetLeaves returns a list of MerkleNode(s) representing the leaves of the
    tree built from the raw content.

func (m MerkleNode) GetLeftChild() *MerkleNode
    GetLeftChild returns the left-child of a MerkleNode.

func (m MerkleNode) GetNodeCount() int
    GetNodeCount returns the total number of nodes present in the tree. (2^n)-1

func (m MerkleNode) GetRawText() string
    GetRawText returns the raw-text of a MerkleNode.

func (m MerkleNode) GetRightChild() *MerkleNode
    GetRightChild returns the right-child of a MerkleNode.

func (m MerkleNode) String() string
    Custom fmt.Print* function for the type MerkleNode.
```

## Usage 
``` go
package main

import (
	"fmt"
	"github.com/anishsujanani/gomerkle"
)

func main() {
	root := gomerkle.MerkleTree("abcdefghijk", 2);

	fmt.Printf("-> Root node - default Print(): \n%v\n", root)
	
	fmt.Printf("-> Root node - raw text:\n%v\n\n", root.GetRawText())
	
	fmt.Printf("-> Left child of root - default Print(): \n%v\n", root.GetLeftChild())
	
	fmt.Printf("-> Right child of root - raw text: \n%v\n\n", (root.GetRightChild()).GetRawText())
	
	fmt.Printf("-> Tree height: \n%v\n\n", root.GetHeight())

	fmt.Printf("-> Node count: \n%v\n\n", root.GetNodeCount()) 

	fmt.Printf("-> Leaves: \n%v\n\n", root.GetLeaves())

	fmt.Printf("-> Breadth First Search (level-wise ordering) - default Print(): \n%v\n\n", root.BreadthFirstSearch())

	fmt.Printf("-> Depth First Search (preorder) - custom print: \n")
	for _, node := range root.DepthFirstSearch("preorder") {
		fmt.Printf("%v ", node.GetRawText())
	}

	fmt.Printf("\n\n-> DepthFirstSearch (inorder) - custom print: \n")
	for _, node := range root.DepthFirstSearch("inorder") {
		fmt.Printf("%v ", node.GetRawText())
	}

	fmt.Printf("\n\n-> Depth First Search (postorder) - custom print: \n%v\n", root.DepthFirstSearch("postorder"))

	t := gomerkle.MerkleTree("abxdefyhijz", 2)

	fmt.Printf("-> Root equality: \n%v\n\n", root.EqualTo(t))

	fmt.Printf("-> Nodes in target tree differing from root tree: \n%v\n", root.GetInconsistentLeaves(t))
}

```

