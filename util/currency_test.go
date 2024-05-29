package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsSupportedCurrency(t *testing.T) {
	isValid := IsSupportedCurrency("USD")
	require.True(t, isValid)
}

func TestIsNotSupportedCurrency(t *testing.T) {
	isValid := IsSupportedCurrency("ASD")
	require.False(t, isValid)
}
