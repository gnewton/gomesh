package jianGoMeSHi

const QUALIFIER_RECORD = "QualifierRecord"

type QualifierRecordSet struct{
	QualifierRecord []QualifierRecord
}

type QualifierRecord struct{
	QualifierUI string
	QualifierName string `xml:">QualifierName>String"`
	DateCreated Date
	DateRevised Date
	DateEstablished Date
	ActiveMeSHYearList []string `xml:">Year"`
	Annotation string
	HistoryNote string
	OnlineNote string
	TreeNumberList TreeNumberList
	TreeNodeAllowedList TreeNodeAllowedList
	RecordOriginatorsList RecordOriginatorsList
	ConceptList ConceptList

}

type TreeNodeAllowedList struct{
	TreeNodeAllowed []TreeNodeAllowedList
}

type TreeNodeAllowed struct{
	TreeNodeAllowed string
}
