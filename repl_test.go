package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/d-darac/pokedex/internal/pokeapi"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  ",
			expected: []string{},
		},
		{
			input:    "  hello  ",
			expected: []string{"hello"},
		},
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  HellO  World  ",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("lengths don't match: '%v' vs '%v'", actual, c.expected)
			continue
		}

		for i := range actual {
			actualWord := actual[i]
			expectedWord := c.expected[i]

			if actualWord != expectedWord {
				t.Errorf("cleanInput(%v) == %v, expected %v", c.input, actualWord, c.expected)
			}
		}
	}
}

func TestStructToMap(t *testing.T) {
	type TestStruct struct {
		String   string
		Int      int
		Float    float64
		Bool     bool
		Map      map[string]interface{}
		Pointer  *int
		Struct   struct{ Foo string }
		Strings  []string
		Ints     []int
		Floats   []float64
		Bools    []bool
		Pointers []*int
		Structs  []struct{ Foo string }
		Slices   [][]string
	}

	var pointer *int
	cases := []struct {
		input    struct{ testStruct TestStruct }
		expected map[string]interface{}
	}{
		{
			input: struct{ testStruct TestStruct }{testStruct: TestStruct{
				String:   "a",
				Int:      1,
				Float:    1.0,
				Bool:     false,
				Map:      map[string]interface{}{"foo": "bar"},
				Pointer:  pointer,
				Struct:   struct{ Foo string }{Foo: "bar"},
				Strings:  []string{"foo", "bar"},
				Ints:     []int{1, 2},
				Floats:   []float64{1.0, 2.0},
				Bools:    []bool{true, false},
				Pointers: []*int{nil, nil},
				Structs:  []struct{ Foo string }{{Foo: "bar"}, {Foo: "bar"}},
				Slices:   [][]string{{"foo", "bar"}, {"foo", "bar"}},
			}},
			expected: map[string]interface{}{
				"String":   "a",
				"Int":      1,
				"Float":    1.0,
				"Bool":     false,
				"Map":      map[string]interface{}{"foo": "bar"},
				"Pointer":  pointer,
				"Struct":   map[string]interface{}{"Foo": "bar"},
				"Strings":  []string{"foo", "bar"},
				"Ints":     []int{1, 2},
				"Floats":   []float64{1.0, 2.0},
				"Bools":    []bool{true, false},
				"Pointers": []*int{nil, nil},
				"Structs":  []map[string]interface{}{{"Foo": "bar"}, {"Foo": "bar"}},
				"Slices": [][]string{
					{"foo", "bar"},
					{"foo", "bar"},
				},
			},
		},
	}

	for _, c := range cases {
		actual, err := structToMap(c.input.testStruct, nil, nil)
		if err != nil {
			t.Errorf("returned an error: %v", err)
		}

		if fmt.Sprint(actual) != fmt.Sprint(c.expected) {
			t.Errorf("output does not match expected value:\nactual: %#v\nexpected: %#v", actual, c.expected)
		}
	}
}

func TestDisplayPokemonDetails(t *testing.T) {
	pokemon := pokeapi.PokeAPIPokemon{
		PokeAPIBaseResource: pokeapi.PokeAPIBaseResource{
			ID:   109,
			Name: "koffing",
		},
		Height: 6,
		Weight: 10,
		Species: struct {
			Name string "json:\"name\""
			URL  string "json:\"url\""
		}{
			Name: "koffing",
			URL:  "https://pokeapi.co/api/v2/pokemon-species/109/",
		},
		Abilities: []struct {
			IsHidden bool "json:\"is_hidden\""
			Slot     int  "json:\"slot\""
			Ability  struct {
				Name string "json:\"name\""
				URL  string "json:\"url\""
			} "json:\"ability\""
		}{
			{
				IsHidden: false,
				Slot:     1,
				Ability: struct {
					Name string "json:\"name\""
					URL  string "json:\"url\""
				}{
					Name: "levitate",
					URL:  "https://pokeapi.co/api/v2/ability/26/",
				},
			},
		},
	}

	cases := []struct {
		input struct {
			pokemonStruct pokeapi.PokeAPIPokemon
			fields        []string
		}
		expected []string
	}{
		{
			input: struct {
				pokemonStruct pokeapi.PokeAPIPokemon
				fields        []string
			}{
				pokemonStruct: pokemon,
				fields:        []string{"Height", "Weight"},
			},
			expected: []string{"Height: 6", "Weight: 10"},
		},
		{
			input: struct {
				pokemonStruct pokeapi.PokeAPIPokemon
				fields        []string
			}{
				pokemonStruct: pokemon,
				fields:        []string{"Height", "Weight", "Species"},
			},
			expected: []string{"Height: 6", "Weight: 10", "Species:", "Name: koffing", "URL: https://pokeapi.co/api/v2/pokemon-species/109/"},
		},
		{
			input: struct {
				pokemonStruct pokeapi.PokeAPIPokemon
				fields        []string
			}{
				pokemonStruct: pokemon,
				fields:        []string{"Height", "Weight", "Abilities"},
			},
			expected: []string{"Height: 6",
				"Weight: 10",
				"Abilities:",
				"IsHidden: false",
				"Slot: 1",
				"Ability:",
				"Name: levitate", "URL: https://pokeapi.co/api/v2/ability/26/"},
		},
		{
			input: struct {
				pokemonStruct pokeapi.PokeAPIPokemon
				fields        []string
			}{
				pokemonStruct: pokemon,
				fields:        []string{"ID", "Name", "Height", "Weight", "Species", "Abilities"},
			},
			expected: []string{"ID: 109",
				"Name: koffing",
				"Height: 6",
				"Weight: 10",
				"Species:",
				"Name: koffing",
				"URL: https://pokeapi.co/api/v2/pokemon-species/109/",
				"Abilities:",
				"IsHidden: false",
				"Slot: 1",
				"Ability:",
				"Name: levitate",
				"URL: https://pokeapi.co/api/v2/ability/26/"},
		},
	}

	errCases := []struct {
		input struct {
			pokemonStruct pokeapi.PokeAPIPokemon
			fields        []string
		}
	}{
		{
			input: struct {
				pokemonStruct pokeapi.PokeAPIPokemon
				fields        []string
			}{
				pokemonStruct: pokemon,
				fields:        []string{"FakeField"},
			},
		},
	}

	for _, c := range cases {
		actual := captureOutput(
			t,
			func() {
				displayRawPokemonDetails(c.input.pokemonStruct, c.input.fields)
			})

		for _, fld := range c.expected {
			if !strings.Contains(actual, fld) {
				t.Errorf("output not equal to expected value:\n%v\n!=\n%v\n", c.expected, actual)
			}
		}
	}

	for _, c := range errCases {
		if err := displayRawPokemonDetails(c.input.pokemonStruct, c.input.fields); err == nil {
			t.Errorf("expected error")
		}
	}
}
