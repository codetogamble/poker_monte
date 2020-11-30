package main

import (
	"math/rand"
	"time"
	"fmt"
)

type Deck struct{
	cards [52]Card
	cursor uint8
	size uint8
}

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
	return deck
}

func shuffleDeck(deck *Deck) {
	rand.Seed(time.Now().UnixNano() / int64(time.Nanosecond))
	// x := rand.Intn(52)
	x := int(deck.size)
	c := &deck.cards
	for i:=2;i!=0;i--{
		rand.Shuffle(x,func(i,j int){
		tmp := c[i]
		c[i] = c[j]
		c[j] = tmp 
		})
	}
	// fmt.Println(x)
}



func main(){

	start_time := time.Now()
	d := unpackNewDeck()
	// fmt.Println(d)
	shuffleDeck(d)
	// fmt.Println(d.cards[:7])
	h := createHand(d.cards[:2],1)
	// fmt.Println(h)
	h.setBoard(d.cards[2:7])
	h.generateCombinations()
	fmt.Println(h.totalcards)
	fmt.Println(h.combinations[0])
	
	// ph2 := createPlayingHand("14_C,13_D,12_D,11_H,10_S")
	// fmt.Println(comparehands(ph1,ph2))
	

	count_array := [10]uint64{}
	total_iter := 100000
	dl := unpackNewDeck()
	for i:=0;i<total_iter;i++{
		// fmt.Println(d)
		shuffleDeck(dl)
		ph1 := createPlayingHand(dl.cards[:5])
		count_array[ph1.lvl]++
	}

	fmt.Println(count_array)

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
	

	part_deck := dl.getPartialDeck(stringToCard("14_C,13_D,12_D,11_H,10_S"))
	fmt.Println(part_deck)
	fmt.Println(len(part_deck))
	fmt.Println(time.Since(start_time))
	// fmt.Println(stack)
	// cards := []Card{Card{val:1,suit:int('C')},Card{val:3,suit:int('C')},Card{val:2,suit:int('C')}}
	// fmt.Println(cards)
}