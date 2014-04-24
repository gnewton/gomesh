package jianGoMeSHi

import (
	"bufio"
	"compress/bzip2"
	"compress/gzip"
	"encoding/xml"
	"io"
	"os"
	"strings"
)


func DescriptorMapFromFile(filename string)(map[string]*DescriptorRecord, error){
	reader, file, err := genericReader(filename)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	return DescriptorMapFromReader(reader)
}

func DescriptorMapFromReader(reader  io.Reader)(map[string]*DescriptorRecord, error){
	descChan, err := DescriptorChannelFromReader(reader)
	if err != nil{
		return nil, err
	}

	descMap := make(map[string]*DescriptorRecord)	
	for mRecord := range descChan {
		descMap[mRecord.DescriptorUI] = mRecord
	}
	return descMap, nil
}

func DescriptorChannelFromFile(filename string)(desChan chan *DescriptorRecord, file *os.File, err error){
	reader, file, err := genericReader(filename)
	//if fl, ok := reader.(*gzip.Reader); ok {
	//defer fl.Close()
	//} 
	if err != nil {
		return nil, nil, err
	}
	
	desChan, err = DescriptorChannelFromReader(reader)
	return desChan, file, err
}

func DescriptorChannelFromReader(reader  io.Reader)(chan *DescriptorRecord, error){
	desriptorChannel := make(chan *DescriptorRecord, 500)
	go decodeDescriptor(desriptorChannel, reader)
	
	return desriptorChannel, nil
}

func SupplementalMapFromFile(filename string)(map[string]*SupplementalRecord, error){
	reader, file, err := genericReader(filename)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	return SupplementalMapFromReader(reader)
}

func SupplementalMapFromReader(reader  io.Reader)(map[string]*SupplementalRecord, error){
	suppChan, err := SupplementalChannelFromReader(reader)
	if err != nil{
		return nil, err
	}

	suppMap := make(map[string]*SupplementalRecord)	
	for sRecord := range suppChan {
		suppMap[sRecord.SupplementalRecordUI] = sRecord
	}
	return suppMap, nil
}

func SupplementalChannelFromFile(filename string)(suppChan chan *SupplementalRecord, file *os.File, err error){
	reader, file, err := genericReader(filename)
	//if fl, ok := reader.(*gzip.Reader); ok {
	//defer fl.Close()
	//} 
	if err != nil {
		return nil, nil, err
	}
	
	suppChan, err = SupplementalChannelFromReader(reader)
	return suppChan, file, err
}



func SupplementalChannelFromReader(reader  io.Reader)(chan *SupplementalRecord, error){
	suppChannel := make(chan *SupplementalRecord, 500)

	go decodeSupplemental(suppChannel, reader)
	
	return suppChannel, nil
}

func decodeSupplemental(suppChannel chan *SupplementalRecord, reader io.Reader){
	decoder := xml.NewDecoder(reader) 
	for { 
		t, _ := decoder.Token() 
		if t == nil { 
			break 
		} 
		switch se := t.(type) { 
		case xml.StartElement: 
			if se.Name.Local == SUPPLEMENTAL_RECORD { 
				var record SupplementalRecord
				decoder.DecodeElement(&record, &se) 
				suppChannel <- &record
			}
		}
	}
	close(suppChannel)
}


func QualifierMapFromFile(filename string)(map[string]*QualifierRecord, error){
	reader, file, err := genericReader(filename)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	return QualifierMapFromReader(reader)
}

func QualifierMapFromReader(reader  io.Reader)(map[string]*QualifierRecord, error){
	qualChan, err := QualifierChannelFromReader(reader)
	if err != nil{
		return nil, err
	}

	qualMap := make(map[string]*QualifierRecord)	
	for qRecord := range qualChan {
		qualMap[qRecord.QualifierUI] = qRecord
	}
	return qualMap, nil
}

func QualifierChannelFromFile(filename string)(qualChan chan *QualifierRecord, file *os.File, err error){
	reader, file, err := genericReader(filename)
	//if fl, ok := reader.(*gzip.Reader); ok {
	//defer fl.Close()
	//} 
	if err != nil {
		return nil, nil, err
	}
	
	qualChan, err = QualifierChannelFromReader(reader)
	return qualChan, file, err
}


func QualifierChannelFromReader(reader  io.Reader)(chan *QualifierRecord, error){
	qualChannel := make(chan *QualifierRecord, 500)

	go decodeQualifier(qualChannel, reader)
	
	return qualChannel, nil
}

func decodeQualifier(qualChannel chan *QualifierRecord, reader io.Reader){
	decoder := xml.NewDecoder(reader) 
	for { 
		t, _ := decoder.Token() 
		if t == nil { 
			break 
		} 
		switch se := t.(type) { 
		case xml.StartElement: 
			if se.Name.Local == QUALIFIER_RECORD { 
				var record QualifierRecord
				decoder.DecodeElement(&record, &se) 
				qualChannel <- &record
			}
		}
	}
	close(qualChannel)
}

func decodeDescriptor(recordChannel chan *DescriptorRecord, reader io.Reader){
//func decodeDescriptor(recordChannel chan interface{}, reader io.Reader){
	decoder := xml.NewDecoder(reader) 

	for { 
		t, _ := decoder.Token() 
		if t == nil { 
			break 
		} 
		switch se := t.(type) { 
		case xml.StartElement: 
			if se.Name.Local == DESCRIPTOR_RECORD { 
				var record DescriptorRecord
				decoder.DecodeElement(&record, &se) 
				recordChannel <- &record
			}
		}
	}
	close(recordChannel)
}

func genericReader(filename string) (io.Reader, *os.File, error){
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	if strings.HasSuffix(filename, "bz2"){
		return bufio.NewReader(bzip2.NewReader(bufio.NewReader(file))), file, err
	}

	if strings.HasSuffix(filename, "gz"){
		reader, err := gzip.NewReader(bufio.NewReader(file))
		if err != nil {
			return nil, nil, err
		}
		return bufio.NewReader(reader), file, err
	}
	return bufio.NewReader(file), file, err
}



type nextChannelItem func() bool

