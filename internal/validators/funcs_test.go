package validators

import (
	"fmt"
	"testing"
)

func TestValidatePositiveInt(t *testing.T) {
	testTable := []struct {
		Name    string
		Val     string
		IsValid bool
	}{
		{
			Name:    `negative int`,
			Val:     `-1000`,
			IsValid: false,
		},
		{
			Name:    `float num`,
			Val:     `100.23`,
			IsValid: false,
		},
		{
			Name:    `big num`,
			Val:     `10000000000000000000000000000000000`,
			IsValid: false,
		},
		{
			Name:    `string`,
			Val:     `somestring`,
			IsValid: false,
		},
		{
			Name:    `num mixed with string`,
			Val:     `123somestring`,
			IsValid: false,
		},
		{
			Name:    `empty string`,
			Val:     ``,
			IsValid: false,
		},
		{
			Name:    `correct int`,
			Val:     `12300`,
			IsValid: true,
		},
	}

	for i, test := range testTable {
		t.Run(fmt.Sprintf("%s_%d", t.Name(), i), func(t *testing.T) {
			t.Log("testing:", test.Name)
			res := ValidatePositiveInt(test.Val)
			if res != test.IsValid {
				t.Logf("should be - %v, got - %v", test.IsValid, res)
				t.FailNow()
			}
		})
	}
}
