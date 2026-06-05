class Solution:
    def hasDuplicate(self, nums: List[int]) -> bool:
        contains_list = set()
        for x in nums:
            if x in contains_list:
                return True
            contains_list.add(x)
        return False

        