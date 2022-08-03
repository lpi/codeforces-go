#### 提示 1

$\textit{nums}_1$ 要是变成 $[0,1,2,\dots, n-1]$ 就会简单不少。

#### 提示 2

枚举 $y$。

---

#### 前置知识：置换

**置换**是一个排列到另一个排列的双射。

以示例 2 为例，定义下列置换 $P(x)$：

$$
\left(\begin{array}{cccc}
x & 0 & 1 & 2 & 3 & 4\\
P(x) & 1 & 2 & 4 & 3 & 0
\end{array}\right)
$$

我们可以把 $[4,0,1,3,2]$ 中的每个元素 $x$ 替换为 $P(x)$，这样可以得到一个新的排列 $[0,1,2,3,4]$。同理可以将 $[4,1,0,2,3]$ 通过置换得到新的排列 $[0,2,1,4,3]$。

---

将 $\textit{nums}_1$ 置换成 $[0,1,2,\dots, n-1]$，设这一置换为 $P(x)$，将 $P(x)$ 也应用到 $\textit{nums}_2$ 上。对于 $\textit{nums}_1$ 和 $\textit{nums}_2$ 中的相同元素，在置换后仍然是相同的，且元素的位置仍然是不变的，因此置换操作不会影响答案个数。

由于 $\textit{nums}_1$ 置换成了 $[0,1,2,\dots, n-1]$，因此置换后的好三元组 $(x,y,z)$ 需满足 $x<y<z$。枚举置换后的 $\textit{nums}_2$ 中的 $y$，问题就变成计算元素 $y$ 的左侧有多少个比 $y$ 小的数，以及右侧有多少个比 $y$ 大的数。这可以用树状数组/线段树/名次树来完成（Python 可以直接用 `SortedList`），下面代码用的是树状数组。

设 $y$ 的下标为 $i$，且其左侧有 $\textit{less}$ 个数比 $y$ 小，由于比 $y$ 大的数有 $n-1-y$ 个（注意 $y$ 的范围为 $[0,n-1]$），减去左侧比 $y$ 大的 $i-\textit{less}$ 个数，因此 $y$ 右侧有 $n-1-y-(i-\textit{less})$ 个数比它大。所以 $y$ 会有

$$
\textit{less}\cdot(n-1-y-(i-\textit{less}))
$$

个好三元组。

累加所有 $y$ 的好三元组个数，即为答案。

注意下面代码使用的是值域在 $[1,n]$ 的树状数组，需要对插入和查询的数额外加一。

- 时间复杂度：$O(n\log n)$。
- 空间复杂度：$O(n)$。

```go [sol1-Go]
func goodTriplets(nums1, nums2 []int) (ans int64) {
	n := len(nums1)
	p := make([]int, n)
	for i, v := range nums1 {
		p[v] = i
	}
	tree := make([]int, n+1)
	for i := 1; i < n-1; i++ {
		for j := p[nums2[i-1]] + 1; j <= n; j += j & -j { // 将 p[nums2[i-1]]+1 加入树状数组
			tree[j]++
		}
		y, less := p[nums2[i]], 0
		for j := y; j > 0; j &= j - 1 { // 计算 less
			less += tree[j]
		}
		ans += int64(less) * int64(n-1-y-(i-less))
	}
	return
}
```

```C++ [sol1-C++]
class Solution {
public:
    long long goodTriplets(vector<int> &nums1, vector<int> &nums2) {
        int n = nums1.size();
        vector<int> p(n);
        for (int i = 0; i < n; ++i)
            p[nums1[i]] = i;
        long long ans = 0;
        vector<int> tree(n + 1);
        for (int i = 1; i < n - 1; ++i) {
            for (int j = p[nums2[i - 1]] + 1; j <= n; j += j & -j) // 将 p[nums2[i-1]]+1 加入树状数组
                ++tree[j];
            int y = p[nums2[i]], less = 0;
            for (int j = y; j; j &= j - 1) // 计算 less
                less += tree[j];
            ans += (long) less * (n - 1 - y - (i - less));
        }
        return ans;
    }
};
```

```Python [sol1-Python3]
class Solution:
    def goodTriplets(self, nums1: List[int], nums2: List[int]) -> int:
        n = len(nums1)
        p = [0] * n
        for i, x in enumerate(nums1):
            p[x] = i
        ans = 0
        tree = [0] * (n + 1)
        for i in range(1, n - 1):
            # 将 p[nums2[i - 1]] + 1 加入树状数组
            j = p[nums2[i - 1]] + 1
            while j <= n:
                tree[j] += 1
                j += j & -j
            # 计算 less
            y, less = p[nums2[i]], 0
            j = y
            while j:
                less += tree[j]
                j &= j - 1
            ans += less * (n - 1 - y - (i - less))
        return ans
```

```java [sol1-Java]
class Solution {
    public long goodTriplets(int[] nums1, int[] nums2) {
        var n = nums1.length;
        var p = new int[n];
        for (var i = 0; i < n; ++i)
            p[nums1[i]] = i;
        var ans = 0L;
        var tree = new int[n + 1];
        for (var i = 1; i < n - 1; ++i) {
            for (var j = p[nums2[i - 1]] + 1; j <= n; j += j & -j) // 将 p[nums2[i-1]]+1 加入树状数组
                ++tree[j];
            var y = p[nums2[i]];
            var less = 0;
            for (var j = y; j > 0; j &= j - 1) // 计算 less
                less += tree[j];
            ans += (long) less * (n - 1 - y - (i - less));
        }
        return ans;
    }
}
```

附 Python `SortedList` 做法：

```Python
from sortedcontainers import SortedList

class Solution:
    def goodTriplets(self, nums1: List[int], nums2: List[int]) -> int:
        n = len(nums1)
        p = [0] * n
        for i, x in enumerate(nums1):
            p[x] = i
        ans = 0
        s = SortedList()
        for i in range(1, n - 1):
            s.add(p[nums2[i - 1]])
            y = p[nums2[i]]
            less = s.bisect_left(y)
            ans += less * (n - 1 - y - (i - less))
        return ans
```

有一道题也用到了这种置换思想：

- [1713. 得到子序列的最少操作次数](https://leetcode-cn.com/problems/minimum-operations-to-make-a-subsequence/)
