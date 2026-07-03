/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    temp1, temp2 := l1, l2
	prevCarry := 0
	result := &ListNode{}
	curr := result

	for temp1 != nil && temp2 != nil {
		sum, newCarry := add(temp1.Val, temp2.Val, prevCarry)
		prevCarry = newCarry
		curr.Next = &ListNode{Val: sum}

		curr = curr.Next
		temp1 = temp1.Next
		temp2 = temp2.Next
	}

	for temp1 != nil {
		sum, newCarry := add(temp1.Val, 0, prevCarry)
		prevCarry = newCarry
		curr.Next = &ListNode{Val: sum}

		curr = curr.Next
		temp1 = temp1.Next
	}

	for temp2 != nil {
		sum, newCarry := add(0, temp2.Val, prevCarry)
		prevCarry = newCarry
		curr.Next = &ListNode{Val: sum}

		curr = curr.Next
		temp2 = temp2.Next
	}

	if prevCarry > 0 {
		curr.Next = &ListNode{Val: prevCarry}
	}

	return result.Next
}

func add(digit1, digit2, prevCarry int) (sum, newCarry int) {
	total := digit1 + digit2 + prevCarry
	newCarry = total / 10
	sum = total % 10

	return sum, newCarry
} 
