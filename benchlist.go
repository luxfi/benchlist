// Copyright (C) 2019-2025, Lux Industries, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package benchlist

import (
	validators "github.com/luxfi/consensus/validator"
	"github.com/luxfi/ids"
)

type Config struct {
	Deprecated bool
}

type Manager interface {
	IsBenched(nodeID ids.NodeID, chainID ids.ID) bool
	GetBenched(chainID ids.ID) []ids.NodeID
	RegisterChain(chainID ids.ID, vdrs validators.Manager) error
	Benchable(chainID ids.ID, nodeID ids.NodeID) Benchable
}

type Benchable interface {
	Benched(chainID ids.ID, nodeID ids.NodeID)
	Unbenched(chainID ids.ID, nodeID ids.NodeID)
}
