# given an integer x, return true if x is a palindrome, and false otherwise.

class Solution(object):
    def isPalindrome(self, x):
        self=str(x)
        new=self[::-1]
        if new==self:
            return True
        else:
            return False



        



        
