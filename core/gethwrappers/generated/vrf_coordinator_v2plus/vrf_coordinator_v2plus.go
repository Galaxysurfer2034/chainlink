// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vrf_coordinator_v2plus

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated"
)

var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

type VRFCoordinatorV2PlusFeeConfig struct {
	FulfillmentFlatFeeLinkPPM uint32
	FulfillmentFlatFeeEthPPM  uint32
}

type VRFCoordinatorV2PlusRequestCommitment struct {
	BlockNum         uint64
	SubId            *big.Int
	CallbackGasLimit uint32
	NumWords         uint32
	Sender           common.Address
	ExtraArgs        []byte
}

type VRFProof struct {
	Pk            [2]*big.Int
	Gamma         [2]*big.Int
	C             *big.Int
	S             *big.Int
	Seed          *big.Int
	UWitness      common.Address
	CGammaWitness [2]*big.Int
	SHashWitness  [2]*big.Int
	ZInv          *big.Int
}

type VRFV2PlusClientRandomWordsRequest struct {
	KeyHash              [32]byte
	SubId                *big.Int
	RequestConfirmations uint16
	CallbackGasLimit     uint32
	NumWords             uint32
	ExtraArgs            []byte
}

var VRFCoordinatorV2PlusMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"blockhashStore\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"internalBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"externalBalance\",\"type\":\"uint256\"}],\"name\":\"BalanceInvariantViolated\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNum\",\"type\":\"uint256\"}],\"name\":\"BlockhashNotInStore\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"coordinatorAddress\",\"type\":\"address\"}],\"name\":\"CoordinatorAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"coordinatorAddress\",\"type\":\"address\"}],\"name\":\"CoordinatorNotRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedToSendEther\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"have\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"want\",\"type\":\"uint32\"}],\"name\":\"GasLimitTooBig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectCommitment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"have\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"InsufficientGasForConsumer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCalldata\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"InvalidConsumer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidExtraArgsTag\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"linkWei\",\"type\":\"int256\"}],\"name\":\"InvalidLinkWeiPrice\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"have\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"min\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"max\",\"type\":\"uint16\"}],\"name\":\"InvalidRequestConfirmations\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSubscription\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LinkAlreadySet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposedOwner\",\"type\":\"address\"}],\"name\":\"MustBeRequestedOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"MustBeSubOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoCorrespondingRequest\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"}],\"name\":\"NoSuchProvingKey\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"have\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"want\",\"type\":\"uint32\"}],\"name\":\"NumWordsTooBig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableFromLink\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PaymentTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingRequestExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"}],\"name\":\"ProvingKeyAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyConsumers\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"minimumRequestConfirmations\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"maxGasLimit\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"stalenessSeconds\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"gasAfterPaymentCalculation\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"fallbackWeiPerUnitLink\",\"type\":\"int256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"fulfillmentFlatFeeLinkPPM\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"fulfillmentFlatFeeEthPPM\",\"type\":\"uint32\"}],\"indexed\":false,\"internalType\":\"structVRFCoordinatorV2Plus.FeeConfig\",\"name\":\"feeConfig\",\"type\":\"tuple\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"coordinatorAddress\",\"type\":\"address\"}],\"name\":\"CoordinatorDeregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"coordinatorAddress\",\"type\":\"address\"}],\"name\":\"CoordinatorRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EthFundsRecovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundsRecovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newCoordinator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"}],\"name\":\"MigrationCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"ProvingKeyDeregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"ProvingKeyRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"outputSeed\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"payment\",\"type\":\"uint96\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"RandomWordsFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"preSeed\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"minimumRequestConfirmations\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"numWords\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RandomWordsRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountLink\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountEth\",\"type\":\"uint256\"}],\"name\":\"SubscriptionCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"SubscriptionConsumerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"SubscriptionConsumerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"SubscriptionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBalance\",\"type\":\"uint256\"}],\"name\":\"SubscriptionFunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldEthBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newEthBalance\",\"type\":\"uint256\"}],\"name\":\"SubscriptionFundedWithEth\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"SubscriptionOwnerTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"SubscriptionOwnerTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BLOCKHASH_STORE\",\"outputs\":[{\"internalType\":\"contractBlockhashStoreInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LINK\",\"outputs\":[{\"internalType\":\"contractLinkTokenInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LINK_ETH_FEED\",\"outputs\":[{\"internalType\":\"contractAggregatorV3Interface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_CONSUMERS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_NUM_WORDS\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_REQUEST_CONFIRMATIONS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"}],\"name\":\"acceptSubscriptionOwnerTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"addConsumer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"cancelSubscription\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"createSubscription\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"deregisterMigratableCoordinator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"publicProvingKey\",\"type\":\"uint256[2]\"}],\"name\":\"deregisterProvingKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"pk\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"gamma\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"c\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seed\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"uWitness\",\"type\":\"address\"},{\"internalType\":\"uint256[2]\",\"name\":\"cGammaWitness\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"sHashWitness\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"zInv\",\"type\":\"uint256\"}],\"internalType\":\"structVRF.Proof\",\"name\":\"proof\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"blockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"numWords\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structVRFCoordinatorV2Plus.RequestCommitment\",\"name\":\"rc\",\"type\":\"tuple\"}],\"name\":\"fulfillRandomWords\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"}],\"name\":\"fundSubscriptionWithEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRequestConfig\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"}],\"name\":\"getSubscription\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"balance\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"ethBalance\",\"type\":\"uint96\"},{\"internalType\":\"uint64\",\"name\":\"reqCount\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"consumers\",\"type\":\"address[]\"},{\"internalType\":\"uint64[]\",\"name\":\"nonces\",\"type\":\"uint64[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"publicKey\",\"type\":\"uint256[2]\"}],\"name\":\"hashOfKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"newCoordinator\",\"type\":\"address\"}],\"name\":\"migrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migrationVersion\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onTokenTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"name\":\"oracleWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"name\":\"oracleWithdrawEth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"}],\"name\":\"ownerCancelSubscription\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"}],\"name\":\"pendingRequestExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"recoverEthFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"recoverFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"registerMigratableCoordinator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"uint256[2]\",\"name\":\"publicProvingKey\",\"type\":\"uint256[2]\"}],\"name\":\"registerProvingKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"removeConsumer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"requestConfirmations\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"numWords\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structVRFV2PlusClient.RandomWordsRequest\",\"name\":\"req\",\"type\":\"tuple\"}],\"name\":\"requestRandomWords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"requestSubscriptionOwnerTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_config\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"minimumRequestConfirmations\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"stalenessSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"gasAfterPaymentCalculation\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_currentSubNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_fallbackWeiPerUnitLink\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_feeConfig\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"fulfillmentFlatFeeLinkPPM\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"fulfillmentFlatFeeEthPPM\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"s_provingKeyHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"s_provingKeys\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"s_requestCommitments\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_totalBalance\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_totalEthBalance\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"minimumRequestConfirmations\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"stalenessSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"gasAfterPaymentCalculation\",\"type\":\"uint32\"},{\"internalType\":\"int256\",\"name\":\"fallbackWeiPerUnitLink\",\"type\":\"int256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"fulfillmentFlatFeeLinkPPM\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"fulfillmentFlatFeeEthPPM\",\"type\":\"uint32\"}],\"internalType\":\"structVRFCoordinatorV2Plus.FeeConfig\",\"name\":\"feeConfig\",\"type\":\"tuple\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"link\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"linkEthFeed\",\"type\":\"address\"}],\"name\":\"setLINKAndLINKETHFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b50604051620062bd380380620062bd833981016040819052620000349162000188565b33806000816200008b5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000be57620000be81620000dc565b505060016002555060601b6001600160601b031916608052620001ba565b6001600160a01b038116331415620001375760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000082565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000602082840312156200019b57600080fd5b81516001600160a01b0381168114620001b357600080fd5b9392505050565b60805160601c6160dd620001e0600039600081816105f80152613cc701526160dd6000f3fe6080604052600436106102f15760003560e01c806386fe91c71161018f578063bec4c08c116100e1578063dc311dd31161008a578063e95704bd11610064578063e95704bd14610949578063ee9d2d3814610970578063f2fde38b1461099d57600080fd5b8063dc311dd3146108e4578063e72f6e3014610916578063e8509bff1461093657600080fd5b8063d98e620e116100bb578063d98e620e1461086e578063da2f26101461088e578063dac83d29146108c457600080fd5b8063bec4c08c1461080e578063caf70c4a1461082e578063cb6317971461084e57600080fd5b8063a4c0ed3611610143578063ad1783611161011d578063ad178361146107ae578063b08c8795146107ce578063b2a7cac5146107ee57600080fd5b8063a4c0ed361461074e578063a8cb447b1461076e578063aa433aff1461078e57600080fd5b80639b1c385e116101745780639b1c385e146106df5780639d40a6fd146106ff578063a21a23e41461073957600080fd5b806386fe91c7146106955780638da5cb5b146106c157600080fd5b806340d6bb821161024857806364d51a2a116101fc5780636b6feccc116101d65780636b6feccc1461061a5780636f64f03f1461066057806379ba50971461068057600080fd5b806364d51a2a146105b157806366316d8d146105c6578063689c4517146105e657600080fd5b806346d8d4861161022d57806346d8d4861461055157806357133e64146105715780635d06b4ab1461059157600080fd5b806340d6bb82146104f657806341af6c871461052157600080fd5b80630ae09540116102aa578063294daa4911610284578063294daa4914610482578063330987b31461049e578063405b84fa146104d657600080fd5b80630ae095401461040257806315c48b84146104225780631b6b6d231461044a57600080fd5b8063043bd6ae116102db578063043bd6ae14610345578063088070f51461036957806308821d58146103e257600080fd5b8062012291146102f657806304104edb14610323575b600080fd5b34801561030257600080fd5b5061030b6109bd565b60405161031a93929190615b68565b60405180910390f35b34801561032f57600080fd5b5061034361033e3660046154ce565b610a39565b005b34801561035157600080fd5b5061035b600e5481565b60405190815260200161031a565b34801561037557600080fd5b50600f546103b19061ffff81169063ffffffff620100008204811691660100000000000081048216916a01000000000000000000009091041684565b6040805161ffff909516855263ffffffff93841660208601529183169184019190915216606082015260800161031a565b3480156103ee57600080fd5b506103436103fd36600461560f565b610bfb565b34801561040e57600080fd5b5061034361041d3660046158cc565b610d8f565b34801561042e57600080fd5b5061043760c881565b60405161ffff909116815260200161031a565b34801561045657600080fd5b5060035461046a906001600160a01b031681565b6040516001600160a01b03909116815260200161031a565b34801561048e57600080fd5b506040516001815260200161031a565b3480156104aa57600080fd5b506104be6104b93660046156e2565b610e7b565b6040516001600160601b03909116815260200161031a565b3480156104e257600080fd5b506103436104f13660046158cc565b6113b7565b34801561050257600080fd5b5061050c6101f481565b60405163ffffffff909116815260200161031a565b34801561052d57600080fd5b5061054161053c366004615664565b611798565b604051901515815260200161031a565b34801561055d57600080fd5b5061034361056c3660046154eb565b61199a565b34801561057d57600080fd5b5061034361058c366004615520565b611b35565b34801561059d57600080fd5b506103436105ac3660046154ce565b611bae565b3480156105bd57600080fd5b50610437606481565b3480156105d257600080fd5b506103436105e13660046154eb565b611c85565b3480156105f257600080fd5b5061046a7f000000000000000000000000000000000000000000000000000000000000000081565b34801561062657600080fd5b506010546106439063ffffffff8082169164010000000090041682565b6040805163ffffffff93841681529290911660208301520161031a565b34801561066c57600080fd5b5061034361067b366004615559565b611e41565b34801561068c57600080fd5b50610343611f59565b3480156106a157600080fd5b506008546104be906801000000000000000090046001600160601b031681565b3480156106cd57600080fd5b506000546001600160a01b031661046a565b3480156106eb57600080fd5b5061035b6106fa3660046157de565b61200a565b34801561070b57600080fd5b506008546107209067ffffffffffffffff1681565b60405167ffffffffffffffff909116815260200161031a565b34801561074557600080fd5b5061035b61244b565b34801561075a57600080fd5b50610343610769366004615586565b6126f0565b34801561077a57600080fd5b506103436107893660046154ce565b6128df565b34801561079a57600080fd5b506103436107a9366004615664565b6129fb565b3480156107ba57600080fd5b5060045461046a906001600160a01b031681565b3480156107da57600080fd5b506103436107e936600461582e565b612a5b565b3480156107fa57600080fd5b50610343610809366004615664565b612be5565b34801561081a57600080fd5b506103436108293660046158cc565b612d52565b34801561083a57600080fd5b5061035b61084936600461562b565b612f25565b34801561085a57600080fd5b506103436108693660046158cc565b612f55565b34801561087a57600080fd5b5061035b610889366004615664565b613272565b34801561089a57600080fd5b5061046a6108a9366004615664565b600b602052600090815260409020546001600160a01b031681565b3480156108d057600080fd5b506103436108df3660046158cc565b613293565b3480156108f057600080fd5b506109046108ff366004615664565b6133cb565b60405161031a96959493929190615d66565b34801561092257600080fd5b506103436109313660046154ce565b6135c0565b610343610944366004615664565b613788565b34801561095557600080fd5b506008546104be90600160a01b90046001600160601b031681565b34801561097c57600080fd5b5061035b61098b366004615664565b600d6020526000908152604090205481565b3480156109a957600080fd5b506103436109b83660046154ce565b6138e0565b600f54600c805460408051602080840282018101909252828152600094859460609461ffff8316946201000090930463ffffffff16939192839190830182828015610a2757602002820191906000526020600020905b815481526020019060010190808311610a13575b50505050509050925092509250909192565b610a416138f1565b60115460005b81811015610bce57826001600160a01b031660118281548110610a6c57610a6c616061565b6000918252602090912001546001600160a01b03161415610bbc576011610a94600184615f59565b81548110610aa457610aa4616061565b600091825260209091200154601180546001600160a01b039092169183908110610ad057610ad0616061565b600091825260209091200180546001600160a01b0319166001600160a01b0392909216919091179055826011610b07600185615f59565b81548110610b1757610b17616061565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b031602179055506011805480610b5657610b5661604b565b6000828152602090819020600019908301810180546001600160a01b03191690559091019091556040516001600160a01b03851681527ff80a1a97fd42251f3c33cda98635e7399253033a6774fe37cd3f650b5282af37910160405180910390a1505050565b80610bc681615fc8565b915050610a47565b50604051635428d44960e01b81526001600160a01b03831660048201526024015b60405180910390fd5b50565b610c036138f1565b604080518082018252600091610c32919084906002908390839080828437600092019190915250612f25915050565b6000818152600b60205260409020549091506001600160a01b031680610c6e57604051631dfd6e1360e21b815260048101839052602401610bef565b6000828152600b6020526040812080546001600160a01b03191690555b600c54811015610d465782600c8281548110610ca957610ca9616061565b90600052602060002001541415610d3457600c805460009190610cce90600190615f59565b81548110610cde57610cde616061565b9060005260206000200154905080600c8381548110610cff57610cff616061565b600091825260209091200155600c805480610d1c57610d1c61604b565b60019003818190600052602060002001600090559055505b80610d3e81615fc8565b915050610c8b565b50806001600160a01b03167f72be339577868f868798bac2c93e52d6f034fef4689a9848996c14ebb7416c0d83604051610d8291815260200190565b60405180910390a2505050565b60008281526006602052604090205482906001600160a01b031680610dc757604051630fb532db60e11b815260040160405180910390fd5b336001600160a01b03821614610dfb57604051636c51fda960e11b81526001600160a01b0382166004820152602401610bef565b600280541415610e3b5760405162461bcd60e51b815260206004820152601f60248201526000805160206160b18339815191526044820152606401610bef565b60028055610e4884611798565b15610e6657604051631685ecdd60e31b815260040160405180910390fd5b610e70848461394d565b505060016002555050565b6000600280541415610ebd5760405162461bcd60e51b815260206004820152601f60248201526000805160206160b18339815191526044820152606401610bef565b6002805560005a90506000610ed28585613ae4565b90506000610ee660808601606087016158f1565b63ffffffff1667ffffffffffffffff811115610f0457610f04616077565b604051908082528060200260200182016040528015610f2d578160200160208202803683370190505b50905060005b610f4360808701606088016158f1565b63ffffffff16811015610fb857826040015181604051602001610f70929190918252602082015260400190565b6040516020818303038152906040528051906020012060001c828281518110610f9b57610f9b616061565b602090810291909101015280610fb081615fc8565b915050610f33565b50602080830180516000908152600d90925260408083208390559051905182917f1fe543e3000000000000000000000000000000000000000000000000000000009161100991908690602401615bdb565b60408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff166001600160e01b0319909416939093179092529150600090611085906110699060608b01908b016158f1565b63ffffffff1661107f60a08b0160808c016154ce565b84613e39565b602089810135600090815260079091526040902054909150600160c01b900467ffffffffffffffff166110b9816001615ed8565b6020808b01356000908152600790915260408120805467ffffffffffffffff93909316600160c01b0277ffffffffffffffffffffffffffffffffffffffffffffffff9093169290921790915561111a61111560a08c018c615dfb565b613e87565b51600f54909150600090611145908a906a0100000000000000000000900463ffffffff163a85613f36565b9050811561124e576020808c01356000908152600790915260409020546001600160601b03808316600160601b90920416101561119557604051631e9acf1760e31b815260040160405180910390fd5b60208b81013560009081526007909152604090208054829190600c906111cc908490600160601b90046001600160601b0316615f70565b82546101009290920a6001600160601b0381810219909316918316021790915589516000908152600b60209081526040808320546001600160a01b03168352600a90915281208054859450909261122591859116615f04565b92506101000a8154816001600160601b0302191690836001600160601b0316021790555061133a565b6020808c01356000908152600790915260409020546001600160601b038083169116101561128f57604051631e9acf1760e31b815260040160405180910390fd5b6020808c0135600090815260079091526040812080548392906112bc9084906001600160601b0316615f70565b82546101009290920a6001600160601b0381810219909316918316021790915589516000908152600b60209081526040808320546001600160a01b03168352600990915281208054859450909261131591859116615f04565b92506101000a8154816001600160601b0302191690836001600160601b031602179055505b60006113556040518060200160405280851515815250613f86565b90508b6020013589602001517f7bad4a5a5566568e03ff9e9fdaedf6b8a760aa3ffc0abad348cf93787f04e17e8b6040015185858a60405161139a9493929190615d2b565b60405180910390a35060016002559b9a5050505050505050505050565b6113c08161400c565b6113e857604051635428d44960e01b81526001600160a01b0382166004820152602401610bef565b6000806000806113f7866133cb565b50945094505093509350336001600160a01b0316826001600160a01b0316146114625760405162461bcd60e51b815260206004820152601660248201527f4e6f7420737562736372697074696f6e206f776e6572000000000000000000006044820152606401610bef565b61146b86611798565b156114b85760405162461bcd60e51b815260206004820152601660248201527f50656e64696e67207265717565737420657869737473000000000000000000006044820152606401610bef565b60006040518060c001604052806114cd600190565b60ff168152602001888152602001846001600160a01b03168152602001838152602001866001600160601b03168152602001856001600160601b031681525090506000816040516020016115219190615af3565b604051602081830303815290604052905061153b88614076565b50506040517fce3f47190000000000000000000000000000000000000000000000000000000081526001600160a01b0388169063ce3f4719906001600160601b0388169061158d908590600401615ae0565b6000604051808303818588803b1580156115a657600080fd5b505af11580156115ba573d6000803e3d6000fd5b505060035460405163a9059cbb60e01b81526001600160a01b038c811660048301526001600160601b038c166024830152909116935063a9059cbb92506044019050602060405180830381600087803b15801561161657600080fd5b505af115801561162a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061164e9190615647565b61169a5760405162461bcd60e51b815260206004820152601260248201527f696e73756666696369656e742066756e647300000000000000000000000000006044820152606401610bef565b60005b835181101561174b578381815181106116b8576116b8616061565b60209081029190910101516040517f8ea981170000000000000000000000000000000000000000000000000000000081526001600160a01b038a8116600483015290911690638ea9811790602401600060405180830381600087803b15801561172057600080fd5b505af1158015611734573d6000803e3d6000fd5b50505050808061174390615fc8565b91505061169d565b50604080516001600160a01b0389168152602081018a90527fd63ca8cb945956747ee69bfdc3ea754c24a4caf7418db70e46052f7850be4187910160405180910390a15050505050505050565b6000818152600660209081526040808320815160608101835281546001600160a01b039081168252600183015416818501526002820180548451818702810187018652818152879693958601939092919083018282801561182257602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611804575b505050505081525050905060005b8160400151518110156119905760005b600c5481101561197d576000611946600c838154811061186257611862616061565b90600052602060002001548560400151858151811061188357611883616061565b60200260200101518860056000896040015189815181106118a6576118a6616061565b6020908102919091018101516001600160a01b03908116835282820193909352604091820160009081208e825282528290205482518083018890529590931685830152606085019390935267ffffffffffffffff9091166080808501919091528151808503909101815260a08401825280519083012060c084019490945260e0808401859052815180850390910181526101009093019052815191012091565b506000818152600d60205260409020549091501561196a5750600195945050505050565b508061197581615fc8565b915050611840565b508061198881615fc8565b915050611830565b5060009392505050565b6002805414156119da5760405162461bcd60e51b815260206004820152601f60248201526000805160206160b18339815191526044820152606401610bef565b60028055336000908152600a60205260409020546001600160601b0380831691161015611a1a57604051631e9acf1760e31b815260040160405180910390fd5b336000908152600a602052604081208054839290611a429084906001600160601b0316615f70565b92506101000a8154816001600160601b0302191690836001600160601b0316021790555080600860148282829054906101000a90046001600160601b0316611a8a9190615f70565b92506101000a8154816001600160601b0302191690836001600160601b031602179055506000826001600160a01b0316826001600160601b031660405160006040518083038185875af1925050503d8060008114611b04576040519150601f19603f3d011682016040523d82523d6000602084013e611b09565b606091505b5050905080611b2b57604051630dcf35db60e41b815260040160405180910390fd5b5050600160025550565b611b3d6138f1565b6003546001600160a01b031615611b80576040517f2d118a6e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600380546001600160a01b039384166001600160a01b03199182161790915560048054929093169116179055565b611bb66138f1565b611bbf8161400c565b15611c01576040517fac8a27ef0000000000000000000000000000000000000000000000000000000081526001600160a01b0382166004820152602401610bef565b601180546001810182556000919091527f31ecc21a745e3968a04e9570e4425bc18fa8019c68028196b546d1669c200c680180546001600160a01b0319166001600160a01b0383169081179091556040519081527fb7cabbfc11e66731fc77de0444614282023bcbd41d16781c753a431d0af016259060200160405180910390a150565b600280541415611cc55760405162461bcd60e51b815260206004820152601f60248201526000805160206160b18339815191526044820152606401610bef565b60028055336000908152600960205260409020546001600160601b0380831691161015611d0557604051631e9acf1760e31b815260040160405180910390fd5b3360009081526009602052604081208054839290611d2d9084906001600160601b0316615f70565b92506101000a8154816001600160601b0302191690836001600160601b03160217905550806008808282829054906101000a90046001600160601b0316611d749190615f70565b82546101009290920a6001600160601b0381810219909316918316021790915560035460405163a9059cbb60e01b81526001600160a01b03868116600483015292851660248201529116915063a9059cbb90604401602060405180830381600087803b158015611de357600080fd5b505af1158015611df7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e1b9190615647565b611e3857604051631e9acf1760e31b815260040160405180910390fd5b50506001600255565b611e496138f1565b604080518082018252600091611e78919084906002908390839080828437600092019190915250612f25915050565b6000818152600b60205260409020549091506001600160a01b031615611ecd576040517f4a0b8fa700000000000000000000000000000000000000000000000000000000815260048101829052602401610bef565b6000818152600b6020908152604080832080546001600160a01b0319166001600160a01b038816908117909155600c805460018101825594527fdf6966c971051c3d54ec59162606531493a51404a002842f56009d7e5cf4a8c7909301849055518381527fe729ae16526293f74ade739043022254f1489f616295a25bf72dfb4511ed73b89101610d82565b6001546001600160a01b03163314611fb35760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610bef565b60008054336001600160a01b0319808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b600060028054141561204c5760405162461bcd60e51b815260206004820152601f60248201526000805160206160b18339815191526044820152606401610bef565b600280556020808301356000908152600690915260409020546001600160a01b031661208b57604051630fb532db60e11b815260040160405180910390fd5b33600090815260056020908152604080832085830135845290915290205467ffffffffffffffff16806120dd576040516379bfd40160e01b815260208401356004820152336024820152604401610bef565b600f5461ffff166120f46060850160408601615813565b61ffff161080612117575060c86121116060850160408601615813565b61ffff16115b1561215d5761212c6060840160408501615813565b600f5460405163539c34bb60e11b815261ffff92831660048201529116602482015260c86044820152606401610bef565b600f5462010000900463ffffffff1661217c60808501606086016158f1565b63ffffffff1611156121e55761219860808401606085016158f1565b600f546040517ff5d7e01e00000000000000000000000000000000000000000000000000000000815263ffffffff9283166004820152620100009091049091166024820152604401610bef565b6101f46121f860a08501608086016158f1565b63ffffffff1611156122575761221460a08401608085016158f1565b6040517f47386bec00000000000000000000000000000000000000000000000000000000815263ffffffff90911660048201526101f46024820152604401610bef565b6000612264826001615ed8565b6040805186356020808301829052338385015280890135606084015267ffffffffffffffff85166080808501919091528451808503909101815260a0808501865281519183019190912060c085019390935260e0808501849052855180860390910181526101009094019094528251920191909120929350906000906122f09061111590890189615dfb565b905060006122fd82613f86565b9050836123086142c5565b60208a013561231d60808c0160608d016158f1565b61232d60a08d0160808e016158f1565b33866040516020016123459796959493929190615c68565b60405160208183030381529060405280519060200120600d600086815260200190815260200160002081905550336001600160a01b0316886020013589600001357feb0e3652e0f44f417695e6e90f2f42c99b65cd7169074c5a654b16b9748c3a4e87878d60400160208101906123bc9190615813565b8e60600160208101906123cf91906158f1565b8f60800160208101906123e291906158f1565b896040516123f596959493929190615c29565b60405180910390a4505033600090815260056020908152604080832098820135835297905295909520805467ffffffffffffffff191667ffffffffffffffff939093169290921790915560016002559392505050565b600060028054141561248d5760405162461bcd60e51b815260206004820152601f60248201526000805160206160b18339815191526044820152606401610bef565b6002805560003361249f600143615f59565b600854604051606093841b6bffffffffffffffffffffffff199081166020830152924060348201523090931b909116605483015260c01b7fffffffffffffffff00000000000000000000000000000000000000000000000016606882015260700160408051601f1981840301815291905280516020909101206008805491925067ffffffffffffffff90911690600061253783615fe3565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055505060008067ffffffffffffffff81111561257957612579616077565b6040519080825280602002602001820160405280156125a2578160200160208202803683370190505b506040805160608082018352600080835260208084018281528486018381528984526007835286842095518654925191516001600160601b039182167fffffffffffffffff00000000000000000000000000000000000000000000000090941693909317600160601b91909216021777ffffffffffffffffffffffffffffffffffffffffffffffff16600160c01b67ffffffffffffffff9092169190910217909355835191820184523382528183018181528285018681528883526006855294909120825181546001600160a01b03199081166001600160a01b0392831617835592516001830180549094169116179091559251805194955090936126ad92600285019201906152f9565b50506040513381528391507f1d3015d7ba850fa198dc7b1a3f5d42779313a681035f77c8c03764c61005518d9060200160405180910390a2509050600160025590565b6002805414156127305760405162461bcd60e51b815260206004820152601f60248201526000805160206160b18339815191526044820152606401610bef565b600280556003546001600160a01b03163314612778576040517f44b0e3c300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b602081146127b2576040517f8129bbcd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006127c082840184615664565b6000818152600660205260409020549091506001600160a01b03166127f857604051630fb532db60e11b815260040160405180910390fd5b600081815260076020526040812080546001600160601b03169186919061281f8385615f04565b92506101000a8154816001600160601b0302191690836001600160601b03160217905550846008808282829054906101000a90046001600160601b03166128669190615f04565b92506101000a8154816001600160601b0302191690836001600160601b03160217905550817f1ced9348ff549fceab2ac57cd3a9de38edaaab274b725ee82c23e8fc8c4eec7a8287846128b99190615ec0565b604080519283526020830191909152015b60405180910390a25050600160025550505050565b6128e76138f1565b6008544790600160a01b90046001600160601b031681811115612927576040516354ced18160e11b81526004810182905260248101839052604401610bef565b818110156129f657600061293b8284615f59565b90506000846001600160a01b03168260405160006040518083038185875af1925050503d806000811461298a576040519150601f19603f3d011682016040523d82523d6000602084013e61298f565b606091505b50509050806129b157604051630dcf35db60e41b815260040160405180910390fd5b604080516001600160a01b0387168152602081018490527f879c9ea2b9d5345b84ccd12610b032602808517cebdb795007f3dcb4df377317910160405180910390a150505b505050565b612a036138f1565b6000818152600660205260409020546001600160a01b0316612a3857604051630fb532db60e11b815260040160405180910390fd5b600081815260066020526040902054610bf89082906001600160a01b031661394d565b612a636138f1565b60c861ffff87161115612a9d5760405163539c34bb60e11b815261ffff871660048201819052602482015260c86044820152606401610bef565b60008213612ac1576040516321ea67b360e11b815260048101839052602401610bef565b604080516080808201835261ffff891680835263ffffffff89811660208086018290528a83168688018190528a84166060978801819052600f805465ffffffffffff19168717620100008602176dffffffffffffffff0000000000001916660100000000000084026dffffffff000000000000000000001916176a010000000000000000000083021790558951601080548c86015192881667ffffffffffffffff1990911617640100000000928816929092029190911790819055600e8c9055895196875286840194909452978501529483019590955291810186905283821660a08201529290911c1660c08201527f777357bb93f63d088f18112d3dba38457aec633eb8f1341e1d418380ad328e789060e00160405180910390a1505050505050565b600280541415612c255760405162461bcd60e51b815260206004820152601f60248201526000805160206160b18339815191526044820152606401610bef565b600280556000818152600660205260409020546001600160a01b0316612c5e57604051630fb532db60e11b815260040160405180910390fd5b6000818152600660205260409020600101546001600160a01b03163314612cd057600081815260066020526040908190206001015490517fd084e9750000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401610bef565b6000818152600660209081526040918290208054336001600160a01b0319808316821784556001909301805490931690925583516001600160a01b0390911680825292810191909152909183917fd4114ab6e9af9f597c52041f32d62dc57c5c4e4c0d4427006069635e216c938691015b60405180910390a250506001600255565b60008281526006602052604090205482906001600160a01b031680612d8a57604051630fb532db60e11b815260040160405180910390fd5b336001600160a01b03821614612dbe57604051636c51fda960e11b81526001600160a01b0382166004820152602401610bef565b600280541415612dfe5760405162461bcd60e51b815260206004820152601f60248201526000805160206160b18339815191526044820152606401610bef565b60028080556000858152600660205260409020015460641415612e4d576040517f05a48e0f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038316600090815260056020908152604080832087845290915290205467ffffffffffffffff1615612e8557610e70565b6001600160a01b03831660008181526005602090815260408083208884528252808320805467ffffffffffffffff19166001908117909155600683528184206002018054918201815584529282902090920180546001600160a01b03191684179055905191825285917f1e980d04aa7648e205713e5e8ea3808672ac163d10936d36f91b2c88ac1575e191015b60405180910390a2505060016002555050565b600081604051602001612f389190615ad2565b604051602081830303815290604052805190602001209050919050565b60008281526006602052604090205482906001600160a01b031680612f8d57604051630fb532db60e11b815260040160405180910390fd5b336001600160a01b03821614612fc157604051636c51fda960e11b81526001600160a01b0382166004820152602401610bef565b6002805414156130015760405162461bcd60e51b815260206004820152601f60248201526000805160206160b18339815191526044820152606401610bef565b6002805561300e84611798565b1561302c57604051631685ecdd60e31b815260040160405180910390fd5b6001600160a01b038316600090815260056020908152604080832087845290915290205467ffffffffffffffff16613089576040516379bfd40160e01b8152600481018590526001600160a01b0384166024820152604401610bef565b6000848152600660209081526040808320600201805482518185028101850190935280835291929091908301828280156130ec57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116130ce575b505050505090506000600182516131039190615f59565b905060005b825181101561320f57856001600160a01b031683828151811061312d5761312d616061565b60200260200101516001600160a01b031614156131fd57600083838151811061315857613158616061565b6020026020010151905080600660008a8152602001908152602001600020600201838154811061318a5761318a616061565b600091825260208083209190910180546001600160a01b0319166001600160a01b0394909416939093179092558981526006909152604090206002018054806131d5576131d561604b565b600082815260209020810160001990810180546001600160a01b03191690550190555061320f565b8061320781615fc8565b915050613108565b506001600160a01b03851660008181526005602090815260408083208a8452825291829020805467ffffffffffffffff19169055905191825287917f32158c6058347c1601b2d12bc696ac6901d8a9a9aa3ba10c27ab0a983e8425a791016128ca565b600c818154811061328257600080fd5b600091825260209091200154905081565b60008281526006602052604090205482906001600160a01b0316806132cb57604051630fb532db60e11b815260040160405180910390fd5b336001600160a01b038216146132ff57604051636c51fda960e11b81526001600160a01b0382166004820152602401610bef565b60028054141561333f5760405162461bcd60e51b815260206004820152601f60248201526000805160206160b18339815191526044820152606401610bef565b600280556000848152600660205260409020600101546001600160a01b03848116911614610e705760008481526006602090815260409182902060010180546001600160a01b0319166001600160a01b03871690811790915582513381529182015285917f21a4dad170a6bf476c31bbcf4a16628295b0e450672eec25d7c93308e05344a19101612f12565b60008181526006602052604081205481908190819060609081906001600160a01b031661340b57604051630fb532db60e11b815260040160405180910390fd5b6000878152600660209081526040918290206002018054835181840281018401909452808452909183018282801561346c57602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161344e575b50505050509150815167ffffffffffffffff81111561348d5761348d616077565b6040519080825280602002602001820160405280156134b6578160200160208202803683370190505b50905060005b825181101561356657600560008483815181106134db576134db616061565b60200260200101516001600160a01b03166001600160a01b03168152602001908152602001600020600089815260200190815260200160002060009054906101000a900467ffffffffffffffff1682828151811061353b5761353b616061565b67ffffffffffffffff909216602092830291909101909101528061355e81615fc8565b9150506134bc565b506000968752600760209081526040808920546006909252909720546001600160601b0380891699600160601b8a0490911698600160c01b900467ffffffffffffffff1697506001600160a01b0390911695509193509150565b6135c86138f1565b6003546040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526000916001600160a01b0316906370a082319060240160206040518083038186803b15801561362557600080fd5b505afa158015613639573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061365d919061567d565b6008549091506801000000000000000090046001600160601b0316818111156136a3576040516354ced18160e11b81526004810182905260248101839052604401610bef565b818110156129f65760006136b78284615f59565b60035460405163a9059cbb60e01b81526001600160a01b0387811660048301526024820184905292935091169063a9059cbb90604401602060405180830381600087803b15801561370757600080fd5b505af115801561371b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061373f9190615647565b50604080516001600160a01b0386168152602081018390527f59bfc682b673f8cbf945f1e454df9334834abf7dfe7f92237ca29ecb9b436600910160405180910390a150505050565b6002805414156137c85760405162461bcd60e51b815260206004820152601f60248201526000805160206160b18339815191526044820152606401610bef565b600280556000818152600660205260409020546001600160a01b031661380157604051630fb532db60e11b815260040160405180910390fd5b60008181526007602052604090208054600160601b90046001600160601b0316903490600c6138308385615f04565b92506101000a8154816001600160601b0302191690836001600160601b0316021790555034600860148282829054906101000a90046001600160601b03166138789190615f04565b92506101000a8154816001600160601b0302191690836001600160601b03160217905550817f3f1ddc3ab1bdb39001ad76ca51a0e6f57ce6627c69f251d1de41622847721cde8234846138cb9190615ec0565b60408051928352602083019190915201612d41565b6138e86138f1565b610bf88161435e565b6000546001600160a01b0316331461394b5760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610bef565b565b60008061395984614076565b60035460405163a9059cbb60e01b81526001600160a01b0387811660048301526001600160601b0385166024830152939550919350919091169063a9059cbb90604401602060405180830381600087803b1580156139b657600080fd5b505af11580156139ca573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906139ee9190615647565b613a0b57604051631e9acf1760e31b815260040160405180910390fd5b6000836001600160a01b0316826001600160601b031660405160006040518083038185875af1925050503d8060008114613a61576040519150601f19603f3d011682016040523d82523d6000602084013e613a66565b606091505b5050905080613a8857604051630dcf35db60e41b815260040160405180910390fd5b604080516001600160a01b03861681526001600160601b038581166020830152841681830152905186917f8c74ce8b8cf87f5eb001275c8be27eb34ea2b62bfab6814fcc62192bb63e81c4919081900360600190a25050505050565b604080516060810182526000808252602082018190528183018190528251808401845291929091613b2d9186906002908390839080828437600092019190915250612f25915050565b6000818152600b60205260409020549091506001600160a01b031680613b6957604051631dfd6e1360e21b815260048101839052602401610bef565b6000828660c00135604051602001613b8b929190918252602082015260400190565b60408051601f1981840301815291815281516020928301206000818152600d90935291205490915080613bea576040517f3688124a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b81613bf8602088018861590c565b6020880135613c0d60608a0160408b016158f1565b613c1d60808b0160608c016158f1565b613c2d60a08c0160808d016154ce565b613c3a60a08d018d615dfb565b604051602001613c51989796959493929190615cb2565b604051602081830303815290604052805190602001208114613c9f576040517fd529142c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000613cb6613cb1602089018961590c565b614408565b905080613dc4576001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001663e9413d38613cf960208a018a61590c565b6040516001600160e01b031960e084901b16815267ffffffffffffffff909116600482015260240160206040518083038186803b158015613d3957600080fd5b505afa158015613d4d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613d71919061567d565b905080613dc457613d85602088018861590c565b6040517f175dadad00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610bef565b6040805160c08a0135602080830191909152818301849052825180830384018152606090920190925280519101206000613e0c613e06368c90038c018c61573a565b8361450f565b90506040518060600160405280888152602001868152602001828152509750505050505050505b92915050565b60005a611388811015613e4b57600080fd5b611388810390508460408204820311613e6357600080fd5b50823b613e6f57600080fd5b60008083516020850160008789f190505b9392505050565b60408051602081019091526000815281613eb05750604080516020810190915260008152613e33565b7f92fd133800000000000000000000000000000000000000000000000000000000613edb8385615f98565b6001600160e01b03191614613f1c576040517f5247fdce00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b613f298260048186615e96565b810190613e809190615696565b60008115613f6457601054613f5d9086908690640100000000900463ffffffff168661457a565b9050613f7e565b601054613f7b908690869063ffffffff16866145e4565b90505b949350505050565b60607f92fd13387c7fe7befbc38d303d6468778fb9731bc4583f17d92989c6fcfdeaaa82604051602401613fbf91511515815260200190565b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff166001600160e01b03199093169290921790915292915050565b6000805b60115481101561406d57826001600160a01b03166011828154811061403757614037616061565b6000918252602090912001546001600160a01b0316141561405b5750600192915050565b8061406581615fc8565b915050614010565b50600092915050565b6000818152600660209081526040808320815160608101835281546001600160a01b0390811682526001830154168185015260028201805484518187028101870186528181528796879694959486019391929083018282801561410257602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116140e4575b505050919092525050506000858152600760209081526040808320815160608101835290546001600160601b03808216808452600160601b8304909116948301859052600160c01b90910467ffffffffffffffff16928201929092529096509094509192505b8260400151518110156141e057600560008460400151838151811061418f5761418f616061565b6020908102919091018101516001600160a01b0316825281810192909252604090810160009081208982529092529020805467ffffffffffffffff19169055806141d881615fc8565b915050614168565b50600085815260066020526040812080546001600160a01b03199081168255600182018054909116905590614218600283018261535e565b50506000858152600760205260408120556008805485919081906142529084906801000000000000000090046001600160601b0316615f70565b92506101000a8154816001600160601b0302191690836001600160601b0316021790555082600860148282829054906101000a90046001600160601b031661429a9190615f70565b92506101000a8154816001600160601b0302191690836001600160601b031602179055505050915091565b60004661a4b18114806142da575062066eed81145b156143575760646001600160a01b031663a3b1b31d6040518163ffffffff1660e01b815260040160206040518083038186803b15801561431957600080fd5b505afa15801561432d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614351919061567d565b91505090565b4391505090565b6001600160a01b0381163314156143b75760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610bef565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60004661a4b181148061441d575062066eed81145b156144ff576101008367ffffffffffffffff166144386142c5565b6144429190615f59565b118061445f57506144516142c5565b8367ffffffffffffffff1610155b1561446d5750600092915050565b6040517f2b407a8200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff84166004820152606490632b407a829060240160206040518083038186803b1580156144c757600080fd5b505afa1580156144db573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613e80919061567d565b505067ffffffffffffffff164090565b60006145438360000151846020015185604001518660600151868860a001518960c001518a60e001518b61010001516146eb565b6003836020015160405160200161455b929190615bc7565b60408051601f1981840301815291905280516020909101209392505050565b600080614585614926565b905060005a6145948888615ec0565b61459e9190615f59565b6145a89085615f3a565b905060006145c163ffffffff871664e8d4a51000615f3a565b9050826145ce8284615ec0565b6145d89190615ec0565b98975050505050505050565b6000806145ef614982565b905060008113614615576040516321ea67b360e11b815260048101829052602401610bef565b600061461f614926565b9050600082825a6146308b8b615ec0565b61463a9190615f59565b6146449088615f3a565b61464e9190615ec0565b61466090670de0b6b3a7640000615f3a565b61466a9190615f26565b9050600061468363ffffffff881664e8d4a51000615f3a565b905061469b816b033b2e3c9fd0803ce8000000615f59565b8211156146d4576040517fe80fa38100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6146de8183615ec0565b9998505050505050505050565b6146f489614a69565b6147405760405162461bcd60e51b815260206004820152601a60248201527f7075626c6963206b6579206973206e6f74206f6e2063757276650000000000006044820152606401610bef565b61474988614a69565b6147955760405162461bcd60e51b815260206004820152601560248201527f67616d6d61206973206e6f74206f6e20637572766500000000000000000000006044820152606401610bef565b61479e83614a69565b6147ea5760405162461bcd60e51b815260206004820152601d60248201527f6347616d6d615769746e657373206973206e6f74206f6e2063757276650000006044820152606401610bef565b6147f382614a69565b61483f5760405162461bcd60e51b815260206004820152601c60248201527f73486173685769746e657373206973206e6f74206f6e206375727665000000006044820152606401610bef565b61484b878a8887614b42565b6148975760405162461bcd60e51b815260206004820152601960248201527f6164647228632a706b2b732a6729213d5f755769746e657373000000000000006044820152606401610bef565b60006148a38a87614c77565b905060006148b6898b878b868989614cdb565b905060006148c7838d8d8a86614dfb565b9050808a146149185760405162461bcd60e51b815260206004820152600d60248201527f696e76616c69642070726f6f66000000000000000000000000000000000000006044820152606401610bef565b505050505050505050505050565b60004661a4b181148061493b575062066eed81145b1561497a57606c6001600160a01b031663c6f7de0e6040518163ffffffff1660e01b815260040160206040518083038186803b15801561431957600080fd5b600091505090565b600f5460048054604080517ffeaf968c00000000000000000000000000000000000000000000000000000000815290516000946601000000000000900463ffffffff169384151593869384936001600160a01b039092169263feaf968c928282019260a09290829003018186803b1580156149fc57600080fd5b505afa158015614a10573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614a349190615936565b509450909250849150508015614a585750614a4f8242615f59565b8463ffffffff16105b15613f7e5750600e54949350505050565b80516000906401000003d01911614ac25760405162461bcd60e51b815260206004820152601260248201527f696e76616c696420782d6f7264696e61746500000000000000000000000000006044820152606401610bef565b60208201516401000003d01911614b1b5760405162461bcd60e51b815260206004820152601260248201527f696e76616c696420792d6f7264696e61746500000000000000000000000000006044820152606401610bef565b60208201516401000003d019908009614b3b8360005b6020020151614e3b565b1492915050565b60006001600160a01b038216614b9a5760405162461bcd60e51b815260206004820152600b60248201527f626164207769746e6573730000000000000000000000000000000000000000006044820152606401610bef565b602084015160009060011615614bb157601c614bb4565b601b5b9050600070014551231950b75fc4402da1732fc9bebe1985876000602002015109865170014551231950b75fc4402da1732fc9bebe19918203925060009190890987516040805160008082526020820180845287905260ff88169282019290925260608101929092526080820183905291925060019060a0016020604051602081039080840390855afa158015614c4f573d6000803e3d6000fd5b5050604051601f1901516001600160a01b039081169088161495505050505050949350505050565b614c7f61537c565b614cac60018484604051602001614c9893929190615ab1565b604051602081830303815290604052614e5f565b90505b614cb881614a69565b613e33578051604080516020810192909252614cd49101614c98565b9050614caf565b614ce361537c565b825186516401000003d0199081900691061415614d425760405162461bcd60e51b815260206004820152601e60248201527f706f696e747320696e2073756d206d7573742062652064697374696e637400006044820152606401610bef565b614d4d878988614eae565b614d995760405162461bcd60e51b815260206004820152601660248201527f4669727374206d756c20636865636b206661696c6564000000000000000000006044820152606401610bef565b614da4848685614eae565b614df05760405162461bcd60e51b815260206004820152601760248201527f5365636f6e64206d756c20636865636b206661696c65640000000000000000006044820152606401610bef565b6145d8868484614fe8565b600060028686868587604051602001614e1996959493929190615a52565b60408051601f1981840301815291905280516020909101209695505050505050565b6000806401000003d01980848509840990506401000003d019600782089392505050565b614e6761537c565b614e70826150af565b8152614e85614e80826000614b31565b6150ea565b602082018190526002900660011415614ea9576020810180516401000003d0190390525b919050565b600082614efd5760405162461bcd60e51b815260206004820152600b60248201527f7a65726f207363616c61720000000000000000000000000000000000000000006044820152606401610bef565b83516020850151600090614f139060029061600b565b15614f1f57601c614f22565b601b5b9050600070014551231950b75fc4402da1732fc9bebe198387096040805160008082526020820180845281905260ff86169282019290925260608101869052608081018390529192509060019060a0016020604051602081039080840390855afa158015614f94573d6000803e3d6000fd5b505050602060405103519050600086604051602001614fb39190615a40565b60408051601f1981840301815291905280516020909101206001600160a01b0392831692169190911498975050505050505050565b614ff061537c565b8351602080860151855191860151600093849384936150119390919061510a565b919450925090506401000003d0198582096001146150715760405162461bcd60e51b815260206004820152601960248201527f696e765a206d75737420626520696e7665727365206f66207a000000000000006044820152606401610bef565b60405180604001604052806401000003d0198061509057615090616035565b87860981526020016401000003d0198785099052979650505050505050565b805160208201205b6401000003d0198110614ea9576040805160208082019390935281518082038401815290820190915280519101206150b7565b6000613e338260026151036401000003d0196001615ec0565b901c6151ea565b60008080600180826401000003d019896401000003d019038808905060006401000003d0198b6401000003d019038a089050600061514a8383858561528c565b909850905061515b88828e886152b0565b909850905061516c88828c876152b0565b9098509050600061517f8d878b856152b0565b90985090506151908882868661528c565b90985090506151a188828e896152b0565b90985090508181146151d6576401000003d019818a0998506401000003d01982890997506401000003d01981830996506151da565b8196505b5050505050509450945094915050565b6000806151f561539a565b6020808252818101819052604082015260608101859052608081018490526401000003d01960a08201526152276153b8565b60208160c0846005600019fa9250826152825760405162461bcd60e51b815260206004820152601260248201527f6269674d6f64457870206661696c7572652100000000000000000000000000006044820152606401610bef565b5195945050505050565b6000806401000003d0198487096401000003d0198487099097909650945050505050565b600080806401000003d019878509905060006401000003d01987876401000003d019030990506401000003d0198183086401000003d01986890990999098509650505050505050565b82805482825590600052602060002090810192821561534e579160200282015b8281111561534e57825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190615319565b5061535a9291506153d6565b5090565b5080546000825590600052602060002090810190610bf891906153d6565b60405180604001604052806002906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b5b8082111561535a57600081556001016153d7565b8035614ea98161608d565b8060408101831015613e3357600080fd5b600082601f83011261541857600080fd5b615420615e49565b80838560408601111561543257600080fd5b60005b6002811015615454578135845260209384019390910190600101615435565b509095945050505050565b600060c0828403121561547157600080fd5b50919050565b803561ffff81168114614ea957600080fd5b803563ffffffff81168114614ea957600080fd5b805169ffffffffffffffffffff81168114614ea957600080fd5b80356001600160601b0381168114614ea957600080fd5b6000602082840312156154e057600080fd5b8135613e808161608d565b600080604083850312156154fe57600080fd5b82356155098161608d565b9150615517602084016154b7565b90509250929050565b6000806040838503121561553357600080fd5b823561553e8161608d565b9150602083013561554e8161608d565b809150509250929050565b6000806060838503121561556c57600080fd5b82356155778161608d565b915061551784602085016153f6565b6000806000806060858703121561559c57600080fd5b84356155a78161608d565b935060208501359250604085013567ffffffffffffffff808211156155cb57600080fd5b818701915087601f8301126155df57600080fd5b8135818111156155ee57600080fd5b88602082850101111561560057600080fd5b95989497505060200194505050565b60006040828403121561562157600080fd5b613e8083836153f6565b60006040828403121561563d57600080fd5b613e808383615407565b60006020828403121561565957600080fd5b8151613e80816160a2565b60006020828403121561567657600080fd5b5035919050565b60006020828403121561568f57600080fd5b5051919050565b6000602082840312156156a857600080fd5b6040516020810181811067ffffffffffffffff821117156156cb576156cb616077565b60405282356156d9816160a2565b81529392505050565b6000808284036101c08112156156f757600080fd5b6101a08082121561570757600080fd5b849350830135905067ffffffffffffffff81111561572457600080fd5b6157308582860161545f565b9150509250929050565b60006101a0828403121561574d57600080fd5b615755615e72565b61575f8484615407565b815261576e8460408501615407565b60208201526080830135604082015260a0830135606082015260c0830135608082015261579d60e084016153eb565b60a08201526101006157b185828601615407565b60c08301526157c4856101408601615407565b60e083015261018084013581830152508091505092915050565b6000602082840312156157f057600080fd5b813567ffffffffffffffff81111561580757600080fd5b613f7e8482850161545f565b60006020828403121561582557600080fd5b613e8082615477565b60008060008060008086880360e081121561584857600080fd5b61585188615477565b965061585f60208901615489565b955061586d60408901615489565b945061587b60608901615489565b9350608088013592506040609f198201121561589657600080fd5b5061589f615e49565b6158ab60a08901615489565b81526158b960c08901615489565b6020820152809150509295509295509295565b600080604083850312156158df57600080fd5b82359150602083013561554e8161608d565b60006020828403121561590357600080fd5b613e8082615489565b60006020828403121561591e57600080fd5b813567ffffffffffffffff81168114613e8057600080fd5b600080600080600060a0868803121561594e57600080fd5b6159578661549d565b945060208601519350604086015192506060860151915061597a6080870161549d565b90509295509295909350565b600081518084526020808501945080840160005b838110156159bf5781516001600160a01b03168752958201959082019060010161599a565b509495945050505050565b8060005b60028110156159ed5781518452602093840193909101906001016159ce565b50505050565b6000815180845260005b81811015615a19576020818501810151868301820152016159fd565b81811115615a2b576000602083870101525b50601f01601f19169290920160200192915050565b615a4a81836159ca565b604001919050565b868152615a6260208201876159ca565b615a6f60608201866159ca565b615a7c60a08201856159ca565b615a8960e08201846159ca565b60609190911b6bffffffffffffffffffffffff19166101208201526101340195945050505050565b838152615ac160208201846159ca565b606081019190915260800192915050565b60408101613e3382846159ca565b602081526000613e8060208301846159f3565b6020815260ff8251166020820152602082015160408201526001600160a01b0360408301511660608201526000606083015160c06080840152615b3960e0840182615986565b905060808401516001600160601b0380821660a08601528060a08701511660c086015250508091505092915050565b60006060820161ffff86168352602063ffffffff86168185015260606040850152818551808452608086019150828701935060005b81811015615bb957845183529383019391830191600101615b9d565b509098975050505050505050565b82815260608101613e8060208301846159ca565b6000604082018483526020604081850152818551808452606086019150828701935060005b81811015615c1c57845183529383019391830191600101615c00565b5090979650505050505050565b86815285602082015261ffff85166040820152600063ffffffff808616606084015280851660808401525060c060a08301526145d860c08301846159f3565b878152866020820152856040820152600063ffffffff80871660608401528086166080840152506001600160a01b03841660a083015260e060c08301526146de60e08301846159f3565b88815267ffffffffffffffff88166020820152866040820152600063ffffffff80881660608401528087166080840152506001600160a01b03851660a083015260e060c08301528260e08301526101008385828501376000838501820152601f909301601f191690910190910198975050505050505050565b8481526001600160601b0384166020820152608060408201526000615d5360808301856159f3565b9050821515606083015295945050505050565b60006001600160601b03808916835260208189168185015267ffffffffffffffff915081881660408501526001600160a01b038716606085015260c06080850152615db460c0850187615986565b84810360a086015285518082528287019183019060005b81811015615de9578351861683529284019291840191600101615dcb565b50909c9b505050505050505050505050565b6000808335601e19843603018112615e1257600080fd5b83018035915067ffffffffffffffff821115615e2d57600080fd5b602001915036819003821315615e4257600080fd5b9250929050565b6040805190810167ffffffffffffffff81118282101715615e6c57615e6c616077565b60405290565b604051610120810167ffffffffffffffff81118282101715615e6c57615e6c616077565b60008085851115615ea657600080fd5b83861115615eb357600080fd5b5050820193919092039150565b60008219821115615ed357615ed361601f565b500190565b600067ffffffffffffffff808316818516808303821115615efb57615efb61601f565b01949350505050565b60006001600160601b03808316818516808303821115615efb57615efb61601f565b600082615f3557615f35616035565b500490565b6000816000190483118215151615615f5457615f5461601f565b500290565b600082821015615f6b57615f6b61601f565b500390565b60006001600160601b0383811690831681811015615f9057615f9061601f565b039392505050565b6001600160e01b03198135818116916004851015615fc05780818660040360031b1b83161692505b505092915050565b6000600019821415615fdc57615fdc61601f565b5060010190565b600067ffffffffffffffff808316818114156160015761600161601f565b6001019392505050565b60008261601a5761601a616035565b500690565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052601260045260246000fd5b634e487b7160e01b600052603160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b0381168114610bf857600080fd5b8015158114610bf857600080fdfe5265656e7472616e637947756172643a207265656e7472616e742063616c6c00a164736f6c6343000806000a",
}

var VRFCoordinatorV2PlusABI = VRFCoordinatorV2PlusMetaData.ABI

var VRFCoordinatorV2PlusBin = VRFCoordinatorV2PlusMetaData.Bin

func DeployVRFCoordinatorV2Plus(auth *bind.TransactOpts, backend bind.ContractBackend, blockhashStore common.Address) (common.Address, *types.Transaction, *VRFCoordinatorV2Plus, error) {
	parsed, err := VRFCoordinatorV2PlusMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VRFCoordinatorV2PlusBin), backend, blockhashStore)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VRFCoordinatorV2Plus{VRFCoordinatorV2PlusCaller: VRFCoordinatorV2PlusCaller{contract: contract}, VRFCoordinatorV2PlusTransactor: VRFCoordinatorV2PlusTransactor{contract: contract}, VRFCoordinatorV2PlusFilterer: VRFCoordinatorV2PlusFilterer{contract: contract}}, nil
}

type VRFCoordinatorV2Plus struct {
	address common.Address
	abi     abi.ABI
	VRFCoordinatorV2PlusCaller
	VRFCoordinatorV2PlusTransactor
	VRFCoordinatorV2PlusFilterer
}

type VRFCoordinatorV2PlusCaller struct {
	contract *bind.BoundContract
}

type VRFCoordinatorV2PlusTransactor struct {
	contract *bind.BoundContract
}

type VRFCoordinatorV2PlusFilterer struct {
	contract *bind.BoundContract
}

type VRFCoordinatorV2PlusSession struct {
	Contract     *VRFCoordinatorV2Plus
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type VRFCoordinatorV2PlusCallerSession struct {
	Contract *VRFCoordinatorV2PlusCaller
	CallOpts bind.CallOpts
}

type VRFCoordinatorV2PlusTransactorSession struct {
	Contract     *VRFCoordinatorV2PlusTransactor
	TransactOpts bind.TransactOpts
}

type VRFCoordinatorV2PlusRaw struct {
	Contract *VRFCoordinatorV2Plus
}

type VRFCoordinatorV2PlusCallerRaw struct {
	Contract *VRFCoordinatorV2PlusCaller
}

type VRFCoordinatorV2PlusTransactorRaw struct {
	Contract *VRFCoordinatorV2PlusTransactor
}

func NewVRFCoordinatorV2Plus(address common.Address, backend bind.ContractBackend) (*VRFCoordinatorV2Plus, error) {
	abi, err := abi.JSON(strings.NewReader(VRFCoordinatorV2PlusABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindVRFCoordinatorV2Plus(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2Plus{address: address, abi: abi, VRFCoordinatorV2PlusCaller: VRFCoordinatorV2PlusCaller{contract: contract}, VRFCoordinatorV2PlusTransactor: VRFCoordinatorV2PlusTransactor{contract: contract}, VRFCoordinatorV2PlusFilterer: VRFCoordinatorV2PlusFilterer{contract: contract}}, nil
}

func NewVRFCoordinatorV2PlusCaller(address common.Address, caller bind.ContractCaller) (*VRFCoordinatorV2PlusCaller, error) {
	contract, err := bindVRFCoordinatorV2Plus(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusCaller{contract: contract}, nil
}

func NewVRFCoordinatorV2PlusTransactor(address common.Address, transactor bind.ContractTransactor) (*VRFCoordinatorV2PlusTransactor, error) {
	contract, err := bindVRFCoordinatorV2Plus(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusTransactor{contract: contract}, nil
}

func NewVRFCoordinatorV2PlusFilterer(address common.Address, filterer bind.ContractFilterer) (*VRFCoordinatorV2PlusFilterer, error) {
	contract, err := bindVRFCoordinatorV2Plus(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusFilterer{contract: contract}, nil
}

func bindVRFCoordinatorV2Plus(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VRFCoordinatorV2PlusMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFCoordinatorV2Plus.Contract.VRFCoordinatorV2PlusCaller.contract.Call(opts, result, method, params...)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.VRFCoordinatorV2PlusTransactor.contract.Transfer(opts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.VRFCoordinatorV2PlusTransactor.contract.Transact(opts, method, params...)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFCoordinatorV2Plus.Contract.contract.Call(opts, result, method, params...)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.contract.Transfer(opts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.contract.Transact(opts, method, params...)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCaller) BLOCKHASHSTORE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFCoordinatorV2Plus.contract.Call(opts, &out, "BLOCKHASH_STORE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) BLOCKHASHSTORE() (common.Address, error) {
	return _VRFCoordinatorV2Plus.Contract.BLOCKHASHSTORE(&_VRFCoordinatorV2Plus.CallOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCallerSession) BLOCKHASHSTORE() (common.Address, error) {
	return _VRFCoordinatorV2Plus.Contract.BLOCKHASHSTORE(&_VRFCoordinatorV2Plus.CallOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCaller) LINK(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFCoordinatorV2Plus.contract.Call(opts, &out, "LINK")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) LINK() (common.Address, error) {
	return _VRFCoordinatorV2Plus.Contract.LINK(&_VRFCoordinatorV2Plus.CallOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCallerSession) LINK() (common.Address, error) {
	return _VRFCoordinatorV2Plus.Contract.LINK(&_VRFCoordinatorV2Plus.CallOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCaller) LINKETHFEED(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFCoordinatorV2Plus.contract.Call(opts, &out, "LINK_ETH_FEED")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) LINKETHFEED() (common.Address, error) {
	return _VRFCoordinatorV2Plus.Contract.LINKETHFEED(&_VRFCoordinatorV2Plus.CallOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCallerSession) LINKETHFEED() (common.Address, error) {
	return _VRFCoordinatorV2Plus.Contract.LINKETHFEED(&_VRFCoordinatorV2Plus.CallOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCaller) MAXCONSUMERS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _VRFCoordinatorV2Plus.contract.Call(opts, &out, "MAX_CONSUMERS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) MAXCONSUMERS() (uint16, error) {
	return _VRFCoordinatorV2Plus.Contract.MAXCONSUMERS(&_VRFCoordinatorV2Plus.CallOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCallerSession) MAXCONSUMERS() (uint16, error) {
	return _VRFCoordinatorV2Plus.Contract.MAXCONSUMERS(&_VRFCoordinatorV2Plus.CallOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCaller) MAXNUMWORDS(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _VRFCoordinatorV2Plus.contract.Call(opts, &out, "MAX_NUM_WORDS")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) MAXNUMWORDS() (uint32, error) {
	return _VRFCoordinatorV2Plus.Contract.MAXNUMWORDS(&_VRFCoordinatorV2Plus.CallOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCallerSession) MAXNUMWORDS() (uint32, error) {
	return _VRFCoordinatorV2Plus.Contract.MAXNUMWORDS(&_VRFCoordinatorV2Plus.CallOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCaller) MAXREQUESTCONFIRMATIONS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _VRFCoordinatorV2Plus.contract.Call(opts, &out, "MAX_REQUEST_CONFIRMATIONS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) MAXREQUESTCONFIRMATIONS() (uint16, error) {
	return _VRFCoordinatorV2Plus.Contract.MAXREQUESTCONFIRMATIONS(&_VRFCoordinatorV2Plus.CallOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCallerSession) MAXREQUESTCONFIRMATIONS() (uint16, error) {
	return _VRFCoordinatorV2Plus.Contract.MAXREQUESTCONFIRMATIONS(&_VRFCoordinatorV2Plus.CallOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCaller) GetRequestConfig(opts *bind.CallOpts) (uint16, uint32, [][32]byte, error) {
	var out []interface{}
	err := _VRFCoordinatorV2Plus.contract.Call(opts, &out, "getRequestConfig")

	if err != nil {
		return *new(uint16), *new(uint32), *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)
	out1 := *abi.ConvertType(out[1], new(uint32)).(*uint32)
	out2 := *abi.ConvertType(out[2], new([][32]byte)).(*[][32]byte)

	return out0, out1, out2, err

}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) GetRequestConfig() (uint16, uint32, [][32]byte, error) {
	return _VRFCoordinatorV2Plus.Contract.GetRequestConfig(&_VRFCoordinatorV2Plus.CallOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCallerSession) GetRequestConfig() (uint16, uint32, [][32]byte, error) {
	return _VRFCoordinatorV2Plus.Contract.GetRequestConfig(&_VRFCoordinatorV2Plus.CallOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCaller) GetSubscription(opts *bind.CallOpts, subId *big.Int) (GetSubscription,

	error) {
	var out []interface{}
	err := _VRFCoordinatorV2Plus.contract.Call(opts, &out, "getSubscription", subId)

	outstruct := new(GetSubscription)
	if err != nil {
		return *outstruct, err
	}

	outstruct.Balance = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EthBalance = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ReqCount = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.Owner = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Consumers = *abi.ConvertType(out[4], new([]common.Address)).(*[]common.Address)
	outstruct.Nonces = *abi.ConvertType(out[5], new([]uint64)).(*[]uint64)

	return *outstruct, err

}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) GetSubscription(subId *big.Int) (GetSubscription,

	error) {
	return _VRFCoordinatorV2Plus.Contract.GetSubscription(&_VRFCoordinatorV2Plus.CallOpts, subId)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCallerSession) GetSubscription(subId *big.Int) (GetSubscription,

	error) {
	return _VRFCoordinatorV2Plus.Contract.GetSubscription(&_VRFCoordinatorV2Plus.CallOpts, subId)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCaller) HashOfKey(opts *bind.CallOpts, publicKey [2]*big.Int) ([32]byte, error) {
	var out []interface{}
	err := _VRFCoordinatorV2Plus.contract.Call(opts, &out, "hashOfKey", publicKey)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) HashOfKey(publicKey [2]*big.Int) ([32]byte, error) {
	return _VRFCoordinatorV2Plus.Contract.HashOfKey(&_VRFCoordinatorV2Plus.CallOpts, publicKey)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCallerSession) HashOfKey(publicKey [2]*big.Int) ([32]byte, error) {
	return _VRFCoordinatorV2Plus.Contract.HashOfKey(&_VRFCoordinatorV2Plus.CallOpts, publicKey)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCaller) MigrationVersion(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _VRFCoordinatorV2Plus.contract.Call(opts, &out, "migrationVersion")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) MigrationVersion() (uint8, error) {
	return _VRFCoordinatorV2Plus.Contract.MigrationVersion(&_VRFCoordinatorV2Plus.CallOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCallerSession) MigrationVersion() (uint8, error) {
	return _VRFCoordinatorV2Plus.Contract.MigrationVersion(&_VRFCoordinatorV2Plus.CallOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFCoordinatorV2Plus.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) Owner() (common.Address, error) {
	return _VRFCoordinatorV2Plus.Contract.Owner(&_VRFCoordinatorV2Plus.CallOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCallerSession) Owner() (common.Address, error) {
	return _VRFCoordinatorV2Plus.Contract.Owner(&_VRFCoordinatorV2Plus.CallOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCaller) PendingRequestExists(opts *bind.CallOpts, subId *big.Int) (bool, error) {
	var out []interface{}
	err := _VRFCoordinatorV2Plus.contract.Call(opts, &out, "pendingRequestExists", subId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) PendingRequestExists(subId *big.Int) (bool, error) {
	return _VRFCoordinatorV2Plus.Contract.PendingRequestExists(&_VRFCoordinatorV2Plus.CallOpts, subId)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCallerSession) PendingRequestExists(subId *big.Int) (bool, error) {
	return _VRFCoordinatorV2Plus.Contract.PendingRequestExists(&_VRFCoordinatorV2Plus.CallOpts, subId)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCaller) SConfig(opts *bind.CallOpts) (SConfig,

	error) {
	var out []interface{}
	err := _VRFCoordinatorV2Plus.contract.Call(opts, &out, "s_config")

	outstruct := new(SConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.MinimumRequestConfirmations = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.MaxGasLimit = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.StalenessSeconds = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.GasAfterPaymentCalculation = *abi.ConvertType(out[3], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) SConfig() (SConfig,

	error) {
	return _VRFCoordinatorV2Plus.Contract.SConfig(&_VRFCoordinatorV2Plus.CallOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCallerSession) SConfig() (SConfig,

	error) {
	return _VRFCoordinatorV2Plus.Contract.SConfig(&_VRFCoordinatorV2Plus.CallOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCaller) SCurrentSubNonce(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _VRFCoordinatorV2Plus.contract.Call(opts, &out, "s_currentSubNonce")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) SCurrentSubNonce() (uint64, error) {
	return _VRFCoordinatorV2Plus.Contract.SCurrentSubNonce(&_VRFCoordinatorV2Plus.CallOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCallerSession) SCurrentSubNonce() (uint64, error) {
	return _VRFCoordinatorV2Plus.Contract.SCurrentSubNonce(&_VRFCoordinatorV2Plus.CallOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCaller) SFallbackWeiPerUnitLink(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFCoordinatorV2Plus.contract.Call(opts, &out, "s_fallbackWeiPerUnitLink")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) SFallbackWeiPerUnitLink() (*big.Int, error) {
	return _VRFCoordinatorV2Plus.Contract.SFallbackWeiPerUnitLink(&_VRFCoordinatorV2Plus.CallOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCallerSession) SFallbackWeiPerUnitLink() (*big.Int, error) {
	return _VRFCoordinatorV2Plus.Contract.SFallbackWeiPerUnitLink(&_VRFCoordinatorV2Plus.CallOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCaller) SFeeConfig(opts *bind.CallOpts) (SFeeConfig,

	error) {
	var out []interface{}
	err := _VRFCoordinatorV2Plus.contract.Call(opts, &out, "s_feeConfig")

	outstruct := new(SFeeConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.FulfillmentFlatFeeLinkPPM = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.FulfillmentFlatFeeEthPPM = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) SFeeConfig() (SFeeConfig,

	error) {
	return _VRFCoordinatorV2Plus.Contract.SFeeConfig(&_VRFCoordinatorV2Plus.CallOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCallerSession) SFeeConfig() (SFeeConfig,

	error) {
	return _VRFCoordinatorV2Plus.Contract.SFeeConfig(&_VRFCoordinatorV2Plus.CallOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCaller) SProvingKeyHashes(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _VRFCoordinatorV2Plus.contract.Call(opts, &out, "s_provingKeyHashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) SProvingKeyHashes(arg0 *big.Int) ([32]byte, error) {
	return _VRFCoordinatorV2Plus.Contract.SProvingKeyHashes(&_VRFCoordinatorV2Plus.CallOpts, arg0)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCallerSession) SProvingKeyHashes(arg0 *big.Int) ([32]byte, error) {
	return _VRFCoordinatorV2Plus.Contract.SProvingKeyHashes(&_VRFCoordinatorV2Plus.CallOpts, arg0)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCaller) SProvingKeys(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _VRFCoordinatorV2Plus.contract.Call(opts, &out, "s_provingKeys", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) SProvingKeys(arg0 [32]byte) (common.Address, error) {
	return _VRFCoordinatorV2Plus.Contract.SProvingKeys(&_VRFCoordinatorV2Plus.CallOpts, arg0)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCallerSession) SProvingKeys(arg0 [32]byte) (common.Address, error) {
	return _VRFCoordinatorV2Plus.Contract.SProvingKeys(&_VRFCoordinatorV2Plus.CallOpts, arg0)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCaller) SRequestCommitments(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _VRFCoordinatorV2Plus.contract.Call(opts, &out, "s_requestCommitments", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) SRequestCommitments(arg0 *big.Int) ([32]byte, error) {
	return _VRFCoordinatorV2Plus.Contract.SRequestCommitments(&_VRFCoordinatorV2Plus.CallOpts, arg0)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCallerSession) SRequestCommitments(arg0 *big.Int) ([32]byte, error) {
	return _VRFCoordinatorV2Plus.Contract.SRequestCommitments(&_VRFCoordinatorV2Plus.CallOpts, arg0)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCaller) STotalBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFCoordinatorV2Plus.contract.Call(opts, &out, "s_totalBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) STotalBalance() (*big.Int, error) {
	return _VRFCoordinatorV2Plus.Contract.STotalBalance(&_VRFCoordinatorV2Plus.CallOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCallerSession) STotalBalance() (*big.Int, error) {
	return _VRFCoordinatorV2Plus.Contract.STotalBalance(&_VRFCoordinatorV2Plus.CallOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCaller) STotalEthBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFCoordinatorV2Plus.contract.Call(opts, &out, "s_totalEthBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) STotalEthBalance() (*big.Int, error) {
	return _VRFCoordinatorV2Plus.Contract.STotalEthBalance(&_VRFCoordinatorV2Plus.CallOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCallerSession) STotalEthBalance() (*big.Int, error) {
	return _VRFCoordinatorV2Plus.Contract.STotalEthBalance(&_VRFCoordinatorV2Plus.CallOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.contract.Transact(opts, "acceptOwnership")
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) AcceptOwnership() (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.AcceptOwnership(&_VRFCoordinatorV2Plus.TransactOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.AcceptOwnership(&_VRFCoordinatorV2Plus.TransactOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactor) AcceptSubscriptionOwnerTransfer(opts *bind.TransactOpts, subId *big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.contract.Transact(opts, "acceptSubscriptionOwnerTransfer", subId)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) AcceptSubscriptionOwnerTransfer(subId *big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.AcceptSubscriptionOwnerTransfer(&_VRFCoordinatorV2Plus.TransactOpts, subId)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactorSession) AcceptSubscriptionOwnerTransfer(subId *big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.AcceptSubscriptionOwnerTransfer(&_VRFCoordinatorV2Plus.TransactOpts, subId)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactor) AddConsumer(opts *bind.TransactOpts, subId *big.Int, consumer common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.contract.Transact(opts, "addConsumer", subId, consumer)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) AddConsumer(subId *big.Int, consumer common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.AddConsumer(&_VRFCoordinatorV2Plus.TransactOpts, subId, consumer)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactorSession) AddConsumer(subId *big.Int, consumer common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.AddConsumer(&_VRFCoordinatorV2Plus.TransactOpts, subId, consumer)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactor) CancelSubscription(opts *bind.TransactOpts, subId *big.Int, to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.contract.Transact(opts, "cancelSubscription", subId, to)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) CancelSubscription(subId *big.Int, to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.CancelSubscription(&_VRFCoordinatorV2Plus.TransactOpts, subId, to)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactorSession) CancelSubscription(subId *big.Int, to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.CancelSubscription(&_VRFCoordinatorV2Plus.TransactOpts, subId, to)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactor) CreateSubscription(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.contract.Transact(opts, "createSubscription")
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) CreateSubscription() (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.CreateSubscription(&_VRFCoordinatorV2Plus.TransactOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactorSession) CreateSubscription() (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.CreateSubscription(&_VRFCoordinatorV2Plus.TransactOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactor) DeregisterMigratableCoordinator(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.contract.Transact(opts, "deregisterMigratableCoordinator", target)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) DeregisterMigratableCoordinator(target common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.DeregisterMigratableCoordinator(&_VRFCoordinatorV2Plus.TransactOpts, target)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactorSession) DeregisterMigratableCoordinator(target common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.DeregisterMigratableCoordinator(&_VRFCoordinatorV2Plus.TransactOpts, target)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactor) DeregisterProvingKey(opts *bind.TransactOpts, publicProvingKey [2]*big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.contract.Transact(opts, "deregisterProvingKey", publicProvingKey)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) DeregisterProvingKey(publicProvingKey [2]*big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.DeregisterProvingKey(&_VRFCoordinatorV2Plus.TransactOpts, publicProvingKey)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactorSession) DeregisterProvingKey(publicProvingKey [2]*big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.DeregisterProvingKey(&_VRFCoordinatorV2Plus.TransactOpts, publicProvingKey)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactor) FulfillRandomWords(opts *bind.TransactOpts, proof VRFProof, rc VRFCoordinatorV2PlusRequestCommitment) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.contract.Transact(opts, "fulfillRandomWords", proof, rc)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) FulfillRandomWords(proof VRFProof, rc VRFCoordinatorV2PlusRequestCommitment) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.FulfillRandomWords(&_VRFCoordinatorV2Plus.TransactOpts, proof, rc)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactorSession) FulfillRandomWords(proof VRFProof, rc VRFCoordinatorV2PlusRequestCommitment) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.FulfillRandomWords(&_VRFCoordinatorV2Plus.TransactOpts, proof, rc)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactor) FundSubscriptionWithEth(opts *bind.TransactOpts, subId *big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.contract.Transact(opts, "fundSubscriptionWithEth", subId)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) FundSubscriptionWithEth(subId *big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.FundSubscriptionWithEth(&_VRFCoordinatorV2Plus.TransactOpts, subId)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactorSession) FundSubscriptionWithEth(subId *big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.FundSubscriptionWithEth(&_VRFCoordinatorV2Plus.TransactOpts, subId)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactor) Migrate(opts *bind.TransactOpts, subId *big.Int, newCoordinator common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.contract.Transact(opts, "migrate", subId, newCoordinator)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) Migrate(subId *big.Int, newCoordinator common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.Migrate(&_VRFCoordinatorV2Plus.TransactOpts, subId, newCoordinator)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactorSession) Migrate(subId *big.Int, newCoordinator common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.Migrate(&_VRFCoordinatorV2Plus.TransactOpts, subId, newCoordinator)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactor) OnTokenTransfer(opts *bind.TransactOpts, arg0 common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.contract.Transact(opts, "onTokenTransfer", arg0, amount, data)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) OnTokenTransfer(arg0 common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.OnTokenTransfer(&_VRFCoordinatorV2Plus.TransactOpts, arg0, amount, data)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactorSession) OnTokenTransfer(arg0 common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.OnTokenTransfer(&_VRFCoordinatorV2Plus.TransactOpts, arg0, amount, data)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactor) OracleWithdraw(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.contract.Transact(opts, "oracleWithdraw", recipient, amount)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) OracleWithdraw(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.OracleWithdraw(&_VRFCoordinatorV2Plus.TransactOpts, recipient, amount)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactorSession) OracleWithdraw(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.OracleWithdraw(&_VRFCoordinatorV2Plus.TransactOpts, recipient, amount)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactor) OracleWithdrawEth(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.contract.Transact(opts, "oracleWithdrawEth", recipient, amount)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) OracleWithdrawEth(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.OracleWithdrawEth(&_VRFCoordinatorV2Plus.TransactOpts, recipient, amount)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactorSession) OracleWithdrawEth(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.OracleWithdrawEth(&_VRFCoordinatorV2Plus.TransactOpts, recipient, amount)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactor) OwnerCancelSubscription(opts *bind.TransactOpts, subId *big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.contract.Transact(opts, "ownerCancelSubscription", subId)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) OwnerCancelSubscription(subId *big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.OwnerCancelSubscription(&_VRFCoordinatorV2Plus.TransactOpts, subId)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactorSession) OwnerCancelSubscription(subId *big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.OwnerCancelSubscription(&_VRFCoordinatorV2Plus.TransactOpts, subId)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactor) RecoverEthFunds(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.contract.Transact(opts, "recoverEthFunds", to)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) RecoverEthFunds(to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.RecoverEthFunds(&_VRFCoordinatorV2Plus.TransactOpts, to)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactorSession) RecoverEthFunds(to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.RecoverEthFunds(&_VRFCoordinatorV2Plus.TransactOpts, to)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactor) RecoverFunds(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.contract.Transact(opts, "recoverFunds", to)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) RecoverFunds(to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.RecoverFunds(&_VRFCoordinatorV2Plus.TransactOpts, to)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactorSession) RecoverFunds(to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.RecoverFunds(&_VRFCoordinatorV2Plus.TransactOpts, to)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactor) RegisterMigratableCoordinator(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.contract.Transact(opts, "registerMigratableCoordinator", target)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) RegisterMigratableCoordinator(target common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.RegisterMigratableCoordinator(&_VRFCoordinatorV2Plus.TransactOpts, target)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactorSession) RegisterMigratableCoordinator(target common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.RegisterMigratableCoordinator(&_VRFCoordinatorV2Plus.TransactOpts, target)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactor) RegisterProvingKey(opts *bind.TransactOpts, oracle common.Address, publicProvingKey [2]*big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.contract.Transact(opts, "registerProvingKey", oracle, publicProvingKey)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) RegisterProvingKey(oracle common.Address, publicProvingKey [2]*big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.RegisterProvingKey(&_VRFCoordinatorV2Plus.TransactOpts, oracle, publicProvingKey)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactorSession) RegisterProvingKey(oracle common.Address, publicProvingKey [2]*big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.RegisterProvingKey(&_VRFCoordinatorV2Plus.TransactOpts, oracle, publicProvingKey)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactor) RemoveConsumer(opts *bind.TransactOpts, subId *big.Int, consumer common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.contract.Transact(opts, "removeConsumer", subId, consumer)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) RemoveConsumer(subId *big.Int, consumer common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.RemoveConsumer(&_VRFCoordinatorV2Plus.TransactOpts, subId, consumer)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactorSession) RemoveConsumer(subId *big.Int, consumer common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.RemoveConsumer(&_VRFCoordinatorV2Plus.TransactOpts, subId, consumer)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactor) RequestRandomWords(opts *bind.TransactOpts, req VRFV2PlusClientRandomWordsRequest) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.contract.Transact(opts, "requestRandomWords", req)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) RequestRandomWords(req VRFV2PlusClientRandomWordsRequest) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.RequestRandomWords(&_VRFCoordinatorV2Plus.TransactOpts, req)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactorSession) RequestRandomWords(req VRFV2PlusClientRandomWordsRequest) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.RequestRandomWords(&_VRFCoordinatorV2Plus.TransactOpts, req)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactor) RequestSubscriptionOwnerTransfer(opts *bind.TransactOpts, subId *big.Int, newOwner common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.contract.Transact(opts, "requestSubscriptionOwnerTransfer", subId, newOwner)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) RequestSubscriptionOwnerTransfer(subId *big.Int, newOwner common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.RequestSubscriptionOwnerTransfer(&_VRFCoordinatorV2Plus.TransactOpts, subId, newOwner)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactorSession) RequestSubscriptionOwnerTransfer(subId *big.Int, newOwner common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.RequestSubscriptionOwnerTransfer(&_VRFCoordinatorV2Plus.TransactOpts, subId, newOwner)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactor) SetConfig(opts *bind.TransactOpts, minimumRequestConfirmations uint16, maxGasLimit uint32, stalenessSeconds uint32, gasAfterPaymentCalculation uint32, fallbackWeiPerUnitLink *big.Int, feeConfig VRFCoordinatorV2PlusFeeConfig) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.contract.Transact(opts, "setConfig", minimumRequestConfirmations, maxGasLimit, stalenessSeconds, gasAfterPaymentCalculation, fallbackWeiPerUnitLink, feeConfig)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) SetConfig(minimumRequestConfirmations uint16, maxGasLimit uint32, stalenessSeconds uint32, gasAfterPaymentCalculation uint32, fallbackWeiPerUnitLink *big.Int, feeConfig VRFCoordinatorV2PlusFeeConfig) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.SetConfig(&_VRFCoordinatorV2Plus.TransactOpts, minimumRequestConfirmations, maxGasLimit, stalenessSeconds, gasAfterPaymentCalculation, fallbackWeiPerUnitLink, feeConfig)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactorSession) SetConfig(minimumRequestConfirmations uint16, maxGasLimit uint32, stalenessSeconds uint32, gasAfterPaymentCalculation uint32, fallbackWeiPerUnitLink *big.Int, feeConfig VRFCoordinatorV2PlusFeeConfig) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.SetConfig(&_VRFCoordinatorV2Plus.TransactOpts, minimumRequestConfirmations, maxGasLimit, stalenessSeconds, gasAfterPaymentCalculation, fallbackWeiPerUnitLink, feeConfig)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactor) SetLINKAndLINKETHFeed(opts *bind.TransactOpts, link common.Address, linkEthFeed common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.contract.Transact(opts, "setLINKAndLINKETHFeed", link, linkEthFeed)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) SetLINKAndLINKETHFeed(link common.Address, linkEthFeed common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.SetLINKAndLINKETHFeed(&_VRFCoordinatorV2Plus.TransactOpts, link, linkEthFeed)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactorSession) SetLINKAndLINKETHFeed(link common.Address, linkEthFeed common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.SetLINKAndLINKETHFeed(&_VRFCoordinatorV2Plus.TransactOpts, link, linkEthFeed)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.contract.Transact(opts, "transferOwnership", to)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.TransferOwnership(&_VRFCoordinatorV2Plus.TransactOpts, to)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.TransferOwnership(&_VRFCoordinatorV2Plus.TransactOpts, to)
}

type VRFCoordinatorV2PlusConfigSetIterator struct {
	Event *VRFCoordinatorV2PlusConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2PlusConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2PlusConfigSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2PlusConfigSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV2PlusConfigSetIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2PlusConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2PlusConfigSet struct {
	MinimumRequestConfirmations uint16
	MaxGasLimit                 uint32
	StalenessSeconds            uint32
	GasAfterPaymentCalculation  uint32
	FallbackWeiPerUnitLink      *big.Int
	FeeConfig                   VRFCoordinatorV2PlusFeeConfig
	Raw                         types.Log
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) FilterConfigSet(opts *bind.FilterOpts) (*VRFCoordinatorV2PlusConfigSetIterator, error) {

	logs, sub, err := _VRFCoordinatorV2Plus.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusConfigSetIterator{contract: _VRFCoordinatorV2Plus.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusConfigSet) (event.Subscription, error) {

	logs, sub, err := _VRFCoordinatorV2Plus.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2PlusConfigSet)
				if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "ConfigSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) ParseConfigSet(log types.Log) (*VRFCoordinatorV2PlusConfigSet, error) {
	event := new(VRFCoordinatorV2PlusConfigSet)
	if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2PlusCoordinatorDeregisteredIterator struct {
	Event *VRFCoordinatorV2PlusCoordinatorDeregistered

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2PlusCoordinatorDeregisteredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2PlusCoordinatorDeregistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2PlusCoordinatorDeregistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV2PlusCoordinatorDeregisteredIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2PlusCoordinatorDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2PlusCoordinatorDeregistered struct {
	CoordinatorAddress common.Address
	Raw                types.Log
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) FilterCoordinatorDeregistered(opts *bind.FilterOpts) (*VRFCoordinatorV2PlusCoordinatorDeregisteredIterator, error) {

	logs, sub, err := _VRFCoordinatorV2Plus.contract.FilterLogs(opts, "CoordinatorDeregistered")
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusCoordinatorDeregisteredIterator{contract: _VRFCoordinatorV2Plus.contract, event: "CoordinatorDeregistered", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) WatchCoordinatorDeregistered(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusCoordinatorDeregistered) (event.Subscription, error) {

	logs, sub, err := _VRFCoordinatorV2Plus.contract.WatchLogs(opts, "CoordinatorDeregistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2PlusCoordinatorDeregistered)
				if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "CoordinatorDeregistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) ParseCoordinatorDeregistered(log types.Log) (*VRFCoordinatorV2PlusCoordinatorDeregistered, error) {
	event := new(VRFCoordinatorV2PlusCoordinatorDeregistered)
	if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "CoordinatorDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2PlusCoordinatorRegisteredIterator struct {
	Event *VRFCoordinatorV2PlusCoordinatorRegistered

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2PlusCoordinatorRegisteredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2PlusCoordinatorRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2PlusCoordinatorRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV2PlusCoordinatorRegisteredIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2PlusCoordinatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2PlusCoordinatorRegistered struct {
	CoordinatorAddress common.Address
	Raw                types.Log
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) FilterCoordinatorRegistered(opts *bind.FilterOpts) (*VRFCoordinatorV2PlusCoordinatorRegisteredIterator, error) {

	logs, sub, err := _VRFCoordinatorV2Plus.contract.FilterLogs(opts, "CoordinatorRegistered")
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusCoordinatorRegisteredIterator{contract: _VRFCoordinatorV2Plus.contract, event: "CoordinatorRegistered", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) WatchCoordinatorRegistered(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusCoordinatorRegistered) (event.Subscription, error) {

	logs, sub, err := _VRFCoordinatorV2Plus.contract.WatchLogs(opts, "CoordinatorRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2PlusCoordinatorRegistered)
				if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "CoordinatorRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) ParseCoordinatorRegistered(log types.Log) (*VRFCoordinatorV2PlusCoordinatorRegistered, error) {
	event := new(VRFCoordinatorV2PlusCoordinatorRegistered)
	if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "CoordinatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2PlusEthFundsRecoveredIterator struct {
	Event *VRFCoordinatorV2PlusEthFundsRecovered

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2PlusEthFundsRecoveredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2PlusEthFundsRecovered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2PlusEthFundsRecovered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV2PlusEthFundsRecoveredIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2PlusEthFundsRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2PlusEthFundsRecovered struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) FilterEthFundsRecovered(opts *bind.FilterOpts) (*VRFCoordinatorV2PlusEthFundsRecoveredIterator, error) {

	logs, sub, err := _VRFCoordinatorV2Plus.contract.FilterLogs(opts, "EthFundsRecovered")
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusEthFundsRecoveredIterator{contract: _VRFCoordinatorV2Plus.contract, event: "EthFundsRecovered", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) WatchEthFundsRecovered(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusEthFundsRecovered) (event.Subscription, error) {

	logs, sub, err := _VRFCoordinatorV2Plus.contract.WatchLogs(opts, "EthFundsRecovered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2PlusEthFundsRecovered)
				if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "EthFundsRecovered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) ParseEthFundsRecovered(log types.Log) (*VRFCoordinatorV2PlusEthFundsRecovered, error) {
	event := new(VRFCoordinatorV2PlusEthFundsRecovered)
	if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "EthFundsRecovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2PlusFundsRecoveredIterator struct {
	Event *VRFCoordinatorV2PlusFundsRecovered

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2PlusFundsRecoveredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2PlusFundsRecovered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2PlusFundsRecovered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV2PlusFundsRecoveredIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2PlusFundsRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2PlusFundsRecovered struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) FilterFundsRecovered(opts *bind.FilterOpts) (*VRFCoordinatorV2PlusFundsRecoveredIterator, error) {

	logs, sub, err := _VRFCoordinatorV2Plus.contract.FilterLogs(opts, "FundsRecovered")
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusFundsRecoveredIterator{contract: _VRFCoordinatorV2Plus.contract, event: "FundsRecovered", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) WatchFundsRecovered(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusFundsRecovered) (event.Subscription, error) {

	logs, sub, err := _VRFCoordinatorV2Plus.contract.WatchLogs(opts, "FundsRecovered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2PlusFundsRecovered)
				if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "FundsRecovered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) ParseFundsRecovered(log types.Log) (*VRFCoordinatorV2PlusFundsRecovered, error) {
	event := new(VRFCoordinatorV2PlusFundsRecovered)
	if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "FundsRecovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2PlusMigrationCompletedIterator struct {
	Event *VRFCoordinatorV2PlusMigrationCompleted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2PlusMigrationCompletedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2PlusMigrationCompleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2PlusMigrationCompleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV2PlusMigrationCompletedIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2PlusMigrationCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2PlusMigrationCompleted struct {
	NewCoordinator common.Address
	SubId          *big.Int
	Raw            types.Log
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) FilterMigrationCompleted(opts *bind.FilterOpts) (*VRFCoordinatorV2PlusMigrationCompletedIterator, error) {

	logs, sub, err := _VRFCoordinatorV2Plus.contract.FilterLogs(opts, "MigrationCompleted")
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusMigrationCompletedIterator{contract: _VRFCoordinatorV2Plus.contract, event: "MigrationCompleted", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) WatchMigrationCompleted(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusMigrationCompleted) (event.Subscription, error) {

	logs, sub, err := _VRFCoordinatorV2Plus.contract.WatchLogs(opts, "MigrationCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2PlusMigrationCompleted)
				if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "MigrationCompleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) ParseMigrationCompleted(log types.Log) (*VRFCoordinatorV2PlusMigrationCompleted, error) {
	event := new(VRFCoordinatorV2PlusMigrationCompleted)
	if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "MigrationCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2PlusOwnershipTransferRequestedIterator struct {
	Event *VRFCoordinatorV2PlusOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2PlusOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2PlusOwnershipTransferRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2PlusOwnershipTransferRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV2PlusOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2PlusOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2PlusOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFCoordinatorV2PlusOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFCoordinatorV2Plus.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusOwnershipTransferRequestedIterator{contract: _VRFCoordinatorV2Plus.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFCoordinatorV2Plus.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2PlusOwnershipTransferRequested)
				if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) ParseOwnershipTransferRequested(log types.Log) (*VRFCoordinatorV2PlusOwnershipTransferRequested, error) {
	event := new(VRFCoordinatorV2PlusOwnershipTransferRequested)
	if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2PlusOwnershipTransferredIterator struct {
	Event *VRFCoordinatorV2PlusOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2PlusOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2PlusOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2PlusOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV2PlusOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2PlusOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2PlusOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFCoordinatorV2PlusOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFCoordinatorV2Plus.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusOwnershipTransferredIterator{contract: _VRFCoordinatorV2Plus.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFCoordinatorV2Plus.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2PlusOwnershipTransferred)
				if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) ParseOwnershipTransferred(log types.Log) (*VRFCoordinatorV2PlusOwnershipTransferred, error) {
	event := new(VRFCoordinatorV2PlusOwnershipTransferred)
	if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2PlusProvingKeyDeregisteredIterator struct {
	Event *VRFCoordinatorV2PlusProvingKeyDeregistered

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2PlusProvingKeyDeregisteredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2PlusProvingKeyDeregistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2PlusProvingKeyDeregistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV2PlusProvingKeyDeregisteredIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2PlusProvingKeyDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2PlusProvingKeyDeregistered struct {
	KeyHash [32]byte
	Oracle  common.Address
	Raw     types.Log
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) FilterProvingKeyDeregistered(opts *bind.FilterOpts, oracle []common.Address) (*VRFCoordinatorV2PlusProvingKeyDeregisteredIterator, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _VRFCoordinatorV2Plus.contract.FilterLogs(opts, "ProvingKeyDeregistered", oracleRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusProvingKeyDeregisteredIterator{contract: _VRFCoordinatorV2Plus.contract, event: "ProvingKeyDeregistered", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) WatchProvingKeyDeregistered(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusProvingKeyDeregistered, oracle []common.Address) (event.Subscription, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _VRFCoordinatorV2Plus.contract.WatchLogs(opts, "ProvingKeyDeregistered", oracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2PlusProvingKeyDeregistered)
				if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "ProvingKeyDeregistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) ParseProvingKeyDeregistered(log types.Log) (*VRFCoordinatorV2PlusProvingKeyDeregistered, error) {
	event := new(VRFCoordinatorV2PlusProvingKeyDeregistered)
	if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "ProvingKeyDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2PlusProvingKeyRegisteredIterator struct {
	Event *VRFCoordinatorV2PlusProvingKeyRegistered

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2PlusProvingKeyRegisteredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2PlusProvingKeyRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2PlusProvingKeyRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV2PlusProvingKeyRegisteredIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2PlusProvingKeyRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2PlusProvingKeyRegistered struct {
	KeyHash [32]byte
	Oracle  common.Address
	Raw     types.Log
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) FilterProvingKeyRegistered(opts *bind.FilterOpts, oracle []common.Address) (*VRFCoordinatorV2PlusProvingKeyRegisteredIterator, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _VRFCoordinatorV2Plus.contract.FilterLogs(opts, "ProvingKeyRegistered", oracleRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusProvingKeyRegisteredIterator{contract: _VRFCoordinatorV2Plus.contract, event: "ProvingKeyRegistered", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) WatchProvingKeyRegistered(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusProvingKeyRegistered, oracle []common.Address) (event.Subscription, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _VRFCoordinatorV2Plus.contract.WatchLogs(opts, "ProvingKeyRegistered", oracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2PlusProvingKeyRegistered)
				if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "ProvingKeyRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) ParseProvingKeyRegistered(log types.Log) (*VRFCoordinatorV2PlusProvingKeyRegistered, error) {
	event := new(VRFCoordinatorV2PlusProvingKeyRegistered)
	if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "ProvingKeyRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2PlusRandomWordsFulfilledIterator struct {
	Event *VRFCoordinatorV2PlusRandomWordsFulfilled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2PlusRandomWordsFulfilledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2PlusRandomWordsFulfilled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2PlusRandomWordsFulfilled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV2PlusRandomWordsFulfilledIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2PlusRandomWordsFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2PlusRandomWordsFulfilled struct {
	RequestId  *big.Int
	OutputSeed *big.Int
	SubID      *big.Int
	Payment    *big.Int
	ExtraArgs  []byte
	Success    bool
	Raw        types.Log
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) FilterRandomWordsFulfilled(opts *bind.FilterOpts, requestId []*big.Int, subID []*big.Int) (*VRFCoordinatorV2PlusRandomWordsFulfilledIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	var subIDRule []interface{}
	for _, subIDItem := range subID {
		subIDRule = append(subIDRule, subIDItem)
	}

	logs, sub, err := _VRFCoordinatorV2Plus.contract.FilterLogs(opts, "RandomWordsFulfilled", requestIdRule, subIDRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusRandomWordsFulfilledIterator{contract: _VRFCoordinatorV2Plus.contract, event: "RandomWordsFulfilled", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) WatchRandomWordsFulfilled(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusRandomWordsFulfilled, requestId []*big.Int, subID []*big.Int) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	var subIDRule []interface{}
	for _, subIDItem := range subID {
		subIDRule = append(subIDRule, subIDItem)
	}

	logs, sub, err := _VRFCoordinatorV2Plus.contract.WatchLogs(opts, "RandomWordsFulfilled", requestIdRule, subIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2PlusRandomWordsFulfilled)
				if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "RandomWordsFulfilled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) ParseRandomWordsFulfilled(log types.Log) (*VRFCoordinatorV2PlusRandomWordsFulfilled, error) {
	event := new(VRFCoordinatorV2PlusRandomWordsFulfilled)
	if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "RandomWordsFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2PlusRandomWordsRequestedIterator struct {
	Event *VRFCoordinatorV2PlusRandomWordsRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2PlusRandomWordsRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2PlusRandomWordsRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2PlusRandomWordsRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV2PlusRandomWordsRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2PlusRandomWordsRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2PlusRandomWordsRequested struct {
	KeyHash                     [32]byte
	RequestId                   *big.Int
	PreSeed                     *big.Int
	SubId                       *big.Int
	MinimumRequestConfirmations uint16
	CallbackGasLimit            uint32
	NumWords                    uint32
	ExtraArgs                   []byte
	Sender                      common.Address
	Raw                         types.Log
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) FilterRandomWordsRequested(opts *bind.FilterOpts, keyHash [][32]byte, subId []*big.Int, sender []common.Address) (*VRFCoordinatorV2PlusRandomWordsRequestedIterator, error) {

	var keyHashRule []interface{}
	for _, keyHashItem := range keyHash {
		keyHashRule = append(keyHashRule, keyHashItem)
	}

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VRFCoordinatorV2Plus.contract.FilterLogs(opts, "RandomWordsRequested", keyHashRule, subIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusRandomWordsRequestedIterator{contract: _VRFCoordinatorV2Plus.contract, event: "RandomWordsRequested", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) WatchRandomWordsRequested(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusRandomWordsRequested, keyHash [][32]byte, subId []*big.Int, sender []common.Address) (event.Subscription, error) {

	var keyHashRule []interface{}
	for _, keyHashItem := range keyHash {
		keyHashRule = append(keyHashRule, keyHashItem)
	}

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VRFCoordinatorV2Plus.contract.WatchLogs(opts, "RandomWordsRequested", keyHashRule, subIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2PlusRandomWordsRequested)
				if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "RandomWordsRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) ParseRandomWordsRequested(log types.Log) (*VRFCoordinatorV2PlusRandomWordsRequested, error) {
	event := new(VRFCoordinatorV2PlusRandomWordsRequested)
	if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "RandomWordsRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2PlusSubscriptionCanceledIterator struct {
	Event *VRFCoordinatorV2PlusSubscriptionCanceled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2PlusSubscriptionCanceledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2PlusSubscriptionCanceled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2PlusSubscriptionCanceled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV2PlusSubscriptionCanceledIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2PlusSubscriptionCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2PlusSubscriptionCanceled struct {
	SubId      *big.Int
	To         common.Address
	AmountLink *big.Int
	AmountEth  *big.Int
	Raw        types.Log
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) FilterSubscriptionCanceled(opts *bind.FilterOpts, subId []*big.Int) (*VRFCoordinatorV2PlusSubscriptionCanceledIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2Plus.contract.FilterLogs(opts, "SubscriptionCanceled", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusSubscriptionCanceledIterator{contract: _VRFCoordinatorV2Plus.contract, event: "SubscriptionCanceled", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) WatchSubscriptionCanceled(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusSubscriptionCanceled, subId []*big.Int) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2Plus.contract.WatchLogs(opts, "SubscriptionCanceled", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2PlusSubscriptionCanceled)
				if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "SubscriptionCanceled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) ParseSubscriptionCanceled(log types.Log) (*VRFCoordinatorV2PlusSubscriptionCanceled, error) {
	event := new(VRFCoordinatorV2PlusSubscriptionCanceled)
	if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "SubscriptionCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2PlusSubscriptionConsumerAddedIterator struct {
	Event *VRFCoordinatorV2PlusSubscriptionConsumerAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2PlusSubscriptionConsumerAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2PlusSubscriptionConsumerAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2PlusSubscriptionConsumerAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV2PlusSubscriptionConsumerAddedIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2PlusSubscriptionConsumerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2PlusSubscriptionConsumerAdded struct {
	SubId    *big.Int
	Consumer common.Address
	Raw      types.Log
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) FilterSubscriptionConsumerAdded(opts *bind.FilterOpts, subId []*big.Int) (*VRFCoordinatorV2PlusSubscriptionConsumerAddedIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2Plus.contract.FilterLogs(opts, "SubscriptionConsumerAdded", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusSubscriptionConsumerAddedIterator{contract: _VRFCoordinatorV2Plus.contract, event: "SubscriptionConsumerAdded", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) WatchSubscriptionConsumerAdded(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusSubscriptionConsumerAdded, subId []*big.Int) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2Plus.contract.WatchLogs(opts, "SubscriptionConsumerAdded", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2PlusSubscriptionConsumerAdded)
				if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "SubscriptionConsumerAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) ParseSubscriptionConsumerAdded(log types.Log) (*VRFCoordinatorV2PlusSubscriptionConsumerAdded, error) {
	event := new(VRFCoordinatorV2PlusSubscriptionConsumerAdded)
	if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "SubscriptionConsumerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2PlusSubscriptionConsumerRemovedIterator struct {
	Event *VRFCoordinatorV2PlusSubscriptionConsumerRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2PlusSubscriptionConsumerRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2PlusSubscriptionConsumerRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2PlusSubscriptionConsumerRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV2PlusSubscriptionConsumerRemovedIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2PlusSubscriptionConsumerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2PlusSubscriptionConsumerRemoved struct {
	SubId    *big.Int
	Consumer common.Address
	Raw      types.Log
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) FilterSubscriptionConsumerRemoved(opts *bind.FilterOpts, subId []*big.Int) (*VRFCoordinatorV2PlusSubscriptionConsumerRemovedIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2Plus.contract.FilterLogs(opts, "SubscriptionConsumerRemoved", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusSubscriptionConsumerRemovedIterator{contract: _VRFCoordinatorV2Plus.contract, event: "SubscriptionConsumerRemoved", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) WatchSubscriptionConsumerRemoved(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusSubscriptionConsumerRemoved, subId []*big.Int) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2Plus.contract.WatchLogs(opts, "SubscriptionConsumerRemoved", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2PlusSubscriptionConsumerRemoved)
				if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "SubscriptionConsumerRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) ParseSubscriptionConsumerRemoved(log types.Log) (*VRFCoordinatorV2PlusSubscriptionConsumerRemoved, error) {
	event := new(VRFCoordinatorV2PlusSubscriptionConsumerRemoved)
	if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "SubscriptionConsumerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2PlusSubscriptionCreatedIterator struct {
	Event *VRFCoordinatorV2PlusSubscriptionCreated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2PlusSubscriptionCreatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2PlusSubscriptionCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2PlusSubscriptionCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV2PlusSubscriptionCreatedIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2PlusSubscriptionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2PlusSubscriptionCreated struct {
	SubId *big.Int
	Owner common.Address
	Raw   types.Log
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) FilterSubscriptionCreated(opts *bind.FilterOpts, subId []*big.Int) (*VRFCoordinatorV2PlusSubscriptionCreatedIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2Plus.contract.FilterLogs(opts, "SubscriptionCreated", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusSubscriptionCreatedIterator{contract: _VRFCoordinatorV2Plus.contract, event: "SubscriptionCreated", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) WatchSubscriptionCreated(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusSubscriptionCreated, subId []*big.Int) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2Plus.contract.WatchLogs(opts, "SubscriptionCreated", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2PlusSubscriptionCreated)
				if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "SubscriptionCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) ParseSubscriptionCreated(log types.Log) (*VRFCoordinatorV2PlusSubscriptionCreated, error) {
	event := new(VRFCoordinatorV2PlusSubscriptionCreated)
	if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "SubscriptionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2PlusSubscriptionFundedIterator struct {
	Event *VRFCoordinatorV2PlusSubscriptionFunded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2PlusSubscriptionFundedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2PlusSubscriptionFunded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2PlusSubscriptionFunded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV2PlusSubscriptionFundedIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2PlusSubscriptionFundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2PlusSubscriptionFunded struct {
	SubId      *big.Int
	OldBalance *big.Int
	NewBalance *big.Int
	Raw        types.Log
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) FilterSubscriptionFunded(opts *bind.FilterOpts, subId []*big.Int) (*VRFCoordinatorV2PlusSubscriptionFundedIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2Plus.contract.FilterLogs(opts, "SubscriptionFunded", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusSubscriptionFundedIterator{contract: _VRFCoordinatorV2Plus.contract, event: "SubscriptionFunded", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) WatchSubscriptionFunded(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusSubscriptionFunded, subId []*big.Int) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2Plus.contract.WatchLogs(opts, "SubscriptionFunded", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2PlusSubscriptionFunded)
				if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "SubscriptionFunded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) ParseSubscriptionFunded(log types.Log) (*VRFCoordinatorV2PlusSubscriptionFunded, error) {
	event := new(VRFCoordinatorV2PlusSubscriptionFunded)
	if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "SubscriptionFunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2PlusSubscriptionFundedWithEthIterator struct {
	Event *VRFCoordinatorV2PlusSubscriptionFundedWithEth

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2PlusSubscriptionFundedWithEthIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2PlusSubscriptionFundedWithEth)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2PlusSubscriptionFundedWithEth)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV2PlusSubscriptionFundedWithEthIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2PlusSubscriptionFundedWithEthIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2PlusSubscriptionFundedWithEth struct {
	SubId         *big.Int
	OldEthBalance *big.Int
	NewEthBalance *big.Int
	Raw           types.Log
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) FilterSubscriptionFundedWithEth(opts *bind.FilterOpts, subId []*big.Int) (*VRFCoordinatorV2PlusSubscriptionFundedWithEthIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2Plus.contract.FilterLogs(opts, "SubscriptionFundedWithEth", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusSubscriptionFundedWithEthIterator{contract: _VRFCoordinatorV2Plus.contract, event: "SubscriptionFundedWithEth", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) WatchSubscriptionFundedWithEth(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusSubscriptionFundedWithEth, subId []*big.Int) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2Plus.contract.WatchLogs(opts, "SubscriptionFundedWithEth", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2PlusSubscriptionFundedWithEth)
				if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "SubscriptionFundedWithEth", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) ParseSubscriptionFundedWithEth(log types.Log) (*VRFCoordinatorV2PlusSubscriptionFundedWithEth, error) {
	event := new(VRFCoordinatorV2PlusSubscriptionFundedWithEth)
	if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "SubscriptionFundedWithEth", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2PlusSubscriptionOwnerTransferRequestedIterator struct {
	Event *VRFCoordinatorV2PlusSubscriptionOwnerTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2PlusSubscriptionOwnerTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2PlusSubscriptionOwnerTransferRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2PlusSubscriptionOwnerTransferRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV2PlusSubscriptionOwnerTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2PlusSubscriptionOwnerTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2PlusSubscriptionOwnerTransferRequested struct {
	SubId *big.Int
	From  common.Address
	To    common.Address
	Raw   types.Log
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) FilterSubscriptionOwnerTransferRequested(opts *bind.FilterOpts, subId []*big.Int) (*VRFCoordinatorV2PlusSubscriptionOwnerTransferRequestedIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2Plus.contract.FilterLogs(opts, "SubscriptionOwnerTransferRequested", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusSubscriptionOwnerTransferRequestedIterator{contract: _VRFCoordinatorV2Plus.contract, event: "SubscriptionOwnerTransferRequested", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) WatchSubscriptionOwnerTransferRequested(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusSubscriptionOwnerTransferRequested, subId []*big.Int) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2Plus.contract.WatchLogs(opts, "SubscriptionOwnerTransferRequested", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2PlusSubscriptionOwnerTransferRequested)
				if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "SubscriptionOwnerTransferRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) ParseSubscriptionOwnerTransferRequested(log types.Log) (*VRFCoordinatorV2PlusSubscriptionOwnerTransferRequested, error) {
	event := new(VRFCoordinatorV2PlusSubscriptionOwnerTransferRequested)
	if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "SubscriptionOwnerTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2PlusSubscriptionOwnerTransferredIterator struct {
	Event *VRFCoordinatorV2PlusSubscriptionOwnerTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2PlusSubscriptionOwnerTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2PlusSubscriptionOwnerTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2PlusSubscriptionOwnerTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFCoordinatorV2PlusSubscriptionOwnerTransferredIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2PlusSubscriptionOwnerTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2PlusSubscriptionOwnerTransferred struct {
	SubId *big.Int
	From  common.Address
	To    common.Address
	Raw   types.Log
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) FilterSubscriptionOwnerTransferred(opts *bind.FilterOpts, subId []*big.Int) (*VRFCoordinatorV2PlusSubscriptionOwnerTransferredIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2Plus.contract.FilterLogs(opts, "SubscriptionOwnerTransferred", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusSubscriptionOwnerTransferredIterator{contract: _VRFCoordinatorV2Plus.contract, event: "SubscriptionOwnerTransferred", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) WatchSubscriptionOwnerTransferred(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusSubscriptionOwnerTransferred, subId []*big.Int) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2Plus.contract.WatchLogs(opts, "SubscriptionOwnerTransferred", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2PlusSubscriptionOwnerTransferred)
				if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "SubscriptionOwnerTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) ParseSubscriptionOwnerTransferred(log types.Log) (*VRFCoordinatorV2PlusSubscriptionOwnerTransferred, error) {
	event := new(VRFCoordinatorV2PlusSubscriptionOwnerTransferred)
	if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "SubscriptionOwnerTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GetSubscription struct {
	Balance    *big.Int
	EthBalance *big.Int
	ReqCount   uint64
	Owner      common.Address
	Consumers  []common.Address
	Nonces     []uint64
}
type SConfig struct {
	MinimumRequestConfirmations uint16
	MaxGasLimit                 uint32
	StalenessSeconds            uint32
	GasAfterPaymentCalculation  uint32
}
type SFeeConfig struct {
	FulfillmentFlatFeeLinkPPM uint32
	FulfillmentFlatFeeEthPPM  uint32
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2Plus) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _VRFCoordinatorV2Plus.abi.Events["ConfigSet"].ID:
		return _VRFCoordinatorV2Plus.ParseConfigSet(log)
	case _VRFCoordinatorV2Plus.abi.Events["CoordinatorDeregistered"].ID:
		return _VRFCoordinatorV2Plus.ParseCoordinatorDeregistered(log)
	case _VRFCoordinatorV2Plus.abi.Events["CoordinatorRegistered"].ID:
		return _VRFCoordinatorV2Plus.ParseCoordinatorRegistered(log)
	case _VRFCoordinatorV2Plus.abi.Events["EthFundsRecovered"].ID:
		return _VRFCoordinatorV2Plus.ParseEthFundsRecovered(log)
	case _VRFCoordinatorV2Plus.abi.Events["FundsRecovered"].ID:
		return _VRFCoordinatorV2Plus.ParseFundsRecovered(log)
	case _VRFCoordinatorV2Plus.abi.Events["MigrationCompleted"].ID:
		return _VRFCoordinatorV2Plus.ParseMigrationCompleted(log)
	case _VRFCoordinatorV2Plus.abi.Events["OwnershipTransferRequested"].ID:
		return _VRFCoordinatorV2Plus.ParseOwnershipTransferRequested(log)
	case _VRFCoordinatorV2Plus.abi.Events["OwnershipTransferred"].ID:
		return _VRFCoordinatorV2Plus.ParseOwnershipTransferred(log)
	case _VRFCoordinatorV2Plus.abi.Events["ProvingKeyDeregistered"].ID:
		return _VRFCoordinatorV2Plus.ParseProvingKeyDeregistered(log)
	case _VRFCoordinatorV2Plus.abi.Events["ProvingKeyRegistered"].ID:
		return _VRFCoordinatorV2Plus.ParseProvingKeyRegistered(log)
	case _VRFCoordinatorV2Plus.abi.Events["RandomWordsFulfilled"].ID:
		return _VRFCoordinatorV2Plus.ParseRandomWordsFulfilled(log)
	case _VRFCoordinatorV2Plus.abi.Events["RandomWordsRequested"].ID:
		return _VRFCoordinatorV2Plus.ParseRandomWordsRequested(log)
	case _VRFCoordinatorV2Plus.abi.Events["SubscriptionCanceled"].ID:
		return _VRFCoordinatorV2Plus.ParseSubscriptionCanceled(log)
	case _VRFCoordinatorV2Plus.abi.Events["SubscriptionConsumerAdded"].ID:
		return _VRFCoordinatorV2Plus.ParseSubscriptionConsumerAdded(log)
	case _VRFCoordinatorV2Plus.abi.Events["SubscriptionConsumerRemoved"].ID:
		return _VRFCoordinatorV2Plus.ParseSubscriptionConsumerRemoved(log)
	case _VRFCoordinatorV2Plus.abi.Events["SubscriptionCreated"].ID:
		return _VRFCoordinatorV2Plus.ParseSubscriptionCreated(log)
	case _VRFCoordinatorV2Plus.abi.Events["SubscriptionFunded"].ID:
		return _VRFCoordinatorV2Plus.ParseSubscriptionFunded(log)
	case _VRFCoordinatorV2Plus.abi.Events["SubscriptionFundedWithEth"].ID:
		return _VRFCoordinatorV2Plus.ParseSubscriptionFundedWithEth(log)
	case _VRFCoordinatorV2Plus.abi.Events["SubscriptionOwnerTransferRequested"].ID:
		return _VRFCoordinatorV2Plus.ParseSubscriptionOwnerTransferRequested(log)
	case _VRFCoordinatorV2Plus.abi.Events["SubscriptionOwnerTransferred"].ID:
		return _VRFCoordinatorV2Plus.ParseSubscriptionOwnerTransferred(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (VRFCoordinatorV2PlusConfigSet) Topic() common.Hash {
	return common.HexToHash("0x777357bb93f63d088f18112d3dba38457aec633eb8f1341e1d418380ad328e78")
}

func (VRFCoordinatorV2PlusCoordinatorDeregistered) Topic() common.Hash {
	return common.HexToHash("0xf80a1a97fd42251f3c33cda98635e7399253033a6774fe37cd3f650b5282af37")
}

func (VRFCoordinatorV2PlusCoordinatorRegistered) Topic() common.Hash {
	return common.HexToHash("0xb7cabbfc11e66731fc77de0444614282023bcbd41d16781c753a431d0af01625")
}

func (VRFCoordinatorV2PlusEthFundsRecovered) Topic() common.Hash {
	return common.HexToHash("0x879c9ea2b9d5345b84ccd12610b032602808517cebdb795007f3dcb4df377317")
}

func (VRFCoordinatorV2PlusFundsRecovered) Topic() common.Hash {
	return common.HexToHash("0x59bfc682b673f8cbf945f1e454df9334834abf7dfe7f92237ca29ecb9b436600")
}

func (VRFCoordinatorV2PlusMigrationCompleted) Topic() common.Hash {
	return common.HexToHash("0xd63ca8cb945956747ee69bfdc3ea754c24a4caf7418db70e46052f7850be4187")
}

func (VRFCoordinatorV2PlusOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (VRFCoordinatorV2PlusOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (VRFCoordinatorV2PlusProvingKeyDeregistered) Topic() common.Hash {
	return common.HexToHash("0x72be339577868f868798bac2c93e52d6f034fef4689a9848996c14ebb7416c0d")
}

func (VRFCoordinatorV2PlusProvingKeyRegistered) Topic() common.Hash {
	return common.HexToHash("0xe729ae16526293f74ade739043022254f1489f616295a25bf72dfb4511ed73b8")
}

func (VRFCoordinatorV2PlusRandomWordsFulfilled) Topic() common.Hash {
	return common.HexToHash("0x7bad4a5a5566568e03ff9e9fdaedf6b8a760aa3ffc0abad348cf93787f04e17e")
}

func (VRFCoordinatorV2PlusRandomWordsRequested) Topic() common.Hash {
	return common.HexToHash("0xeb0e3652e0f44f417695e6e90f2f42c99b65cd7169074c5a654b16b9748c3a4e")
}

func (VRFCoordinatorV2PlusSubscriptionCanceled) Topic() common.Hash {
	return common.HexToHash("0x8c74ce8b8cf87f5eb001275c8be27eb34ea2b62bfab6814fcc62192bb63e81c4")
}

func (VRFCoordinatorV2PlusSubscriptionConsumerAdded) Topic() common.Hash {
	return common.HexToHash("0x1e980d04aa7648e205713e5e8ea3808672ac163d10936d36f91b2c88ac1575e1")
}

func (VRFCoordinatorV2PlusSubscriptionConsumerRemoved) Topic() common.Hash {
	return common.HexToHash("0x32158c6058347c1601b2d12bc696ac6901d8a9a9aa3ba10c27ab0a983e8425a7")
}

func (VRFCoordinatorV2PlusSubscriptionCreated) Topic() common.Hash {
	return common.HexToHash("0x1d3015d7ba850fa198dc7b1a3f5d42779313a681035f77c8c03764c61005518d")
}

func (VRFCoordinatorV2PlusSubscriptionFunded) Topic() common.Hash {
	return common.HexToHash("0x1ced9348ff549fceab2ac57cd3a9de38edaaab274b725ee82c23e8fc8c4eec7a")
}

func (VRFCoordinatorV2PlusSubscriptionFundedWithEth) Topic() common.Hash {
	return common.HexToHash("0x3f1ddc3ab1bdb39001ad76ca51a0e6f57ce6627c69f251d1de41622847721cde")
}

func (VRFCoordinatorV2PlusSubscriptionOwnerTransferRequested) Topic() common.Hash {
	return common.HexToHash("0x21a4dad170a6bf476c31bbcf4a16628295b0e450672eec25d7c93308e05344a1")
}

func (VRFCoordinatorV2PlusSubscriptionOwnerTransferred) Topic() common.Hash {
	return common.HexToHash("0xd4114ab6e9af9f597c52041f32d62dc57c5c4e4c0d4427006069635e216c9386")
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2Plus) Address() common.Address {
	return _VRFCoordinatorV2Plus.address
}

type VRFCoordinatorV2PlusInterface interface {
	BLOCKHASHSTORE(opts *bind.CallOpts) (common.Address, error)

	LINK(opts *bind.CallOpts) (common.Address, error)

	LINKETHFEED(opts *bind.CallOpts) (common.Address, error)

	MAXCONSUMERS(opts *bind.CallOpts) (uint16, error)

	MAXNUMWORDS(opts *bind.CallOpts) (uint32, error)

	MAXREQUESTCONFIRMATIONS(opts *bind.CallOpts) (uint16, error)

	GetRequestConfig(opts *bind.CallOpts) (uint16, uint32, [][32]byte, error)

	GetSubscription(opts *bind.CallOpts, subId *big.Int) (GetSubscription,

		error)

	HashOfKey(opts *bind.CallOpts, publicKey [2]*big.Int) ([32]byte, error)

	MigrationVersion(opts *bind.CallOpts) (uint8, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	PendingRequestExists(opts *bind.CallOpts, subId *big.Int) (bool, error)

	SConfig(opts *bind.CallOpts) (SConfig,

		error)

	SCurrentSubNonce(opts *bind.CallOpts) (uint64, error)

	SFallbackWeiPerUnitLink(opts *bind.CallOpts) (*big.Int, error)

	SFeeConfig(opts *bind.CallOpts) (SFeeConfig,

		error)

	SProvingKeyHashes(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error)

	SProvingKeys(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error)

	SRequestCommitments(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error)

	STotalBalance(opts *bind.CallOpts) (*big.Int, error)

	STotalEthBalance(opts *bind.CallOpts) (*big.Int, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AcceptSubscriptionOwnerTransfer(opts *bind.TransactOpts, subId *big.Int) (*types.Transaction, error)

	AddConsumer(opts *bind.TransactOpts, subId *big.Int, consumer common.Address) (*types.Transaction, error)

	CancelSubscription(opts *bind.TransactOpts, subId *big.Int, to common.Address) (*types.Transaction, error)

	CreateSubscription(opts *bind.TransactOpts) (*types.Transaction, error)

	DeregisterMigratableCoordinator(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error)

	DeregisterProvingKey(opts *bind.TransactOpts, publicProvingKey [2]*big.Int) (*types.Transaction, error)

	FulfillRandomWords(opts *bind.TransactOpts, proof VRFProof, rc VRFCoordinatorV2PlusRequestCommitment) (*types.Transaction, error)

	FundSubscriptionWithEth(opts *bind.TransactOpts, subId *big.Int) (*types.Transaction, error)

	Migrate(opts *bind.TransactOpts, subId *big.Int, newCoordinator common.Address) (*types.Transaction, error)

	OnTokenTransfer(opts *bind.TransactOpts, arg0 common.Address, amount *big.Int, data []byte) (*types.Transaction, error)

	OracleWithdraw(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error)

	OracleWithdrawEth(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error)

	OwnerCancelSubscription(opts *bind.TransactOpts, subId *big.Int) (*types.Transaction, error)

	RecoverEthFunds(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	RecoverFunds(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	RegisterMigratableCoordinator(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error)

	RegisterProvingKey(opts *bind.TransactOpts, oracle common.Address, publicProvingKey [2]*big.Int) (*types.Transaction, error)

	RemoveConsumer(opts *bind.TransactOpts, subId *big.Int, consumer common.Address) (*types.Transaction, error)

	RequestRandomWords(opts *bind.TransactOpts, req VRFV2PlusClientRandomWordsRequest) (*types.Transaction, error)

	RequestSubscriptionOwnerTransfer(opts *bind.TransactOpts, subId *big.Int, newOwner common.Address) (*types.Transaction, error)

	SetConfig(opts *bind.TransactOpts, minimumRequestConfirmations uint16, maxGasLimit uint32, stalenessSeconds uint32, gasAfterPaymentCalculation uint32, fallbackWeiPerUnitLink *big.Int, feeConfig VRFCoordinatorV2PlusFeeConfig) (*types.Transaction, error)

	SetLINKAndLINKETHFeed(opts *bind.TransactOpts, link common.Address, linkEthFeed common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterConfigSet(opts *bind.FilterOpts) (*VRFCoordinatorV2PlusConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*VRFCoordinatorV2PlusConfigSet, error)

	FilterCoordinatorDeregistered(opts *bind.FilterOpts) (*VRFCoordinatorV2PlusCoordinatorDeregisteredIterator, error)

	WatchCoordinatorDeregistered(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusCoordinatorDeregistered) (event.Subscription, error)

	ParseCoordinatorDeregistered(log types.Log) (*VRFCoordinatorV2PlusCoordinatorDeregistered, error)

	FilterCoordinatorRegistered(opts *bind.FilterOpts) (*VRFCoordinatorV2PlusCoordinatorRegisteredIterator, error)

	WatchCoordinatorRegistered(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusCoordinatorRegistered) (event.Subscription, error)

	ParseCoordinatorRegistered(log types.Log) (*VRFCoordinatorV2PlusCoordinatorRegistered, error)

	FilterEthFundsRecovered(opts *bind.FilterOpts) (*VRFCoordinatorV2PlusEthFundsRecoveredIterator, error)

	WatchEthFundsRecovered(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusEthFundsRecovered) (event.Subscription, error)

	ParseEthFundsRecovered(log types.Log) (*VRFCoordinatorV2PlusEthFundsRecovered, error)

	FilterFundsRecovered(opts *bind.FilterOpts) (*VRFCoordinatorV2PlusFundsRecoveredIterator, error)

	WatchFundsRecovered(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusFundsRecovered) (event.Subscription, error)

	ParseFundsRecovered(log types.Log) (*VRFCoordinatorV2PlusFundsRecovered, error)

	FilterMigrationCompleted(opts *bind.FilterOpts) (*VRFCoordinatorV2PlusMigrationCompletedIterator, error)

	WatchMigrationCompleted(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusMigrationCompleted) (event.Subscription, error)

	ParseMigrationCompleted(log types.Log) (*VRFCoordinatorV2PlusMigrationCompleted, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFCoordinatorV2PlusOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*VRFCoordinatorV2PlusOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFCoordinatorV2PlusOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*VRFCoordinatorV2PlusOwnershipTransferred, error)

	FilterProvingKeyDeregistered(opts *bind.FilterOpts, oracle []common.Address) (*VRFCoordinatorV2PlusProvingKeyDeregisteredIterator, error)

	WatchProvingKeyDeregistered(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusProvingKeyDeregistered, oracle []common.Address) (event.Subscription, error)

	ParseProvingKeyDeregistered(log types.Log) (*VRFCoordinatorV2PlusProvingKeyDeregistered, error)

	FilterProvingKeyRegistered(opts *bind.FilterOpts, oracle []common.Address) (*VRFCoordinatorV2PlusProvingKeyRegisteredIterator, error)

	WatchProvingKeyRegistered(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusProvingKeyRegistered, oracle []common.Address) (event.Subscription, error)

	ParseProvingKeyRegistered(log types.Log) (*VRFCoordinatorV2PlusProvingKeyRegistered, error)

	FilterRandomWordsFulfilled(opts *bind.FilterOpts, requestId []*big.Int, subID []*big.Int) (*VRFCoordinatorV2PlusRandomWordsFulfilledIterator, error)

	WatchRandomWordsFulfilled(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusRandomWordsFulfilled, requestId []*big.Int, subID []*big.Int) (event.Subscription, error)

	ParseRandomWordsFulfilled(log types.Log) (*VRFCoordinatorV2PlusRandomWordsFulfilled, error)

	FilterRandomWordsRequested(opts *bind.FilterOpts, keyHash [][32]byte, subId []*big.Int, sender []common.Address) (*VRFCoordinatorV2PlusRandomWordsRequestedIterator, error)

	WatchRandomWordsRequested(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusRandomWordsRequested, keyHash [][32]byte, subId []*big.Int, sender []common.Address) (event.Subscription, error)

	ParseRandomWordsRequested(log types.Log) (*VRFCoordinatorV2PlusRandomWordsRequested, error)

	FilterSubscriptionCanceled(opts *bind.FilterOpts, subId []*big.Int) (*VRFCoordinatorV2PlusSubscriptionCanceledIterator, error)

	WatchSubscriptionCanceled(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusSubscriptionCanceled, subId []*big.Int) (event.Subscription, error)

	ParseSubscriptionCanceled(log types.Log) (*VRFCoordinatorV2PlusSubscriptionCanceled, error)

	FilterSubscriptionConsumerAdded(opts *bind.FilterOpts, subId []*big.Int) (*VRFCoordinatorV2PlusSubscriptionConsumerAddedIterator, error)

	WatchSubscriptionConsumerAdded(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusSubscriptionConsumerAdded, subId []*big.Int) (event.Subscription, error)

	ParseSubscriptionConsumerAdded(log types.Log) (*VRFCoordinatorV2PlusSubscriptionConsumerAdded, error)

	FilterSubscriptionConsumerRemoved(opts *bind.FilterOpts, subId []*big.Int) (*VRFCoordinatorV2PlusSubscriptionConsumerRemovedIterator, error)

	WatchSubscriptionConsumerRemoved(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusSubscriptionConsumerRemoved, subId []*big.Int) (event.Subscription, error)

	ParseSubscriptionConsumerRemoved(log types.Log) (*VRFCoordinatorV2PlusSubscriptionConsumerRemoved, error)

	FilterSubscriptionCreated(opts *bind.FilterOpts, subId []*big.Int) (*VRFCoordinatorV2PlusSubscriptionCreatedIterator, error)

	WatchSubscriptionCreated(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusSubscriptionCreated, subId []*big.Int) (event.Subscription, error)

	ParseSubscriptionCreated(log types.Log) (*VRFCoordinatorV2PlusSubscriptionCreated, error)

	FilterSubscriptionFunded(opts *bind.FilterOpts, subId []*big.Int) (*VRFCoordinatorV2PlusSubscriptionFundedIterator, error)

	WatchSubscriptionFunded(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusSubscriptionFunded, subId []*big.Int) (event.Subscription, error)

	ParseSubscriptionFunded(log types.Log) (*VRFCoordinatorV2PlusSubscriptionFunded, error)

	FilterSubscriptionFundedWithEth(opts *bind.FilterOpts, subId []*big.Int) (*VRFCoordinatorV2PlusSubscriptionFundedWithEthIterator, error)

	WatchSubscriptionFundedWithEth(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusSubscriptionFundedWithEth, subId []*big.Int) (event.Subscription, error)

	ParseSubscriptionFundedWithEth(log types.Log) (*VRFCoordinatorV2PlusSubscriptionFundedWithEth, error)

	FilterSubscriptionOwnerTransferRequested(opts *bind.FilterOpts, subId []*big.Int) (*VRFCoordinatorV2PlusSubscriptionOwnerTransferRequestedIterator, error)

	WatchSubscriptionOwnerTransferRequested(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusSubscriptionOwnerTransferRequested, subId []*big.Int) (event.Subscription, error)

	ParseSubscriptionOwnerTransferRequested(log types.Log) (*VRFCoordinatorV2PlusSubscriptionOwnerTransferRequested, error)

	FilterSubscriptionOwnerTransferred(opts *bind.FilterOpts, subId []*big.Int) (*VRFCoordinatorV2PlusSubscriptionOwnerTransferredIterator, error)

	WatchSubscriptionOwnerTransferred(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusSubscriptionOwnerTransferred, subId []*big.Int) (event.Subscription, error)

	ParseSubscriptionOwnerTransferred(log types.Log) (*VRFCoordinatorV2PlusSubscriptionOwnerTransferred, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
