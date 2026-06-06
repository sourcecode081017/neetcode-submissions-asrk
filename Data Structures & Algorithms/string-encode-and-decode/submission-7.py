class Solution:

    def encode(self, strs: List[str]) -> str:
        if not strs:
            return "nullstring"
        return "~".join(strs)

    def decode(self, s: str) -> List[str]:
        if s == "nullstring":
            return []
        if s == "":
            return [""]
        return s.split("~")
