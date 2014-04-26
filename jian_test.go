
package jianGoMeSHi

import (
	"testing"
//	"fmt"
//	"compress/gzip"
)





func makeNextSupplemental(suppChannel chan *SupplementalRecord)(nextChannelItem){
	return func() bool{
		_, err := <- suppChannel
		return err
	}
}


func makeNextDescriptor(descChannel chan *DescriptorRecord)(nextChannelItem){
	return func() bool{
		_, err := <- descChannel
		return err
	}
}

func makeNextQualifier(qualChannel chan *QualifierRecord)(nextChannelItem){
	return func() bool{
		_, err := <- qualChannel
		return err
	}
}


func TestReadOneDescription(t *testing.T){
	descFileName := "./testData/desc2014_1record.xml"
	descChannel, file, err := DescriptorChannelFromFile(descFileName)
	defer file.Close()

	if err != nil {
		t.Fatal("Error occured", err)
		return
	}

	if countChannel(makeNextDescriptor(descChannel)) != 1{
		t.Fail()
	}
}

func TestReadManyDescriptions(t *testing.T){
	descFileName := "./testData/desc2014_29records.xml.bz2"
	descChannel, file, err := DescriptorChannelFromFile(descFileName)
	defer file.Close()

	if err != nil {
		t.Fatal("Error occured", err)
		return
	}
	if countChannel(makeNextDescriptor(descChannel)) != 29{
		t.Fail()
	}
}

func TestReadManyQualifiers(t *testing.T){
	qualFileName := "./testData/qual2014_8records.xml.bz2"
	qualChannel, file, err := QualifierChannelFromFile(qualFileName)
	defer file.Close()

	if err != nil {
		t.Fatal("Error occured", err)
		return
	}
	if countChannel(makeNextQualifier(qualChannel)) != 8{
		t.Fail()
	}
}

func TestReadOneQualifier(t *testing.T){
	qualFileName := "./testData/qual2014_1record.xml.bz2"
	qualChannel, file, err := QualifierChannelFromFile(qualFileName)
	defer file.Close()

	if err != nil {
		t.Fatal("Error occured", err)
		return
	}
	if countChannel(makeNextQualifier(qualChannel)) != 1{
		t.Fail()
	}
}

func TestReadManySupplementalRecords(t *testing.T){
	suppFileName := "./testData/supp2014_4records.xml"
	suppChannel, file, err := SupplementalChannelFromFile(suppFileName)
	defer file.Close()

	if err != nil {
		t.Fatal("Error occured", err)
		return
	}
	if countChannel(makeNextSupplemental(suppChannel)) != 4{
		t.Fail()
	}
}

func TestReadOneSupplementalRecord(t *testing.T){
	suppFileName := "./testData/supp2014_1record.xml"
	suppChannel, file, err := SupplementalChannelFromFile(suppFileName)
	defer file.Close()

	if err != nil {
		t.Fatal("Error occured", err)
		return
	}
	if countChannel(makeNextSupplemental(suppChannel)) != 1{
		t.Fail()
	}
}


func countChannel(nextDescriptor nextChannelItem)int{
	counter := 0
	for{
		val := nextDescriptor()
		if !val {
			break
		}
		counter += 1
	}
	return counter
}

