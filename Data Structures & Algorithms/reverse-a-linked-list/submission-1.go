/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
    var rev *ListNode

	for temp := head; temp != nil; {
		next := temp.Next
		temp.Next = rev
		rev = temp
		temp = next
	}

	return rev
}
