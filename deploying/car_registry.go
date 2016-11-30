// 
// A sample Blockchain application to store car registration number
// 

package main

import (
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
	
	return nil, nil
}

// Invoke - modify the chaincode data

func (t *CarRegistration) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	
	fmt.Printf("In CarRegistration.Invoke: %s, %v", function, args);
	
	return nil, nil
}


// Main function - required to keep the compiler happy

func main() {
	
	err := shim.Start(new(CarRegistration))
	if err != nil {
		fmt.Printf("Error starting CarRegistration chaincode: %s", err)
	}
}
