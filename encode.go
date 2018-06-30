package main

import (
	"fmt"
	"io/ioutil"
	"log"
  "os"
  "os/exec"
	"image"
	"image/color"
	"image/png"
)

func main() {
	//clear terminal and assign arguements
  system("clear")
  inFile := os.Args[1]
  var outFile = os.Args[2]

	//open input file
  inBytes, err := ioutil.ReadFile(inFile)
  if err != nil {
		fmt.Println("Could not open input file!")
		os.Exit(2)
	}
  inText := string(inBytes)

	//pad with null chars
	for l := 0; l < (len(inText) % 3); l++ {
		inText += string(0)
	}

	//create new image
	img := image.NewRGBA(image.Rect(0, 0, (len(inText)/3), 1))

	//make triple val array
  val := make([]int, 3)

	//loop though charectars
  n:= 0
  for i := 0; i <= len(inText); i++ {
		//every three chars draw pixel
    if i % 3 == 0 && i > 0 {
      n = 0
      //fmt.Println(i/3)
			fmt.Println(val)
			img.Set(((i / 3)-1), 0, color.RGBA{uint8(val[0]), uint8(val[1]), uint8(val[2]), 255})

    }
		//get ascii value of char
    if i < len(inText) {
    c := string(inText[i])
    val[n] = int([]rune(c)[0])
  }
    n++
  }

//export image
f, _ := os.OpenFile(outFile, os.O_WRONLY|os.O_CREATE, 0600)
defer f.Close()
png.Encode(f, img)















}

//run system command
func system(cmd string, arg ...string) {
    out, err := exec.Command(cmd, arg...).Output()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(out))
}
