package main

import (
	"context"
	"fmt"
)

type SatanKey string

func main() {
	ctx0 := context.Background()

	var ctx1 context.Context

	ctx1 = context.WithValue(ctx0, SatanKey("name"), "benjamin")
	ctx1 = context.WithValue(ctx1, SatanKey("age"), 12)
	fmt.Println(ctx1.Value(SatanKey("name")))
	fmt.Println(ctx1.Value(SatanKey("age")))
}
