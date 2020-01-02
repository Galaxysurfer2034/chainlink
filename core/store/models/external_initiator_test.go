package models_test

import (
	"testing"

	"chainlink/core/auth"
	"chainlink/core/internal/cltest"
	"chainlink/core/store/models"

	"github.com/stretchr/testify/assert"
)

func TestNewExternalInitiator(t *testing.T) {
	eia := auth.NewToken()
	assert.Len(t, eia.AccessKey, 32)
	assert.Len(t, eia.Secret, 64)

	url := cltest.WebURL(t, "http://localhost:8888")
	eir := &models.ExternalInitiatorRequest{
		Name: "bitcoin",
		URL:  &url,
	}
	ei, err := models.NewExternalInitiator(eia, eir)
	assert.NoError(t, err)
	assert.NotEqual(t, ei.HashedSecret, eia.Secret)
	assert.Equal(t, ei.AccessKey, eia.AccessKey)
}
