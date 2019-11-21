package main

import (
	"fmt"
	"merkle_tree"
)

func main() {
	var a []byte =[]byte("2F62B3AC357EDDF37DD3D45D37104F446509EE8575ED4A20F48C0623427D0855")
	//fmt.Println(a)
	i := 0
	txhex:=[][]byte{}
	for i=0;i<len(a) ;i+=2  {
		row1 := []byte{a[i]}
		row2 := []byte{a[i+1]}
		txhex=append(txhex, row1)
		txhex=append(txhex, row2)
	}
	//fmt.Println(txhex)
	tree := merkle_tree.NewMerkleTree(txhex)
	data:=tree.Rootnode.Left.Left.Left.Left.Left
	fmt.Println(data)
}