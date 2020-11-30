package main

import (
	// "fmt"
	"sort"
)

type Card struct{
	val uint8
	suit uint8
}

type Hand struct{
	totalcards []Card
	mycards []Card
	board []Card
	combinations [][]Card
	gametype uint8
}

type PlayingHand struct{
	playing_hand []Card
	sorted_vals []uint8
	lvl uint8
	hfhv_val uint8
	shfhv_val uint8
	hfhv uint8
	shfhv uint8
	name string
}

type byCards []Card

func (s byCards) Len() int {
    return len(s)
}
func (s byCards) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
func (s byCards) Less(i, j int) bool {
    return s[i].val > s[j].val
}

func createHand(mycards []Card, gametype uint8) *Hand{
	// fmt.Println("Cards are ",cards)
	// fmt.Println("Sorted Cards are ",cards)
	h := new(Hand)
	h.gametype = gametype
	if gametype == 1{
		h.setTexasHand(mycards)
	}
	return h
}

// func (h *Hand) setTexasCards(cards []Card){
// 	// fmt.Println("texas")	
// 	// h.mycards = cards[:2]
// 	h.mycards = make([]Card,2)
// 	h.board = make([]Card,5)
// 	copy(h.mycards,cards[:2])
// 	copy(h.board,cards[2:])
// 	sort.Sort(byCards(cards))
// 	h.totalcards = cards

// }

func (hand *Hand) setTexasHand(mycards []Card){
	hand.mycards = mycards
}

func (hand *Hand) generateCombinations(){
	stack := [][]Card{}
	if(len(hand.totalcards) == 0){
		return
	}
	total_c := hand.totalcards
	for i:=0;i<7;i++{
		for j:=i+1;j<7;j++{
			intst := []Card{}
			intst = append(intst,total_c[:i]...)
			intst = append(intst,total_c[i+1:j]...)
			intst = append(intst,total_c[j+1:]...)
			stack = append(stack,[][]Card{intst}...)
			// fmt.Println(stack)
		}
	}
	// fmt.Println(stack)
	hand.combinations = stack
}


func (h *Hand) setBoard(cards []Card){
	h.board = cards
	h.totalcards = append(h.mycards,h.board...)
	sort.Sort(byCards(h.totalcards))
}

func createPlayingHand(cards []Card) *PlayingHand{
	// cards := stringToCard(cs)
	ph := new(PlayingHand)
	ph.playing_hand = cards
	sort.Sort(byCards(ph.playing_hand))
	x := []uint8{ph.playing_hand[0].val,ph.playing_hand[1].val,ph.playing_hand[2].val,ph.playing_hand[3].val,ph.playing_hand[4].val}
	freqset_x := [15]uint8{0,0,0,0,0,0,0,0,0,0,0,0,0,0}
	set_num_len := uint8(0)
	tmp_var := uint8(0)
	ph.sorted_vals = x

	for i:=0;i<len(x);i++{
		freqset_x[x[i]]++
		if x[i]==14{
			freqset_x[1]++
		}
		if tmp_var!=x[i] {
			tmp_var = x[i]
			set_num_len++
		}
		// count_x := [len(x)]int
	}
	ph.hfhv = 0
	ph.shfhv = 0
	ph.hfhv_val = 0
	ph.shfhv_val = 0
	seq_flag := false
	seq_count := uint8(0)
	// fmt.Println(len(freqset_x))

	for i:=len(freqset_x)-1;i!=1;i--{
		if freqset_x[i]>ph.hfhv {
			ph.shfhv = ph.hfhv
			ph.hfhv = freqset_x[i]
			ph.shfhv_val = ph.hfhv_val
			ph.hfhv_val =uint8(i)
		}else if freqset_x[i]>ph.shfhv{
			ph.shfhv = freqset_x[i]
			ph.shfhv_val = uint8(i)
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

	if freqset_x[14]==1 && freqset_x[2]==1 && freqset_x[3]==1 && freqset_x[4]==1 && freqset_x[5]==1 {
		seq_flag = true
	}



	y := []uint8{ph.playing_hand[0].suit,ph.playing_hand[1].suit,ph.playing_hand[2].suit,ph.playing_hand[3].suit,ph.playing_hand[4].suit}
	tmp := y[0]
	flush_flag := true
	for i:=0;i<5;i++{
		if(tmp != y[i]){
			flush_flag = false
		}
	}

	if seq_flag == true && flush_flag == true {
		ph.lvl = 9
		ph.name = "straight flush"
		if ph.hfhv_val==14 && ph.shfhv_val==5{
			ph.hfhv_val=5
			ph.shfhv_val=4
		}
	}else if ph.hfhv == 4{
		ph.lvl = 8
		ph.name = "four of a kind"
	}else if ph.hfhv==3 && ph.shfhv == 2{
		ph.lvl = 7
		ph.name = "full house"
	}else if flush_flag == true{
		ph.lvl = 6
		ph.name = "flush"
	}else if seq_flag == true{
		ph.lvl = 5
		ph.name = "straight"
		if ph.hfhv_val==14 && ph.shfhv_val==5{
			ph.hfhv_val=5
			ph.shfhv_val=4
		}
	}else if ph.hfhv==3{
		ph.lvl = 4
		ph.name = "three of a kind"
	}else if ph.hfhv==2 && ph.shfhv==2{
		ph.lvl = 3
		ph.name = "two pair"
	}else if ph.hfhv==2{
		ph.lvl = 2
		ph.name = "pair"
	}else{
		ph.lvl = 1
		ph.name = "highcard"
	}
	// fmt.Println("playing hand : ",ph.playing_hand)
	// fmt.Println("cards vals : ",x)
	// fmt.Println("frequency set : ",freqset_x)
	// fmt.Println("values set length : ",set_num_len)
	// fmt.Println("highest freq : ",ph.hfhv)
	// fmt.Println("second highest freq : ",ph.shfhv)
	// fmt.Println("highest freq with highest val : ",ph.hfhv_val)
	// fmt.Println("SECOND highest freq with highest val : ",ph.shfhv_val)
	// fmt.Println(seq_flag)
	// fmt.Println(ph.name)
	return ph
}

// func (hand *Hand) createPlaying(cards []Card){
// 	// fmt.Println(hand)
// 	local_c := cards
// 	sort.Sort(byCards(local_c))
// 	if(hand.gametype == 1){
// 		// sort.Sort(byCards(local_c))
// 		fmt.Println("inside createPlaying",local_c)
// 	}
// }

func comparehands(ph1 *PlayingHand, ph2 *PlayingHand) int {
	if ph1.lvl > ph2.lvl {
		return 1
	}else if ph1.lvl < ph2.lvl {
		return -1
	}else {
		if ph1.lvl== 9 || ph1.lvl==5{
			if ph1.hfhv_val>ph2.hfhv_val{
				return 1
			}else if ph1.hfhv_val<ph2.hfhv_val{
				return -1
			}else{
				return 0
			}

		}else if ph1.lvl==8 || ph1.lvl==7 || ph1.lvl==3 {
			if ph1.hfhv_val>ph2.hfhv_val{
				return 1
			}else if ph1.hfhv_val<ph2.hfhv_val{
				return 0
			}else{
				if ph1.shfhv_val > ph2.shfhv_val{
					return 1
				}else if ph1.shfhv_val < ph2.shfhv_val{
					return -1
				}else{
					return 0
				}
			}
		}else if ph1.lvl==6 || ph1.lvl==1 || ph1.lvl==4 || ph1.lvl==2{
			for i:=0;i<5;i++{
				if ph1.sorted_vals[i]>ph2.sorted_vals[i]{
					return 1
				}else if ph1.sorted_vals[i]<ph2.sorted_vals[i]{
					return -1
				}
			}
			return 0
		}


	}
	return 0
}