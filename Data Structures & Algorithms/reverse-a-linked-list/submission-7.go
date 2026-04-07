/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {

	var temp *ListNode
	temp = nil
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = temp
		temp = curr
		curr = next
	}

	return temp

}
