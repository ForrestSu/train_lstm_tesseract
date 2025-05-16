package groundtruth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseTruth(t *testing.T) {
	code := parseTruth("eng_000_P5SY.png")
	assert.Equal(t, "P5SY", code)
}
