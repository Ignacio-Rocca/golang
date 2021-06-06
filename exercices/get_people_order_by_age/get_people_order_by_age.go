package get_people_order_by_age

import (
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type ByAgeAsc []Person

func (b ByAgeAsc) Len() int { return len(b) }

func (b ByAgeAsc) Less(i, j int) bool { return b[i].Age < b[j].Age }

func (b ByAgeAsc) Swap(i, j int) { b[i], b[j] = b[j], b[i] }

type ByAgeDesc []Person

func (b ByAgeDesc) Len() int { return len(b) }

func (b ByAgeDesc) Less(i, j int) bool { return b[i].Age > b[j].Age }

func (b ByAgeDesc) Swap(i, j int) { b[i], b[j] = b[j], b[i] }


type OrderBy int
const(
	Asc OrderBy = iota
	Desc
)


func orderPeopleByAge(people []Person, ord OrderBy) []Person {
	if ord == Asc {
		sort.Sort(ByAgeAsc(people))
		return people
	}
	sort.Sort(ByAgeDesc(people))
	return people
}
