from poker_card import Card


def getSortedDeck():
	sorted_deck = []
	suits = ['S','C','H','D']
	for suit in suits:
		for val in range(2,15):
			card = Card(val,suit)
			sorted_deck.append(card)
	return sorted_deck