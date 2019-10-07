package services

import (
	"testing"
	"time"

	"github.com/guilhermesteves/github-repos-sort-api/internal/services"
	"github.com/stretchr/testify/assert"
)

func TestListRpostInGithub(t *testing.T) {
	result, err := services.ListReposInGithub("go", 1)

	assert.Nil(t, err)
	assert.NotNil(t, result)

	time.Sleep(100 * time.Millisecond)
}
