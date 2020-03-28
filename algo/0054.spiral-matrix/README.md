# [54. Spiral Matrix](https://leetcode-cn.com/problems/spiral-matrix/)

## 题目


给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。

示例 1:

```
输入:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]
输出: [1,2,3,6,9,8,7,4,5]
```

示例 2:

```
输入:
[
  [1, 2, 3, 4],
  [5, 6, 7, 8],
  [9,10,11,12]
]
输出: [1,2,3,4,8,12,11,10,9,5,6,7]
```

## 解题思路


按题意，顺时针顺序对矩阵进行遍历；每次遍历都转一圈，每一行或列遍历完后，更新边界值，并且判断是否结束。

定义好上下左右边界初始值；up = 0; d = m-1; l = 0;
假设当前点为(i, j)
遍历顺序： 左->右(u, j+1),完成后up++； 上->下(i+1, r),完成后r--；右->左(d, j-1),完成后d--；下->上(i-1, l),完成后l++



## 总结

