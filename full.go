package gomesh

import ()

func SelfLinkDescriptor(m map[string]*DescriptorRecord) {
	var nd *Node
	root := nd.Init()

	for _, v := range m {
		//linkSeeRelated(m, v)
		linkConceptRelations(m, v)
		makeTree(root, v)
	}
	root.Traverse(0, nil)
}

func MakeTree(m map[string]*DescriptorRecord) (root *Node) {
	var nd *Node
	root = nd.Init()
	for _, v := range m {
		//linkSeeRelated(m, v)
		linkConceptRelations(m, v)
		makeTree(root, v)
	}
	return addTopLevels(root)
}

func makeTree(root *Node, desc *DescriptorRecord) {
	addDescriptor(root, desc)
}

func linkSeeRelated(m map[string]*DescriptorRecord, desc *DescriptorRecord) {
	seeRelatedList := desc.SeeRelatedList
	if &seeRelatedList != nil {
		seeRelatedDescriptors := seeRelatedList.SeeRelatedDescriptor
		for _, srd := range seeRelatedDescriptors {
			refDesc, ok := m[srd.DescriptorReferredTo.DescriptorUI]
			if ok {
				srd.DescriptorReferredTo.DescriptorRecord = refDesc
				//fmt.Println(desc.DescriptorUI, desc.DescriptorName, "links to", srd.DescriptorReferredTo.DescriptorUI, srd.DescriptorReferredTo.DescriptorName)
			}
		}
	}
}

func linkConceptRelations(m map[string]*DescriptorRecord, desc *DescriptorRecord) {
	conceptList := desc.ConceptList
	if &conceptList != nil {
		concepts := conceptList.Concept
		for _, concept := range concepts {
			conceptRelationList := concept.ConceptRelationList
			if &conceptRelationList != nil {
				_ = conceptRelationList.ConceptRelation
				//conceptRelations := conceptRelationList.ConceptRelation
				//for _, conceptRelation := range conceptRelations{
				//fmt.Println(concept.ConceptName, " -- ", conceptRelation.RelationName, conceptRelation.Concept1UI, conceptRelation.Concept2UI)
				//}
			}
		}
	}
}
