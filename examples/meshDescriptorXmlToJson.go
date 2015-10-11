package main

import (
	"encoding/json"
	"fmt"
	"github.com/gnewton/gomesh"
	"os"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	if len(os.Args) != 2 {
		usage()
		os.Exit(42)
	}
	descFilename := os.Args[1]

	descChan, file, err := gomesh.DescriptorChannelFromFile(descFilename)
	defer file.Close()
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	for value := range descChan {
		b, err := json.Marshal(value)
		if err != nil {
			fmt.Println("error:", err)
		}
		os.Stdout.Write(b)
		os.Stdout.Write([]byte("\n"))
	}
}

func usage() {
	fmt.Println("Convert MeSH descriptor record XML to json")
	fmt.Println("\nUsage: meshDescriptorXmlToJson.go <filename>")
	fmt.Println("\t Try filename=../testData/desc2014_29records.xml.bz2")
}
