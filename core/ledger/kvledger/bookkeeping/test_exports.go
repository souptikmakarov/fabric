/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package bookkeeping

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestEnv provides the bookkeeper provider env for testing
type TestEnv struct {
	t            testing.TB
	TestProvider *Provider
	dbPath       string
}

// NewTestEnv construct a TestEnv for testing
func NewTestEnv(t testing.TB) *TestEnv {
	dbPath, err := ioutil.TempDir("", "bookkeep")
	require.NoError(t, err)
	provider, err := NewProvider(dbPath)
	require.NoError(t, err)
	return &TestEnv{t, provider, dbPath}
}

// Cleanup cleansup the  store env after testing
func (te *TestEnv) Cleanup() {
	te.TestProvider.Close()
	os.RemoveAll(te.dbPath)
}
