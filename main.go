package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type DynamoDBTable struct {
	AttributeDefinitions  []map[string]string
	KeySchema             []map[string]string
	ProvisionedThroughput map[string]int
	TableName             string
}

func main() {
	decoder := json.NewDecoder(os.Stdin)

	for {
		var dynamoDBTable DynamoDBTable
		if err := decoder.Decode(&dynamoDBTable); err != nil {
			// Exit the loop if EOF is reached or other errors occur
			break
		}

		// Append "DynamoDBTable" to the TableName
		dynamoDBTable.TableName += "DynamoDBTable"

		// Create a map in the AWS SAM CLI DynamoDB table resource format
		tableResource := map[string]interface{}{
			dynamoDBTable.TableName: map[string]interface{}{
				"Type": "AWS::DynamoDB::Table",
				"Properties": map[string]interface{}{
					"TableName":             dynamoDBTable.TableName,
					"AttributeDefinitions":  dynamoDBTable.AttributeDefinitions,
					"KeySchema":             dynamoDBTable.KeySchema,
					"ProvisionedThroughput": dynamoDBTable.ProvisionedThroughput,
				},
			},
		}

		// Convert the map to YAML
		yamlOutput, err := yaml.Marshal(tableResource)
		if err != nil {
			fmt.Println("Error converting to YAML:", err)
			return
		}

		// Write YAML output to standard output (stdout)
		fmt.Print(string(yamlOutput))
	}
}
