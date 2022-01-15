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
	var rawContentChunks []string     = getChunks(Content, LeafSize)
	var pendingInsertion []MerkleNode = formUnlinkedMerkleNodes(rawContentChunks)

	return consumePendingInsertionIntoTree(pendingInsertion)
}

func consumePendingInsertionIntoTree(pendingInsertion []MerkleNode) MerkleNode {
	var newLevelNodes []MerkleNode

	// if we ever have an odd number of nodes, we need to balance
	if lenPendingInsertion := len(pendingInsertion); lenPendingInsertion % 2 == 1 && lenPendingInsertion != 1 {
		pendingInsertion = append(pendingInsertion, MerkleNode{})
	}

	// else iterate in pairs, form hash, link left and right child
	for i := 0; i < len(pendingInsertion); i += 2 {
		var leftChild MerkleNode  = pendingInsertion[i]
		var rightChild MerkleNode = pendingInsertion[i+1]

		var rawContentChunks []string = []string{leftChild.GetRawText() + rightChild.GetRawText()}

		var newNode MerkleNode = formUnlinkedMerkleNodes(rawContentChunks)[0]
		newNode.setLeftChild(leftChild)
		newNode.setRightChild(rightChild)

		newLevelNodes = append(newLevelNodes, newNode)
	}

	// all nodes for the current level, retrieved from pendingInsertion have been consumed
	// set the new level just formed (newLevelNodes) as the new pendingInsertion and recursively call this func
	// if we have more nodes to process
	pendingInsertion = newLevelNodes
	if len(pendingInsertion) > 1 {
		return consumePendingInsertionIntoTree(pendingInsertion)
	}
	// else we have formed the root
	return pendingInsertion[0]
}

func formUnlinkedMerkleNodes(rawContentChunks []string) []MerkleNode {
	var nodeList []MerkleNode
	for _, chunk := range rawContentChunks {
		nodeList = append(nodeList, MerkleNode{
			hash:       computeHash([]byte(chunk)),
			rawtext:    chunk,
			leftchild:  nil,
			rightchild: nil})
	}
	return nodeList
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

func (m MerkleNode) GetLeftChild() *MerkleNode {
	return m.leftchild
}

func (m MerkleNode) GetRightChild() *MerkleNode {
	return m.rightchild
}

func (m *MerkleNode) GetHash() string {
	return m.hash
}

func (m *MerkleNode) setLeftChild(lc MerkleNode) {
	m.leftchild = &lc
}

func (m *MerkleNode) setRightChild(rc MerkleNode) {
	m.rightchild = &rc
}

func computeHash(chunk_bytes []byte) string {
	return fmt.Sprintf("%x", (sha256.Sum256(chunk_bytes)))
}

func (m MerkleNode) String() string {
	return fmt.Sprintf(
		"RawText:\t\t%v\nHash:\t\t\t%v\nLeftChild.RawText:\t%v\nRightChild.RawText\t%v\n",
		m.GetRawText(), m.GetHash(), (m.GetLeftChild()).GetRawText(), (m.GetRightChild()).GetRawText())
}
