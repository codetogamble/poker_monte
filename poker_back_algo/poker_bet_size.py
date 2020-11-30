def percentage_bet(n_cards):
	p = (93-n_cards)*(n_cards)
	pr = float(p)/(2162-p)
	# print float(p)/2162
	return pr*100

# for x in range(0,15):
# 	pb = percentage_bet(x)
# 	print ("for "+str(x)+" outs, pecentage bet is : " + str(pb))
# 	if(x>0):
# 		print("difference is " + str(pb - prev_b))
# 	prev_b = pb
