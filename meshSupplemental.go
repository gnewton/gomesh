package gomesh

const SUPPLEMENTAL_RECORD = "SupplementalRecord"

type SupplementalRecordSet struct {
	SupplementalRecord []SupplementalRecord
}

type SupplementalRecord struct {
	SupplementalRecordUI    string
	SupplementalRecordName  string                  `xml:"SupplementalRecordName>String"`
	DateCreated             Date                    `json:",omitempty"`
	DateEstablished         Date                    `json:",omitempty"`
	DateRevised             Date                    `json:",omitempty"`
	ActiveMeSHYearList      []string                `xml:">Year"`
	Note                    string                  `json:",omitempty"`
	Frequency               string                  `json:",omitempty"`
	HeadingMappedToList     HeadingMappedToList     `json:",omitempty"`
	IndexingInformationList IndexingInformationList `json:",omitempty"`
	SourceList              SourceList              `json:",omitempty"`
	RecordOriginatorsList   RecordOriginatorsList   `json:",omitempty"`
	ConceptList             ConceptList             `json:",omitempty"`
}

type HeadingMappedToList struct {
	HeadingMappedTo []HeadingMappedTo `json:",omitempty"`
}

type HeadingMappedTo struct {
	DescriptorReferredTo *DescriptorReferredTo `json:",omitempty"`
}

type IndexingInformationList struct {
	IndexingInformation []IndexingInformation `json:",omitempty"`
}

type IndexingInformation struct {
	DescriptorReferredTo *DescriptorReferredTo `json:",omitempty"`
}

type SourceList struct {
	Source []Source `json:",omitempty"`
}

type Source struct {
	Source string `json:",omitempty"`
}
