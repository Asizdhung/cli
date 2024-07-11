package jsonfieldstest

import (
	"strings"
	"testing"

	"github.com/cli/cli/v2/pkg/cmdutil"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// May God have mercy on my soul
func jsonFieldsFor(cmd *cobra.Command) []string {
	stringFields, ok := cmd.Annotations["help:json-fields"]
	if !ok {
		return nil
	}

	return strings.Split(stringFields, ",")
}

type CommandConstructionFunc[T any] func(f *cmdutil.Factory, runF func(*T) error) *cobra.Command

func ExpectCommandToHaveJSONFields[T any](t *testing.T, fn CommandConstructionFunc[T], expectedFields []string) {
	actualFields := jsonFieldsFor(fn(&cmdutil.Factory{}, nil))
	assert.Equal(t, len(actualFields), len(expectedFields), "expected number of fields to match")
	require.ElementsMatch(t, expectedFields, actualFields)
}
