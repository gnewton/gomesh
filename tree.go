package jianGoMeSHi

import (
	"fmt"
	"strings"
)


type Node struct{
	foo string `json:"foo"`
	treeNumber string `json:"treeNumber`
	nodeNumber string `json:`
	children map[string]*Node `json:"-"`
	allDescriptors map[string]bool `json:"-"`
	descriptor *DescriptorRecord `json:"-"`
}

func (node *Node) traverse(depth int){
	if node == nil{
		return
	}
	if depth == 1{
		fmt.Println("---------------------------------------------------------------------")
	}
	spc := strings.Repeat(" ", depth)
	var str string
	if node.descriptor != nil{
		str = node.descriptor.DescriptorUI + " " + node.descriptor.DescriptorName
	}
	fmt.Println(spc, node.nodeNumber, str)
	for _, child := range node.children{
		child.traverse(depth +1)
	}
}

func (node *Node) Init()(*Node){
	if node == nil{
		node = new(Node)
		node.children = make(map[string]*Node)
		node.allDescriptors = make(map[string]bool)
	}
	return node
}

func addDescriptor(root *Node, rec *DescriptorRecord){
	if rec == nil{
		return
	}

	if root == nil{
		return
	}

	if rec.TreeNumberList.TreeNumber != nil{
		for _, treeNumber := range rec.TreeNumberList.TreeNumber {
			addTreeNumber(root, rec, treeNumber)
		}
	}
	
}

const TREE_SEPARATOR = "."
func addTreeNumber(root *Node, rec *DescriptorRecord, treeNumber string){
	_, ok := root.allDescriptors[treeNumber]
	if ok{
		return
	}
	parts := strings.Split(treeNumber, TREE_SEPARATOR)

	thisTree := ""
	node := root
	for index,part := range parts{
		if index > 0{
			thisTree += TREE_SEPARATOR
		}
		thisTree += part

		child, ok := node.children[part]
		if !ok{
			var nd *Node
			child = nd.Init()
			child.treeNumber = thisTree
			child.nodeNumber = part
			child.children["foo"] = nil
			node.children[child.nodeNumber] = child
		}
		node = child
	}
	root.allDescriptors[treeNumber] = true
	node.descriptor = rec
}