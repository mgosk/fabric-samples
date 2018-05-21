Publishing chaincode in dev mode

    CORE_PEER_ADDRESS=peer:7052 CORE_CHAINCODE_ID_NAME=dee:0 ./dee


installing chaincode

    peer chaincode install -p chaincodedev/chaincode/dee -n dee -v 0
    peer chaincode instantiate -n dee -v 0 -c '{"Args":["a","10"]}' -C myc
    
