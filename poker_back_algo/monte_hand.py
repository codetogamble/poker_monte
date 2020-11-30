from poker_card import Card
from poker_hand import *
from get_card_deck import getSortedDeck
import random
from get_prob_hand import getPartialDeck
from get_best_hand import getBestHand
import sys

inputList = sys.argv
deck = getSortedDeck()
# print deck
# def getProbMonte(incomplete_hand,iterations):

if(len(inputList) == 5):
	counts = [0,0,0,0,0,0,0,0,0]
	count = 0
	card1 = inputList[1]
	card2 = inputList[2]
	iterations = int(inputList[3])
	attempts = int(inputList[4])
	input_card_1 = Card(int(card1[0]),card1[1])
	input_card_2 = Card(int(card2[0]),card2[1])
	input_hand = [input_card_1,input_card_2]
	part_deck = getPartialDeck(input_hand)
	# print iterations
	while(count < iterations):
		# print count
		random.shuffle(part_deck)
		whole_hand = input_hand + part_deck[:attempts]
		newHand = getBestHand(whole_hand)
		# newHand = Hand(new_hand)
		level = newHand.rank['level']
		if(level == 1):
			counts[0] = counts[0] + 1
		elif(level == 2):
			counts[1] = counts[1] + 1
		elif(level == 3):
			counts[2] = counts[2] + 1
		elif(level == 4):
			counts[3] = counts[3] + 1
		elif(level == 5):
			counts[4] = counts[4] + 1
		elif(level == 6):
			counts[5] = counts[5] + 1
		elif(level == 7):
			counts[6] = counts[6] + 1	
		elif(level == 8):
			counts[7] = counts[7] + 1
		elif(level == 9):
			counts[8] = counts[8] + 1
		count = count + 1
	print counts
	for x in range(0,len(counts)):
		print "for level " + str(x+1) + " prob is" + str(float(counts[x])*100/iterations) + "%"
		