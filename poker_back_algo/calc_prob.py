from get_card_deck import getSortedDeck
from poker_card import Card
from poker_hand import *
import random
from get_best_hand import getBestHand
import time

a = time.time()

low_hand = Hand([Card(2,'D'),Card(3,'S'),Card(4,'S'),Card(5,'S'),Card(7,'S')])

iterations = 10000
SFCount = 0
FKCount = 0
FHCount = 0
FLCount = 0
STCount = 0
TKCount = 0
TPCount =0
PCount = 0
HCCount=0
count = 0
while(count<iterations):
	# print(count)
	deck = getSortedDeck()
	random.shuffle(deck)
	sev_cards = deck[:5]
	# bHand = getBestHand(sev_cards)
	bHand = Hand(sev_cards)
	if(bHand.rank['level'] == 9):
		SFCount = SFCount + 1
	elif(bHand.rank['level'] == 8):
		FKCount = FKCount + 1
	elif(bHand.rank['level'] == 7):
		FHCount = FHCount + 1
	elif(bHand.rank['level'] == 6):
		FLCount = FLCount + 1
	elif(bHand.rank['level'] == 5):
		STCount = STCount + 1
	elif(bHand.rank['level'] == 4):
		TKCount = TKCount + 1
	elif(bHand.rank['level'] == 3):
		TPCount = TPCount + 1
	elif(bHand.rank['level'] == 2):
		PCount = PCount + 1
	else:
		HCCount = HCCount + 1
	count = count + 1

print("probability of straight flush is " + str((float(SFCount)*100.0/float(iterations))) + "%")
print("probability of four of kind is " + str((float(FKCount)*100.0/float(iterations))) + "%")
print("probability of full house is " + str((float(FHCount)*100.0/float(iterations))) + "%")
print("probability of flush is " + str((float(FLCount)*100.0/float(iterations))) + "%")
print("probability of straight is " + str((float(STCount)*100.0/float(iterations))) + "%")
print("probability of 3 of a kind is " + str((float(TKCount)*100.0/float(iterations))) + "%")
print("probability of two pair is " + str((float(TPCount)*100.0/float(iterations))) + "%")
print("probability of pair is " + str((float(PCount)*100.0/float(iterations))) + "%")
print("probability of HighCard is " + str((float(HCCount)*100.0/float(iterations))) + "%")

b = time.time()

print(b-a)