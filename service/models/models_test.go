package models

import (
	"github.com/gsinekliev/eval-service/service/eval"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitEmptyErrorStore(t *testing.T) {
	var actual = InitErrorStore()
	assert.Empty(t, actual)
	assert.IsType(t, ErrorStore{}, actual)
}

func TestAddErrorToStore(t *testing.T) {
	var store = InitErrorStore()
	err := Error{Expression: "expr",
		Endpoint: "endp1", Frequency: 1, ErrorType: eval.ErrorDivisionByZero}
	store.AddError(err)

	assert.Equal(t, 1, len(store))
	assert.Equal(t, err, store["endp1expr"])
}

func TestAddMultipleItemsToStoreIncreasesFrequency(t *testing.T) {
	var store = InitErrorStore()
	err := Error{Expression: "expr",
		Endpoint: "endp1", Frequency: 1, ErrorType: eval.ErrorDivisionByZero}
	store.AddError(err)
	store.AddError(err)

	assert.Equal(t, 1, len(store))
	assert.Equal(t, 2, store["endp1expr"].Frequency)
	assert.Equal(t, err.Endpoint, store["endp1expr"].Endpoint)
	assert.Equal(t, err.Expression, store["endp1expr"].Expression)
	assert.Equal(t, err.ErrorType, store["endp1expr"].ErrorType)
}
