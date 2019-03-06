package leetcode096

/*
   F(i, n) 代表不同BST的个数，i表示根结点，n代表1--n
   G(n) 代表1--n不同BST个数，且G(0)=G(1)=1。 所以得
   公式1: G(n) = F(1, n) + F(2, n) + ... + F(n, n)
   因为是二叉搜索树，所以左节点都小于i右节点都大于i 所以得
   公式2: F(i, n) = G(i-1) * G(n-i)
   公式2带入公式1得
   公式3: G(n) = G(0) * G(n-1) + G(1) * G(n-2) + ... + G(n-1) * G(0)
*/

func numTrees(n int) int {
	G := make([]int, n+1)
	G[0], G[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			G[i] += G[j] * G[i-j-1]
		}
	}
	return G[n]
}
