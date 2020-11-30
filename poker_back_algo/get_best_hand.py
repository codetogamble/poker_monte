from poker_card import Card
from poker_hand import *

low_hand = Hand([Card(2,'D'),Card(3,'S'),Card(4,'S'),Card(5,'S'),Card(7,'S')])
def getBestHand(seven_cards):
	best_hand = low_hand
	if(len(seven_cards) == 7):
		for x in range(0,7):
			for y in range(x+1,7):
				hand = seven_cards[0:x] + seven_cards[x+1:y] + seven_cards[y+1:]
				newHand = Hand(hand)
				if(newHand.compareHand(best_hand) == 1):
					best_hand = newHand
		return best_hand
	elif(len(seven_cards) == 6):
		for x in range(0,6):
			hand = seven_cards[0:x] + seven_cards[x+1:]
			newHand = Hand(hand)
			if(newHand.compareHand(best_hand) == 1):
				best_hand = newHand
		return best_hand
	elif(len(seven_cards) == 5):
		return Hand(seven_cards)


