func topKFrequent(nums []int, k int) []int {
    count := make(map[int]int)
    for _, num := range nums {
        count[num]++
    }

    pq := NewPriorityQueue[int]()

    for num, freq := range count {
        pq.Push(num, freq)

        if pq.Len() > k {
            pq.Pop()
        }
    }

    res := make([]int, k)
    for i := 0; i < k; i++ {
        v, _, _ := pq.Pop()
        res[i] = v
    }
    return res
}

type Item[T any] struct {
    Value T
    Priority int
}

type itemHeap[T any] []Item[T]

func (h itemHeap[T]) Len() int {
    return len(h)
}

func (h itemHeap[T]) Less(i, j int) bool {
    return h[i].Priority < h[j].Priority
}

func (h itemHeap[T]) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *itemHeap[T]) Push(x any) {
    *h = append(*h, x.(Item[T]))
}

func (h *itemHeap[T]) Pop() any {
    old := *h
    n := len(old)
    item := old[n-1]
    *h = old[:n-1]
    return item
}

type PriorityQueue[T any] struct {
    h itemHeap[T]
}

func NewPriorityQueue[T any]() *PriorityQueue[T] {
    pq := &PriorityQueue[T]{}
    heap.Init(&pq.h)
    return pq
}

func (pq *PriorityQueue[T]) Push(value T, priority int) {
    heap.Push(&pq.h, Item[T]{
        Value: value,
        Priority: priority,
    })
}

func (pq *PriorityQueue[T]) Pop() (T, int, bool) {
    var zero T

    if pq.Len() == 0 {
        return zero, 0, false
    }

    item := heap.Pop(&pq.h).(Item[T])
    return item.Value, item.Priority, true
}

func (pq *PriorityQueue[T]) Len() int {
    return pq.h.Len()
}