package types

import (
	"context"

	"github.com/smartcontractkit/chainlink-common/pkg/types/ccip"
)

type CCIPCommitProvider interface {
	PluginProvider

	NewOnRampReader(ctx context.Context, addr ccip.Address) (ccip.OnRampReader, error)
	NewOffRampReader(ctx context.Context, addr ccip.Address) (ccip.OffRampReader, error)
	NewCommitStoreReader(ctx context.Context, addr ccip.Address) (ccip.CommitStoreReader, error)
	NewPriceRegistryReader(ctx context.Context, addr ccip.Address) (ccip.PriceRegistryReader, error)
	NewPriceGetter(ctx context.Context) (ccip.PriceGetter, error)
	SourceNativeToken(ctx context.Context) (ccip.Address, error)
}

type CCIPExecProvider interface {
	PluginProvider

	NewOnRampReader(ctx context.Context, addr ccip.Address) (ccip.OnRampReader, error)
	NewOffRampReader(ctx context.Context, addr ccip.Address) (ccip.OffRampReader, error)
	NewCommitStoreReader(ctx context.Context, addr ccip.Address) (ccip.CommitStoreReader, error)
	NewPriceRegistryReader(ctx context.Context, addr ccip.Address) (ccip.PriceRegistryReader, error)
	NewTokenDataReader(ctx context.Context, tokenAddress ccip.Address) (ccip.TokenDataReader, error)
	SourceNativeToken(ctx context.Context) (ccip.Address, error)
	NewTokenPoolBatchedReader(ctx context.Context) (ccip.TokenPoolBatchedReader, error)
}

type CCIPCommitFactoryGenerator interface {
	NewCommitFactory(ctx context.Context, provider CCIPCommitProvider) (ReportingPluginFactory, error)
}

type CCIPExecFactoryGenerator interface {
	NewExecFactory(ctx context.Context, provider CCIPExecProvider) (ReportingPluginFactory, error)
}
type CCIPFactoryGenerator interface {
	CCIPCommitFactoryGenerator
	CCIPExecFactoryGenerator
}