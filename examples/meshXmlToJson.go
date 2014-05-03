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

	descChan, file, err := jianGoMeSHi.DescriptorChannelFromFile("/home/newtong/2014/mesh/desc2014.xml.bz2")
	//descChan, file, err := jianGoMeSHi.DescriptorChannelFromFile("../testData/desc2014_29records.xml.bz2")
	defer file.Close()
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	for value := range descChan{
		b, err := json.Marshal(value)
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		os.Stdout.Write(b)
		os.Stdout.Write([]byte("\n"))
	}
}

	