/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	
	for curr := head; curr != nil; curr = curr.Next.Next {
		node := &Node{
			Val: curr.Val,
			Next: curr.Next,
		}

		curr.Next = node
	}

	
	for curr := head; curr != nil; curr = curr.Next.Next {
		if curr.Random != nil {
			curr.Next.Random = curr.Random.Next
		}
	}

	copyHead := head.Next

	for curr := head; curr != nil; curr = curr.Next {
		copy := curr.Next
		curr.Next = copy.Next

		if copy.Next != nil {
			copy.Next = copy.Next.Next
		}
	}

	return copyHead
}
