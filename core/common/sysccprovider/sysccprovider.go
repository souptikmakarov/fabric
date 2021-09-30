/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package sysccprovider

import (
	"github.com/hyperledger/fabric/common/channelconfig"
	"github.com/hyperledger/fabric/common/policies"
	"github.com/hyperledger/fabric/core/ledger"
)

// SystemChaincodeProvider provides an abstraction layer that is
// used for different packages to interact with code in the
// system chaincode package without importing it; more methods
// should be added below if necessary
type SystemChaincodeProvider interface {
	// GetQueryExecutorForLedger returns a query executor for the
	// ledger of the supplied channel.
	// That's useful for system chaincodes that require unfettered
	// access to the ledger
	GetQueryExecutorForLedger(cid string) (ledger.QueryExecutor, error)

	// GetApplicationConfig returns the configtxapplication.SharedConfig for the channel
	// and whether the Application config exists
	GetApplicationConfig(cid string) (channelconfig.Application, bool)

	// Returns the policy manager associated to the passed channel
	// and whether the policy manager exists
	PolicyManager(channelID string) (policies.Manager, bool)
}

// ChaincodeInstance is unique identifier of chaincode instance
type ChaincodeInstance struct {
	ChannelID        string
	ChaincodeName    string
	ChaincodeVersion string
}

func (ci *ChaincodeInstance) String() string {
	return ci.ChannelID + "." + ci.ChaincodeName + "#" + ci.ChaincodeVersion
}
