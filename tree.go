package jianGoMeSHi

import (
//	"fmt"
	"strings"
//	"log"
)


type Node struct{
	TreeNumber string `json:"treeNumber`
	NodeNumber string  `json:"-"`
	Name string `json:",omitempty"`
	Children []IdEntry `json:",omitempty"`
	ChildrenMap map[string]*Node `json:"-"`
	AllDescriptors map[string]bool `json:"-"`
	DescriptorUrl string `json:",omitempty"`
	Descriptor *DescriptorRecord `json:"-"`
}

var TOP_LEVEL = map[string]string{
	"A": "Anatomy", 
	"B": "Organisms", 
	"C": "Diseases", 
	"D": "Chemicals and Drugs", 
	"E": "Analytical, Diagnostic and Therapeutic Techniques and Equipment", 
	"F": "Psychiatry and Psychology", 
	"G": "Phenomena and Processes", 
	"H": "Disciplines and Occupations", 
	"I": "Anthropology, Education, Sociology and Social Phenomena", 
	"J": "Technology, Industry, Agriculture", 
	"K": "Humanities", 
	"L": "Information Science", 
	"M": "Named Groups", 
	"N": "Health Care", 
	"V": "Publication Characteristics", 
	"Z": "Geographicals",  
}

type IdEntry struct{
	Id string
	Url string
	Label string `json:",omitempty"`
}


type useNode func(*Node)

func (node *Node) Traverse(depth int, nodeUser useNode){
	if node == nil{
		return
	}
	if node.Descriptor != nil{
		//str = node.Descriptor.DescriptorUI + " " + node.Descriptor.DescriptorName
		node.Name = node.Descriptor.DescriptorName
	}
	for _, childNode := range node.ChildrenMap{
		childNode.Traverse(depth +1, nodeUser)
	}
	if nodeUser != nil{
		nodeUser(node)
	}
}

func (node *Node) Init()(*Node){
	if node == nil{
		node = new(Node)
		node.ChildrenMap = make(map[string]*Node)
		node.AllDescriptors = make(map[string]bool)
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
			addTreeNumber(root, rec, treeNumber.TreeNumber)
		}
	}
	
}

const TREE_SEPARATOR = "."
func addTreeNumber(root *Node, rec *DescriptorRecord, treeNumber string){
	_, ok := root.AllDescriptors[treeNumber]
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

		child, ok := node.ChildrenMap[part]
		if !ok{
			var nd *Node
			child = nd.Init()
			child.TreeNumber = thisTree
			//fmt.Println("---- " + thisTree)
			child.NodeNumber = part
			node.ChildrenMap[child.NodeNumber] = child
		}
		node = child
	}
	root.AllDescriptors[treeNumber] = true
	node.Descriptor = rec
}

func addTopLevels(root *Node) *Node{
	var nd *Node
	top := nd.Init()
	top.Children = make([]IdEntry, len(TOP_LEVEL))
	i := 0
	for key, label := range TOP_LEVEL {
		entry := new(IdEntry)
		entry.Label = label
		entry.Id = key
		top.Children[i] = *entry
		//log.Println(key, label)
		//log.Printf("%+v\n", top)

		childNode := nd.Init()
		childNode.TreeNumber = key
		childNode.Name = label
		top.ChildrenMap[key] = childNode
		///log.Println(root.ChildrenMap)
		findTopChildren(childNode, root.ChildrenMap)

		i += 1
	}
	//log.Printf("*** %+v\n", top)
	return top
}

func findTopChildren(node *Node, rootChildren map[string]*Node){
	for key,child := range rootChildren {
		//fmt.Println("~~ ", string(key[0]), node.TreeNumber)
		if string(key[0]) == node.TreeNumber{
			//fmt.Println("+++++++++++++++++++++ ", child.TreeNumber, node.Name)
			node.ChildrenMap[child.TreeNumber] = child
		}
	}
}