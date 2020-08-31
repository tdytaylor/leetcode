### Given a binary tree, find its maximum depth.

The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

Note: A leaf is a node with no children.

Example:

Given binary tree [3,9,20,null,null,15,7],  
```mermaid
graph TB
3((3))
9((9))
20((20))
15((15))
7((7))
3-->9
3-->20
20-->15
20-->7
```