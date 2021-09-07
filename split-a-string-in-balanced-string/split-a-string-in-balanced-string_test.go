package splitastringinbalancedstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBalancedStringSplit(t *testing.T) {
	assert.Equal(t, 4, balancedStringSplit("RLRRLLRLRL"))
	assert.Equal(t, 3, balancedStringSplit("RLLLLRRRLR"))
	assert.Equal(t, 1, balancedStringSplit("LLLLRRRR"))
	assert.Equal(t, 2, balancedStringSplit("RLRRRLLRLL"))
}
