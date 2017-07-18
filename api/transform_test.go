package api

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestTransform(t *testing.T) {
	sample := `Lifeguards have whistles.`
	ss := extractKeywords(sample)
	assert.Len(t, ss, 2)
}

