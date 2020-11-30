def getFreq(v,x):
	count = 0
	for val in x:
		if(val == v):
			count = count + 1
	return count

def getFreqDict(x):
	dL = []
	for val in x:
		freq=getFreq(val,x)
		tup = (val,freq)
		dL.append(tup)
	return dict(list(set(dL)))
