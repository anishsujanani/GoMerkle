# GoMerkle
A merkle-tree implementation in Go.

## Motivation:
- Learning Go
- Merkle trees are great

## API
```go
// docs pending
```

## Usage 
``` go
package main

import (
	"fmt"
	"github.com/anishsujanani/gomerkle"
)

func main() {
	root := gomerkle.MerkleTree("abcdefgh", 2);

	fmt.Printf("-> Root node - default Print(): \n%v\n", root)
	
	fmt.Printf("-> Root node - raw text:\n%v\n\n", root.GetRawText())
	
	fmt.Printf("-> Left child of root - default Print(): \n%v\n", root.GetLeftChild())
	
	fmt.Printf("-> Right child of root - raw text: \n%v\n\n", (root.GetRightChild()).GetRawText())
	
	fmt.Printf("-> Tree height: \n%v\n\n", root.GetHeight())

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
}
```

