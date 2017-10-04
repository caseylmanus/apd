package apd

import (
	"bytes"
	"encoding/json"
	"testing"
)

type jsonTestStruct struct {
	Value NullDecimal
}

func TestDecimalJsonUnmarshal(t *testing.T) {
	data := `{ "Value": 124.5}`
	var target jsonTestStruct
	bytes := []byte(data)
	err := json.Unmarshal(bytes, &target)
	if err != nil {
		t.Fatal("failed to unmarshal json")
	}
	if target.Value.Decimal.String() != "124.5" {
		t.Fatal("Value not expecteded")
	}
	if !target.Value.Valid {
		t.Fatal("Valid should be set to true")
	}
}
func TestDecimalJsonMarshal(t *testing.T) {
	expected := []byte(`{"Value":124.5}`)
	var target jsonTestStruct
	d, _, _ := NewFromString("124.5")
	target.Value.SetValid(*d)
	actual, err := json.Marshal(&target)
	if err != nil {
		t.Fatal("Failed to marshal json")
	}
	if !bytes.Equal(actual, expected) {
		t.Fatalf("Expected %v, got %v", string(expected), string(actual))
	}
}
