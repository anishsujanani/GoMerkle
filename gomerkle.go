package gomerkle

import (
	"crypto/sha256"
	"fmt"
)

type MerkleNode struct {
	hash       string
	rawtext    string
	leftchild  *MerkleNode
	rightchild *MerkleNode
}

func MerkleTree(Content string, LeafSize int) MerkleNode {
	return createNode(Content, LeafSize)
}

func createNode(content string, leafSize int) MerkleNode {
	chunks := getChunks(content, leafSize)

	for _, chunk := range chunks {
		chunkHash := getChunkHash([]byte(chunk))
		fmt.Println(chunk, chunkHash)
	}

	return MerkleNode{
		hash:       "h_root",
		rawtext:    content,
		leftchild:  nil,
		rightchild: nil}
}

func getChunks(content string, leafSize int) []string {
	var chunks []string
	var contentLength = len(content)

	for i := 0; i < contentLength; i += leafSize {
		var to int
		if to = i + leafSize; to > contentLength {
			to = contentLength
		}
		chunks = append(chunks, string(content[i:to]))
	}

	return chunks
}

func (m MerkleNode) GetRawText() string {
	return m.rawtext
}

func (m MerkleNode) GetHash() string {
	return m.hash
}

func (m MerkleNode) GetLeftChild() *MerkleNode {
	return m.leftchild
}

func (m MerkleNode) GetRightChild() *MerkleNode {
	return m.rightchild
}

func (m *MerkleNode) setHash() {
}

func (m *MerkleNode) setLeftChild(lc MerkleNode) {
	m.leftchild = &lc
}

func (m *MerkleNode) setRightChild(rc MerkleNode) {
	m.rightchild = &rc
}

func (m MerkleNode) String() string {
	return fmt.Sprintf(
		"Hash: %v\nRawText:%v\nLeftChild.Hash=%v\nRightChild.Hash=%v",
		m.hash, m.rawtext, m.leftchild, m.rightchild)
}

func getChunkHash(chunk_bytes []byte) string {
	return fmt.Sprintf("%x", (sha256.Sum256(chunk_bytes)))
}
