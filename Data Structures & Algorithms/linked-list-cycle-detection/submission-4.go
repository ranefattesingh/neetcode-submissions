/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return false
	}

	fp := head.Next.Next
	sp := head

	for fp != nil && fp.Next != nil {
		sp = sp.Next
		fp = fp.Next.Next

		if fp == sp {
			return true
		}
	}

	return false
}
