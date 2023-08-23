package main

import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	nums     []int
	indexMap map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{[]int{}, map[int]int{}}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.indexMap[val]; ok {
		return false
	}
	this.nums = append(this.nums, val)
	this.indexMap[val] = len(this.nums) - 1

	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.indexMap[val]; !ok {
		return false
	}

	index := this.indexMap[val]

	// 把最后一个元素移到要删除的元素的位置
	this.nums[index] = this.nums[len(this.nums)-1]
	this.indexMap[this.nums[index]] = index
	this.nums = this.nums[:len(this.nums)-1]
	delete(this.indexMap, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}

func main() {
	randomizedSet := Constructor()
	randomizedSet.Insert(1)   // Inserts 1 to the set. Returns true as 1 was inserted successfully.
	randomizedSet.Remove(2)   // Returns false as 2 does not exist in the set.
	randomizedSet.Insert(2)   // Inserts 2 to the set, returns true. Set now contains [1,2].
	fmt.Println(randomizedSet.GetRandom())
	randomizedSet.Remove(1)   // Removes 1 from the set, returns true. Set now contains [2].
	randomizedSet.Insert(2)   // 2 was already in the set, so return false.
	fmt.Println(randomizedSet.GetRandom())
}
