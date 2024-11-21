package util_test

import (
	"fmt"
	"scrapperjaltup/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateUniqueID(t *testing.T) {
	t.Parallel()

	for i := 0; i < 1000; i++ {
		id := util.GenerateUniqueID(10)
		assert.Len(t, id, 10)

		fmt.Printf("id = %v\n", id)
	}
}
