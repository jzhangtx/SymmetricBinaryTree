Symmetric Binary Tree (Please see the docx version of README for better understanding)


A binary tree is considered symmetric if it is a mirror image of itself, i.e, it is symmetric around its root node.

Given the root node of a binary tree, determine whether it's symmetric.

symmetric-binary-tree
Testing
Input Format
The first line contains an integer T denoting the number of test cases.

For each test case, the input has 2 lines:

The first line contains an integer n denoting the number of nodes in the tree (including the NULL nodes).
The second line contains n space-separated integers that will form the binary tree. The integers follow level order traversal of the tree where -1 indicates a NULL node.
Output Format
For each test case, the output contains a line with 1 or 0 based on whether the tree is symmetric or not respectively.

Sample Input
4
7
1 2 2 4 -1 -1 4
7
6 4 4 -1 2 -1 2
7
1 2 3 4 -1 -1 4
1
6
Expected Output
1
0
0
1