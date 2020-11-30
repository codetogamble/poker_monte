# Card class has 2 member variables i.e value and suit

class Card:
	value = 0
	suit = 'N'

	def __str__(self):
		return str(self.value) + " " + self.suit

	def __init__(self, value,suit):
		self.value = value
		self.suit = suit
		self.validate()

	def __eq__(self,other):
		return (self.value == other.value and self.suit == other.suit)

	def __repr__(self):
		return str(self)

	def validate(self):
		if(self.value>14 or self.value < 2):
			raise ValueError("invalid card")
		elif(self.suit != 'S' and self.suit != 'C' and self.suit != 'D' and self.suit != 'H'):
			raise ValueError("invalid card")