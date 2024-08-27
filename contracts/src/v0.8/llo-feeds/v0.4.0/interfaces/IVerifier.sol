// SPDX-License-Identifier: MIT
pragma solidity 0.8.19;

import {Common} from "../../libraries/Common.sol";

interface IVerifier {
  /**
   * @notice Verifies that the data encoded has been signed correctly using the signatures included within the payload.
   * @param signedReport The encoded data to be verified.
   * @param parameterPayload The encoded parameters to be used in the verification and billing process.
   * @param sender The address that requested to verify the contract.Used for logging and applying the fee.
   * @dev Verification is typically only done through the proxy contract so we can't just use msg.sender.
   * @return verifierResponse The encoded verified response.
   */
  function verify(
    bytes calldata signedReport,
    bytes calldata parameterPayload,
    address sender
  ) external payable returns (bytes memory verifierResponse);

  /**
   * @notice Bulk verifies that the data encoded has been signed correctly using the signatures included within the payload.
   * @param signedReports The encoded data to be verified.
   * @param parameterPayload The encoded parameters to be used in the verification and billing process.
   * @param sender The address that requested to verify the contract. Used for logging and applying the fee.
   * @dev Verification is typically only done through the proxy contract so we can't just use msg.sender.
   * @return verifiedReports The encoded verified responses.
   */
  function verifyBulk(
    bytes[] calldata signedReports,
    bytes calldata parameterPayload,
    address sender
  ) external payable returns (bytes[] memory verifiedReports);
}
