package main

import (
        "github.com/ant0ine/go-json-rest/rest"
	"github.com/gnewton/jianGoMeSHi"
        "net/http"
        "log"
        "os"
        "runtime"
        "sort"
        "fmt"
        "strings"
)

//const URL_HOST = "http://localhost"
var URL_HOST string
const PORT = "8080"
const PATH = "/mesh"

var BASE_URL string

// Change these to match the MeSH files you download from http://www.nlm.nih.gov/mesh/ist.html
// The XML files are desc2014, supp2014, qual2014
// You can compress these with either gz or bz2 & this app will transparently uncompress them. Or you can leave them as-is.
//
const DESCRIPTOR_XML_FILE = "../testData/desc2014_29records.xml.bz2"
//const DESCRIPTOR_XML_FILE = "/home/newtong/2014/mesh/desc2014.xml.bz2"
const QUALIFIER_XML_FILE = "../testData/qual2014_8records.xml.bz2"
//const QUALIFIER_XML_FILE = "/home/newtong/2014/mesh/qual2014.xml.bz2"
const SUPPLEMENTAL_XML_FILE = "../testData/supp2014_4records.xml"
//const SUPPLEMENTAL_XML_FILE = "/home/newtong/2014/mesh/supp2014.xml.bz2"

const DESCRIPTOR = "descriptor"
const QUALIFIER = "qualifier"
const SUPPLEMENTAL = "supplemental"
const TREE = "tree"

var NOUNS = []string{DESCRIPTOR, QUALIFIER, SUPPLEMENTAL, TREE}

var descMap map[string]*jianGoMeSHi.DescriptorRecord
var descMap2 map[string]*LocalDesc
var descSlice  []*jianGoMeSHi.IdEntry



var suppMap map[string]*jianGoMeSHi.SupplementalRecord
var suppSlice  []*jianGoMeSHi.IdEntry

var qualMap map[string]*jianGoMeSHi.QualifierRecord
var qualSlice  []*jianGoMeSHi.IdEntry

var root *jianGoMeSHi.Node
var treeMap map[string]*jianGoMeSHi.Node

var allNouns []jianGoMeSHi.IdEntry


type LocalDesc jianGoMeSHi.DescriptorRecord

func GetAll(w rest.ResponseWriter, req *rest.Request) {
	if req.Request.Method != "GET"{
		return
	}
	w.WriteJson(allNouns)
}


func GetAllDescriptors(w rest.ResponseWriter, req *rest.Request) {
	log.Println("-------- ", req.PathParam("start"))
	if req.Request.Method != "GET"{
		return
	}
	w.WriteJson(descSlice)
}


func GetDescriptor(w rest.ResponseWriter, req *rest.Request) {
	if req.Request.Method != "GET"{
		return
	}
	descriptorUI := req.PathParam("id")

	//descriptor, ok := descMap[descriptorUI]
	_, ok := descMap[descriptorUI]
	if ok{
		//var descriptorWithUrl jianGoMeSHi.DescriptorRecord
		
		//descriptorWithUrl = *descriptor
		//w.WriteJson(populateWithUrl(descriptor, "http://localhost:8080/descriptor"))
		w.WriteJson(descMap2[descriptorUI])
	}else{
		rest.NotFound(w,req)
	}
}


func GetAllSupplementals(w rest.ResponseWriter, req *rest.Request) {
	if req.Request.Method != "GET"{
		return
	}
	w.WriteJson(suppSlice)
}


func GetSupplemental(w rest.ResponseWriter, req *rest.Request) {
	if req.Request.Method != "GET"{
		return
	}
	supplementalUI := req.PathParam("id")
	
	supplemental, ok := suppMap[supplementalUI]
	if ok{
		w.WriteJson(supplemental)
	}else{
		rest.NotFound(w,req)
	}

}


func GetAllQualifiers(w rest.ResponseWriter, req *rest.Request) {
	if req.Request.Method != "GET"{
		return
	}
	w.WriteJson(qualSlice)
}


func GetQualifier(w rest.ResponseWriter, req *rest.Request) {
	if req.Request.Method != "GET"{
		return
	}
	qualifierUI := req.PathParam("id")
	
	qualifier, ok := qualMap[qualifierUI]
	if ok{
		w.WriteJson(qualifier)
	}else{
		rest.NotFound(w,req)
	}
}


func GetTrees(w rest.ResponseWriter, req *rest.Request) {
	if req.Request.Method != "GET"{
		return
	}
	//var nd *jianGoMeSHi.Node
	//child := nd.Init()
	//w.WriteJson(root)
	//log.Println(root)
	log.Printf("%+v\n", root)
	w.WriteJson(root.Children)
	//log.Println(root.Children["D02"].Children["705"])
	//w.WriteJson("hello")
}


func GetTree(w rest.ResponseWriter, req *rest.Request) {
	if req.Request.Method != "GET"{
		return
	}
	treeNumber := req.PathParam("a")
	
	if req.PathParam("b") != ""{
		treeNumber = treeNumber + "." + req.PathParam("b")
	}
	if req.PathParam("c") != ""{
		treeNumber = treeNumber + "." + req.PathParam("c")
	}
	if req.PathParam("d") != ""{
		treeNumber = treeNumber + "." + req.PathParam("d")
	}
	if req.PathParam("e") != ""{
		treeNumber = treeNumber + "." + req.PathParam("e")
	}
	if req.PathParam("f") != ""{
		treeNumber = treeNumber + "." + req.PathParam("f")
	}
	if req.PathParam("g") != ""{
		treeNumber = treeNumber + "." + req.PathParam("f")
	}
	//log.Println("---- id", treeNumber)
	node, ok := treeMap[treeNumber]
	log.Println(node)
	if ok{
		w.WriteJson(node)
	}else{
		rest.NotFound(w,req)
	}
}

func (desc *LocalDesc) setTreeNumberUrls(baseUrl string){
	if desc.TreeNumberList.TreeNumber != nil{
		for i:=0; i<len(desc.TreeNumberList.TreeNumber); i++{
			tn := &(desc.TreeNumberList.TreeNumber[i])
			tn.Url = baseUrl + "/" + TREE + "/" + treeToUrlPath(tn.TreeNumber)
		}
	}
}

func treeToUrlPath(treeNumber string)string{
	return strings.Replace(treeNumber, ".", "/", -1)
}

func (desc *LocalDesc) setDescUrls(baseUrl string){
	if desc.PharmacologicalActionList.PharmacologicalAction != nil{
		for i:=0; i<len(desc.PharmacologicalActionList.PharmacologicalAction); i++{
			ref := &(desc.PharmacologicalActionList.PharmacologicalAction[i])
			ref.DescriptorReferredTo.Url = baseUrl + "/" + DESCRIPTOR + "/" + ref.DescriptorReferredTo.DescriptorUI
		}
	}

	if desc.SeeRelatedList.SeeRelatedDescriptor != nil{
		for i:=0; i<len(desc.SeeRelatedList.SeeRelatedDescriptor); i++{
			ref := &(desc.SeeRelatedList.SeeRelatedDescriptor[i])
			ref.DescriptorReferredTo.Url = baseUrl + "/" + DESCRIPTOR + "/" + ref.DescriptorReferredTo.DescriptorUI
		}
	}
}

func main() {
	URL_HOST,_ := os.Hostname()
	BASE_URL ="http://" + URL_HOST + ":" + PORT + PATH
	log.Println(os.Hostname())
	runtime.GOMAXPROCS(runtime.NumCPU())
	err := loadData()
	if err != nil{
		fmt.Println(err)
		return 
	}
        handler := rest.ResourceHandler{}
        handler.SetRoutes(
		&rest.Route{"HEAD", PATH, GetAll},

		&rest.Route{"GET", PATH, GetAll},

                &rest.Route{"GET", PATH + "/" + DESCRIPTOR + "\\?start=/:start/", GetAllDescriptors},
		&rest.Route{"GET", PATH + "/" + DESCRIPTOR, GetAllDescriptors},
		&rest.Route{"GET", PATH + "/" + DESCRIPTOR + "/", GetAllDescriptors},
                &rest.Route{"GET", PATH + "/" + DESCRIPTOR + "/:id", GetDescriptor},


		&rest.Route{"GET", PATH + "/" + SUPPLEMENTAL + "/:id", GetSupplemental},
		&rest.Route{"GET", PATH + "/" + SUPPLEMENTAL, GetAllSupplementals},
		&rest.Route{"GET", PATH + "/" + SUPPLEMENTAL + "/", GetAllSupplementals},

		&rest.Route{"GET", PATH + "/" + QUALIFIER + "/:id", GetQualifier},
		&rest.Route{"GET", PATH + "/" + QUALIFIER + "/", GetAllQualifiers},
		&rest.Route{"GET", PATH + "/" + QUALIFIER, GetAllQualifiers},

		&rest.Route{"GET", PATH + "/" + TREE, GetTrees},
		&rest.Route{"GET", PATH + "/" + TREE + "/", GetTrees},
		&rest.Route{"GET", PATH + "/" + TREE+ "/:a", GetTree},
		&rest.Route{"GET", PATH + "/" + TREE+ "/:a/", GetTree},
		&rest.Route{"GET", PATH + "/" + TREE+ "/:a/:b", GetTree},
		&rest.Route{"GET", PATH + "/" + TREE+ "/:a/:b/:c", GetTree},
		&rest.Route{"GET", PATH + "/" + TREE+ "/:a/:b/:c/:d", GetTree},
		&rest.Route{"GET", PATH + "/" + TREE+ "/:a/:b/:c/:d/:e", GetTree},
		&rest.Route{"GET", PATH + "/" + TREE+ "/:a/:b/:c/:d/:e/:f", GetTree},
		&rest.Route{"GET", PATH + "/" + TREE+ "/:a/:b/:c/:d/:e/:f/:g", GetTree},
        )
        http.ListenAndServe(":" + PORT, &handler)
}


func loadData()(error){
	treeMap = make(map[string]*jianGoMeSHi.Node)
	var err error
	log.Println("Start Loading MeSH XML...")

	log.Println("\tLoading Supplemental MeSH XML from file: ", SUPPLEMENTAL_XML_FILE)
	suppMap, err = jianGoMeSHi.SupplementalMapFromFile(SUPPLEMENTAL_XML_FILE)
	if err != nil{
		return err
	}
	index := 0

	suppSlice = make([]*jianGoMeSHi.IdEntry, len(suppMap))

	for supp := range suppMap{
		newEntry := new(jianGoMeSHi.IdEntry)
		newEntry.Id = suppMap[supp].SupplementalRecordUI
		newEntry.Url = BASE_URL + "/" + SUPPLEMENTAL + "/" + newEntry.Id
		suppSlice[index] = newEntry
		index += 1
	}


	log.Println("\tLoading Qualifier MeSH XML from file:", QUALIFIER_XML_FILE)
	qualMap, err = jianGoMeSHi.QualifierMapFromFile(QUALIFIER_XML_FILE)
	if err != nil{
		return err
	}

	qualSlice = make([]*jianGoMeSHi.IdEntry, len(qualMap))
	index = 0
	for qual := range qualMap{
		newEntry := new(jianGoMeSHi.IdEntry)
		newEntry.Id = qualMap[qual].QualifierUI
		newEntry.Url = BASE_URL + "/" + QUALIFIER + "/" + newEntry.Id
		qualSlice[index] = newEntry
		index += 1
	}

	log.Println("\tLoading Descriptor MeSH XML from file: ", DESCRIPTOR_XML_FILE)
	descMap, err = jianGoMeSHi.DescriptorMapFromFile(DESCRIPTOR_XML_FILE)
	if err != nil{
		return err
	}
	log.Println("Building name map")
	_ = jianGoMeSHi.MeshDescriptorNameMap(descMap)

	descSlice = make([]*jianGoMeSHi.IdEntry, len(descMap))
	index = 0
	descMap2 = make(map[string]*LocalDesc)

	for desc := range descMap{
		newEntry := new(jianGoMeSHi.IdEntry)
		descriptorRecord := descMap[desc]
		var localDesc = (*LocalDesc)(descriptorRecord)
		localDesc.setDescUrls(BASE_URL)
		localDesc.setTreeNumberUrls(BASE_URL)
		
		descMap2[desc] = localDesc
		newEntry.Id = descMap[desc].DescriptorUI
		newEntry.Url = BASE_URL + "/" + DESCRIPTOR + "/" + newEntry.Id
		descSlice[index] = newEntry
		index += 1
	}

	sort.Sort(ById(descSlice))
	sort.Sort(ById(qualSlice))
	sort.Sort(ById(suppSlice))

	root = jianGoMeSHi.MakeTree(descMap)
	log.Printf("*** 1  %+v\n", root)
	root.Traverse(0, AddUrlInfo)
	log.Printf("*** 2  %+v\n", root)
	//log.Println(root)

	sort.Sort(ByIdX(root.Children))
	

	log.Println("Done Loading MeSH XML...")

	allNouns = make([]jianGoMeSHi.IdEntry, len(NOUNS))
	for i,noun := range NOUNS{
		allNouns[i].Id = "/" + noun
		allNouns[i].Url = BASE_URL + "/" + noun
	}


	return nil
}

func AddUrlInfo(node *jianGoMeSHi.Node){
	//fmt.Println("AddUrlInfo", node.TreeNumber)
	treeMap[node.TreeNumber] = node
	if node.Children == nil{
		node.Children = make([]jianGoMeSHi.IdEntry, len(node.ChildrenMap))
		if node.Descriptor != nil{
			node.DescriptorUrl = BASE_URL + "/" + DESCRIPTOR + "/" + node.Descriptor.DescriptorUI
		}
	}
	i :=0
	for _,childNode := range node.ChildrenMap{
		node.Children[i].Id = childNode.TreeNumber
		node.Children[i].Url = BASE_URL + "/" + TREE + "/" + treeToUrlPath(childNode.TreeNumber)
		node.Children[i].Label = childNode.Name
		i++
	}
}


//sort slices

type ByIdX []jianGoMeSHi.IdEntry

type ById []*jianGoMeSHi.IdEntry

func (a ById) Len() int           { return len(a) }
func (a ByIdX) Len() int           { return len(a) }
func (a ById) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByIdX) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ById) Less(i, j int) bool { return a[i].Id < a[j].Id }
func (a ByIdX) Less(i, j int) bool { return a[i].Id < a[j].Id }

