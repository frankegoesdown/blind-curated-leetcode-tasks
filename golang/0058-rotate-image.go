package main

func main() {

}

func rotate(m [][]int) {
	n := len(m)
	for i := 0; i < n/2; i++ {
		for j := i; j < n-i-1; j++ {
			temp := m[i][j]
			m[i][j] = m[n-j-1][i]
			m[n-j-1][i] = m[n-i-1][n-j-1]
			m[n-i-1][n-j-1] = m[j][n-i-1]
			m[j][n-i-1] = temp
		}
	}
}
