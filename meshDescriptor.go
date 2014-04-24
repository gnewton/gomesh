package jianGoMeSHi


const DESCRIPTOR_RECORD = "DescriptorRecord"

type Abbreviation  struct{
	Abbreviation string `xml:",chardata"`
}

//type ActiveMeSHYearList struct{
//	Year int
//}

type AllowableQualifier struct{
	QualifierReferredTo QualifierReferredTo
	Abbreviation Abbreviation
}

type AllowableQualifiersList struct{
	AllowableQualifier []AllowableQualifier
}

type Concept struct{
	CASN1Name string
	ConceptName string `xml:">String"`
	ConceptRelationList ConceptRelationList
	ConceptUI string
	PreferredConceptYN string `xml:",attr"`
	RegistryNumber string
	RelatedRegistryNumberList RelatedRegistryNumberList 
	ScopeNote string
	SemanticTypeList SemanticTypeList 
	TermList TermList
}

type ConceptList struct{
	Concept []Concept
}

type ConceptRelation struct{
	RelationName string `xml:",attr"`
	Concept1UI string
	Concept2UI string
}

type ConceptRelationList struct{
	ConceptRelation []ConceptRelation
}


type Date struct{
	Year string
	Month string
	Day string
}


type DescriptorRecord struct {
	ActiveMeSHYearList []string `xml:">Year"`
	AllowableQualifiersList AllowableQualifiersList
	Annotation string
	ConceptList ConceptList
	ConsiderAlso string
	DateCreated Date
	DateEstablished Date
	DateRevised Date
	DescriptorName string `xml:"DescriptorName>String"`
	DescriptorUI string
	HistoryNote string
	EntryCombinationList EntryCombinationList
	PharmacologicalActionList PharmacologicalActionList
	PreviousIndexingList PreviousIndexingList
	RecordOriginatorsList RecordOriginatorsList
	SeeRelatedList SeeRelatedList
	TreeNumberList TreeNumberList
}

type DescriptorRecordSet struct{
	DescriptorRecord []DescriptorRecord
}

type DescriptorReferredTo struct{
	DescriptorUI string
	DescriptorName string `xml:"DescriptorName>String"`
	DescriptorRecord *DescriptorRecord  `xml:"-"`
}

type ECIN struct{
	DescriptorReferredTo DescriptorReferredTo
}

type ECOUT struct{
	DescriptorReferredTo DescriptorReferredTo
}

type EntryCombination struct{
	ECIN ECIN
	ECOUT ECOUT
}

type EntryCombinationList struct{
	EntryCombination []EntryCombination
}

type PharmacologicalAction struct{
	DescriptorReferredTo DescriptorReferredTo
}

type PharmacologicalActionList struct{
	PharmacologicalAction []PharmacologicalAction
}

type  PreviousIndexingList struct{
	PreviousIndexing []string
}

type QualifierReferredTo struct{
	QualifierUI string
	QualifierName string `xml:"QualifierName>String"`
	QualifierRecord *QualifierRecord `xml:"-"`
}



type RecordOriginator struct{
	RecordOriginator string
	RecordMaintainer string
	RecordAuthorizer string
}

type RecordOriginatorsList struct{
	RecordOriginator []RecordOriginator
}

type RelatedRegistryNumber struct{
	RelatedRegistryNumber string  `xml:",chardata"`
}

type RelatedRegistryNumberList struct{
	RelatedRegistryNumber []RelatedRegistryNumber
}

type SeeRelatedDescriptor struct{
	DescriptorReferredTo DescriptorReferredTo
}

type SeeRelatedList struct{
	SeeRelatedDescriptor []SeeRelatedDescriptor
}

type SemanticType struct{
	SemanticTypeUI string
	SemanticTypeName string
}

type SemanticTypeList struct{
	SemanticType []SemanticType
}

type Term struct{
	ConceptPreferredTermYN string `xml:",attr"`
	DateCreated Date
	EntryVersion string
	IsPermutedTermYN string `xml:",attr"`
	LexicalTag string `xml:",attr"`
	PrintFlagYN string `xml:",attr"`
	RecordPreferredTermYN string `xml:",attr"`
	String string
	TermUI string
	ThesaurusIDlist ThesaurusIDlist
}

type TermList struct{
	Term []Term
}

type ThesaurusIDlist struct{
	ThesaurusID []string
}

type TreeNumberList struct{
	TreeNumber []string
}
