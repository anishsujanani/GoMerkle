# GoMerkle
A merkle-tree implementation in Go.

## Motivation:
- Learning Go
- Merkle trees are great

## Usage 
``` go
package main

import (
	"fmt"
	"github.com/anishsujanani/gomerkle"
)

func main() {
	root := gomerkle.MerkleTree("abcdefghijklmnopqrstuvwxyz", 4); // content, leaf size
	fmt.Println(root)
	fmt.Println(root.GetLeftChild())
	fmt.Println(root.GetRightChild())
}
```
Output:
```
RawText:		    abcdefghijklmnopqrstuvwxyz  
Hash:			    71c480df93d6ae2f1efad1447c66c9525e316218cf51fc8d9ed832f2daf18b73  
LeftChild.RawText:	abcdefghijklmnop  
RightChild.RawText	qrstuvwxyz  
  
RawText:		    abcdefghijklmnop  
Hash:			    f39dac6cbaba535e2c207cd0cd8f154974223c848f727f98b3564cea569b41cf  
LeftChild.RawText:	abcdefgh  
RightChild.RawText	ijklmnop  

RawText:		    qrstuvwxyz  
Hash:			    f20641c45ff3e440f7bbbf2a2fb1538808fb8e80f929c24096ddbcd280bc1e8d  
LeftChild.RawText:	qrstuvwx  
RightChild.RawText	yz  
```

## API
```go
// docs pending
```
