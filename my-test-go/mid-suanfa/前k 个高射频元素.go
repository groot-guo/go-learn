package mid_suanfa

import (
	"container/heap"
	"math/rand"
	"sort"
	"time"
)

func topKFrequent(nums []int, k int) []int {
	res := make([]int, 0, len(nums))

	keyMap := make(map[int]int)
	for i := range nums {
		if _, ok := keyMap[nums[i]]; !ok {
			res = append(res, nums[i])
		}
		keyMap[nums[i]]++
	}

	left, right := 0, 1
	for left < len(res)-1 {
		countI := res[left]
		countJ := res[right]
		if keyMap[countI] >= keyMap[countJ] {
			right++
		} else {
			res[left], res[right] = res[right], res[left]
			right++
		}
		if right >= len(res) {
			left++
			right = left + 1
		}
	}

	if k > len(res) {
		return res
	}
	return res[:k]
}

// 速度最快
func topKFrequent22(nums []int, k int) []int {
	frequency := make(map[int]int)
	for _, num := range nums {
		frequency[num]++
	}
	type pair struct {
		k int
		v int
	}
	frequencies := make([]pair, 0, len(frequency))
	for k, v := range frequency {
		frequencies = append(frequencies, pair{k, v})
	}
	sort.Slice(frequencies, func(i, j int) bool {
		return frequencies[i].v > frequencies[j].v
	})
	ans := make([]int, 0, k)
	for i := 0; i < k; i++ {
		ans = append(ans, frequencies[i].k)
	}
	return ans
}

// 堆

type IHeap [][2]int

func (h IHeap) Len() int           { return len(h) }
func (h IHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h IHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent2244(nums []int, k int) []int {
	keyMap := map[int]int{}
	for _, num := range nums {
		keyMap[num]++
	}
	h := &IHeap{}
	heap.Init(h)
	for key, value := range keyMap {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[k-i-1] = heap.Pop(h).([2]int)[0]
	}

	return res
}

// 快速排序

func qsort(values [][]int, start, end int, ret []int, retIndex, k int) {
	rand.Seed(time.Now().UnixNano())
	picked := rand.Int()%(end-start+1) + start
	values[picked], values[start] = values[start], values[picked]

	pivot := values[start][1]
	index := start

	for i := start + 1; i <= end; i++ {
		if values[i][1] >= pivot {
			values[index+1], values[i] = values[i], values[index+1]
			index++
		}
	}

	values[start], values[index] = values[index], values[start]
	if k <= index-start {
		qsort(values, start, index-1, ret, retIndex, k)
	} else {
		for i := start; i <= index; i++ {
			ret[retIndex] = values[i][0]
			retIndex++
		}
		if k > index-start+1 {
			qsort(values, index+1, end, ret, retIndex, k-(index-start+1))
		}
	}
}
