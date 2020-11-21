package main

func main() {

}

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] && dfs(board, i, j, 0, word) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, i int, j int, idx int, word string) bool {
	if idx == len(word) {
		return true
	}
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) ||
		board[i][j] != word[idx] {
		return false
	}

	temp := board[i][j]
	board[i][j] = ' '
	found := dfs(board, i+1, j, idx+1, word) || dfs(board, i-1, j, idx+1, word) ||
		dfs(board, i, j+1, idx+1, word) || dfs(board, i, j-1, idx+1, word)
	board[i][j] = temp
	return found
}
