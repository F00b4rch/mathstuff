#!/usr/env/python
import argparse

parser = argparse.ArgumentParser(description='Taking a number to test syracuse.')
parser.add_argument('--number', default='42', type=int, help='an int')

args = parser.parse_args()

def main() :
    sum = 0
    limit = 1
    num = args.number
    print("Starting loop")

    while True :
        sum = sum +1 
		# init
        isEven = evenNumber(num)
        print("turn " + str(sum) + " number is " + str(num))

        if isEven == True :
            num = even(num)
            if num == limit :
                print("Ending here with final value " + str(num) + "after " + str(sum))
                break
			
        elif isEven == False :
            num = nEven(num)
            if num == limit :
                print("Ending here with final value " + str(num) + "after " + str(sum))
                break
		
        else :
            print ("ending")



# Testing even number
def evenNumber(n) :
    if n%2 == 0 :
	    return True
    return False

# Even number
def even(n) :
	res = n / 2
	#print(str(res) + " is new value\n")
	return res

# Not even number
def nEven(n) :
	res = n*3 + 1
	#print(str(res) + " is new value\n")
	return res

main()
