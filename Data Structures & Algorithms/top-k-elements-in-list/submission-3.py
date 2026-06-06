 
'''class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        count_map = dict(Counter(nums))
        pq = []
        print("count map:", count_map)
        for k, v in count_map.items():
            heapq.heappush(pq,(v, k))
            if len(pq) > k:
                heapq.heappop(pq)
        res = []
        for i in range(k):
            res.append(heapq.heappop(pq)[1])
        return res '''

class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        count = {}
        for num in nums:
            count[num] = 1 + count.get(num, 0)

        heap = []
        for num in count.keys():
            heapq.heappush(heap, (count[num], num))
            if len(heap) > k:
                heapq.heappop(heap)

        res = []
        for i in range(k):
            res.append(heapq.heappop(heap)[1])
        return res