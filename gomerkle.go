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
	if lenPendingInsertion := len(pendingInsertion); lenPendingInsertion%2 == 1 && lenPendingInsertion != 1 {
		pendingInsertion = append(pendingInsertion, MerkleNode{rawtext: "$"})
	}

	// else iterate in pairs, form hash, link left and right child
	for i := 0; i < len(pendingInsertion); i += 2 {
		var leftChild MerkleNode  = pendingInsertion[i]
		var rightChild MerkleNode = pendingInsertion[i+1]

		var rawContentChunks []string = []string{leftChild.GetHash() + rightChild.GetHash()}

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

func computeHash(chunk_bytes []byte) string {
	return fmt.Sprintf("%x", (sha256.Sum256(chunk_bytes)))
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

// Going down left subtrees only since we always insert a left child first
func (m MerkleNode) GetHeight() int {
	if m.GetLeftChild() == nil {
		return 1
	}
	return 1 + (m.GetLeftChild()).GetHeight()
}

func (m MerkleNode) DepthFirstSearch(order string) []MerkleNode {
	var nodeList []MerkleNode
	return m.dfs(&nodeList, order)
}

func (m MerkleNode) dfs(nodeList *[]MerkleNode, order string) []MerkleNode {
	processLeft := func(nodeList *[]MerkleNode, order string) {
		if m.GetLeftChild() != nil {
			(m.GetLeftChild()).dfs(nodeList, order)
		}
	}
	processRight := func(nodeList *[]MerkleNode, order string) {
		if m.GetRightChild() != nil {
			(m.GetRightChild()).dfs(nodeList, order)
		}
	}
	processCurrent := func(nodeList *[]MerkleNode, m MerkleNode) {
		*nodeList = append(*nodeList, m)
	}

	switch order {
	case "preorder":
		processCurrent(nodeList, m)
		processLeft(nodeList, order)
		processRight(nodeList, order)

	case "inorder":
		processLeft(nodeList, order)
		processCurrent(nodeList, m)
		processRight(nodeList, order)
	case "postorder":
		processLeft(nodeList, order)
		processRight(nodeList, order)
		processCurrent(nodeList, m)
	}

	return *nodeList
}

func (m MerkleNode) BreadthFirstSearch() []MerkleNode {
	var nodeList []MerkleNode = []MerkleNode{m}
	var nodesByLevel []MerkleNode

	for {
		if len(nodeList) == 0 {
			break
		}

		var current MerkleNode = nodeList[0]
		nodesByLevel = append(nodesByLevel, current)

		if lc := current.GetLeftChild(); lc != nil {
			nodeList = append(nodeList, *lc)
		}
		if rc := current.GetRightChild(); rc != nil {
			nodeList = append(nodeList, *rc)
		}

		nodeList = nodeList[1:]
	}

	return nodesByLevel
}

func (m MerkleNode) String() string {
	var lcRawText string = "no_left_child"
	var rcRawText string = "no_right_child"

	if m.GetLeftChild() != nil {
		lcRawText = (m.GetLeftChild()).GetRawText()
	}
	if m.GetRightChild() != nil {
		rcRawText = (m.GetRightChild()).GetRawText()
	}
	return fmt.Sprintf(
		"RawText:\t\t%v\nHash:\t\t\t%v\nLeftChild.RawText:\t%v\nRightChild.RawText\t%v\n",
		m.GetRawText(), m.GetHash(), lcRawText, rcRawText)
}
