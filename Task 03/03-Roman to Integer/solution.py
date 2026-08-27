# Convert the given Roman Number to Integer
class Solution(object):
    def romanToInt(self, s):
        n=0
        values={"I":1, "V":5, "X":10, "L":50, "C":100, "D":500, "M":1000}
        for i in range(len(s)):
            if i+1<len(s) and values[s[i]]<values[s[i+1]]:
                n-=values[s[i]]
            else:
                n+=values[s[i]]
        return n
