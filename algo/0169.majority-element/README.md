# [169. Majority Element](https://leetcode-cn.com/problems/majority-element/)

## 题目

给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。

 

示例 1:

输入: [3,2,3]
输出: 3
示例 2:

输入: [2,2,1,1,1,2,2]
输出: 2


## 解题思路

思路一：先排序，直接通过 `nums[len(nums)/2]` 获取

思路二：通过 map 空间换时间


细细分析，这个问题属于多数投票算法。


## 总结