/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/
package factory

import (
	"github.com/hyperledger/fabric/bccsp"
	"github.com/hyperledger/fabric/bccsp/sw"
	"github.com/pkg/errors"
)

const (
	// SoftwareBasedFactoryName is the name of the factory of the software-based BCCSP implementation
	SoftwareBasedFactoryName = "SW"
)

// SWFactory is the factory of the software-based BCCSP.
type SWFactory struct{}

// Name returns the name of this factory
func (f *SWFactory) Name() string {
	return SoftwareBasedFactoryName
}

// Get returns an instance of BCCSP using Opts.
func (f *SWFactory) Get(config *FactoryOpts) (bccsp.BCCSP, error) {
	// Validate arguments
	if config == nil || config.SW == nil {
		return nil, errors.New("Invalid config. It must not be nil.")
	}

	swOpts := config.SW

	var ks bccsp.KeyStore
	switch {
	case swOpts.FileKeystore != nil:
		fks, err := sw.NewFileBasedKeyStore(nil, swOpts.FileKeystore.KeyStorePath, false)
		if err != nil {
			return nil, errors.Wrapf(err, "Failed to initialize software key store")
		}
		ks = fks
	default:
		// Default to ephemeral key store
		ks = sw.NewDummyKeyStore()
	}

	return sw.NewWithParams(swOpts.Security, swOpts.Hash, ks)
}

// SwOpts contains options for the SWFactory
type SwOpts struct {
	// Default algorithms when not specified (Deprecated?)
	Security     int               `json:"security" yaml:"Security"`
	Hash         string            `json:"hash" yaml:"Hash"`
	FileKeystore *FileKeystoreOpts `json:"filekeystore,omitempty" yaml:"FileKeyStore,omitempty"`
}

// Pluggable Keystores, could add JKS, P12, etc..
type FileKeystoreOpts struct {
	KeyStorePath string `yaml:"KeyStore"`
}
