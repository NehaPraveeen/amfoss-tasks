# Given two integers dividend and divisor, divide two integers without using multiplication, division, and mod operator.
class Solution(object):
    def divide(self, n, x):
        its_neg= (x>0 and n<0) or (x<0 and n>0)
        x=abs(x)
        n=abs(n)
        quo=0
    
        while n>=x:
            newx=x
            count=1

            while n>=(newx+newx):   #The code logic is completely correct, but it got a Time Limit Exceeded (TLE) error because subtracting 1 from 2,147,483,648 requires the loop to 
                newx=newx+newx      #run 2 billion times. LeetCode cuts off programs that take that long.To fix this while keeping the code easy to understand, we can use a "doubling trick" with addition. 
                count=count+count   #Instead of subtracting x one by one, we can double our divisor (x + x) so we subtract much larger chunks at a time.
            n=n-newx
            quo=quo+count
        
        if its_neg:
            quo=-quo
        
        if quo<-2**31:
            return -2**31
        if quo>(2**31)-1:
            return (2**31)-1
        return quo
        


