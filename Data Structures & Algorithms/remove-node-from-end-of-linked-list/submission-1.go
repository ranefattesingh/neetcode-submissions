/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	count := 0
	
	for curr := head; curr != nil; curr = curr.Next {
		count++
	}

	removeIndex := count - n
	if removeIndex == 0 {
		return head.Next
	}

	curr := head
	for i := 0; i < count-1; i++ {
		if i + 1 == removeIndex {
			curr.Next = curr.Next.Next
			break
		}

		curr = curr.Next
	}

	return head
}
