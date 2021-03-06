package backend

import (
	"testing"

	"github.com/go-fish/gojson/errors"
	"github.com/stretchr/testify/assert"
)

func TestDecodeStringBasic(t *testing.T) {
	decoder := NewDecoder()
	decoder.SetData([]byte(`"TestString"`))
	defer decoder.Release()

	v, err := decoder.DecodeString()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, "TestString", v, "v must be equal to the value expected")
}

func TestDecoderStringComplex(t *testing.T) {
	decoder := NewDecoder()
	decoder.SetData([]byte(`  "string with spaces and \"escape\"d \"quotes\" and escaped line returns \\n and escaped \\\\ escaped char"`))
	defer decoder.Release()

	v, err := decoder.DecodeString()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, "string with spaces and \"escape\"d \"quotes\" and escaped line returns \\n and escaped \\\\ escaped char", v, "v is not equal to the value expected")
}

func TestDecoderStringNull(t *testing.T) {
	decoder := NewDecoder()
	decoder.SetData([]byte(`  null`))
	defer decoder.Release()

	v, err := decoder.DecodeString()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, "", v, "v must be equal to ''")
}

func TestDecoderStringQuotaNull(t *testing.T) {
	decoder := NewDecoder()
	decoder.SetData([]byte("  \n\"null\""))
	defer decoder.Release()

	v, err := decoder.DecodeString()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, "null", v, "v must be equal to 'null'")
}

func TestDecoderStringInvalidJSON(t *testing.T) {
	decoder := NewDecoder()
	decoder.SetData([]byte(`  "invalid JSONs`))
	defer decoder.Release()

	_, err := decoder.DecodeString()
	assert.NotNil(t, err, "Err must not be nil")
	assert.Equal(t, errors.NewParseError(decoder.data[decoder.length-1], decoder.length-1), err, "err message must equal to the value expected")
}

func TestDecoderStringInvalidType(t *testing.T) {
	decoder := NewDecoder()
	decoder.SetData([]byte(`  123333`))
	defer decoder.Release()

	_, err := decoder.DecodeString()
	assert.NotNil(t, err, "Err must not be nil")
	assert.IsType(t, errors.NewParseError(decoder.data[0], 0), err, "err message must euqal to the value expected")
}

func TestReadString(t *testing.T) {
	decoder := NewDecoder()
	decoder.SetData([]byte(`    "string with spaces and \"escape\"d \"quotes\" and escaped line returns \\n and escaped \\\\ escaped char"`))
	defer decoder.Release()

	data, err := decoder.ReadString()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, `"string with spaces and \"escape\"d \"quotes\" and escaped line returns \\n and escaped \\\\ escaped char"`, string(data), "data must be equal to the value expected")
}
