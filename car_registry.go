// 
// A sample Blockchain application to store car registration number
// 

package main

import (
    "bytes" 
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// A few constants

const TABLE_REGISTRATION 	= "REGISTRATION"

const COLUMN_REGISTRATION	= "REGISTRATION"
const COLUMN_OWNER			= "OWNER"


type CarRegistration struct {
}

// Init - Initialize the chaincode

func (t *CarRegistration) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	
	fmt.Printf("In CarRegistration.Init: %s, %v", function, args);
	
	if len(args) != 0 {
		return nil, errors.New("Incorrect number of arguments. Expecting 0.")
	}
	
	// Creating the table to store the Car Reg and Owner
	
	var columnDefsRegTable []*shim.ColumnDefinition
	columnOneRegTableDef := shim.ColumnDefinition{
		Name: COLUMN_REGISTRATION, Type: shim.ColumnDefinition_STRING, Key: true}
	columnTwoRegTableDef := shim.ColumnDefinition{
		Name: COLUMN_OWNER, Type: shim.ColumnDefinition_STRING, Key: false}

	columnDefsRegTable = append(columnDefsRegTable, &columnOneRegTableDef)
	columnDefsRegTable = append(columnDefsRegTable, &columnTwoRegTableDef)
	
	err := stub.CreateTable(TABLE_REGISTRATION, columnDefsRegTable)
	if err != nil {
		return nil, fmt.Errorf("Error creating the registration table: %s", err)
	}
	
	return nil, nil
}

// Query - query the chaincode data

func (t *CarRegistration) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	
	fmt.Printf("In CarRegistration.Query: %s, %v", function, args);
	
	if function == "GetCarOwner" {
        // update balance
        return t.GetCarOwner(stub, args)
    
	} else {
		return nil, errors.New("Received unknown function invocation")
	}
}

// Invoke - modify the chaincode data

func (t *CarRegistration) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	
	fmt.Printf("In CarRegistration.Invoke: %s, %v", function, args);
	
	if function == "RegisterCar" {
        return t.RegisterCar(stub, args)
    
	} else if function == "TransferCar" {
        return t.TransferCar(stub, args)
	
	} else {
		return nil, errors.New("Received unknown function invocation")
	}
}

// Business Logic functions

// Retrieve the owner of a car

func (t *CarRegistration) GetCarOwner(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	
	fmt.Printf("In CarRegistration.GetCarOwner: %v", args);
	
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1: car registration")
	}
	
 	// Getting the row from the registration
 	
 	var key []shim.Column
 	reg := shim.Column{ Value: &shim.Column_String_{ String_: args[0] } }
 	key = append(key, reg)
 	row, err := stub.GetRow(TABLE_REGISTRATION, key)
 	
 	if err != nil {
	 	return nil, errors.New(fmt.Sprintf("Failed to retrieve record for reg '%s': %s", args[0], err))
 	}
 	
 	// Extracting owner if found and returning
 	
 	var jsonResp string;
 	var buffer bytes.Buffer;
 	if (len(row.Columns) != 0) {
 		
 		owner := row.Columns[1].GetString_()
 		
 		buffer.WriteString("{ \"registered\": true, \"owner\": " + owner + " }")
 		//buffer.WriteString(owner)
 		//buffer.WriteString(" }")
 		jsonResp = buffer.String()

 	} else {
 		jsonResp = "{ \"registered\": false }"
 	}
 	
 	return []byte(jsonResp), err

}

// Initially registers a car against an owner

func (t *CarRegistration) RegisterCar(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	
	fmt.Printf("In CarRegistration.RegisterCar: %v", args);
	
	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2: car registration and name of the owner")
	}
	
	// Registers the car - you may want to check first that the car is not already registered
	
	_, err := stub.InsertRow(TABLE_REGISTRATION, shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: args[0]}},
			&shim.Column{Value: &shim.Column_String_{String_: args[1]}},
		},
	})
		
	if err != nil {
		return nil, err
	}
	
	return nil, nil
}

// Transfers a car from an owner to another

func (t *CarRegistration) TransferCar(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	
	fmt.Printf("In CarRegistration.TransferCar: %v", args);
	
	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2: car registration and name of the new owner")
	}
	
	// Transfers the car ownership
	
	updated, err := stub.ReplaceRow(TABLE_REGISTRATION, shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: args[0]}},
			&shim.Column{Value: &shim.Column_String_{String_: args[1]}},
		},
	})
		
	if err != nil {
		return nil, err
	}
	
	if updated != true {
		return nil, errors.New(fmt.Sprintf("No car with registration number '%s'", args[0]))
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
