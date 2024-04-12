package mid_suanfa

import (
	"math/rand"
	"runtime/debug"
)

type RandomizedSet struct {
	existedMap map[int]bool
	res        map[int]bool
}

func Constructor23() RandomizedSet {
	return RandomizedSet{
		existedMap: make(map[int]bool),
		res:        make(map[int]bool),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if this.existedMap[val] {
		return false
	}
	this.existedMap[val] = true
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if !this.existedMap[val] {
		return false
	}
	delete(this.existedMap, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	if len(this.res) == len(this.existedMap) {
		this.res = make(map[int]bool)
	}
	for i := range this.existedMap {
		if this.res[i] {
			continue
		}
		this.res[i] = true
		return i
	}

	return 0
}

type RandomizedSet33 struct {
	nums    []int
	indices map[int]int
}

func init() { debug.SetGCPercent(-1) }

func Constructor333() RandomizedSet33 {
	return RandomizedSet33{[]int{}, map[int]int{}}
}

func (rs *RandomizedSet33) Insert(val int) bool {
	if _, ok := rs.indices[val]; ok {
		return false
	}
	rs.indices[val] = len(rs.nums)
	rs.nums = append(rs.nums, val)
	return true
}

func (rs *RandomizedSet33) Remove(val int) bool {
	id, ok := rs.indices[val]
	if !ok {
		return false
	}
	last := len(rs.nums) - 1
	rs.nums[id] = rs.nums[last]
	rs.indices[rs.nums[id]] = id
	rs.nums = rs.nums[:last]
	delete(rs.indices, val)
	return true
}

func (rs *RandomizedSet33) GetRandom() int {
	return rs.nums[rand.Intn(len(rs.nums))]
}
