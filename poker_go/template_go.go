package main

import(
	"fmt"
	"image/png"
	"os"
)

func main(){
	file,_ := os.Open("./dev_images/chrome_x.png")
	img,_ := png.Decode(file)
	fmt.Println(file)
	fmt.Println(img)
}