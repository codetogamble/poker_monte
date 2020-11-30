//This file is for creating and testing algos (image processing)
//Threshold values for add is 80 & for recon is 70
//

package main

import(
	"fmt"
	"image/png"
	"os"
	"reflect"
	"github.com/disintegration/imaging"
	"image"
	"time"
	"image/color"
	"github.com/disintegration/gift"
	// "gonum.org/v1/gonum/mat"
)

//406 122 centre adda track (sub x,y)

//289 583 delaer chat beg

//1000 130 centre of multitable button

//950 130 centre of single table focus

//920 130 centre of multitablef focus

//155 180 join similar

//354 64 870 80 tab name



func convolveImage(main_img *image.NRGBA, template *image.NRGBA,scale int) (int,int){
	
	downsize_scale := scale
	template = imaging.Resize(template,template.Bounds().Max.X/downsize_scale,0,imaging.NearestNeighbor)
	main_img = imaging.Resize(main_img,main_img.Bounds().Max.X/downsize_scale,0,imaging.NearestNeighbor)

	dis_file,_ := os.Create("./dev_images/template_by2.png")
	defer dis_file.Close()
	png.Encode(dis_file,template)

	dis_file2,_ := os.Create("./dev_images/mainimg_by2.png")
	defer dis_file2.Close()
	png.Encode(dis_file2,main_img)

	fmt.Println(main_img.Bounds())
	fmt.Println(template.Bounds())


	// fmt.Printf("%+v\n",template.Pix)
	starti := main_img.Bounds().Min.X
	startj := main_img.Bounds().Min.Y
	endi := main_img.Bounds().Max.X
	endj := main_img.Bounds().Max.Y

	// fmt.Println(starti,startj,endi,endj)

	template_size := template.Bounds().Max.X
	// fmt.Println(startj-1+template_size)
	// fmt.Println(starti-1+template_size)
	var prev_sum_diff uint32
	prev_sum_diff = 1000000
	min_x := 0
	min_y := 0
	for j:=startj;j!=endj+1-template_size;j++{
		for i:=starti;i!=endi+1-template_size;i++{
			var sum_diff uint32
			
			sum_diff = 0
			// fmt.Println("inside 2 loop")

			for jj:= j;jj!=j+template_size;jj++{
				for ii:=i;ii!=i+template_size;ii++{
					temj := jj-j
					temi := ii-i

					pixel_val_main := main_img.At(ii,jj).(color.NRGBA).R
					pixel_val_tem := template.At(temi,temj).(color.NRGBA).R
					// fmt.Println("pixel_val_main : ",pixel_val_main)
					// fmt.Println("pixel_val_tem : ",pixel_val_tem)
					// fmt.Println("uint32 pixel_val_main : ",uint32(pixel_val_main))
					// fmt.Println("uint32 pixel_val_tem : ",uint32(pixel_val_tem))
					// fmt.Println(main_img.At(ii,jj).(color.NRGBA).R)
					// fmt.Println(reflect.TypeOf(main_img.At(ii,jj)))
					// fmt.Println(ii,jj,temi,temj)
					var diff uint32
					if pixel_val_main>=pixel_val_tem {
						diff = uint32(pixel_val_main) - uint32(pixel_val_tem)
						}else{
							diff = uint32(pixel_val_tem) - uint32(pixel_val_main)
						}
					// fmt.Println("diff is ",diff)
					sum_diff = sum_diff + diff
				}
			}

			// fmt.Println(main_img.At(i,j))
			// fmt.Println("sum is ",sum_diff)
			if sum_diff < prev_sum_diff {
				prev_sum_diff = sum_diff
				min_x = i
				min_y = j
				match_ratio := float64(prev_sum_diff)/float64(template.Bounds().Max.X*template.Bounds().Max.Y*255)
				if match_ratio < 0.1{
					fmt.Println("sum was : ",prev_sum_diff)
					fmt.Println("percent mismatch : ",float64(prev_sum_diff)*100/float64(template.Bounds().Max.X*template.Bounds().Max.Y*255))
					return min_x*downsize_scale,min_y*downsize_scale
				}
			}
			sum_diff = 0
			// color_gray := main_img.At(i,j)
			// color_gray = color.GrayModel.Convert(color_gray)
			// fmt.Printf("%+v\n",color_gray)	
		}
	}
	fmt.Println("sum was : ",prev_sum_diff)
	fmt.Println("percent mismatch : ",float64(prev_sum_diff)*100/float64(template.Bounds().Max.X*template.Bounds().Max.Y*255))
	fmt.Println("x , y min was : ",min_x*downsize_scale,min_y*downsize_scale)
	// return -1,-1
	return min_x*downsize_scale,min_y*downsize_scale
}

func main(){

	
	// file,_ := os.Open("./dev_images/detect_reconnect_conv.png")
	file,_ := os.Open("./dev_images/gray_main.png")
	// file,_ := os.Open("./dev_images/gray_disconnected.png")
	defer file.Close()
	img,_ := png.Decode(file)
	// fmt.Println(reflect.TypeOf(img))

	// crop_img := imaging.Crop(img,image.Rect(406,122,422,138))
	
	
	
	file2,_ := os.Open("./dev_images/adda_tracker_2.png")
	// file2,_ := os.Open("./dev_images/detect_reconnect.png")
	defer file2.Close()
	adda_img,_ := png.Decode(file2)
	adda_img_gray := imaging.Grayscale(adda_img)
	img_gray := imaging.Grayscale(img)
	// adda_img_gray = adda_img_gray.(*image.Gray)
	fmt.Println(reflect.TypeOf(adda_img_gray))
	fmt.Println(reflect.TypeOf(img_gray))
	// fmt.Printf("%+v\n",adda_img_gray)

	thresh_img := image.NewNRGBA(img_gray.Bounds())
	thresh_adda_img := image.NewNRGBA(adda_img_gray.Bounds())
	thresh_filter := gift.Threshold(70).(gift.Filter)
	thresh_filter.Draw(thresh_img,img_gray,nil)
	thresh_filter.Draw(thresh_adda_img,adda_img_gray,nil)
	
	start_time := time.Now()
	// x,y := convolveImage(img_gray,adda_img_gray,1)
	x,y := convolveImage(thresh_img,thresh_adda_img,2)
	// x,y := convolveImage(img_gray,adda_img_gray,2)
	fmt.Println("x y : ",x,y)
	fmt.Println(time.Since(start_time))

	// img_gray.SetNRGBA(x,y,color.NRGBA{255,0,0,255})
	// img_gray.SetNRGBA(x+1,y,color.NRGBA{255,0,0,255})
	// img_gray.SetNRGBA(x,y+1,color.NRGBA{255,0,0,255})
	// img_gray.SetNRGBA(x+1,y+1,color.NRGBA{255,0,0,255})
	// img_gray.SetNRGBA(x+2,y+2,color.NRGBA{255,0,0,255})
	// img_gray.SetNRGBA(x+1,y+2,color.NRGBA{255,0,0,255})
	// img_gray.SetNRGBA(x+2,y+1,color.NRGBA{255,0,0,255})
	// img_gray.SetNRGBA(x+1,y+1,color.NRGBA{255,0,0,255})

	

	// dis_file,_ := os.Create("./dev_images/test_tmp.png")
	// defer dis_file.Close()
	// png.Encode(dis_file,img_gray)



	// fmt.Println(reflect.TypeOf(thresh_filter))

	// crop_img := imaging.Crop(img_gray,image.Rect(460,200,900,480))
	// crop_img := imaging.Crop(img_gray,image.Rect(289,584,589,663))
	// crop_file,_ := os.Create("./dev_images/read_chat.png")
	// defer crop_file.Close()
	// png.Encode(crop_file,crop_img)	

}