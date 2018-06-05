package jsontest

import (
	"fmt"
	"strconv"

	"github.com/gogo/protobuf/jsonpb"
)

// Uint64Alias is used to test custom type serialization.
type Uint64Alias uint64

// MarshalJSON renders uint as a single hex string. The value is returned enclosed in quotes.
func (t Uint64Alias) MarshalJSON() ([]byte, error) {
	println("Uint64Alias.MarshalJSON called")
	return []byte(fmt.Sprintf(`"%x"`, t)), nil
}

// UnmarshalJSON populates Uint64Alias from a hex string. Called by gogo/protobuf/jsonpb.
// There appears to be a bug in gogoproto, as this function is only called for numeric values.
// https://github.com/gogo/protobuf/issues/411#issuecomment-393856837
func (t *Uint64Alias) UnmarshalJSON(b []byte) error {
	println("Uint64Alias.UnmarshalJSON called")
	q, err := fromString(string(b))
	if err != nil {
		return err
	}
	*t = q
	return nil
}

// UnmarshalJSONPB populates Uint64Alias from a quoted hex string. Called by gogo/protobuf/jsonpb.
// The input value is a quoted string.
func (t *Uint64Alias) UnmarshalJSONPB_disabled(_ *jsonpb.Unmarshaler, b []byte) error {
	if len(b) < 3 {
		return fmt.Errorf("SpanID JSON string cannot be shorter than 3 chars: %s", string(b))
	}
	if b[0] != '"' || b[len(b)-1] != '"' {
		return fmt.Errorf("SpanID JSON string must be enclosed in quotes: %s", string(b))
	}
	return t.UnmarshalJSON(b[1 : len(b)-1])
}

func fromString(s string) (Uint64Alias, error) {
	id, err := strconv.ParseUint(s, 16, 64)
	if err != nil {
		return Uint64Alias(0), err
	}
	return Uint64Alias(id), nil
}

// MarshalJSONPB renders struct as a single hex string. The value is returned enclosed in quotes.
func (t AStruct) MarshalJSONPB(*jsonpb.Marshaler) ([]byte, error) {
	println("AStruct.MarshalJSON called")
	return []byte(`"` + t.String() + `"`), nil
}

// UnmarshalJSONPB is fake, always returns empty struct for now.
func (t *AStruct) UnmarshalJSONPB(_ *jsonpb.Unmarshaler, b []byte) error {
	println("AStruct.UnmarshalJSONPB called")
	*t = AStruct{}
	return nil
}

func (t AStruct) String() string {
	return fmt.Sprintf("%016x%016x", t.High, t.Low)
}
