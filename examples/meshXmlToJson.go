package main
 
import (
	"encoding/json"
	"fmt"
	"github.com/gnewton/jianGoMeSHi"
	"os"
	"runtime"
)

func main() {
	numCpu := runtime.NumCPU()
	runtime.GOMAXPROCS(numCpu)

	descChan, file, err := jianGoMeSHi.DescriptorChannelFromFile("../testData/desc2014_29records.xml.bz2")
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

	