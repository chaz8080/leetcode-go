package insertdeletegetrandomo1

import "math/rand"

type RandomizedSet struct {
	valMap  map[int]int
	keyList []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{valMap: map[int]int{}, keyList: make([]int, 0, 500)}
}

func (rs *RandomizedSet) Insert(val int) bool {
	if _, ok := rs.valMap[val]; ok {
		return false
	}
	rs.valMap[val] = len(rs.keyList)
	rs.keyList = append(rs.keyList, val)
	return true
}

func (rs *RandomizedSet) Remove(val int) bool {
	if v, ok := rs.valMap[val]; ok {
		rs.keyList[v] = rs.keyList[len(rs.keyList)-1]
		rs.valMap[rs.keyList[len(rs.keyList)-1]] = v
		rs.keyList = rs.keyList[:len(rs.keyList)-1]
		delete(rs.valMap, val)
		return true
	}
	return false
}

func (rs *RandomizedSet) GetRandom() int {
	return rs.keyList[rand.Intn(len(rs.keyList))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
