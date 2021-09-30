/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package msp

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTLSCAs(t *testing.T) {
	// testdata/tls contains TLS a root CA and an intermediate CA
	thisMSP := getLocalMSP(t, "testdata/tls")

	id, err := thisMSP.GetDefaultSigningIdentity()
	require.NoError(t, err)

	err = thisMSP.Validate(id.GetPublicVersion())
	require.NoError(t, err)

	tlsRootCerts := thisMSP.GetTLSRootCerts()
	require.Len(t, tlsRootCerts, 1)
	tlsRootCerts2, err := getPemMaterialFromDir("testdata/tls/tlscacerts")
	require.NoError(t, err)
	require.Len(t, tlsRootCerts2, 1)
	require.Equal(t, tlsRootCerts2[0], tlsRootCerts[0])

	tlsIntermediateCerts := thisMSP.GetTLSIntermediateCerts()
	require.Len(t, tlsIntermediateCerts, 1)
	tlsIntermediateCerts2, err := getPemMaterialFromDir("testdata/tls/tlsintermediatecerts")
	require.NoError(t, err)
	require.Len(t, tlsIntermediateCerts2, 1)
	require.Equal(t, tlsIntermediateCerts2[0], tlsIntermediateCerts[0])
}
