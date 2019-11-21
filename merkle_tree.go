package merkle_tree

import (
	"crypto/sha256"
	_ "crypto/sha256"
)

type MerkleNode struct{     //MerkleNode structure (recursive)
	Left *MerkleNode		//left and right fields referencing other merkle node structure
	Right *MerkleNode
	Data []byte				//to store data associated with each feilds
}

type MerkleTree struct {	//Merklee tree structure
	Rootnode *MerkleNode	//merkle root containing pointer to merkle Node
}
func NewMerkleNode(left, right *MerkleNode, data []byte) *MerkleNode { 	//function to create a new merkle node
	node := MerkleNode{}

	if left == nil && right == nil {	//if left and right node exists
		hash := sha256.Sum256(data)		//make a hash
		node.Data = hash[:]				//put the hash into the data field
	} else {
		prevHashes := append(left.Data , right.Data ...)	//else take the data of left and right node and concatinate them
		hash := sha256.Sum256(prevHashes)		//again make a hash
		node.Data = hash[:]				//put hash into the data feild
	}
	node.Left = left  	//assign left node to the Left
	node.Right = right	//assign right node to the Right

	return &node		//return reference to the node
}

func NewMerkleTree(data [][]byte) *MerkleTree {		//function to create Merkle tree
	var nodes []MerkleNode							//nodes for the merkle tree

	if len(data)%2 != 0 {				//check if the data passed is even
		data = append(data, data[len(data) -1])	//if not even, take the last data twice to make it even
	}

	for _, dat :=range data {
		node :=NewMerkleNode(nil,nil, dat)	//data nil for left and right branch of leaf nodes put into new Merkle Node
		nodes = append(nodes, *node)	//append data to the nodes
	}

	for i :=0; i< (len(data)/2); i++{	//for loop goes up the left side of tree
		var level []MerkleNode			//levels as array of merkle node
		if(len(nodes))==1{				//terminating condition when root node is formed
			break
		}
		for j :=0; j <len(nodes);j+= 2 {	//for loop going up the length of the nodes and incremented by 2 each time
			node := NewMerkleNode(&nodes[j],&nodes[j+1], nil)	//let the left value [j] and right value [j+1] and nil for data
			level = append(level, *node) 		//append nodes to the level structure
			//print(level)
		}
		//fmt.Println(level)
		//println(level)
		nodes = level		//replace nodes variable with level structure


	}
	//fmt.Println(nodes)
	tree := MerkleTree{&nodes[0]}	//create new tree with node index zero into the merkle tree for merkle root
	return &tree //return tree
}
