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
def getProbMonte(incomplete_hand):
	counts = [0,0,0,0,0,0,0,0,0]
	count = 0
	iterations = 10000
	part_deck = getPartialDeck(incomplete_hand)
	# print iterations
	while(count < iterations):
		# print count
		random.shuffle(part_deck)
		whole_hand = incomplete_hand + part_deck[:5]
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
	# print counts
	prob_list = []
	for x in range(0,len(counts)):
		probPerc = float(counts[x])*100/iterations
		# print "for level " + str(x+1) + " prob is" + str(probPerc) + "%"
		prob_list.append(probPerc)
	return prob_list

print(getProbMonte([Card(5,'S'),Card(12,'D')]))
# f = open('poker_probs_river.txt','w')
# for x in range(2,15):
# 	for y in range(x+1,15):
# 		# print str(x) + " " + str(y)
# 		list_cards = [Card(x,'S'),Card(y,'D')]
# 		prob_list = getProbMonte(list_cards)
# 		f.write(str(list_cards)+",")
# 		for prob in prob_list:
# 			f.write(str(prob)+",")
# 		f.write("\n")
# for x in range(2,15):
# 	for y in range(x+1,15):
# 		# print str(x) + " " + str(y)
# 		list_cards = [Card(x,'S'),Card(y,'S')]
# 		prob_list = getProbMonte(list_cards)
# 		f.write(str(list_cards)+",")
# 		for prob in prob_list:
# 			f.write(str(prob)+",")
# 		f.write("\n")
# for x in range(2,15):
# 	list_cards = [Card(x,'S'),Card(x,'D')]
# 	prob_list = getProbMonte(list_cards)
# 	f.write(str(list_cards)+",")
# 	for prob in prob_list:
# 		f.write(str(prob)+",")
# 	f.write("\n")

# f.close()