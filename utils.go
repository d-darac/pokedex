package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

func structToMap(strct interface{}, data map[string]interface{}, fieldNames []string) (map[string]interface{}, error) {
	if data == nil {
		data = make(map[string]interface{})
	}

	dataType := reflect.TypeOf(strct)
	if dataType.Kind() != reflect.Struct {
		return nil, fmt.Errorf("not a struct")
	}

	dataValue := reflect.ValueOf(strct)
	strctFields := reflect.VisibleFields(dataType)

	if fieldNames == nil {
		for _, strctFld := range strctFields {
			fieldNames = append(fieldNames, strctFld.Name)
		}
	}

	for _, fieldName := range fieldNames {
		nested := splitNested(fieldName)
		parentFieldName := strings.Title(nested[0])

		var index *int
		fld, i, isIndexed := splitIndexed(nested[0])
		if isIndexed {
			parentFieldName = strings.Title(fld)
			index = i
		}

		strctFld, exists := dataValue.Type().FieldByName(parentFieldName)

		if !exists {
			return nil, fmt.Errorf("no such field: %s", parentFieldName)
		}

		if strctFld.Type.Kind() == reflect.Struct {
			if strctFld.Anonymous {
				_, err := structToMap(dataValue.FieldByName(parentFieldName).Interface(), data, nil)
				if err != nil {
					return nil, err
				}
				continue
			}

			if _, ok := data[parentFieldName]; !ok {
				data[parentFieldName] = make(map[string]interface{})
			}

			var flds []string
			if len(nested) > 1 {
				flds = nested[1:]
			}

			_, err := structToMap(dataValue.FieldByName(parentFieldName).Interface(), data[parentFieldName].(map[string]interface{}), flds)
			if err != nil {
				return nil, err
			}

			continue
		}

		if strctFld.Type.Kind() == reflect.Slice {
			sliceValue := dataValue.FieldByName(parentFieldName)
			sliceLen := sliceValue.Len()
			if _, ok := data[parentFieldName]; !ok {
				data[parentFieldName] = make([]map[string]interface{}, sliceLen)
			}

			if sliceLen > 0 {
				elType := sliceValue.Index(0).Type()
				if elType.Kind() == reflect.Struct {
					if index != nil {
						if *index >= sliceLen {
							return nil, fmt.Errorf("'%s' index '%d' out of range", parentFieldName, *index)
						}

						element := sliceValue.Index(*index)

						var flds []string
						if len(nested) > 1 {
							flds = nested[1:]
						}

						m, err := structToMap(element.Interface(), data[parentFieldName].([]map[string]interface{})[*index], flds)
						if err != nil {
							return nil, err
						}

						if len(nested) > 1 {
							for k, v := range data[parentFieldName].([]map[string]interface{})[*index] {
								if data[parentFieldName].([]map[string]interface{})[*index] == nil {
									data[parentFieldName].([]map[string]interface{})[*index] = make(map[string]interface{})
								}
								data[parentFieldName].([]map[string]interface{})[*index][k] = v
							}
						}
						data[parentFieldName].([]map[string]interface{})[*index] = m
						continue
					}

					for i := range sliceLen {
						element := sliceValue.Index(i)

						var flds []string
						if len(nested) > 1 {
							flds = nested[1:]
						}

						m, err := structToMap(element.Interface(), data[parentFieldName].([]map[string]interface{})[i], flds)
						if err != nil {
							return nil, err
						}

						if len(nested) > 1 {
							for k, v := range data[parentFieldName].([]map[string]interface{})[i] {
								if data[parentFieldName].([]map[string]interface{})[i] == nil {
									data[parentFieldName].([]map[string]interface{})[i] = make(map[string]interface{})
								}
								data[parentFieldName].([]map[string]interface{})[i][k] = v
							}
						}
						data[parentFieldName].([]map[string]interface{})[i] = m
					}
				} else {
					if index != nil {
						if *index >= sliceLen {
							return nil, fmt.Errorf("'%s' index '%d' out of range", parentFieldName, *index)
						}

						switch elType.Kind() {
						case reflect.String:
							data[parentFieldName] = make([]string, sliceLen)
							data[parentFieldName].([]string)[*index] = dataValue.FieldByName(parentFieldName).Index(*index).String()
						case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
							data[parentFieldName] = make([]int64, sliceLen)
							data[parentFieldName].([]int64)[*index] = dataValue.FieldByName(parentFieldName).Index(*index).Int()
						case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
							data[parentFieldName] = make([]uint64, sliceLen)
							data[parentFieldName].([]uint64)[*index] = dataValue.FieldByName(parentFieldName).Index(*index).Uint()
						case reflect.Float32, reflect.Float64:
							data[parentFieldName] = make([]float64, sliceLen)
							data[parentFieldName].([]float64)[*index] = dataValue.FieldByName(parentFieldName).Index(*index).Float()
						case reflect.Bool:
							data[parentFieldName] = make([]bool, sliceLen)
							data[parentFieldName].([]bool)[*index] = dataValue.FieldByName(parentFieldName).Index(*index).Bool()
						case reflect.Pointer:
							data[parentFieldName] = make([]interface{}, 0)
						default:
							return nil, fmt.Errorf("type '%s' not handled", elType.Kind())
						}
					}
					data[parentFieldName] = dataValue.FieldByName(parentFieldName).Interface()
				}
				continue
			}
			data[parentFieldName] = []interface{}{}
		}

		data[parentFieldName] = dataValue.FieldByName(parentFieldName).Interface()
	}

	return data, nil
}

func splitNested(nestedFieldName string) []string {
	before, after, found := strings.Cut(nestedFieldName, ".")

	if !found {
		return []string{before}
	}

	return []string{before, after}
}

func splitIndexed(fieldWithIndex string) (string, *int, bool) {
	re := regexp.MustCompile(`\[(\d)\]`)
	match := re.FindStringSubmatch(fieldWithIndex)
	if len(match) == 0 {
		return strings.SplitN(fieldWithIndex, "[", 2)[0], nil, false
	}
	index, err := strconv.Atoi(match[1])
	if err != nil {
		return "", nil, false
	}
	return strings.SplitN(fieldWithIndex, "[", 2)[0], &index, true
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
