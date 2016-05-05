package gomesh2014

const QUALIFIER_RECORD = "QualifierRecord"

type QualifierRecordSet struct {
	QualifierRecord []*QualifierRecord
}

type QualifierRecord struct {
	QualifierUI           string                `json:",omitempty"`
	QualifierName         string                `xml:"QualifierName>String" json:",omitempty"`
	DateCreated           Date                  `json:",omitempty"`
	DateRevised           Date                  `json:",omitempty"`
	DateEstablished       Date                  `json:",omitempty"`
	ActiveMeSHYearList    []string              `xml:">Year" json:",omitempty"`
	Annotation            string                `json:",omitempty"`
	HistoryNote           string                `json:",omitempty"`
	OnlineNote            string                `json:",omitempty"`
	TreeNumberList        TreeNumberList        `json:",omitempty"`
	TreeNodeAllowedList   TreeNodeAllowedList   `json:",omitempty"`
	RecordOriginatorsList RecordOriginatorsList `json:",omitempty"`
	ConceptList           ConceptList           `json:",omitempty"`
}

type TreeNodeAllowedList struct {
	TreeNodeAllowed []TreeNodeAllowed `json:",omitempty"`
	//TreeNodeAllowed []string `xml:">TreeNodeAllowed" json:",omitempty"`
}

type TreeNodeAllowed struct {
	TreeNodeAllowed string `xml:",chardata" json:",omitempty"`
}
