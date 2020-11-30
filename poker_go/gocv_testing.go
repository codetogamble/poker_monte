package main

import(
	"fmt"
	"gocv.io/x/gocv"

)

func main() {
	fmt.Println("yo")
	webcam, _ := gocv.VideoCaptureDevice(0)
	window := gocv.NewWindow("Hello")	
	img := gocv.NewMat()

	for {
		webcam.Read(&img)
		window.IMShow(img)
		window.WaitKey(1)
	}
}