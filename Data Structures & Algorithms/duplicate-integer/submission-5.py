class Solution:
    def hasDuplicate(self, nums: List[int]) -> bool:
        contains_list = {}
        for x in nums:
            if x in contains_list:
                return True
            contains_list[x] = True
        return False

        