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

func TestSign(t *testing.T) {
	t.Parallel()

	expectedKey := &mocks2.MockKey{}
	expectetDigest := []byte{1, 2, 3, 4}
	expectedOpts := &mocks2.SignerOpts{}
	expectetValue := []byte{0, 1, 2, 3, 4}
	expectedErr := errors.New("Expected Error")

	signers := make(map[reflect.Type]Signer)
	signers[reflect.TypeOf(&mocks2.MockKey{})] = &mocks.Signer{
		KeyArg:    expectedKey,
		DigestArg: expectetDigest,
		OptsArg:   expectedOpts,
		Value:     expectetValue,
		Err:       nil,
	}
	csp := CSP{Signers: signers}
	value, err := csp.Sign(expectedKey, expectetDigest, expectedOpts)
	require.Equal(t, expectetValue, value)
	require.Nil(t, err)

	signers = make(map[reflect.Type]Signer)
	signers[reflect.TypeOf(&mocks2.MockKey{})] = &mocks.Signer{
		KeyArg:    expectedKey,
		DigestArg: expectetDigest,
		OptsArg:   expectedOpts,
		Value:     nil,
		Err:       expectedErr,
	}
	csp = CSP{Signers: signers}
	value, err = csp.Sign(expectedKey, expectetDigest, expectedOpts)
	require.Nil(t, value)
	require.Contains(t, err.Error(), expectedErr.Error())
}
