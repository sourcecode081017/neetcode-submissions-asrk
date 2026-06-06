class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        product_prefix = [1] * len(nums)
        p = 1
        for i in range(len(nums)):
            product_prefix[i] = p
            p *= nums[i]
        postfix = 1
        for i in range(len(nums) - 1, -1, -1):
            product_prefix[i] *= postfix
            postfix *= nums[i]
        return product_prefix
        
        