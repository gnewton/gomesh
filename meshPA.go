package jianGoMeSHi

const PHARMACOLOGICAL_RECORD = "PharmacologicalAction"

type PharmacologicalActionSet struct{
	PharmacologicalAction []PharmacologicalAction
}


type Substance struct{
	RecordUI string
	DescriptorUrl string `json:",omitempty"`
	SupplementalUrl string `json:",omitempty"`
	RecordName string `xml:"RecordName>String"`
}

