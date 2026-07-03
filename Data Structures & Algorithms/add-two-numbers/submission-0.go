/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    temp1, temp2 := l1, l2
	carry := 0
	result := &ListNode{}
	curr := result

	for temp1 != nil && temp2 != nil {
		sum := temp1.Val + temp2.Val + carry
		carry = 0
		if sum >= 10 {
			carry = sum / 10
			sum = sum % 10
		}

		curr.Next = &ListNode{Val: sum}

		curr = curr.Next
		temp1 = temp1.Next
		temp2 = temp2.Next
	}

	for temp1 != nil {
		sum := temp1.Val + carry
		carry = 0
		if sum >= 10 {
			carry = sum / 10
			sum = sum % 10
		}

		curr.Next = &ListNode{Val: sum}
		curr = curr.Next
		temp1 = temp1.Next
	}

	for temp2 != nil {
		sum := temp2.Val + carry
		carry = 0
		if sum >= 10 {
			carry = sum / 10
			sum = sum % 10
		}

		curr.Next = &ListNode{Val: sum}
		curr = curr.Next
		temp2 = temp2.Next
	}

	if carry > 0 {
		curr.Next = &ListNode{Val: carry}
	}

	return result.Next
}
