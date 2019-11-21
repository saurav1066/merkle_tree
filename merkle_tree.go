package merkle_tree

import (
	"crypto/sha256"
	_ "crypto/sha256"
)

type MerkleNode struct{
	Left *MerkleNode
	Right *MerkleNode
	Data []byte
}

type MerkleTree struct {
	Rootnode *MerkleNode
}
func NewMerkleNode(left, right *MerkleNode, data []byte) *MerkleNode {
	node := MerkleNode{}

	if left == nil && right == nil {
		hash := sha256.Sum256(data)
		node.Data = hash[:]
	} else {
		prevHashes := append(left.Data , right.Data ...)
		hash := sha256.Sum256(prevHashes)
		node.Data = hash[:]
	}
	node.Left = left
	node.Right = right

	return &node
}

func NewMerkleTree(data [][]byte) *MerkleTree {
	var nodes []MerkleNode

	if len(data)%2 != 0 {
		data = append(data, data[len(data) -1])
	}

	for _, dat :=range data {
		node :=NewMerkleNode(nil,nil, dat)
		nodes = append(nodes, *node)
	}

	for i :=0; i< (len(data)/2); i++{
		var level []MerkleNode
		if(len(nodes))==1{
			break
		}
		for j :=0; j <len(nodes);j+= 2 {
			node := NewMerkleNode(&nodes[j],&nodes[j+1], nil)
			level = append(level, *node)
			//print(level)
		}
		//fmt.Println(level)
		//println(level)
		nodes = level


	}
	//fmt.Println(nodes)
	tree := MerkleTree{&nodes[0]}
	return &tree
}
