package db

import (
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/ahnsv/goCommerce/models"
	"github.com/aws/aws-sdk-go/aws"
	"fmt"

)

type Repository struct{}

const SERVER = ""

func (r *Repository) GetProducts() []Products {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)

	 params := &dynamodb.ScanInput{ 
		TableName: aws.String("SP_product_register")
	}
	result, err := svc.Scan(params)
	if err != nil {
	fmt.Errorf("failed to make Query API call, %v", err)
	} 
	 obj := []Products{}
	 err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &obj)
	if err != nil {
	fmt.Errorf("failed to unmarshal Query result items, %v", err)
	}
	return obj
}

