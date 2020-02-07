package problem037

type Set map[interface{}]bool

func NewSet(values ...interface{}) *Set {
	newSet := make(Set)
	for _, value := range values {
		newSet[value] = true
	}
	return &newSet
}

func (thisSet *Set) add(value interface{}) *Set {
	(*thisSet)[value] = true
	return thisSet
}

func (thisSet *Set) copy() *Set {
	cloneSet := make(Set)
	for value := range *thisSet {
		cloneSet.add(value)
	}
	return &cloneSet
}

func (thisSet *Set) has(value interface{}) bool {
	return (*thisSet)[value]
}

func (thisSet *Set) equal(other *Set) bool {
	if len(*thisSet) != len(*other) {
		return false
	}
	for value := range *other {
		if !thisSet.has(value) {
			return false
		}
	}
	return true
}

func (thisSet *Set) PowerSet() *SetOfSet {
	powerSet, tempPowerSet := NewSetOfSet(NewSet()), NewSetOfSet()
	for element := range *thisSet {
		for subset := range *powerSet {
			cloneSubset := subset.copy()
			cloneSubset.add(element)
			tempPowerSet.add(subset).add(cloneSubset)
		}
		powerSet, tempPowerSet = tempPowerSet, powerSet
	}
	return powerSet
}

type SetOfSet map[*Set]bool

func NewSetOfSet(values ...*Set) *SetOfSet {
	newSetOfSet := make(SetOfSet)
	for _, value := range values {
		newSetOfSet[value] = true
	}
	return &newSetOfSet
}

func (thisSet *SetOfSet) add(value *Set) *SetOfSet {
	(*thisSet)[value] = true
	return thisSet
}

func (thisSet *SetOfSet) has(value *Set) bool {
	for element := range *thisSet {
		if element.equal(value) {
			return true
		}
	}
	return false
}

func (thisSet *SetOfSet) equal(other *SetOfSet) bool {
	if len(*thisSet) != len(*other) {
		return false
	}
	for value := range *other {
		if !thisSet.has(value) {
			return false
		}
	}
	return true
}
