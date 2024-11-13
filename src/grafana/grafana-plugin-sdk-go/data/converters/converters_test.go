package converters_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStringConversions(t *testing.T) {
	val, err := AnyToNullableString.Converter(12.3)
	require.NoError(t, err)
	require.Equal(t, "12.3", *(val.(*string)))

	ptr := &val
	val, err = AnyToNullableString.Converter(ptr)
	require.NoError(t, err)
	require.Equal(t, fmt.Sprintf("%p", ptr), *(val.(*string))) // pointer printed as a pointer?

	val, err = AnyToNullableString.Converter(nil)
	require.NoError(t, err)
	require.Nil(t, val)
}

func TestNumericConversions(t *testing.T) {
	val, err := Float64ToNullableFloat64.Converter(12.34)
	require.NoError(t, err)
	require.Equal(t, 12.34, *(val.(*float64)))
}

func TestJSONConversions(t *testing.T) {
	val, err := JSONValueToFloat64.Converter(12.34)
	require.NoError(t, err)
	require.Equal(t, 12.34, val)

	val, err = JSONValueToFloat64.Converter(12)
	require.NoError(t, err)
	require.Equal(t, float64(12), val)

	val, err = JSONValueToFloat64.Converter(int64(12))
	require.NoError(t, err)
	require.Equal(t, float64(12), val)

	val, err = JSONValueToFloat64.Converter("12.34")
	require.NoError(t, err)
	require.Equal(t, 12.34, val)
}
