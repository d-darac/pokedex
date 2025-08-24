package main

import (
	"bytes"
	"io"
	"os"
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
		expected string
	}{
		{
			input: struct {
				pokemonStruct pokeapi.PokeAPIPokemon
				fields        []string
			}{
				pokemonStruct: pokemon,
				fields:        []string{"Height", "Weight"},
			},
			expected: `Height: 6
Weight: 10
`,
		},
		{
			input: struct {
				pokemonStruct pokeapi.PokeAPIPokemon
				fields        []string
			}{
				pokemonStruct: pokemon,
				fields:        []string{"Height", "Weight", "Species"},
			},
			expected: `Height: 6
Weight: 10
Species:
  Name: koffing
  URL: https://pokeapi.co/api/v2/pokemon-species/109/
`,
		},
		{
			input: struct {
				pokemonStruct pokeapi.PokeAPIPokemon
				fields        []string
			}{
				pokemonStruct: pokemon,
				fields:        []string{"Height", "Weight", "Abilities"},
			},
			expected: `Height: 6
Weight: 10
Abilities:
  IsHidden: false
  Slot: 1
  Ability:
    Name: levitate
    URL: https://pokeapi.co/api/v2/ability/26/
`,
		},
		{
			input: struct {
				pokemonStruct pokeapi.PokeAPIPokemon
				fields        []string
			}{
				pokemonStruct: pokemon,
				fields:        []string{"ID", "Name", "Height", "Weight", "Species", "Abilities"},
			},
			expected: `ID: 109
Name: koffing
Height: 6
Weight: 10
Species:
  Name: koffing
  URL: https://pokeapi.co/api/v2/pokemon-species/109/
Abilities:
  IsHidden: false
  Slot: 1
  Ability:
    Name: levitate
    URL: https://pokeapi.co/api/v2/ability/26/
`,
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
		//displayPokemonDetails(c.input.pokemonMap, c.input.fields)

		actual := captureOutput(
			t,
			func() {
				displayPokemonDetails(c.input.pokemonStruct, c.input.fields)
			})

		if c.expected != actual {
			t.Errorf("output not equal to expected value:\n%v\n!=\n%v\n", c.expected, actual)
		}
	}

	for _, c := range errCases {
		if err := displayPokemonDetails(c.input.pokemonStruct, c.input.fields); err == nil {
			t.Errorf("expected error")
		}
	}
}

func captureOutput(t *testing.T, f func()) string {
	// Save the original stdout
	originalStdout := os.Stdout

	// Create a new buffer and redirect stdout
	var bf bytes.Buffer
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("Failed to create pipe: %v", err)
	}
	os.Stdout = w

	// Call the function that's being tested
	f()

	// Stop writing and restore stdout
	w.Close()
	os.Stdout = originalStdout
	io.Copy(&bf, r)

	return bf.String()
}
