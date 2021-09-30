/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/
package sw

import (
	"testing"

	"github.com/hyperledger/fabric/bccsp/mocks"
	"github.com/stretchr/testify/require"
)

func TestNewDummyKeyStore(t *testing.T) {
	t.Parallel()

	ks := NewDummyKeyStore()
	require.NotNil(t, ks)
}

func TestDummyKeyStore_GetKey(t *testing.T) {
	t.Parallel()

	ks := NewDummyKeyStore()
	_, err := ks.GetKey([]byte{0, 1, 2, 3, 4})
	require.Error(t, err)
}

func TestDummyKeyStore_ReadOnly(t *testing.T) {
	t.Parallel()

	ks := NewDummyKeyStore()
	require.True(t, ks.ReadOnly())
}

func TestDummyKeyStore_StoreKey(t *testing.T) {
	t.Parallel()

	ks := NewDummyKeyStore()
	err := ks.StoreKey(&mocks.MockKey{})
	require.Error(t, err)
}
