class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        n = len(nums)
        left_product = [0] * n
        right_product = [0] * n
        left_product[0] = right_product[n - 1] = 1
        for i in range(1, n):
            left_product[i] = left_product[i - 1] * nums[i - 1]
        print("left product:", left_product)

        for i in range(n - 2, -1, -1):
            right_product[i] = nums[i + 1] * right_product[i + 1]
        print("right product:", right_product)
        for i in range(n):
            nums[i] = left_product[i] * right_product[i]
        return nums
        