/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package sw

import (
	"errors"
	"reflect"
	"testing"

	mocks2 "github.com/hyperledger/fabric/bccsp/mocks"
	"github.com/hyperledger/fabric/bccsp/sw/mocks"
	"github.com/stretchr/testify/require"
)

func TestKeyDeriv(t *testing.T) {
	t.Parallel()

	expectedKey := &mocks2.MockKey{BytesValue: []byte{1, 2, 3}}
	expectedOpts := &mocks2.KeyDerivOpts{EphemeralValue: true}
	expectetValue := &mocks2.MockKey{BytesValue: []byte{1, 2, 3, 4, 5}}
	expectedErr := errors.New("Expected Error")

	keyDerivers := make(map[reflect.Type]KeyDeriver)
	keyDerivers[reflect.TypeOf(&mocks2.MockKey{})] = &mocks.KeyDeriver{
		KeyArg:  expectedKey,
		OptsArg: expectedOpts,
		Value:   expectetValue,
		Err:     expectedErr,
	}
	csp := CSP{KeyDerivers: keyDerivers}
	value, err := csp.KeyDeriv(expectedKey, expectedOpts)
	require.Nil(t, value)
	require.Contains(t, err.Error(), expectedErr.Error())

	keyDerivers = make(map[reflect.Type]KeyDeriver)
	keyDerivers[reflect.TypeOf(&mocks2.MockKey{})] = &mocks.KeyDeriver{
		KeyArg:  expectedKey,
		OptsArg: expectedOpts,
		Value:   expectetValue,
		Err:     nil,
	}
	csp = CSP{KeyDerivers: keyDerivers}
	value, err = csp.KeyDeriv(expectedKey, expectedOpts)
	require.Equal(t, expectetValue, value)
	require.Nil(t, err)
}

func TestECDSAPublicKeyKeyDeriver(t *testing.T) {
	t.Parallel()

	kd := ecdsaPublicKeyKeyDeriver{}

	_, err := kd.KeyDeriv(&mocks2.MockKey{}, nil)
	require.Error(t, err)
	require.Contains(t, err.Error(), "Invalid opts parameter. It must not be nil.")

	_, err = kd.KeyDeriv(&ecdsaPublicKey{}, &mocks2.KeyDerivOpts{})
	require.Error(t, err)
	require.Contains(t, err.Error(), "Unsupported 'KeyDerivOpts' provided [")
}

func TestECDSAPrivateKeyKeyDeriver(t *testing.T) {
	t.Parallel()

	kd := ecdsaPrivateKeyKeyDeriver{}

	_, err := kd.KeyDeriv(&mocks2.MockKey{}, nil)
	require.Error(t, err)
	require.Contains(t, err.Error(), "Invalid opts parameter. It must not be nil.")

	_, err = kd.KeyDeriv(&ecdsaPrivateKey{}, &mocks2.KeyDerivOpts{})
	require.Error(t, err)
	require.Contains(t, err.Error(), "Unsupported 'KeyDerivOpts' provided [")
}

func TestAESPrivateKeyKeyDeriver(t *testing.T) {
	t.Parallel()

	kd := aesPrivateKeyKeyDeriver{}

	_, err := kd.KeyDeriv(&mocks2.MockKey{}, nil)
	require.Error(t, err)
	require.Contains(t, err.Error(), "Invalid opts parameter. It must not be nil.")

	_, err = kd.KeyDeriv(&aesPrivateKey{}, &mocks2.KeyDerivOpts{})
	require.Error(t, err)
	require.Contains(t, err.Error(), "Unsupported 'KeyDerivOpts' provided [")
}
