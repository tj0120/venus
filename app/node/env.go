package node

import (
	"context"
	"github.com/filecoin-project/venus/app/submodule/blockservice"
	"github.com/filecoin-project/venus/app/submodule/blockstore"
	"github.com/filecoin-project/venus/app/submodule/chain"
	"github.com/filecoin-project/venus/app/submodule/config"
	"github.com/filecoin-project/venus/app/submodule/discovery"
	"github.com/filecoin-project/venus/app/submodule/messaging"
	"github.com/filecoin-project/venus/app/submodule/network"
	"github.com/filecoin-project/venus/app/submodule/proofverification"
	"github.com/filecoin-project/venus/app/submodule/storagenetworking"
	"github.com/filecoin-project/venus/app/submodule/syncer"
	"github.com/filecoin-project/venus/app/submodule/wallet"
	cmds "github.com/ipfs/go-ipfs-cmds"
)

// Env is the environment for command API handlers.
type Env struct {
	ctx                  context.Context
	InspectorAPI         *Inspector
	BlockServiceAPI      *blockservice.BlockServiceAPI
	BlockStoreAPI        *blockstore.BlockstoreAPI
	ChainAPI             *chain.ChainAPI
	ConfigAPI            *config.ConfigAPI
	DiscoveryAPI         *discovery.DiscoveryAPI
	MessagingAPI         *messaging.MessagingAPI
	NetworkAPI           *network.NetworkAPI
	ProofVerificationAPI *proofverification.ProofVerificationApi
	StorageNetworkingAPI *storagenetworking.StorageNetworkingAPI
	SyncerAPI            *syncer.SyncerAPI
	WalletAPI            *wallet.WalletAPI
}

var _ cmds.Environment = (*Env)(nil)

// NewClientEnv returns a new environment for command API clients.
// This environment lacks direct access to any internal APIs.
func NewClientEnv(ctx context.Context) *Env {
	return &Env{ctx: ctx}
}

// Context returns the context of the environment.
func (ce *Env) Context() context.Context {
	return ce.ctx
}
