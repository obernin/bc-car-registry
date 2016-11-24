// 
// A sample Blockchain application to store car registration number
// 

package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/op/go-logging"
)

var logger = logging.MustGetLogger("CarRegistration")


type CarRegistration struct {
}

// Init - Initialize the chaincode

func (t *CarRegistration) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	
	logger.Infof("In CarRegistration.Init: %s, %v", function, args);
	
	return nil, nil
}

// Query - query the chaincode data

func (t *CarRegistration) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	
	logger.Infof("In CarRegistration.Query: %s, %v", function, args);
	
	return nil, nil
}

// Invoke - modify the chaincode data

func (t *CarRegistration) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	
	logger.Infof("In CarRegistration.Invoke: %s, %v", function, args);
	
	return nil, nil
}


// Main function - required to keep the compiler happy

func main() {
	
	err := shim.Start(new(CarRegistration))
	if err != nil {
		logger.Errorf("Error starting CarRegistration chaincode: %s", err)
	}
}
