package main

import (
	"context"
	"fmt"
	"testing"
)
func TestHandler(t *testing.T) {
    t.Run("success request", func(t *testing.T) {
        k,err := handler(context.TODO(), Event{Email: "japorres18@gmail.com"})
		fmt.Println(k)
		fmt.Println(err)
    })
}