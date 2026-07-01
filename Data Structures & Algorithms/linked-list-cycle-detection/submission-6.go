/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
    curr := head
	third := head

	for third != nil && third.Next != nil {
		third = third.Next.Next
		curr = curr.Next
		if third == curr {
			return true
		}
	}

	return false
}
