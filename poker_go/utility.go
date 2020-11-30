package main

import(
	"fmt"
	"strings"
	"strconv"
)

func stringToCard(cards string) []Card{
		css := strings.Split(cards,",")
		fmt.Println(css)
		cl := []Card{}
		fmt.Println("cs","val","suit","i","err")
		for i,cs := range css{
			tmpc := strings.Split(cs,"_")
			val,err := strconv.ParseInt(tmpc[0],10,8)
			// suit,err := strconv.ParseInt(tmpc[1],10,8)
			suit := []byte(tmpc[1])
			cl = append(cl,[]Card{Card{val:uint8(val),suit:uint8(suit[0])}}...)
			fmt.Println(cs,val,suit,i,err)
		}
		fmt.Println(cl)
		return cl
}

func test(){
	x := []int{1,1,2,3,4,5,6,6,6,7,7}

	freqset_x := [14]int{0,0,0,0,0,0,0,0,0,0,0,0,0,0}
	set_num_len := 0
	tmp_var := 0

	for i:=0;i<len(x);i++{
		freqset_x[x[i]]++
		if tmp_var!=x[i] {
			tmp_var = x[i]
			set_num_len++
		}
		// count_x := [len(x)]int
	}
	hfhv := 0
	shfhv := 0
	hfhv_val := 0
	shfhv_val := 0
	seq_flag := false
	seq_count := 0
	fmt.Println(len(freqset_x))

	for i:=len(freqset_x)-1;i!=-1;i--{
		if freqset_x[i]>hfhv {
			shfhv = hfhv
			hfhv = freqset_x[i]
			shfhv_val = hfhv_val
			hfhv_val = i
		}else if freqset_x[i]>shfhv{
			shfhv = freqset_x[i]
			shfhv_val = i
		}
		if freqset_x[i]!=0{
			seq_count++
			if seq_count>4 {
				seq_flag = true
			}
		}else{
			seq_count = 0
		}
	}


	fmt.Println(x)
	fmt.Println(freqset_x)
	fmt.Println(set_num_len)
	fmt.Println(hfhv)
	fmt.Println(shfhv)
	fmt.Println(hfhv_val)
	fmt.Println(shfhv_val)
	fmt.Println(seq_flag)
}