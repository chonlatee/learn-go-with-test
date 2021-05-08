package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {

	t.Run("with all type of field", func(t *testing.T) {
		cases := []struct {
			Name          string
			Input         interface{}
			ExpectedCalls []string
		}{
			{
				"Struct with one string field",
				struct {
					Name string
				}{"Chris"},
				[]string{"Chris"},
			},
			{
				"Struct with two string field",
				struct {
					Name string
					City string
				}{"Chris", "London"},
				[]string{"Chris", "London"},
			},
			{
				"Struct with non string field",
				struct {
					Name string
					Age  int
				}{"Chris", 33},
				[]string{"Chris"},
			},
			{
				"Nested fields",
				Person{
					"Chris",
					Profile{33, "London"},
				},
				[]string{"Chris", "London"},
			},
			{
				"Pointers to thing",
				&Person{
					"Chris",
					Profile{33, "London"},
				},
				[]string{"Chris", "London"},
			},
			{
				"Slices",
				[]Profile{
					{33, "London"},
					{34, "Reykjavik"},
				},
				[]string{"London", "Reykjavik"},
			},
			{
				"Arrays",
				[2]Profile{
					{33, "London"},
					{34, "Reykjavik"},
				},
				[]string{"London", "Reykjavik"},
			},
		}

		for _, test := range cases {
			t.Run(test.Name, func(t *testing.T) {
				var got []string
				walk(test.Input, func(input string) {
					got = append(got, input)
				})

				if !reflect.DeepEqual(got, test.ExpectedCalls) {
					t.Errorf("got %v, want %v", got, test.ExpectedCalls)
				}
			})
		}
	})

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string

		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContain(t, got, "Bar")
		assertContain(t, got, "Boz")
	})

}

func assertContain(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}

	if !contains {
		t.Errorf("Expected %+v to contain %q but it didn't ", haystack, needle)
	}
}

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}
