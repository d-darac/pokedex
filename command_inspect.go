package main

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/d-darac/pokedex/internal/pokeapi"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("missing argument pokemon_name\nusage:\ncatch <pokemon_name>")
	}

	pokemonName := args[0]

	if pokemon, exists := cfg.caughtPokemon[pokemonName]; !exists {
		return errors.New("you have not caught that pokemon")
	} else {
		if err := displayPokemonDetails(pokemon, []string{"Name", "Types"}); err != nil {
			return err
		}
	}

	return nil
}

func displayPokemonDetails(pokemon pokeapi.PokeAPIPokemon, fields []string) error {
	lines, err := constructLines(pokemon, fields, nil)
	if err != nil {
		return err
	}

	for _, line := range lines {
		fmt.Println(line)
	}

	return nil
}

func constructLines(collection interface{}, fieldsNames []string, nestLevel *int) ([]string, error) {
	dataType := reflect.TypeOf(collection)
	dataValue := reflect.ValueOf(collection)
	finalLines := []string{}

	lastParentFieldName := ""

	if nestLevel == nil {
		nestLevel = new(int)
		*nestLevel = 0
	}

	if dataType.Kind() != reflect.Struct {
		return nil, fmt.Errorf("provaded value is not a struct")
	}

	for _, fieldName := range fieldsNames {
		nested := splitNested(fieldName)
		parentFieldName := nested[0]

		structField, exists := dataType.FieldByName(parentFieldName)

		if !exists {
			return nil, fmt.Errorf("no such field: %s", fieldName)
		}

		fieldValue := dataValue.FieldByName(parentFieldName)

		if fieldValue.Kind() != reflect.Struct && fieldValue.Kind() != reflect.Slice {
			line, err := lineForPrimitives(fieldValue, fieldName)
			if err != nil {
				return nil, err
			}

			withPadding := fmt.Sprintf("%s%s", getPadding(*nestLevel, 2), line)
			finalLines = append(finalLines, withPadding)
		}

		if structField.Type.Kind() == reflect.Struct {
			if parentFieldName != lastParentFieldName && !structField.Anonymous {
				line := fmt.Sprintf("%s:", parentFieldName)
				withPadding := fmt.Sprintf("%s%s", getPadding(*nestLevel, 2), line)
				finalLines = append(finalLines, withPadding)
				*nestLevel++
			}

			if len(nested) > 1 {
				lines, err := constructLines(fieldValue.Interface(), nested[1:], nestLevel)
				if err != nil {
					return nil, err
				}

				finalLines = append(finalLines, lines...)
				// selects all child fields of a nested struct in case the `fields` slice contains
				// just the name of the nested struct field, and not any specific children
			} else {
				flds := []string{}
				numFields := structField.Type.NumField()
				for i := range numFields {
					flds = append(flds, structField.Type.Field(i).Name)
				}

				msgs, err := constructLines(fieldValue.Interface(), flds, nestLevel)
				if err != nil {
					return nil, err
				}

				finalLines = append(finalLines, msgs...)
			}

			if !structField.Anonymous {
				*nestLevel--
				lastParentFieldName = parentFieldName
			}
		}

		if structField.Type.Kind() == reflect.Slice {
			line := fmt.Sprintf("%s:", parentFieldName)
			withPadding := fmt.Sprintf("%s%s", getPadding(*nestLevel, 2), line)
			finalLines = append(finalLines, withPadding)

			lines, err := linesForSlices(nestLevel, fieldValue, fieldName, nil)
			if err != nil {
				return nil, err
			}

			finalLines = append(finalLines, lines...)
		}
	}

	return finalLines, nil
}

func lineForPrimitives(fieldValue reflect.Value, fieldName string) (string, error) {
	line := ""
	switch fieldValue.Kind() {
	case reflect.String:
		line = fmt.Sprintf("%s: %v", fieldName, fieldValue.String())
	case reflect.Bool:
		line = fmt.Sprintf("%s: %v", fieldName, fieldValue.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		line = fmt.Sprintf("%s: %v", fieldName, fieldValue.Int())
	case reflect.Float32, reflect.Float64:
		line = fmt.Sprintf("%s: %v", fieldName, fieldValue.Float())
	case reflect.Pointer:
		line = fmt.Sprintf("%s: null", fieldName)
	default:
		return line, fmt.Errorf("type: %s not covered", fieldValue.Kind().String())
	}
	return line, nil
}

func linesForSlices(nestLevel *int, fieldValue reflect.Value, fieldName string, nested *[]string) ([]string, error) {
	*nestLevel++
	lines := []string{}

	for i := range fieldValue.Len() {
		elemValue := fieldValue.Index(i)

		if elemValue.Kind() != reflect.Struct && elemValue.Kind() != reflect.Slice {
			line, err := lineForPrimitives(elemValue, fieldName)
			if err != nil {
				return nil, err
			}

			withPadding := fmt.Sprintf("%s%s", getPadding(*nestLevel, 2), line)
			lines = append(lines, withPadding)

		}

		if elemValue.Kind() == reflect.Struct {
			fields := []string{}
			numFields := elemValue.NumField()
			for i := range numFields {
				fields = append(fields, elemValue.Type().Field(i).Name)
			}

			lns, err := constructLines(elemValue.Interface(), fields, nestLevel)
			if err != nil {
				return nil, err
			}

			lines = append(lines, lns...)
		}
	}

	*nestLevel--
	return lines, nil
}

func splitNested(nestedFieldName string) []string {
	before, after, found := strings.Cut(nestedFieldName, ".")

	if !found {
		return []string{before}
	}

	return []string{before, after}
}

func getPadding(n, indent int) string {
	ind := ""
	padding := ""

	for range indent {
		ind += " "
	}

	for range n {
		padding += ind
	}

	return padding
}
