from poker_card import Card
from poker_hand import *
from get_card_deck import getSortedDeck
import random
import math

low_hand = Hand([Card(2,'D'),Card(3,'S'),Card(4,'S'),Card(5,'S'),Card(7,'S')])
# print low_hand.rank

incHand = [Card(3,'S'),Card(5,'S')]
def getProbHand(incomplete_hand):
	numCards = len(incomplete_hand)
	count = 0
	iterations = 10000
	poss_set_7 = []
	poss_set_6 = []
	poss_set_5 = []
	poss_set_4 = []
	poss_set_3 = []
	poss_set_2 = []
	totalNumberHands = 0
	if(numCards == 3):
		totalNumberHands = 1176
	elif(numCards == 2):
		totalNumberHands = 19600
	elif(numCards == 4):
		totalNumberHands = 48
	while(count < iterations):
		part_deck = getPartialDeck(incomplete_hand)
		random.shuffle(part_deck)
		add_cards = part_deck[:(5-numCards)]
		hand = Hand(add_cards + incomplete_hand)
		level = hand.rank['level']
		
		if(hand.rank['level'] == 3):
			exists = inCardList(add_cards,poss_set_3)
			if(not exists):
				poss_set_3.append(add_cards)
		elif(hand.rank['level'] == 2):
			exists = inCardList(add_cards,poss_set_2)
			if(not exists):
				poss_set_2.append(add_cards)
		elif(hand.rank['level'] == 4):
			exists = inCardList(add_cards,poss_set_4)
			if(not exists):
				poss_set_4.append(add_cards)
		elif(hand.rank['level'] == 7):
			exists = inCardList(add_cards,poss_set_7)
			if(not exists):
				poss_set_7.append(add_cards)
		elif(hand.rank['level'] == 6):
			exists = inCardList(add_cards,poss_set_6)
			if(not exists):
				poss_set_6.append(add_cards)
		elif(hand.rank['level'] == 5):
			exists = inCardList(add_cards,poss_set_5)
			if(not exists):
				poss_set_5.append(add_cards)
		count = count + 1

	# print "number for full house is : " + str(len(poss_set_7)) #+ " " + str(poss_set_7)
	# print "number for flush is : " + str(len(poss_set_6))
	# print "number for straight is : " + str(len(poss_set_5))
	# print "number for three of a kind is : " + str(len(poss_set_4))

	return [len(poss_set_7)*100/float(totalNumberHands),len(poss_set_6)*100/float(totalNumberHands),len(poss_set_5)*100/float(totalNumberHands),len(poss_set_4)*100/float(totalNumberHands),len(poss_set_3)*100/float(totalNumberHands),len(poss_set_2)*100/float(totalNumberHands)]


def getPartialDeck(list_removed_cards):
	deck = getSortedDeck()
	for card in list_removed_cards:
		# print card
		deck.remove(card)
	return deck


def inCardList(list_minor,list_major):
	for lista in list_major:
		if all(x in lista for x in list_minor):
			return True
	return False

def nCr(n,r):
	f = math.factorial
	return f(n)/(f(n-r)*f(r))

# print nCr(47,2)

# print getProbHand(incHand)

# lista = [Card(4,'D'),Card(6,'S'),Card(6,'D')]
# listb = [Card(6,'S'),Card(6,'D'),Card(4,'D')]

