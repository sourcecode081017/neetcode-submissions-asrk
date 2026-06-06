class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        graph = defaultdict(list)
        for s in strs:
            char_items = [0] * 26
            for c in s:
                char_items[ord(c) - ord('a')] += 1
            graph[tuple(char_items)].append(s)
        return list(graph.values())