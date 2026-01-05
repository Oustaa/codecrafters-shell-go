package utils

import "testing"

func TestAdd(t *testing.T) {
	testCases := [][2]string{
		{"\"hello 'there'\"", "hello 'there'"},
		{"\"hello i am oussma'", "\"hello i am oussma'"},
		{"\"He is''''' oussama\"", "He is''''' oussama"},
		{"He is          oussama", "He is oussama"},
		{"'hello     example' 'shell''script' world''test", "hello     example shellscript worldtest"},
	}

	for _, testCase := range testCases {
		result := FormatMessage(testCase[0])
		expected := testCase[1]
		if result != expected {
			t.Errorf("FormatMessage(%s) = %s; expected %s", testCase[0], result, expected)
		}

	}
}
