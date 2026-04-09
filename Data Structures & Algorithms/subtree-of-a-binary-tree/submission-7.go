/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {



    if root == nil && subRoot == nil {
        return true
    } else if root == nil || subRoot == nil {
        return false
    } else if root.Val != subRoot.Val {
        return isSubtree(root.Left,subRoot) || isSubtree(root.Right,subRoot)   
    } else {
        res := SameTree(root,subRoot)
        if !res {
           return isSubtree(root.Left,subRoot) || isSubtree(root.Right,subRoot)  
        }

        return res
    }

    
}


func SameTree(root1 *TreeNode,root2 *TreeNode) bool {
    if root1 == nil && root2 == nil {
        return true
    } else if root1 == nil || root2 == nil {
        return false
    } else if root1.Val != root2.Val {
        return false
    } else {
        return SameTree(root1.Left,root2.Left) && SameTree(root1.Right,root2.Right) 
    }

}
