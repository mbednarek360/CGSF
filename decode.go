package main

//import packages
import (
	"fmt"
    "os"
	"image/png"
	"log"
	"strconv"
	"os/exec"
)

func main() {
	//check for arguements
	if len(os.Args) != 3 {
		system("clear")
		fmt.Println("Please specify an input and output file path!")
		os.Exit(2)
	}

	//assign arguements
	inFile := os.Args[1]
	outFile := os.Args[2]

	//print message
	system("clear")
	fmt.Println("Reading input file...")

	//open input file
	f, err := os.Open(inFile)
		if err != nil {
			system("clear")
			fmt.Println("Could not read input file!")
			os.Exit(2)
		}
		defer f.Close()

	//decode input as png
	img, err := png.Decode(f)
	if err != nil {
			system("clear")
		fmt.Println("Invalid input format!")
		os.Exit(2)
}

//print message
	system("clear")
	fmt.Println("Decompressing...")

//calculate width of input
		b := img.Bounds()
		width := b.Max.X

//create output file
		out, err := os.Create(outFile)
		if err != nil {
			system("clear")
			fmt.Println("Could not write to output file!")
			os.Exit(2)
	}
		defer out.Close()

//loop through pixels
		for x := 0; x < width; x++ {
			//get rgb value of pixel
			r, g, b, _ := img.At(x, 0).RGBA()

			//end of string for reds
			if r == 0 {
				break;
			} else {
				//write char to output file
				out.WriteString(string(r/257))
			}

			//end of string for blues
			if g == 0 {
				break;
			} else {
				//write char to output file
				out.WriteString(string(g/257))
			}

			//end of string for greens
			if b == 0 {
				break;
			} else {
				//write char to output file
				out.WriteString(string(b/257))
			}
	  }

	//print message
		system("clear")
		fmt.Println("Finished decompression to " + outFile + ".")

}
//get os
func isWindows() bool {
    return os.PathSeparator == '\\' && os.PathListSeparator == ';'
}

//convert float to string
func FloatToString(input_num float64) string {
    return strconv.FormatFloat(input_num, 'f', 6, 64)
}

//run system command
func system(cmd string, arg ...string) {
    out, err := exec.Command(cmd, arg...).Output()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(out))
}