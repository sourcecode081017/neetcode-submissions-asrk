class Solution:
    def numIslands(self, grid: List[List[str]]) -> int:
        ROWS = len(grid)
        COLS = len(grid[0])
        num_isles = 0
        for i in range(0, ROWS):
            for j in range(0, COLS):
                if grid[i][j] == '1':
                    self._dfs(i, j, grid)
                    num_isles += 1
        return num_isles
    def _dfs(self, i, j, grid: List[List[str]]) -> None:
        #check bounds
        if i < 0 or i >= len(grid) or j < 0 or j >= len(grid[0]) or grid[i][j] != '1':
            return
        # visit grid
        grid[i][j] = '#'
        self._dfs(i + 1, j, grid)
        self._dfs(i - 1, j, grid)
        self._dfs(i, j + 1, grid)
        self._dfs(i, j - 1, grid)
