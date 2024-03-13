package decoder

import (
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ModelQuery struct {
	ID        string    `query:"id"`
	SomeInt   int32     `query:"some_int"`
	SomeFloat float32   `query:"some_float"`
	SomeBool  bool      `query:"some_bool"`
	SomeDate  time.Time `query:"some_date"`
}

type ModelQueryInvalid struct {
	id string `query:"id"`
}

type DecodeSuiteTest struct {
	suite.Suite
}

func TestRunDecodeSuite(t *testing.T) {
	suite.Run(t, new(DecodeSuiteTest))
}

func (s *DecodeSuiteTest) TestDecodeQueryParams() {
	values := url.Values{}
	values.Add("id", "new_id")
	values.Add("some_int", "1")
	values.Add("some_float", "3.1314")
	values.Add("some_bool", "true")
	values.Add("some_date", "2024-02-01T10:54:12Z")

	s.T().Run("should decode query params", func(t *testing.T) {
		model := New[ModelQuery]()
		response := model.DecodeQueryParams(values)
		assert.Equal(t, "new_id", response.ID)
		assert.Equal(t, int32(1), response.SomeInt)
		assert.Equal(t, float32(3.1314), response.SomeFloat)
		assert.Equal(t, true, response.SomeBool)
		assert.Equal(t, time.Date(2024, 02, 01, 10, 54, 12, 0, time.UTC), response.SomeDate)
	})
}

func (s *DecodeSuiteTest) TestDecodeQueryParamsWithCustomTag() {
	values := url.Values{}
	values.Add("id", "new_id")
	values.Add("some_int", "1")
	values.Add("some_float", "3.1314")
	values.Add("some_bool", "true")
	values.Add("some_date", "2024-02-01T10:54:12Z")

	s.T().Run("should decode query params", func(t *testing.T) {
		model := New[ModelQuery]()
		response := model.DecodeQueryParamsWithCustomTag(values, "query")
		assert.Equal(t, "new_id", response.ID)
		assert.Equal(t, int32(1), response.SomeInt)
		assert.Equal(t, float32(3.1314), response.SomeFloat)
		assert.Equal(t, true, response.SomeBool)
		assert.Equal(t, time.Date(2024, 02, 01, 10, 54, 12, 0, time.UTC), response.SomeDate)
	})
}
