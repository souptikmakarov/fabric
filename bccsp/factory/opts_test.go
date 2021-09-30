/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/
package factory

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFactoryOptsFactoryName(t *testing.T) {
	require.Equal(t, GetDefaultOpts().FactoryName(), "SW")
}
