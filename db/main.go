package db

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// register table type init
type Data struct {
	URL         string `json:"url"`
	ProductInfo struct {
		ID        int      `json:"id"`
		Title     string   `json:"title"`
		Subtitle  string   `json:"subtitle"`
		Inventory string   `json:"inventory"`
		Options   []string `json:"options"`
		Price     string   `json:"price"`
		Image     []string `json:"image"`
	} `json:"product_info"`
	RegisterInfo struct {
		Date     string `json:"date"`
		Userinfo struct {
			UserName     string `json:"user_name"`
			UserUsername string `json:"user_username"`
			UserEmail    string `json:"user_email"`
		} `json:"userinfo"`
		RegisterID string `json:"register_id"`
	} `json:"register_info"`
}

// Get table items from JSON file
func getItems() []Data {
	raw, err := ioutil.ReadFile("./initData.json")

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var items []Data
	json.Unmarshal(raw, &items)
	return items
}

func Init(svc *dynamodb.DynamoDB) error {
	// Get table items from initData.json
	items := getItems()

	for _, item := range items {
		av, err := dynamodbattribute.MarshalMap(item)

		if err != nil {
			fmt.Println("Got error marshalling map:")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		// Create item in table Movies
		input := &dynamodb.PutItemInput{
			Item:      av,
			TableName: aws.String("SP_product_register"),
		}

		_, err = svc.PutItem(input)

		if err != nil {
			fmt.Println("Got error calling PutItem:")
			fmt.Println(err.Error())
			return err
		}

		fmt.Println("Successfully added '", item.ProductInfo.Title, "' (", item.ProductInfo.ID, ") to SP_product_register table")
	}
	return nil
}

func main() {
	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)

	if err != nil {
		fmt.Println("Error creating session:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	err = Init(svc)
	if err != nil {
		fmt.Println("Error has been created when initiating DB")
		os.Exit(1)
	}
}
