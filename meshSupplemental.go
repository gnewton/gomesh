package jianGoMeSHi

const SUPPLEMENTAL_RECORD = "SupplementalRecord"

type SupplementalRecordSet struct{
	SupplementalRecord []SupplementalRecord
}


type SupplementalRecord struct{
	SupplementalRecordUI string
	SupplementalRecordName string
	DateCreated Date
	DateEstablished Date
	DateRevised Date
	ActiveMeSHYearList []string `xml:">Year"`
	Note string
	Frequency string
	HeadingMappedToList HeadingMappedToList
	IndexingInformationList IndexingInformationList
	SourceList SourceList
	RecordOriginatorsList RecordOriginatorsList
	ConceptList ConceptList
}

type HeadingMappedToList struct {
	HeadingMappedTo []HeadingMappedTo
}

type HeadingMappedTo struct{
	DescriptorReferredTo DescriptorReferredTo
}

type IndexingInformationList struct{
	IndexingInformation []IndexingInformation
}

type  IndexingInformation struct{
	DescriptorReferredTo DescriptorReferredTo
}

type  SourceList struct{
	Source []Source
}

type  Source struct{
	Source string
}

