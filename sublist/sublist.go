package sublist

type Relation string

const Equal = Relation("equal")
const Unequal = Relation("unequal")
const SubList = Relation("sublist")
const SuperList = Relation("superlist")

func Sublist(listOne, listTwo []int) Relation {
	lenOne := len(listOne)
	lenTwo := len(listTwo)
	if lenOne < lenTwo {
		diffLength := lenTwo - lenOne
		for i := 0; i <= diffLength; i++ {
			if compareSlices(listOne, listTwo[i:]) {
				return SubList
			}
		}
		return Unequal
	} else if lenOne == lenTwo {
		if compareSlices(listOne, listTwo) {
			return Equal
		}
		return Unequal
	} else {
		diffLength := lenOne - lenTwo
		for i := 0; i <= diffLength; i++ {
			if compareSlices(listTwo, listOne[i:]) {
				return SuperList
			}
		}
		return Unequal
	}
}

func compareSlices(slice1, slice2 []int) bool {
	for index, elem := range slice1 {
		if elem != slice2[index] {
			return false
		}
	}
	return true
}