package main

import (
	//"encoding/json"
	"fmt"
	//"bytes"
	//"strings"
	//"time"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"github.com/segmentio/ksuid"
	"encoding/json"
	"strconv"
)

type SmartContract struct {
}

type SimpleOrder struct {
	ObjectType string `json:"objectType"` //docType is used to distinguish the various types of objects in state database
	Id         string `json:"id"`         //order ID
	Resource   string `json:"resource"`   //traded resource energy at YYYY-MM-DDTHH
	BuyOrSell  string `json:"buyOrSell"`
	Amount     int    `json:"amount"`
	Owner      string `json:"owner"`
}

func main() {
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) peer.Response {
	//function, args := APIstub.GetFunctionAndParameters()
	//TODO implement body
	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) submitSimpleOrder(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 4  {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	objectType := "simpleOrder"
	id := genId()
	resource := args[0]
	buyOrSell := args[1]
	amount, err := strconv.Atoi(args[2])
	owner := args[3]

	// ==== Check if order already exists ====
	orderAsBytes, err := APIstub.GetState(id)
	if err != nil {
		return shim.Error("Failed to get order: " + err.Error())
	} else if orderAsBytes != nil {
		fmt.Println("This order already exists: " + id)
		return shim.Error("This order already exists: " + id)
	}

	// ==== Create order object and marshal to JSON ====
	order := &SimpleOrder{ObjectType: objectType, Id: id, Resource: resource, BuyOrSell: buyOrSell, Amount: amount, Owner: owner}
	orderJSONasBytes, err := json.Marshal(order)
	if err != nil {
		return shim.Error(err.Error())
	}

	// === Save order to state ===
	err = APIstub.PutState(id, orderJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	//TODO add to index
	//TODO emit event here

	return shim.Success(nil)
}


//private functions

//https://blog.kowalczyk.info/article/JyRZ/generating-good-unique-ids-in-go.html
func genId() string{
	id := ksuid.New()
	fmt.Printf("github.com/segmentio/ksuid:  %s\n", id.String())
	return id.String()
}