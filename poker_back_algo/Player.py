from pymouse import PyMouseEvent

print(dir(PyMouseEvent.click))


# class ListenInterrupt(Exception):
#     pass

class handler(PyMouseEvent):
    def click(self, x, y, button,press):
        print("there was a click at", x, y, "with this button:", button , " press is " ,press)
        if(not press):
        	raise Exception("Clicked at ",str(x)+" , "+str(y))
        return x, y, button

    def move(self, x, y):
        print("the mouse was moved to", x, y)
        return x, y

test = handler() # start listening
print(dir(test))
try:
    test.run()
except Exception as e:
	print(e)
	pass
    

# print("sadas")
# del test # stop listening
# print(test.is_alive())

# def get_click_data():
# 	t = handler()
# 	print(t.is_alive())
# 	t.run()
# 	print(t.is_alive())


# get_click_data()