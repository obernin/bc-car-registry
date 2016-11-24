// 
// A sample Blockchain application to store car registration number
// 

package main

import (
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)


type CarRegistration struct {
}

// Init - Initialize the chaincode

func (t *CarRegistration) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	
	fmt.Printf("In CarRegistration.Init: %s, %v", function, args);
	
	return nil, nil
}

// Query - query the chaincode data

func (t *CarRegistration) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	
	fmt.Printf("In CarRegistration.Query: %s, %v", function, args);
	
	if len(args) != 1 {
		return nil, errors.New("{ \"Error\": \"Incorrect number of arguments. Expecting 1: name of the key to read\" }")
	}
	
	valAsbytes, err := stub.GetState(args[0])
	if err != nil {
		return nil, errors.New(fmt.Sprintf("{ \"Error\": \"Error getting state for key '%s': %s\" }", args[0], err))
	}
	
	return valAsbytes, nil
}

// Invoke - modify the chaincode data

func (t *CarRegistration) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	
	fmt.Printf("In CarRegistration.Invoke: %s, %v", function, args);
	
	if len(args) != 2 {
		return nil, errors.New("{ \"Error\": \"Incorrect number of arguments. Expecting 2: name of the key and value to write\" }")
	}
	
	err := stub.PutState(args[0], []byte(args[1]))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("{ \"Error\": \"Error setting state for key '%s': %s\" }", args[0], err))
	}
	
	return nil, nil
}


// Main function - required to keep the compiler happy

func main() {
	
	err := shim.Start(new(CarRegistration))
	if err != nil {
		fmt.Printf("Error starting CarRegistration chaincode: %s", err)
	}
}
