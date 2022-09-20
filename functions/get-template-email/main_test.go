package main

import (
	"context"
	"os"
	"testing"
)

func TestHandler(t *testing.T) {
	t.Run("success request", func(t *testing.T) {
		os.Setenv("REGION", "us-east-1")
		os.Setenv("BucketName", "ofvi-api-templatebucket-ze9qqqz1ddzw")

		_, err := handler(context.TODO(), Event{UserType: "Usuario Protecta"})
		if err != nil {
			t.Fatal("Error")
		}

	})

}
