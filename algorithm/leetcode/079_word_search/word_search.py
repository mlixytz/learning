class Solution:
    def exist(self, board, word):
        for i, _ in enumerate(board):
            for j, _ in enumerate(board[i]):
                if word[0] == board[i][j]:
                    old = board[i][j]
                    board[i][j] = -1
                    if self.traceback(board, word[1:], i, j):
                        return True
                    board[i][j] = old
        return False
    
    def traceback(self, board, word, i, j):
        if not word:
            return True
        
        if i < len(board) - 1 and word[0] == board[i+1][j]:
            old = board[i+1][j]
            board[i+1][j] = -1
            if self.traceback(board, word[1:], i+1, j):
                return True
            board[i+1][j] = old

        if i > 0 and word[0] == board[i-1][j]:
            old = board[i-1][j]
            board[i-1][j] = -1
            if self.traceback(board, word[1:], i-1, j):
                return True
            board[i-1][j] = old

        if j < len(board[0]) - 1 and word[0] == board[i][j+1]:
            old = board[i][j+1]
            board[i][j+1] = -1
            if self.traceback(board, word[1:], i, j+1):
                return True
            board[i][j+1] = old

        if j > 0 and word[0] == board[i][j-1]:
            old = board[i][j-1]
            board[i][j-1] = -1
            if self.traceback(board, word[1:], i, j-1):
                return True
            board[i][j-1] = old


if __name__ == '__main__':
    s = Solution()
    board = [['A','B','C','E'],['S','F','C','S'],['A','D','E','E']]
    word = 'ABCCED'
    print(s.exist(board, word))
