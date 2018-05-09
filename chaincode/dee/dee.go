package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}


type Offer struct {
	ObjectType string `json:"docType"` //docType is used to distinguish the various types of objects in state database
	Id         string   `json:"id"`    //the fieldtags are needed to keep case from bouncing around
	Size       int    `json:"size"`
	Owner      string `json:"owner"`
}


type Trade struct {
	ObjectType string `json:"docType"` //docType is used to distinguish the various types of objects in state database

}



//type marble struct {
//	ObjectType string `json:"docType"` //docType is used to distinguish the various types of objects in state database
//	Name       string `json:"name"`    //the fieldtags are needed to keep case from bouncing around
//	Color      string `json:"color"`
//	Size       int    `json:"size"`
//	Owner      string `json:"owner"`
//}
