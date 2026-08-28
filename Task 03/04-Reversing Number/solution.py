# Reverse the number given (for both negetive and positive numbers)
class Solution(object):
    def reverse(self, x):
        rev=0
        if x>=0:
            sign=1
        else:
            sign=-1
        x=abs(x)
        while x>0:
            digit=x%10
            rev=(rev*10)+digit
            x=x//10

        rev=rev*sign

        if rev<-2**31 or rev>2**31 -1:
            return 0
        return rev


        
