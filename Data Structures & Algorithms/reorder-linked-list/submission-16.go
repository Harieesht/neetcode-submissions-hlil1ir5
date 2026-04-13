/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {

	slow := head
	fast := head

	//find the mid point using slow ad fast pointer
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	curr := slow.Next
	slow.Next = nil // cutting off the second set of nodes

	// now reverse the second half of the nodes

	var prev *ListNode
	var next *ListNode
	prev = nil
	

	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	first := head
	second := prev

	for second != nil{
		firstNext := first.Next
		secondNext := second.Next

		first.Next = second
		second.Next = firstNext

		first = firstNext
		second = secondNext

	}

  
}
