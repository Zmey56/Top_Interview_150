// Implement the RandomizedSet class:
//
// RandomizedSet() Initializes the RandomizedSet object.
// bool insert(int val) Inserts an item val into the set if not present. Returns true if the item was not present, false otherwise.
// bool remove(int val) Removes an item val from the set if present. Returns true if the item was present, false otherwise.
// int getRandom() Returns a random element from the current set of elements (it's guaranteed that at least one element
// exists when this method is called). Each element must have the same probability of being returned.

// You must implement the functions of the class such that each function works in average O(1) time complexity.

package RandomizedSet

import (
	"math/rand"
	"time"
)

type RandomizedSet struct {
	rand map[int]struct{}
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		rand: make(map[int]struct{}),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	exist := false
	if _, ok := this.rand[val]; !ok {
		exist = true
		this.rand[val] = struct{}{}
	}
	return exist
}

func (this *RandomizedSet) Remove(val int) bool {
	exist := false
	if _, ok := this.rand[val]; ok {
		exist = true
		delete(this.rand, val)
	}
	return exist

}

func (this *RandomizedSet) GetRandom() int {
	lenRand := len(this.rand)
	randNum := rand.Intn(lenRand)
	i := 0
	for k := range this.rand {
		if i == randNum {
			return k
		}
		i++
	}
	return -1
}

type RandomizedSetVer2 struct {
	randGen *rand.Rand
	nums    []int
	indices map[int]int
}

func ConstructorVer2() RandomizedSetVer2 {
	return RandomizedSetVer2{
		randGen: rand.New(rand.NewSource(time.Now().UnixNano())),
		nums:    []int{},
		indices: make(map[int]int),
	}
}

func (this *RandomizedSetVer2) InsertVer2(val int) bool {
	if _, exists := this.indices[val]; exists {
		return false
	}
	this.indices[val] = len(this.nums)
	this.nums = append(this.nums, val)
	return true
}

func (this *RandomizedSetVer2) RemoveVer2(val int) bool {
	index, exists := this.indices[val]
	if !exists {
		return false
	}
	lastElement := this.nums[len(this.nums)-1]
	this.nums[index] = lastElement
	this.indices[lastElement] = index
	this.nums = this.nums[:len(this.nums)-1]
	delete(this.indices, val)
	return true
}

func (this *RandomizedSetVer2) GetRandomVer2() int {
	return this.nums[this.randGen.Intn(len(this.nums))]
}

type RandomizedSetVer3 struct {
	randGen *rand.Rand
	nums    []int
	indices map[int]int
	length  int
}

func ConstructorVer3() RandomizedSetVer3 {
	return RandomizedSetVer3{
		randGen: rand.New(rand.NewSource(time.Now().UnixNano())),
		nums:    []int{},
		indices: make(map[int]int),
		length:  0,
	}
}

func (this *RandomizedSetVer3) InsertVer3(val int) bool {
	if _, exists := this.indices[val]; exists {
		return false
	}
	this.indices[val] = this.length
	this.nums = append(this.nums, val)
	this.length++
	return true
}

func (this *RandomizedSetVer3) RemoveVer3(val int) bool {
	index, exists := this.indices[val]
	if !exists {
		return false
	}
	lastElement := this.nums[this.length-1]
	this.nums[index] = lastElement
	this.indices[lastElement] = index
	this.nums = this.nums[:this.length-1]
	this.length--
	delete(this.indices, val)
	return true
}

func (this *RandomizedSetVer3) GetRandomVer3() int {
	return this.nums[this.randGen.Intn(this.length)]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
