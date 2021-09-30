/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package blockcutter

import (
	"testing"

	"github.com/hyperledger/fabric/orderer/common/blockcutter"
)

func TestBlockCutterInterface(t *testing.T) {
	_ = blockcutter.Receiver(NewReceiver())
}
