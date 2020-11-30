package main

import (
	"fmt"
	"math"
	"runtime/pprof"	
	"os"
	"log"
	"flag"
	"time"
)


type shape interface{
	area() float64
}

type rect struct {
	width, heigth float64
}

type circ struct {
	radius float64
}

func (r *rect) area() float64{
	return r.width*r.heigth
}

func (c *circ) area() float64{
	return math.Pi*c.radius*c.radius
}

func printArea(s shape){
	fmt.Printf("area is %+v\n",s.area())
}

var cpuprofile = flag.String("cpuprofile", "got", "write cpu profile to `file`")



func main(){
	flag.Parse()
	fmt.Println(*cpuprofile)
    if *cpuprofile != "" {
        f, err := os.Create(*cpuprofile)
        if err != nil {
            log.Fatal("could not create CPU profile: ", err)
        }
        if err := pprof.StartCPUProfile(f); err != nil {
            log.Fatal("could not start CPU profile: ", err)
        }
        defer pprof.StopCPUProfile()
        // defer f.Close()
    }

	fmt.Println("yo")
	time.Sleep(2000*time.Millisecond)
	var s shape

	r1 := rect{width:2.0,heigth:3.0}
	c1 := circ{radius:4.0}

	ar := r1.area()
	ac := c1.area()

	s = &r1

	var a uint8 
	var b uint8

	a = 3
	b = 5
	fmt.Println("a - b is ",a-b)

	// printArea(s1)
	// printArea(r1)
	// printArea(c1)

	// r_copy := s1.(rect)

	fmt.Printf("%+v\n",ar)
	fmt.Printf("%+v\n",ac)
	fmt.Printf("%+v\n",s.area())
	// fmt.Printf("%+v\n",r_copy)

}

