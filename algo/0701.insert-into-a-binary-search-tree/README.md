# [701 Insert Into a Binary Search Tree](https://leetcode-cn.com/problems/insert-into-a-binary-search-tree/)

## 题目

给定二叉搜索树（BST）的根节点和要插入树中的值，将值插入二叉搜索树。 返回插入后二叉搜索树的根节点。 保证原始二叉搜索树中不存在新值。

注意，可能存在多种有效的插入方式，只要树在插入后仍保持为二叉搜索树即可。 你可以返回任意有效的结果。

例如, 

给定二叉搜索树:

```
        4
       / \
      2   7
     / \
    1   3

和 插入的值: 5
```

你可以返回这个二叉搜索树:

```
         4
       /   \
      2     7
     / \   /
    1   3 5
```
    
或者这个树也是有效的:

```
         5
       /   \
      2     7
     / \   
    1   3
         \
          4
```


## 解题思路


思路一：新插入的值作为叶子节点的子节点插入，递归遍历对比各节点并插入。


思路二：迭代



## 总结