package main

import (
        "github.com/ant0ine/go-json-rest/rest"
	"github.com/gnewton/jianGoMeSHi"
        "net/http"
        "log"
        "runtime"
        "fmt"
)



// Change these to match the MeSH files you download from http://www.nlm.nih.gov/mesh/filelist.html
// The XML files are desc2014, supp2014, qual2014
// You can compress these with either gz or bz2 & this app will transparently uncompress them. Or you can leave them as-is.
//
const DESCRIPTOR_XML_FILE = "../testData/desc2014_29records.xml.bz2"
const QUALIFIER_XML_FILE = "../testData/qual2014_8records.xml.bz2"
const SUPPLEMENTAL_XML_FILE = "../testData/supp2014_4records.xml"


var descMap map[string]*jianGoMeSHi.DescriptorRecord
var descSlice  []string
var suppMap map[string]*jianGoMeSHi.SupplementalRecord
var suppSlice  []string
var qualMap map[string]*jianGoMeSHi.QualifierRecord
var qualSlice  []string

var root *jianGoMeSHi.Node

func GetAllDescriptors(w rest.ResponseWriter, req *rest.Request) {
	w.WriteJson(descSlice)
}

func GetDescriptor(w rest.ResponseWriter, req *rest.Request) {
	descriptorUI := req.PathParam("id")
	
	descriptor, ok := descMap[descriptorUI]
	if ok{
		w.WriteJson(descriptor)
	}else{
		rest.NotFound(w,req)
	}
}

func GetAllSupplementals(w rest.ResponseWriter, req *rest.Request) {
	w.WriteJson(suppSlice)
}


func GetSupplemental(w rest.ResponseWriter, req *rest.Request) {
	supplementalUI := req.PathParam("id")
	
	supplemental, ok := suppMap[supplementalUI]
	if ok{
		w.WriteJson(supplemental)
	}else{
		rest.NotFound(w,req)
	}

}

func GetAllQualifiers(w rest.ResponseWriter, req *rest.Request) {
	w.WriteJson(qualSlice)
}

func GetQualifier(w rest.ResponseWriter, req *rest.Request) {
	qualifierUI := req.PathParam("id")
	
	qualifier, ok := qualMap[qualifierUI]
	if ok{
		w.WriteJson(qualifier)
	}else{
		rest.NotFound(w,req)
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	err := loadData()
	if err != nil{
		fmt.Println(err)
		return 
	}
        handler := rest.ResourceHandler{}
        handler.SetRoutes(
		&rest.Route{"GET", "/descriptor", GetAllDescriptors},
		&rest.Route{"GET", "/descriptor/", GetAllDescriptors},
                &rest.Route{"GET", "/descriptor/:id", GetDescriptor},
		&rest.Route{"GET", "/supplemental/:id", GetSupplemental},
		&rest.Route{"GET", "/supplemental", GetAllSupplementals},
		&rest.Route{"GET", "/supplemental/", GetAllSupplementals},
		&rest.Route{"GET", "/qualifier/:id", GetQualifier},
		&rest.Route{"GET", "/qualifier/", GetAllQualifiers},
		&rest.Route{"GET", "/qualifier", GetAllQualifiers},
        )
        http.ListenAndServe(":8080", &handler)
}

func loadData()(error){
	var err error
	log.Println("Start Loading MeSH XML...")

	log.Println("\tLoading Supplemental MeSH XML from file: ", SUPPLEMENTAL_XML_FILE)
	suppMap, err = jianGoMeSHi.SupplementalMapFromFile(SUPPLEMENTAL_XML_FILE)
	if err != nil{
		return err
	}
	suppSlice = make([]string, len(suppMap))
	index := 0
	for supp := range suppMap{
		suppSlice[index] = supp
		index += 1
	}


	log.Println("\tLoading Qualifier MeSH XML from file:", QUALIFIER_XML_FILE)
	qualMap, err = jianGoMeSHi.QualifierMapFromFile(QUALIFIER_XML_FILE)
	if err != nil{
		return err
	}

	qualSlice = make([]string, len(qualMap))
	index = 0
	for qual := range qualMap{
		qualSlice[index] = qual
		index += 1
	}

	log.Println("\tLoading Descriptor MeSH XML from file: ", DESCRIPTOR_XML_FILE)
	descMap, err = jianGoMeSHi.DescriptorMapFromFile(DESCRIPTOR_XML_FILE)
	if err != nil{
		return err
	}

	descSlice = make([]string, len(descMap))
	index = 0
	for desc := range descMap{
		descSlice[index] = desc
		index += 1
	}


	root = jianGoMeSHi.MakeTree(descMap)

	log.Println("Done Loading MeSH XML...")
	return nil
}