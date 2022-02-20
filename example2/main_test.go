/*
 */

package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGoodDbl(t *testing.T) {

	require := require.New(t)

	number := dbl(2)
	require.Greater(number, 3)
}
