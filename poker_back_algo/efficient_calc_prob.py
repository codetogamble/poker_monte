from get_card_deck import getSortedDeck
from poker_card import Card
from poker_hand import *
import random
from get_best_hand import getBestHand
from get_prob_hand import getPartialDeck
import time

a = time.time()

low_hand = Hand([Card(2,'D'),Card(3,'S'),Card(4,'S'),Card(5,'S'),Card(7,'S')])

def getHighestHand(list_of_hands):
	winHand = low_hand
	for hand in list_of_hands:
		if(hand.compareHand(winHand) == 1):
			winHand = hand
	return winHand

def printProbHands(list_of_cards_pair,board):
	numPlayers = len(list_of_cards_pair)
	main_list = []
	for cards in list_of_cards_pair:
		main_list.append(cards[0])
		main_list.append(cards[1])
	# print main_list
	for card in board:
		main_list.append(card)
	deck = getPartialDeck(main_list)
	# print deck
	board_left = 5 - len(board)
	iterations = 10000
	count = 0
	winCount = [0] * (numPlayers+1)
	probCount = []
	while(count < iterations):
		random.shuffle(deck)
		flop = deck[:board_left]
		won_index = -1
		winHand = low_hand
		for cards in range(0,numPlayers):
			hand = flop + list_of_cards_pair[cards] + board
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
cards_1 = [Card(10,'D'),Card(9,'C')]
cards_2 = [Card(5,'S'),Card(5,'D')]
cards_3 = [Card(13,'H'),Card(14,'D')]
cards_4 = [Card(3,'C'),Card(9,'S')]
cards_5 = [Card(2,'C'),Card(8,'H')]
board = [Card(4,'C'),Card(3,'H'),Card(7,'S'),Card(2,'S')]
# board = []
printProbHands([cards_1,cards_2,cards_3,cards_4,cards_5],board)

b = time.time()

print(b-a)


