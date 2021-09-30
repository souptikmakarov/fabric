/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package multichannel

import (
	"testing"

	"github.com/hyperledger/fabric/orderer/consensus"
)

func TestConsenterSupportInterface(t *testing.T) {
	_ = consensus.ConsenterSupport(&ConsenterSupport{})
}
