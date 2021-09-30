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

func TestEncrypt(t *testing.T) {
	t.Parallel()

	expectedKey := &mocks2.MockKey{}
	expectedPlaintext := []byte{1, 2, 3, 4}
	expectedOpts := &mocks2.EncrypterOpts{}
	expectedCiphertext := []byte{0, 1, 2, 3, 4}
	expectedErr := errors.New("no error")

	encryptors := make(map[reflect.Type]Encryptor)
	encryptors[reflect.TypeOf(&mocks2.MockKey{})] = &mocks.Encryptor{
		KeyArg:       expectedKey,
		PlaintextArg: expectedPlaintext,
		OptsArg:      expectedOpts,
		EncValue:     expectedCiphertext,
		EncErr:       expectedErr,
	}

	csp := CSP{Encryptors: encryptors}

	ct, err := csp.Encrypt(expectedKey, expectedPlaintext, expectedOpts)
	require.Equal(t, expectedCiphertext, ct)
	require.Equal(t, expectedErr, err)
}
