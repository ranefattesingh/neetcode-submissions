type Node struct {
	num   int
	count int
}

type PriorityQueue []Node

func (h *PriorityQueue) Pop() any {
	a := *h
	lastElement := a[len(a)-1]
	*h = a[:len(a)-1]
	return lastElement
}

func (h *PriorityQueue) Push(x any) {
	*h = append(*h, x.(Node))
}

func (h PriorityQueue) Less(i, j int) bool {
	return h[i].count < h[j].count
}

func (h PriorityQueue) Len() int {
	return len(h)
}

func (h PriorityQueue) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func topKFrequent(nums []int, k int) []int {
	store := make(map[int]int)
	for _, v := range nums {
		store[v]++
	}

	queue := make(PriorityQueue, 0, len(store))

	heap.Init(&queue)

	for hk, v := range store {
		heap.Push(&queue, Node{num: hk, count: v})
		if queue.Len() > k {
			heap.Pop(&queue)
		}
	}

	topFrequent := make([]int, 0, k)
	for _, v := range queue {
		topFrequent = append(topFrequent, v.num)
	}

	return topFrequent
}