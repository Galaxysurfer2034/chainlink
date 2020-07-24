pragma solidity ^0.6.0;

import './AggregatorValidatorInterface.sol';
import '../interfaces/FlagsInterface.sol';
import '../Owned.sol';
import '../vendor/SafeMath.sol';
import '../SignedSafeMath.sol';

/**
 * @title The Deviation Flagging Validator contract
 * @notice Checks the current value against the previous value, and makes sure
 * that it does not deviate outside of some relative range. If the deviation
 * threshold is passed then the validator raises a flag on the designated
 * flag contract.
 */
contract DeviationFlaggingValidator is Owned, AggregatorValidatorInterface {
  using SafeMath for uint256;
  using SignedSafeMath for int256;

  uint32 constant public THRESHOLD_MULTIPLIER = 100000;

  uint32 public flaggingThreshold;
  FlagsInterface public flags;

  event FlaggingThresholdUpdated(
    uint24 indexed previous,
    uint24 indexed current
  );
  event FlagsAddressUpdated(
    address indexed previous,
    address indexed current
  );

  int256 constant private INT256_MIN = -2**255;

  /**
   * @notice sets up the validator with its threshold and flag address.
   * @param _flags sets the address of the flags contract
   * @param _flaggingThreshold sets the threshold that will trigger a flag to be
   * raised. Setting the value of 100,000 is equivalent to tolerating a 100%
   * change compared to the previous price.
   */
  constructor(
    address _flags,
    uint24 _flaggingThreshold
  )
    public
  {
    setFlagsAddress(_flags);
    setFlaggingThreshold(_flaggingThreshold);
  }

  /**
   * @notice checks whether the parameters count as valid by comparing the
   * difference change to the flagging threshold.
   * @param _previousRoundId is ignored.
   * @param _previousAnswer is used as the median of the difference with the
   * current answer to determine if the deviation threshold has been exceeded.
   * @param _roundId is ignored.
   * @param _answer is the latest answer which is compared for a ratio of change
   * to make sure it has not execeeded the flagging threshold.
   */
  function validate(
    uint256 _previousRoundId,
    int256 _previousAnswer,
    uint256 _roundId,
    int256 _answer
  )
    external
    override
    returns (bool)
  {
    if (!isValid(_previousRoundId, _previousAnswer, _roundId, _answer)) {
      flags.raiseFlags(arrayifyMsgSender());
      return false;
    }

    return true;
  }

  /**
   * @notice checks whether the parameters count as valid by comparing the
   * difference change to the flagging threshold and raises a flag on the
   * flagging contract if so.
   * @param _previousAnswer is used as the median of the difference with the
   * current answer to determine if the deviation threshold has been exceeded.
   * @param _answer is the current answer which is compared for a ratio of
   * change * to make sure it has not execeeded the flagging threshold.
   */
  function isValid(
    uint256 ,
    int256 _previousAnswer,
    uint256 ,
    int256 _answer
  )
    public
    view
    returns (bool)
  {
    if (_previousAnswer == 0) return true;

    (int256 change, bool changeOk) = safeSub(_previousAnswer, _answer);
    (int256 ratioNumerator, bool numOk) = safeMul(change, THRESHOLD_MULTIPLIER);
    (int256 ratio, bool ratioOk) = safeDiv(ratioNumerator, _previousAnswer);

    return abs(ratio) <= flaggingThreshold && changeOk && numOk && ratioOk;
  }

  /**
   * @notice updates the flagging threshold
   * @param _flaggingThreshold sets the threshold that will trigger a flag to be
   * raised. Setting the value of 100,000 is equivalent to tolerating a 100%
   * change compared to the previous price.
   */
  function setFlaggingThreshold(uint24 _flaggingThreshold)
    public
    onlyOwner()
  {
    uint24 previousFT = uint24(flaggingThreshold);

    if (previousFT != _flaggingThreshold) {
      flaggingThreshold = _flaggingThreshold;

      emit FlaggingThresholdUpdated(previousFT, _flaggingThreshold);
    }
  }

  /**
   * @notice updates the flagging contract address for raising flags
   * @param _flags sets the address of the flags contract
   */
  function setFlagsAddress(address _flags)
    public
    onlyOwner()
  {
    address previous = address(flags);

    if (previous != _flags) {
      flags = FlagsInterface(_flags);

      emit FlagsAddressUpdated(previous, _flags);
    }
  }


  // PRIVATE

  function arrayifyMsgSender()
    private
    returns (address[] memory)
  {
      address[] memory addresses = new address[](1);
      addresses[0] = msg.sender;
      return addresses;
  }

  function abs(
    int256 value
  )
    private
    pure
    returns (uint256)
  {
    return uint256(value < 0 ? value.mul(-1): value);
  }

  /**
   * @dev Subtracts two signed integers, returns false 2nd param on overflow.
   * Modified version of OpenZeppelin's SignedSafeMath.
   */
  function safeSub(int256 a, int256 b) internal pure returns (int256, bool) {
    int256 c = a - b;
    if (!((b >= 0 && c <= a) || (b < 0 && c > a))) {
      return (0, false);
    }

    return (c, true);
  }


  /**
   * @dev Multiplies two signed integers, returns false 2nd param on overflow.
   * Modified version of OpenZeppelin's SignedSafeMath.
   */
  function safeMul(int256 a, int256 b) internal pure returns (int256, bool) {
    // Gas optimization: this is cheaper than requiring 'a' not being zero, but the
    // benefit is lost if 'b' is also tested.
    // See: https://github.com/OpenZeppelin/openzeppelin-contracts/pull/522
    if (a == 0) {
      return (0, true);
    }

    if (a == -1 && b == INT256_MIN) {
      return (0, false);
    }

    int256 c = a * b;
    if (!(c / a == b)) {
      return (0, false);
    }

    return (c, true);
  }

  /**
   * @dev Divides two signed integers, returns false 2nd param on overflow.
   * Modified version of OpenZeppelin's SignedSafeMath.
   */
  function safeDiv(int256 a, int256 b) internal pure returns (int256, bool) {
    // don't need to check for 0 as that is checked in isValid
    if (b == -1 && a == INT256_MIN) {
      return (0, false);
    }

    int256 c = a / b;

    return (c, true);
  }

}

