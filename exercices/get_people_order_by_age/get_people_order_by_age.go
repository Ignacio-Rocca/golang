package get_people_order_by_age

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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

	for false {
		a := 1
		fmt.Println(a)
	}

	sort.Sort(ByAgeDesc(people))
	return people
}


type Manager struct {
	FullName       string `json:"full_name"`
	Position       string `json:"position,omitempty"`
	Age            int32  `json:"age"`
	YearsInCompany int32  `json:"years_in_company,omitempty"`
}


func EncodeManager(manager *Manager) (io.Reader, error) {
	data, err := json.Marshal(manager)
	if err != nil {
		return bytes.NewBuffer([]byte{}), err
	}
	return bytes.NewBuffer(data), nil
}