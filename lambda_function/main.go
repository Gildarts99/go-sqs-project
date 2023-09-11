package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// func encrypt(plaintext string) []byte {
// 	secretKey, exist := os.LookupEnv("ENCRYPTION_KEY")
// 	if !exist {
// 		log.Fatal("environment variable ENCRYPTION_KEY not set")
// 	}

// 	block, err := aes.NewCipher([]byte(secretKey))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	aesgcm, err := cipher.NewGCM(block)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Never use more than 2^32 random nonces with a given key because of the risk of a repeat.
// 	nonce := make([]byte, aesgcm.NonceSize())
// 	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
// 		log.Fatal(err)
// 	}

// 	ciphertext := aesgcm.Seal(nonce, nonce, []byte(plaintext), nil)
// 	return ciphertext
// }

func handler(ctx context.Context, sqsEvent events.SQSEvent) (events.APIGatewayProxyResponse, error) {
	// // get the S3_BUCKET and AWS_REGION from an environment variable
	// bucket, exist := os.LookupEnv("S3_BUCKET")
	// if !exist {
	// 	log.Fatal("environment variable S3_BUCKET not set")
	// }

	// sdkConfig, err := config.LoadDefaultConfig(context.TODO())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// s3Client := s3.NewFromConfig(sdkConfig)

	// // for each record, ceaser cipher it and put it in s3 and a nobject with the path being the message id
	// for _, record := range sqsEvent.Records {
	// 	encrypted := encrypt(record.Body)

	// 	_, err := s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
	// 		Bucket: &bucket,
	// 		Key:    &record.MessageId,
	// 		Body:   strings.NewReader(string(encrypted)),
	// 	})

	// 	if err != nil {
	// 		log.Fatalln("Error: ", err)
	// 	}
	// }

	fmt.Println("hello")

	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "\"Success!\"",
	}

	return response, nil
}

func main() {
	lambda.Start(handler)
}
