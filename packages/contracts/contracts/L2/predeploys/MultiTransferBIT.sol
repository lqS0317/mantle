pragma solidity ^0.8.0;

contract MultiTransferBITContract {
    event MultiTransfer(
        address indexed _from,
        uint256 indexed _value,
        address _to,
        uint256 _amount
    );

    function multiTransfer(address[] calldata _addresses, uint256 _amount) public payable {
        uint256 toReturn = msg.value;
        for (uint256 i = 0; i < _addresses.length; i++) {
            payable(_addresses[i]).transfer(_amount);
            toReturn = toReturn - _amount;
            emit MultiTransfer(msg.sender, msg.value, _addresses[i], _amount);
        }
        payable(msg.sender).transfer(toReturn);
    }
}
