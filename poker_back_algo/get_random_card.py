import random
from poker_card import Card

def getCard():
	value = random.randint(2,14)
	suits = ['S','D','H','C']
	suit = random.choice(suits)
	card = Card(value,suit)
	return card