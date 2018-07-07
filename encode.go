package main

//import packages
import (
	"fmt"
	"io/ioutil"
	"log"
    "os"
    "os/exec"
	"image"
	"image/color"
	"image/png"
	"strconv"
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
  inBytes, err := ioutil.ReadFile(inFile)
  if err != nil {
		system("clear")
		fmt.Println("Could not read input file!")
		os.Exit(2)
	}
  inText := string(inBytes)

	//pad with null chars
	for l := 0; l < (len(inText) % 3); l++ {
		inText += string(0)
	}

	//print message
		system("clear")
		fmt.Println("Compressing...")

	//create new image
	img := image.NewRGBA(image.Rect(0, 0, (len(inText)/3), 1))

	//make triple val array
  val := make([]int, 3)

	//loop though charectars
  n:= 0
  for i := 0; i <= len(inText); i++ {
		//every four chars draw pixel
    if i % 3 == 0 && i > 0 {
      n = 0
			img.Set(((i / 3)-1), 0, color.RGBA{uint8(val[0]), uint8(val[1]), uint8(val[2]), 255})
    }
		//get ascii value of char
    if i < len(inText) {
    c := string(inText[i])
    val[n] = int([]rune(c)[0])
  }
    n++
  }

//create output file
f, err := os.OpenFile(outFile, os.O_WRONLY|os.O_CREATE, 0600)
if err != nil {
	system("clear")
	fmt.Println("Could not write to output file!")
	os.Exit(2)
}
defer f.Close()

//encode as png
png.Encode(f, img)

//print message
//system("clear")
fmt.Println("Finished compression to " + outFile + ".")
}

//run system command
func system(cmd string, arg ...string) {
    out, err := exec.Command(cmd, arg...).Output()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(out))
}

//get os
func isWindows() bool {
    return os.PathSeparator == '\\' && os.PathListSeparator == ';'
}

//convert float to string
func FloatToString(input_num float64) string {
    return strconv.FormatFloat(input_num, 'f', 6, 64)
}