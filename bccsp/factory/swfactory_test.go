/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/
package factory

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSWFactoryName(t *testing.T) {
	f := &SWFactory{}
	require.Equal(t, f.Name(), SoftwareBasedFactoryName)
}

func TestSWFactoryGetInvalidArgs(t *testing.T) {
	f := &SWFactory{}

	_, err := f.Get(nil)
	require.Error(t, err, "Invalid config. It must not be nil.")

	_, err = f.Get(&FactoryOpts{})
	require.Error(t, err, "Invalid config. It must not be nil.")

	opts := &FactoryOpts{
		SW: &SwOpts{},
	}
	_, err = f.Get(opts)
	require.Error(t, err, "CSP:500 - Failed initializing configuration at [0,]")
}

func TestSWFactoryGet(t *testing.T) {
	f := &SWFactory{}

	opts := &FactoryOpts{
		SW: &SwOpts{
			Security: 256,
			Hash:     "SHA2",
		},
	}
	csp, err := f.Get(opts)
	require.NoError(t, err)
	require.NotNil(t, csp)

	opts = &FactoryOpts{
		SW: &SwOpts{
			Security:     256,
			Hash:         "SHA2",
			FileKeystore: &FileKeystoreOpts{KeyStorePath: os.TempDir()},
		},
	}
	csp, err = f.Get(opts)
	require.NoError(t, err)
	require.NotNil(t, csp)
}
