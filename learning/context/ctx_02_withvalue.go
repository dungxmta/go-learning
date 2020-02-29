package main

import (
	"context"
	"fmt"
)

// pass func's params with context
func main() {
	ctx := context.WithValue(context.Background(), "key1", "val1")
	First(ctx)
	fmt.Println("end main!")
}

func First(ctx context.Context) {
	if v := ctx.Value("key1"); v != nil {
		ctx = context.WithValue(ctx, "key2", "val2")
	}
	Second(ctx)
}

func Second(ctx context.Context) {
	v1 := ctx.Value("key1")
	v2 := ctx.Value("key2")

	fmt.Println(v1, " | ", v2)
}
