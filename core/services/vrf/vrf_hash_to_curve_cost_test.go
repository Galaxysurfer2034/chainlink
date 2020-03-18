package vrf

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"strings"
	"testing"

	"chainlink/core/services/signatures/secp256k1"
	"chainlink/core/services/vrf/generated/solidity_verifier_wrapper"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type contract struct {
	contract *bind.BoundContract
	address  common.Address
	abi      *abi.ABI
	backend  *backends.SimulatedBackend
}

// deployVRFContract returns a deployed VRF contract, with some extra attributes
// which are useful for gas measurements.
func deployVRFContract(t *testing.T) (contract, common.Address) {
	x, y := secp256k1.Coordinates(Generator)
	key := ecdsa.PrivateKey{
		PublicKey: ecdsa.PublicKey{Curve: crypto.S256(), X: x, Y: y},
		D:         big.NewInt(1),
	}
	auth := bind.NewKeyedTransactor(&key)
	genesisData := core.GenesisAlloc{auth.From: {Balance: bi(1000000000)}}
	gasLimit := eth.DefaultConfig.Miner.GasCeil
	backend := backends.NewSimulatedBackend(genesisData, gasLimit)
	parsed, err := abi.JSON(strings.NewReader(
		solidity_verifier_wrapper.VRFTestHelperABI))
	require.NoError(t, err, "could not parse VRF ABI")
	address, _, vRFContract, err := bind.DeployContract(auth, parsed,
		common.FromHex(solidity_verifier_wrapper.VRFTestHelperBin), backend)
	require.NoError(t, err, "failed to deploy VRF contract to simulated blockchain")
	backend.Commit()
	return contract{vRFContract, address, &parsed, backend}, crypto.PubkeyToAddress(
		key.PublicKey)
}

func measureHashToCurveGasCost(t *testing.T, contract contract,
	owner common.Address, input int64) (gasCost, numOrdinates uint64) {
	rawData, err := contract.abi.Pack("hashToCurve_", pair(secp256k1.Coordinates(Generator)),
		big.NewInt(input))
	require.NoError(t, err, "failed to construct hashToCurve message for VRF contract")
	callMsg := ethereum.CallMsg{From: owner, To: &contract.address, Data: rawData}
	estimate, err := contract.backend.EstimateGas(context.TODO(), callMsg)
	require.NoError(t, err, "failed to estimate gas for on-chain hashToCurve calculation")
	_, err = HashToCurve(Generator, big.NewInt(input),
		func(*big.Int) { numOrdinates += 1 })
	require.NoError(t, err, "corresponding golang HashToCurve calculation failed")
	return estimate, numOrdinates
}

var baseCost uint64 = 25000
var marginalCost uint64 = 15555

func HashToCurveGasCostBound(numOrdinates uint64) uint64 {
	return baseCost + marginalCost*numOrdinates
}

func TestMeasureHashToCurveGasCost(t *testing.T) {
	contract, owner := deployVRFContract(t)
	numSamples := int64(numSamples())
	for i := int64(0); i < numSamples; i += 1 {
		gasCost, numOrdinates := measureHashToCurveGasCost(t, contract, owner, i)
		assert.Less(t, gasCost, HashToCurveGasCostBound(numOrdinates),
			"on-chain hashToCurve gas cost exceeded estimate function")
	}
	require.Less(t, HashToCurveGasCostBound(128), uint64(2.017e6),
		"estimate for on-chain hashToCurve gas cost with 128 iterations is greater "+
			"than stated in the VRF.sol documentation")
}
