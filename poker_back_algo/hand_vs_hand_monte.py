from get_card_deck import getSortedDeck
from poker_card import Card
from poker_hand import *
import random
from get_best_hand import getBestHand
from get_prob_hand import getPartialDeck

low_hand = Hand([Card(2,'D'),Card(3,'S'),Card(4,'S'),Card(5,'S'),Card(7,'S')])

def getHighestHand(list_of_hands):
	winHand = low_hand
	for hand in list_of_hands:
		if(hand.compareHand(winHand) == 1):
			winHand = hand
	return winHand

def printProbHands(list_of_cards_pair,attempts):
	numPlayers = len(list_of_cards_pair)
	main_list = []
	for cards in list_of_cards_pair:
		main_list.append(cards[0])
		main_list.append(cards[1])
	# print main_list
	deck = getPartialDeck(main_list)
	# print deck
	iterations = 10000
	count = 0
	winCount = [0] * (numPlayers+1)
	probCount = []
	while(count < iterations):
		random.shuffle(deck)
		flop = deck[:attempts]
		won_index = -1
		winHand = low_hand
		for cards in range(0,numPlayers):
			hand = flop + list_of_cards_pair[cards]
			newHand = getBestHand(hand)
			# print newHand
			# print newHand.rank['hand']
			if(newHand.compareHand(winHand) == 1):
				winHand = newHand
				won_index = cards
			elif(newHand.compareHand(winHand) == 0):
				won_index = numPlayers
		# print winHand
		# print won_index
		winCount[won_index] = winCount[won_index] + 1
		count = count + 1
	# print winCount
	for wc in winCount:
		probCount.append(float(wc)*100/iterations)
	print(probCount)
cards_1 = [Card(10,'S'),Card(11,'C')]
cards_2 = [Card(2,'H'),Card(3,'D')]
# cards_3 = [Card(7,'D'),Card(6,'D')]
# cards_4 = [Card(6,'D'),Card(7,'D')]
# cards_5 = [Card(9,'H'),Card(10,'H')]
# cards_6 = [Card(11,'S'),Card(12,'D')]
printProbHands([cards_1,cards_2],5)#,cards_3,cards_4,cards_5,cards_6],3)


