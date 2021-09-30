/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package policy

import (
	"github.com/golang/protobuf/proto"
	protomsp "github.com/hyperledger/fabric-protos-go/msp"
	"github.com/hyperledger/fabric/msp"
	"github.com/pkg/errors"
)

const (
	// Admins is the label for the local MSP admins
	Admins = "Admins"
	// Members is the label for the local MSP members
	Members = "Members"
)

type MSPPrincipalGetter interface {
	// Get returns an MSP principal for the given role
	Get(role string) (*protomsp.MSPPrincipal, error)
}

func NewLocalMSPPrincipalGetter(localMSP msp.MSP) MSPPrincipalGetter {
	return &localMSPPrincipalGetter{
		localMSP: localMSP,
	}
}

type localMSPPrincipalGetter struct {
	localMSP msp.MSP
}

func (m *localMSPPrincipalGetter) Get(role string) (*protomsp.MSPPrincipal, error) {
	mspid, err := m.localMSP.GetIdentifier()
	if err != nil {
		return nil, errors.WithMessage(err, "could not extract local msp identifier")
	}

	switch role {
	case Admins:
		principalBytes, err := proto.Marshal(&protomsp.MSPRole{Role: protomsp.MSPRole_ADMIN, MspIdentifier: mspid})
		if err != nil {
			return nil, errors.Wrap(err, "marshalling failed")
		}

		return &protomsp.MSPPrincipal{
			PrincipalClassification: protomsp.MSPPrincipal_ROLE,
			Principal:               principalBytes,
		}, nil
	case Members:
		principalBytes, err := proto.Marshal(&protomsp.MSPRole{Role: protomsp.MSPRole_MEMBER, MspIdentifier: mspid})
		if err != nil {
			return nil, errors.Wrap(err, "marshalling failed")
		}

		return &protomsp.MSPPrincipal{
			PrincipalClassification: protomsp.MSPPrincipal_ROLE,
			Principal:               principalBytes,
		}, nil
	default:
		return nil, errors.Errorf("MSP Principal role [%s] not recognized", role)
	}
}
