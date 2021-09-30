/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package sw

import (
	"crypto/sha256"
	"errors"
	"reflect"
	"testing"

	mocks2 "github.com/hyperledger/fabric/bccsp/mocks"
	"github.com/hyperledger/fabric/bccsp/sw/mocks"
	"github.com/stretchr/testify/require"
)

func TestHash(t *testing.T) {
	t.Parallel()

	expectetMsg := []byte{1, 2, 3, 4}
	expectedOpts := &mocks2.HashOpts{}
	expectetValue := []byte{1, 2, 3, 4, 5}
	expectedErr := errors.New("Expected Error")

	hashers := make(map[reflect.Type]Hasher)
	hashers[reflect.TypeOf(&mocks2.HashOpts{})] = &mocks.Hasher{
		MsgArg:  expectetMsg,
		OptsArg: expectedOpts,
		Value:   expectetValue,
		Err:     nil,
	}
	csp := CSP{Hashers: hashers}
	value, err := csp.Hash(expectetMsg, expectedOpts)
	require.Equal(t, expectetValue, value)
	require.Nil(t, err)

	hashers = make(map[reflect.Type]Hasher)
	hashers[reflect.TypeOf(&mocks2.HashOpts{})] = &mocks.Hasher{
		MsgArg:  expectetMsg,
		OptsArg: expectedOpts,
		Value:   nil,
		Err:     expectedErr,
	}
	csp = CSP{Hashers: hashers}
	value, err = csp.Hash(expectetMsg, expectedOpts)
	require.Nil(t, value)
	require.Contains(t, err.Error(), expectedErr.Error())
}

func TestGetHash(t *testing.T) {
	t.Parallel()

	expectedOpts := &mocks2.HashOpts{}
	expectetValue := sha256.New()
	expectedErr := errors.New("Expected Error")

	hashers := make(map[reflect.Type]Hasher)
	hashers[reflect.TypeOf(&mocks2.HashOpts{})] = &mocks.Hasher{
		OptsArg:   expectedOpts,
		ValueHash: expectetValue,
		Err:       nil,
	}
	csp := CSP{Hashers: hashers}
	value, err := csp.GetHash(expectedOpts)
	require.Equal(t, expectetValue, value)
	require.Nil(t, err)

	hashers = make(map[reflect.Type]Hasher)
	hashers[reflect.TypeOf(&mocks2.HashOpts{})] = &mocks.Hasher{
		OptsArg:   expectedOpts,
		ValueHash: expectetValue,
		Err:       expectedErr,
	}
	csp = CSP{Hashers: hashers}
	value, err = csp.GetHash(expectedOpts)
	require.Nil(t, value)
	require.Contains(t, err.Error(), expectedErr.Error())
}

func TestHasher(t *testing.T) {
	t.Parallel()

	hasher := &hasher{hash: sha256.New}

	msg := []byte("Hello World")
	out, err := hasher.Hash(msg, nil)
	require.NoError(t, err)
	h := sha256.New()
	h.Write(msg)
	out2 := h.Sum(nil)
	require.Equal(t, out, out2)

	hf, err := hasher.GetHash(nil)
	require.NoError(t, err)
	require.Equal(t, hf, sha256.New())
}
