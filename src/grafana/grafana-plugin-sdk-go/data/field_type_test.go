package data_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

// The slice data in the Field is a not exported, so methods on the Field are used to to manipulate its data.
type simpleFieldInfo struct {
	Name      string
	FieldType FieldType
}

func TestFieldTypeConversion(t *testing.T) {
	f := FieldTypeBool
	s := f.ItemTypeString()
	require.Equal(t, "bool", s)
	c, ok := FieldTypeFromItemTypeString(s)
	require.True(t, ok, "must parse ok")
	require.Equal(t, f, c)

	_, ok = FieldTypeFromItemTypeString("????")
	require.False(t, ok, "unknown type")

	c, ok = FieldTypeFromItemTypeString("float")
	require.True(t, ok, "must parse ok")
	require.Equal(t, FieldTypeFloat64, c)

	obj := &simpleFieldInfo{
		Name:      "hello",
		FieldType: FieldTypeFloat64,
	}
	body, err := json.Marshal(obj)
	require.NoError(t, err)

	objCopy := &simpleFieldInfo{}
	err = json.Unmarshal(body, &objCopy)
	require.NoError(t, err)

	require.Equal(t, obj.FieldType, objCopy.FieldType)
}
