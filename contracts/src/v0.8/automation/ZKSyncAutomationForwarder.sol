// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.16;

import {IAutomationRegistryConsumer} from "./interfaces/IAutomationRegistryConsumer.sol";
import {GAS_BOUND_CALLER, IGasBoundCaller} from "./interfaces/zksync/IGasBoundCaller.sol";

uint256 constant PERFORM_GAS_CUSHION = 5_000;

/**
 * @title ZKSyncAutomationForwarder is a relayer that sits between the registry and the customer's target contract
 * @dev The purpose of the forwarder is to give customers a consistent address to authorize against,
 * which stays consistent between migrations. The Forwarder also exposes the registry address, so that users who
 * want to programmatically interact with the registry (ie top up funds) can do so.
 */
contract ZKSyncAutomationForwarder {
  error InvalidCaller(address);

  // @notice the user's target contract address
  address private immutable i_target;

  // @notice the shared logic address
  address private immutable i_logic;

  IAutomationRegistryConsumer private s_registry;

  // for testing purpose, remove later
  event GasDetails(
    uint256 indexed pubdataUsed,
    uint256 indexed gasPerPubdataByte,
    uint256 indexed executionGasUsed,
    uint256 gasprice
  );

  constructor(address target, address registry, address logic) {
    s_registry = IAutomationRegistryConsumer(registry);
    i_target = target;
    i_logic = logic;
  }

  /**
   * @notice forward is called by the registry and forwards the call to the target
   * @param gasAmount is the amount of gas to use in the call
   * @param data is the 4 bytes function selector + arbitrary function data
   * @return success indicating whether the target call succeeded or failed
   */
  function forward(
    uint256 gasAmount,
    bytes memory data
  ) external returns (bool success, uint256 gasUsed, uint256 l1GasUsed) {
    if (msg.sender != address(s_registry)) revert InvalidCaller(msg.sender);
    address target = i_target;

    uint256 g1 = gasleft();
    bytes memory returnData;
    (success, returnData) = GAS_BOUND_CALLER.delegatecall{gas: gasAmount}(
      abi.encodeWithSelector(IGasBoundCaller.gasBoundCall.selector, target, gasAmount, data)
    );
    uint256 pubdataGasSpent;
    if (success) {
      (, pubdataGasSpent) = abi.decode(returnData, (bytes, uint256));
    }
    uint256 g2 = gasleft();
    gasUsed = g1 - g2;
    return (success, gasUsed, pubdataGasSpent);
  }

  function getTarget() external view returns (address) {
    return i_target;
  }

  fallback() external payable {
    // copy to memory for assembly access
    address logic = i_logic;
    // copied directly from OZ's Proxy contract
    assembly {
      // Copy msg.data. We take full control of memory in this inline assembly
      // block because it will not return to Solidity code. We overwrite the
      // Solidity scratch pad at memory position 0.
      calldatacopy(0, 0, calldatasize())
      // out and outsize are 0 because we don't know the size yet.
      let result := delegatecall(gas(), logic, 0, calldatasize(), 0, 0)
      // Copy the returned data.
      returndatacopy(0, 0, returndatasize())
      switch result
      // delegatecall returns 0 on error.
      case 0 {
        revert(0, returndatasize())
      }
      default {
        return(0, returndatasize())
      }
    }
  }
}
