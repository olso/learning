pragma solidity ^0.4.23;

import "openzeppelin-solidity/contracts/token/ERC20/ERC20.sol";

contract ProgressbarToken is ERC20 {
    string public constant name = "ProgressbarToken";
    string public constant symbol = "PBT";
    uint8 public constant decimals = 8;
    uint256 public constant INITIAL_SUPPLY = 10000 * (10 ** uint256(decimals));

    constructor() public {
        _mint(msg.sender, INITIAL_SUPPLY);
    }
}