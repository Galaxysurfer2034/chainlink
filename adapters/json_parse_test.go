package adapters_test

import (
	"testing"

	"github.com/smartcontractkit/chainlink/adapters"
	"github.com/smartcontractkit/chainlink/store/models"
	"github.com/stretchr/testify/assert"
)

// Ensures the value for the field "last" returns "11779.99" as the value
func TestParseExistingPath(t *testing.T) {
	t.Parallel()
	input := models.RunResultWithValue(`{"high": "11850.00", "last": "11779.99", "timestamp": "1512487535", "bid": "11779.89", "vwap": "11525.17", "volume": "12916.67066094", "low": "11100.00", "ask": "11779.99", "open": 11613.07}`)

	adapter := adapters.JsonParse{Path: []string{"last"}}
	result := adapter.Perform(input, nil)
	assert.Equal(t, "11779.99", result.Value())
	assert.Nil(t, result.GetError())
}

// Ensures that the given path "doesnotexist" is not present in the JSON object
func TestParseNonExistingPath(t *testing.T) {
	t.Parallel()
	input := models.RunResultWithValue(`{"high": "11850.00", "last": "11779.99", "timestamp": "1512487535", "bid": "11779.89", "vwap": "11525.17", "volume": "12916.67066094", "low": "11100.00", "ask": "11779.99", "open": 11613.07}`)

	adapter := adapters.JsonParse{Path: []string{"doesnotexist"}}
	result := adapter.Perform(input, nil)
	assert.Equal(t, true, result.NullValue())
	assert.Nil(t, result.GetError())

	adapter = adapters.JsonParse{Path: []string{"doesnotexist", "noreally"}}
	result = adapter.Perform(input, nil)
	assert.Equal(t, true, result.NullValue())
	assert.NotNil(t, result.GetError())
}

func TestParseNullValue(t *testing.T) {
	t.Parallel()
	input := models.RunResult{}

	adapter := adapters.JsonParse{Path: []string{"last"}}
	result := adapter.Perform(input, nil)
	assert.Equal(t, true, result.NullValue())
	assert.NotNil(t, result.Error)
}
