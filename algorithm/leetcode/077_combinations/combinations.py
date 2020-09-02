from copy import copy

# my solution
class Solution:
    def combine(self, n, k):
        res = []
        for i in range(1, n-k+2):
            self.backtrace(i, n, k, [], res)
        return res

    def backtrace(self, cur, n, k, tmp_comb, res):
        tmp_comb.append(cur)
        if len(tmp_comb) == k:
            res.append(tmp_comb)
            return
        if cur > n:
            return
        for i in range(cur+1, n-k+2+len(tmp_comb)):
            self.backtrace(i, n, k, copy(tmp_comb), res)


# clean solution
class Solution1:
    def combine(self, n, k, start=1)
    if k == 0:
        return [[]]
    return [
        [i, *j]
        for i in range(start, n+1) 
        for j in self.combine(n, k-1, i+1)]


# build-in pkg solution
class Solution1:
    def combine(self, n, k, start=1)
        from itertools import combinations
        return combinations(range(1, n+1), k)


if __name__ == '__main__':
    s = Solution()
    print(s.combine(1, 1))