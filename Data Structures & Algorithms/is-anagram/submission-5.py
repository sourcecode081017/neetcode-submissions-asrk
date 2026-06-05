class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if s == t:
            return True
        if len(s) != len(t):
            return False
        char_list = [0] * 26
        for i in range(len(s)):
            char_list[ord(s[i]) - ord('a')] += 1
            char_list[ord(t[i]) - ord('a')] -= 1
        for x in char_list:
            if x != 0:
                return False
        return True
        