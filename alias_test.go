package aliastest

import (
	"bytes"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yurishkuro/gogoproto-custom-type/jsontest"
)

func TestMarshalSimple(t *testing.T) {
	tests := []struct {
		obj *jsontest.Simple
		out string
	}{
		{&jsontest.Simple{Val: 0x42}, `{"first":"00000000000000000000000000000000","val":"42"}`},
		{&jsontest.Simple{Val: 0x42f}, `{"first":"00000000000000000000000000000000","val":"42f"}`},
	}
	for _, test := range tests {
		t.Run(test.out, func(t *testing.T) {
			out := new(bytes.Buffer)
			require.NoError(t, new(jsonpb.Marshaler).Marshal(out, test.obj))
			assert.Equal(t, test.out, out.String())
			var val jsontest.Simple
			require.NoError(t, jsonpb.Unmarshal(bytes.NewReader([]byte(test.out)), &val))
			assert.Equal(t, *test.obj, val)
		})
	}
}

func TestMarshalSimple2(t *testing.T) {
	tests := []struct {
		obj *jsontest.Simple2
		out string
	}{
		{&jsontest.Simple2{Val: 0x42}, `{"val":"42"}`},
		{&jsontest.Simple2{Val: 0x42f}, `{"val":"42f"}`},
	}
	for _, test := range tests {
		t.Run(test.out, func(t *testing.T) {
			out := new(bytes.Buffer)
			require.NoError(t, new(jsonpb.Marshaler).Marshal(out, test.obj))
			assert.Equal(t, test.out, out.String())
			var val jsontest.Simple2
			require.NoError(t, jsonpb.Unmarshal(bytes.NewReader([]byte(test.out)), &val))
			assert.Equal(t, *test.obj, val)
		})
	}
}
