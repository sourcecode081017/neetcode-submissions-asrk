class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        sum_map = {}
        for i, num in enumerate(nums):
            if (target - num) in sum_map:
                if i > sum_map[target - num]:
                    return [sum_map[target - num], i]
                return [i, sum_map[target - num]]
            sum_map[num] = i
        return []
        