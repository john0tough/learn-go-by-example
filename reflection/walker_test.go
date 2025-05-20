package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			Name: "struct with one string field",
			Input: struct {
				Name string
			}{
				Name: "Chris",
			},
			ExpectedCalls: []string{"Chris"},
		},
		{
			Name: "struct with two string field",
			Input: struct {
				Name string
				City string
			}{
				Name: "Chris",
				City: "London",
			},
			ExpectedCalls: []string{"Chris", "London"},
		},
		{
			Name: "struct with non string field",
			Input: struct {
				Name string
				Age  int
			}{
				Name: "Chris",
				Age:  35,
			},
			ExpectedCalls: []string{"Chris"},
		},
		{
			Name: "nested fields",
			Input: Person{
				Name: "Chris",
				Profile: Profile{
					Age:  35,
					City: "London",
				},
			},
			ExpectedCalls: []string{"Chris", "London"},
		},
		{
			Name: "pointer to things",
			Input: &Person{
				Name: "Chris",
				Profile: Profile{
					Age:  35,
					City: "London",
				},
			},
			ExpectedCalls: []string{"Chris", "London"},
		},
		{
			Name: "slices",
			Input: []Profile{
				{
					Age:  35,
					City: "London",
				},
				{
					Age:  45,
					City: "Halacho",
				},
			},
			ExpectedCalls: []string{"London", "Halacho"},
		},
		{
			"arrays",
			[2]Profile{
				{33, "London"},
				{34, "Reykjavík"},
			},
			[]string{"London", "Reykjavík"},
		},
	}

	for _, tt := range cases {
		input := tt.Input
		expected := tt.ExpectedCalls
		var got []string
		Walk(input, func(input string) {
			got = append(got, input)
		})

		if len(got) == 0 {
			t.Errorf("wrong number of function calls, got %d want %d", len(got), 1)
		}
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got %v, want %v", got, expected)
		}
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baa",
		}

		var got []string
		Walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Moo")
		assertContains(t, got, "Baa")
	})
	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)
		go func() {
			defer close(aChannel)
			aChannel <- Profile{33, "Berlin"}
			aChannel <- Profile{39, "Valladolid"}
		}()

		var got []string
		want := []string{"Berlin", "Valladolid"}

		Walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("with functions", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{33, "Berlin"}, Profile{34, "Katowice"}
		}

		var got []string
		want := []string{"Berlin", "Katowice"}

		Walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %v contains %q, but it did not", haystack, needle)
	}
}
