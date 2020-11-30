from get_freq_list import getFreqDict

class Hand:
	def __str__(self):
		s = ""
		for x in self.hand:
			s = s + str(x.value) + "-" + x.suit + " "
		return s

	def __repr__(self):
		return str(self)

	# def __eq__(self,other):
	# 	return (self.valueList == other.valueList and self.suitList == other.suitList)
	
	def __init__(self,hand):
		self.valueList = []
		self.suitList = []
		self.hand=[]
		self.sortedVL = []
		self.Nc = 0
		self.Nr = 0
		self.Vmax = 0
		self.SSD = 0
		self.freqDict = {}
		self.rank = {}
		self.hand = hand
		self.maxF = 0
		self.maxFValue = 0
		for card in self.hand:
			self.valueList.append(card.value)
			self.suitList.append(card.suit)
		self.Nc = len(set(self.valueList))
		self.Nr = len(set(self.suitList))
		self.sortedVL = sorted(self.valueList)
		self.Vmax = sorted(self.valueList)[-1]
		self.Vmax2 = sorted(self.valueList)[-2]
		SSD_1 = self.sortedVL[-2]-self.sortedVL[0]
		if(SSD_1 == 3 and ((self.sortedVL[-1]-self.sortedVL[-2]) == 1 or self.sortedVL[-1]-self.sortedVL[-2] == 9)):
			self.SSD = 4
		else:
			self.SSD = self.sortedVL[-1] - self.sortedVL[0]
		self.freqDict = getFreqDict(self.sortedVL)
		# print self.freqDict
		# print (self.freqDict).keys()
		for key in (self.freqDict).keys():
			if(self.maxF <= self.freqDict[key]):
				self.maxF = self.freqDict[key]
				if(key > self.maxFValue):
				 	self.maxFValue = key
		self.calcRank()


	def getNc(self):
		return self.Nc

	def getNr(self):
		return self.Nr

	def getVmax(self):
		return self.Vmax

	def getSortedValueList(self):
		return self.sortedVL

	def getSSD(self):
		return self.SSD

	def getVmax2(self):
		return self.Vmax2

	def getFreqDict(self):
		return self.freqDict

	def getRank(self):
		return self.rank

	def getMaxFreq(self):
		return self.maxF

	def calcRank(self):
		if(self.Nc == 5 and self.Nr == 1 and self.SSD == 4):
			self.rank = {'level':9,'sorted_value_list':self.sortedVL,'hand':'straight flush'}
		elif(self.Nc == 2 and self.maxF == 4):
			self.rank = {'level':8,'sorted_value_list':self.sortedVL,'hand':'four of a kind'}
		elif(self.Nc == 2 and self.maxF == 3):
			self.rank = {'level':7,'sorted_value_list':self.sortedVL,'hand':'full house'}
		elif(self.Nr == 1):
			self.rank = {'level':6,'sorted_value_list':self.sortedVL,'hand':'flush'}
		elif(self.Nc == 5 and self.SSD == 4):
			self.rank = {'level':5,'sorted_value_list':self.sortedVL,'hand':'straight'}
		elif(self.Nc == 3 and self.maxF == 3):
			self.rank = {'level':4,'sorted_value_list':self.sortedVL,'hand':'three of a kind'}
		elif(self.Nc == 3 and self.maxF == 2):
			self.rank = {'level':3,'sorted_value_list':self.sortedVL,'hand':'two pair'}
		elif(self.Nc == 4):
			self.rank = {'level':2,'sorted_value_list':self.sortedVL,'hand':'pair'}
		else:
			self.rank = {'level':1,'sorted_value_list':self.sortedVL,'hand':'highcard'}

	def compareHand(self,hand):
		if(self.rank['level'] > hand.rank['level']):
			return 1
		elif(self.rank['level'] < hand.rank['level']):
			return -1
		else:
			if(self.rank['level'] == 9 or self.rank['level'] == 5):
				if(self.sortedVL == hand.sortedVL):
					return 0
				elif(self.sortedVL[3] > hand.sortedVL[3]):
					return 1
				elif(self.sortedVL[3] < hand.sortedVL[3]):
					return -1
			elif(self.rank['level'] == 8 or self.rank['level'] == 7):
				if(self.sortedVL == hand.sortedVL):
					return 0
				elif(self.maxFValue > hand.maxFValue):
					return 1
				elif(self.maxFValue < hand.maxFValue):
					return -1
				else:
					if(sum(self.sortedVL) > sum(hand.sortedVL)):
						return 1
					elif(sum(self.sortedVL) < sum(hand.sortedVL)):
						return -1
					else:
						return 0
			elif(self.rank['level'] == 3 or self.rank['level'] == 2):
				if(self.sortedVL == hand.sortedVL):
					return 0
				pairs_self = []
				pairs_other = []
				for x in(self.freqDict).keys():
					if(self.freqDict[x] == 2):
						pairs_self.append(x)
				for x in(hand.freqDict).keys():
					if(hand.freqDict[x] == 2):
						pairs_other.append(x)
				if(max(pairs_self) > max(pairs_other)):
					return 1
				elif(max(pairs_self) < max(pairs_other)):
					return -1
				elif(max(pairs_self) == max(pairs_other)):
					if(min(pairs_self) > min(pairs_other)):
						return 1
					elif(min(pairs_self) < min(pairs_other)):
						return -1
					else:
						for x in range(0,5):
							if(self.sortedVL[4-x] > hand.sortedVL[4-x]):
								return 1
								break
							elif(self.sortedVL[4-x] < hand.sortedVL[4-x]):
								return -1
								break
						
				
			else:
				if(self.sortedVL == hand.sortedVL):
					return 0
				else:
					for x in range(0,5):
						if(self.sortedVL[4-x] > hand.sortedVL[4-x]):
							return 1
							break
						elif(self.sortedVL[4-x] < hand.sortedVL[4-x]):
							return -1
							break


