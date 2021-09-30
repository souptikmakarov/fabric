/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package support

import (
	"sync"

	"github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric/common/channelconfig"
	"github.com/hyperledger/fabric/core/ledger"
	"github.com/hyperledger/fabric/msp"
)

type Support struct {
	LedgerVal     ledger.PeerLedger
	MSPManagerVal msp.MSPManager
	ApplyVal      error
	ACVal         channelconfig.ApplicationCapabilities

	sync.Mutex
	capabilitiesInvokeCount int
	mspManagerInvokeCount   int
}

func (ms *Support) Capabilities() channelconfig.ApplicationCapabilities {
	ms.Lock()
	defer ms.Unlock()
	ms.capabilitiesInvokeCount++
	return ms.ACVal
}

// Ledger returns LedgerVal
func (ms *Support) Ledger() ledger.PeerLedger {
	return ms.LedgerVal
}

// MSPManager returns MSPManagerVal
func (ms *Support) MSPManager() msp.MSPManager {
	ms.Lock()
	defer ms.Unlock()
	ms.mspManagerInvokeCount++
	return ms.MSPManagerVal
}

// Apply returns ApplyVal
func (ms *Support) Apply(configtx *common.ConfigEnvelope) error {
	return ms.ApplyVal
}

func (ms *Support) GetMSPIDs() []string {
	return []string{"SampleOrg"}
}

func (ms *Support) CapabilitiesInvokeCount() int {
	ms.Lock()
	defer ms.Unlock()
	return ms.capabilitiesInvokeCount
}

func (ms *Support) MSPManagerInvokeCount() int {
	ms.Lock()
	defer ms.Unlock()
	return ms.mspManagerInvokeCount
}
