// SPDX-License-Identifier: LGPL-3.0
pragma solidity 0.8.28;

interface IPrecompile {
    function Echo(string memory) external view returns (string memory);

    function Echo(uint256) external view returns (uint256);

    function HashPacked(uint256, bytes2, address) external returns (bytes32);

    struct Wrapper {
        int256 val;
    }

    function Extract(Wrapper memory) external returns (int256);

    function Self() external returns (address);

    function RevertWith(bytes memory) external;
}
