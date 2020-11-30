package main

import (
	"math/rand"
	"time"
	"fmt"
	// "sync"
	"runtime/pprof"
	"runtime"
	"os"
	"log"
	"flag"
	"net/http"
)

import _ "net/http/pprof"

type Deck struct{
	cards []Card
	cursor uint8
	size uint8
}

//EDIT NOT GOOD FOR FUTURE
//
func (deck *Deck) getPartialDeck(rem_cards []Card) []Card{
	cards_total := []Card{}
	for _,card := range(deck.cards){
		rem_flag := false
		for _,card2 := range(rem_cards){
			if card2 == card {
				rem_flag = true
				break
			}
		}
		if(!rem_flag){
			cards_total = append(cards_total,[]Card{card}...)
		}
	}

	return cards_total
}

func unpackNewDeck() *Deck{
	C := uint8('C') //67
	D := uint8('D') //68
	H := uint8('H') //72
	S := uint8('S') //83
	deck := new(Deck)
	deck.cursor = 0
	deck.size = 52
	deck.cards 	= make([]Card,52)
	deck.cards[0] = Card{val:14,suit:C}
	deck.cards[13] = Card{val:14,suit:D}
	deck.cards[26] = Card{val:14,suit:H}
	deck.cards[39] = Card{val:14,suit:S}
	for i:=1; i<13; i++{
		deck.cards[i] = Card{val:uint8(i+1),suit:C}
		deck.cards[i+13] = Card{val:uint8(i+1),suit:D}
		deck.cards[i+26] = Card{val:uint8(i+1),suit:H}
		deck.cards[i+39] = Card{val:uint8(i+1),suit:S}
	}

	// deck.cards = xcards[:]
	// deck.cards = xcards
	return deck
}

// func (deck *Deck)shuffleDeck() {
// 	// x := rand.Intn(52)
// 	x := int(deck.size)
// 	c := &deck.cards
// 	for ii:=2;ii!=0;ii--{
// 		rand.Shuffle(x,func(i,j int){
// 		tmp := (c)[i]
// 		(c)[i] = (c)[j]
// 		(c)[j] = tmp
// 		})
// 	}
// 	// fmt.Println(x)
// }

func (deck *Deck)shuffleDeck2(src_rand *rand.Rand){
	c := &deck.cards
	// perm := rand.Perm(len(c))
	// // _=perm
	// for i, v := range perm {
 //    	tmp := c[i]
	// 	c[i] = c[v]
	// 	c[v] = tmp
	// }
	_ = src_rand
	for i:=32;i!=0;i--{
		x := src_rand.Int()%52
		v := src_rand.Int()%52
		// x:=0
		// v:=1
		tmp := (*c)[x]
		(*c)[x] = (*c)[v]
		(*c)[v] = tmp
	}
}

func giveProbRot(iter int,count_chan chan [10]uint64, routineid int){
	fmt.Println("START routineid : ",routineid)
	src := rand.NewSource(time.Now().UnixNano() + int64(routineid))
	src_rand := rand.New(src)
	count_array := [10]uint64{0,0,0,0,0,0,0,0,0,0}
	total_iter := iter

	dl := unpackNewDeck()
	for i:=total_iter;i!=0;i-- {
		// fmt.Println(d)
		// mutex.Lock()

		dl.shuffleDeck2(src_rand)
		// mutex.Unlock()
		ph1 := new(PlayingHand)
		ph1.createPlayingHand(dl.cards[:5])
		count_array[ph1.lvl]++
	}


	fmt.Println(count_array)
	count_chan <- count_array
}

func giveProb(iter int) [10]uint64{
	src := rand.NewSource(time.Now().UnixNano())
	src_rand := rand.New(src)
	count_array := [10]uint64{0,0,0,0,0,0,0,0,0,0}
	total_iter := iter
	dl := unpackNewDeck()
	for i:=total_iter;i!=0;i-- {
		// dl.shuffleDeck()//slow
		dl.shuffleDeck2(src_rand)
		ph1 := new(PlayingHand)
		ph1.createPlayingHand(dl.cards[:5])
		count_array[ph1.lvl]++
	}


	fmt.Println(count_array)
	return count_array
}

var cpuprofile = flag.String("cpuprofile", "probpoker.prof", "write cpu profile to file")
var memprofile = flag.String("memprofile", "probpoker_mem.prof", "write memory profile to `file`")

func main(){
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	rand.Seed(time.Now().UnixNano())
	flag.Parse()
    if *cpuprofile != "" {
        f, err := os.Create(*cpuprofile)
        if err != nil {
            log.Fatal(err)
            fmt.Println("error while creating cpu profile")
        }
        pprof.StartCPUProfile(f)
        defer pprof.StopCPUProfile()
    }



    parts := 8
	// prev_pr := runtime.GOMAXPROCS(4)
	// fmt.Println("previous setting for gomaxprocs was ",prev_pr)
	start_time := time.Now()
	d := unpackNewDeck()
	// fmt.Println(d)
	// d.shuffleDeck()
	// fmt.Println(d.cards[:7])
	h := createHand(d.cards[:2],1)
	// fmt.Println(h)
	h.setBoard(d.cards[2:7])
	h.generateCombinations()
	fmt.Println(h.totalcards)
	fmt.Println(h.combinations[0])

	total_iter := 4000000



	count_chan := make(chan [10]uint64,parts)
	for gri := 1;gri<parts+1;gri++{
		go giveProbRot(total_iter/parts,count_chan,gri)
	}


	main_count_array := [10]uint64{0,0,0,0,0,0,0,0,0,0}

	for gri := 0;gri<parts;gri++{
		count_array := <- count_chan
		main_count_array[0] = main_count_array[0]+count_array[0]
		main_count_array[1] = main_count_array[1]+count_array[1]
		main_count_array[2] = main_count_array[2]+count_array[2]
		main_count_array[3] = main_count_array[3]+count_array[3]
		main_count_array[4] = main_count_array[4]+count_array[4]
		main_count_array[5] = main_count_array[5]+count_array[5]
		main_count_array[6] = main_count_array[6]+count_array[6]
		main_count_array[7] = main_count_array[7]+count_array[7]
		main_count_array[8] = main_count_array[8]+count_array[8]
		main_count_array[9] = main_count_array[9]+count_array[9]
	}


	count_array := main_count_array
	// count_array := giveProb(total_iter)




	fmt.Println("probability of HighCard is ",float64(count_array[1])/float64(total_iter)*100)
	fmt.Println("probability of pair is ",float64(count_array[2])/float64(total_iter)*100)
	fmt.Println("probability of two pair is ",float64(count_array[3])/float64(total_iter)*100)
	fmt.Println("probability of three of a kind is ",float64(count_array[4])/float64(total_iter)*100)
	fmt.Println("probability of straight is ",float64(count_array[5])/float64(total_iter)*100)
	fmt.Println("probability of flush is ",float64(count_array[6])/float64(total_iter)*100)
	fmt.Println("probability of full house is ",float64(count_array[7])/float64(total_iter)*100)
	fmt.Println("probability of four of a kind is ",float64(count_array[8])/float64(total_iter)*100)
	fmt.Println("probability of straight flush is ",float64(count_array[9])/float64(total_iter)*100)

	fmt.Println("iterations : ",total_iter)


	// part_deck := dl.getPartialDeck(stringToCard("14_C,13_D,12_D,11_H,10_S"))
	// fmt.Println(part_deck)
	// fmt.Println(len(part_deck))
	fmt.Println(time.Since(start_time))
	if *memprofile != "" {
        f, err := os.Create(*memprofile)
        if err != nil {
            log.Fatal("could not create memory profile: ", err)
        }
        runtime.GC() // get up-to-date statistics
        if err := pprof.WriteHeapProfile(f); err != nil {
            log.Fatal("could not write memory profile: ", err)
        }
        f.Close()
    }
	// fmt.Println(stack)
	// cards := []Card{Card{val:1,suit:int('C')},Card{val:3,suit:int('C')},Card{val:2,suit:int('C')}}
	// fmt.Println(cards)
}
