package exec

import (
	"context"

	ethcommon "github.com/ethereum/go-ethereum/common"

	derive "github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	opeth "github.com/ethereum-optimism/optimism/op-service/eth"

	abci "github.com/cometbft/cometbft/abci/types"
)

var _ derive.Engine = (*WrapperEngine)(nil)

type WrapperEngine struct {
	AbciApp abci.Application
}

// PayloadByHash implements Engine
// nolint // TODO: implement
func (we *WrapperEngine) PayloadByHash(ctx context.Context, hash ethcommon.Hash) (*opeth.ExecutionPayloadEnvelope, error) {
	return nil, nil
}

// PayloadByNumber implements Engine
// nolint // TODO: implement
func (we *WrapperEngine) PayloadByNumber(ctx context.Context, number uint64) (*opeth.ExecutionPayloadEnvelope, error) {
	return nil, nil
}

// L2BlockRefByLabel implements Engine
// nolint // TODO: implement
func (we *WrapperEngine) L2BlockRefByLabel(ctx context.Context, label opeth.BlockLabel) (opeth.L2BlockRef, error) {
	return opeth.L2BlockRef{}, nil
}

// L2BlockRefByHash implements Engine
// nolint // TODO: implement
func (we *WrapperEngine) L2BlockRefByHash(ctx context.Context, l2Hash ethcommon.Hash) (opeth.L2BlockRef, error) {
	return opeth.L2BlockRef{}, nil
}

// L2BlockRefByNumber implements Engine
// nolint // TODO: implement
func (we *WrapperEngine) L2BlockRefByNumber(ctx context.Context, num uint64) (opeth.L2BlockRef, error) {
	return opeth.L2BlockRef{}, nil
}

// SystemConfigByL2Hash implements Engine
// nolint // TODO: implement
func (we *WrapperEngine) SystemConfigByL2Hash(ctx context.Context, hash ethcommon.Hash) (opeth.SystemConfig, error) {
	return opeth.SystemConfig{}, nil
}

// GetPayload implements Engine
// nolint // TODO: implement
func (we *WrapperEngine) GetPayload(ctx context.Context, payloadId opeth.PayloadID) (*opeth.ExecutionPayloadEnvelope, error) {
	return nil, nil
}

// ForkchoiceUpdate implements Engine
// nolint // TODO: implement
func (we *WrapperEngine) ForkchoiceUpdate(ctx context.Context, state *opeth.ForkchoiceState, attr *opeth.PayloadAttributes) (*opeth.ForkchoiceUpdatedResult, error) {
	return nil, nil
}

// NewPayload implements Engine
// nolint // TODO: implement
func (we *WrapperEngine) NewPayload(ctx context.Context, payload *opeth.ExecutionPayload, parentBeaconBlockRoot *ethcommon.Hash) (*opeth.PayloadStatusV1, error) {
	return nil, nil
}
