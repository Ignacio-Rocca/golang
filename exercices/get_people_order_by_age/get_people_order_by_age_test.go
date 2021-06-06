package get_people_order_by_age

import (
	"reflect"
	"testing"
)

func Test_orderPeopleByAge(t *testing.T) {
	type args struct {
		people []Person
		ord    OrderBy
	}
	tests := []struct {
		name string
		args args
		want []Person
	}{

		{
			name: "success desc",
			args: args{
				people: []Person{
					{
						Name: "Lia Mur",
						Age: 26,
					},
					{
						Name: "Ignacio Rocca",
						Age: 27,
					},
					{
						Name: "Martina Rocca",
						Age: 21,
					},
				},
				ord: Desc,
			},
			want: []Person{
				{
					Name: "Ignacio Rocca",
					Age: 27,
				},
				{
					Name: "Lia Mur",
					Age: 26,
				},
				{
					Name: "Martina Rocca",
					Age: 21,
				},

			},
		},
		{
			name: "success asc",
			args: args{
				people: []Person{
					{
						Name: "Lia Mur",
						Age: 26,
					},
					{
						Name: "Ignacio Rocca",
						Age: 27,
					},
					{
						Name: "Martina Rocca",
						Age: 21,
					},
				},
				ord: Asc,
			},
			want: []Person{
				{
					Name: "Martina Rocca",
					Age: 21,
				},
				{
					Name: "Lia Mur",
					Age: 26,
				},
				{
					Name: "Ignacio Rocca",
					Age: 27,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := orderPeopleByAge(tt.args.people, tt.args.ord); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("orderPeopleByAge() = %v, want %v", got, tt.want)
			}
		})
	}
}
