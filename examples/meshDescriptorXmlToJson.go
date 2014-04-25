package main
 
import (
	"encoding/json"
	"fmt"
	"github.com/gnewton/jianGoMeSHi"
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

	descChan, file, err := jianGoMeSHi.DescriptorChannelFromFile(descFilename)
	defer file.Close()
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	for value := range descChan{
		b, err := json.Marshal(value)
		if err != nil {
			fmt.Println("error:", err)
		}
		os.Stdout.Write(b)
	}
}

func usage(){
	
}

	