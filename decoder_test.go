package decoder

import (
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ModelQuery struct {
	ID          string    `query:"id"`
	SomeInt     int32     `query:"some_int"`
	SomeFloat   float32   `query:"some_float"`
	SomeBool    bool      `query:"some_bool"`
	SomeDate    time.Time `query:"some_date"`
	IntSlice    []int32   `query:"int_slice"`
	FloatSlice  []float32 `query:"float_slice"`
	StringSlice []string  `query:"string_slice"`
	BoolSlice   []bool    `query:"bool_slice"`
	UIntSlice   []uint32  `query:"uint_slice"`
}

type DecodeSuiteTest struct {
	suite.Suite
}

func TestRunDecodeSuite(t *testing.T) {
	suite.Run(t, new(DecodeSuiteTest))
}

func (s *DecodeSuiteTest) TestDecodeQueryParams() {
	s.T().Run("should decode query params", func(t *testing.T) {
		response := DecodeQueryParams[ModelQuery](buildUrlValues())
		assertResult(t, response)
	})
	s.T().Run("should decode query params", func(t *testing.T) {
		response := DecodeQueryParamsWithCustomTag[ModelQuery](buildUrlValues(), "query")
		assertResult(t, response)
	})
}

func buildUrlValues() url.Values {
	values := url.Values{}
	values.Add("id", "new_id")
	values.Add("some_int", "1")
	values.Add("some_float", "3.1314")
	values.Add("some_bool", "true")
	values.Add("some_date", "2024-02-01T10:54:12Z")
	values.Add("int_slice", "1,2,3")
	values.Add("float_slice", "1.1,2.2,3.3")
	values.Add("string_slice", "a,b,c")
	values.Add("bool_slice", "true,false,true")
	values.Add("uint_slice", "1,2,3")
	return values
}

func assertResult(t *testing.T, response ModelQuery) {
	assert.Equal(t, "new_id", response.ID)
	assert.Equal(t, int32(1), response.SomeInt)
	assert.Equal(t, float32(3.1314), response.SomeFloat)
	assert.Equal(t, true, response.SomeBool)
	assert.Equal(t, time.Date(2024, 02, 01, 10, 54, 12, 0, time.UTC), response.SomeDate)
	assert.ElementsMatch(t, []int32{1, 2, 3}, response.IntSlice)
	assert.ElementsMatch(t, []float32{1.1, 2.2, 3.3}, response.FloatSlice)
	assert.ElementsMatch(t, []string{"a", "b", "c"}, response.StringSlice)
	assert.ElementsMatch(t, []bool{true, false, true}, response.BoolSlice)
	assert.ElementsMatch(t, []uint32{1, 2, 3}, response.UIntSlice)
}
