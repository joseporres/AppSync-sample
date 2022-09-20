package main

import (
	"context"
	"fmt"
	"testing"
)
func TestHandler(t *testing.T) {
    t.Run("success request", func(t *testing.T) {
        k,err := handler(context.TODO(), Event{Username: "e57b838a-e1db-4205-81f2-566d14b029a2"})
		fmt.Println(k)
		fmt.Println(err)
    })
}