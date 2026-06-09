class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        row_map = defaultdict(set)
        col_map = defaultdict(set)
        sq_map = defaultdict(set)
        for i in range(0, 9):
            for j in range(0, 9):
                if board[i][j] == '.':
                    continue
                if (board[i][j] in row_map[i]
                    or board[i][j] in col_map[j]
                    or board[i][j] in sq_map[(i // 3, j // 3)]):
                    return False
                
                col_map[j].add(board[i][j])
                row_map[i].add(board[i][j])
                sq_map[(i // 3, j // 3)].add(board[i][j])
        return True
        