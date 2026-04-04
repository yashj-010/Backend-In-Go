package main

import (
	"context"
	"fmt"
)

func userinfo(ctx context.Context){
	user := ctx.Value("user").(string)
	role := ctx.Value("role").(string)
	fmt.Println("User:",user)
	fmt.Println("Role:",role)
	return
}

func ctx_data() {
	ctx := context.Background()
	ctx = context.WithValue(ctx,"user","yash")
	ctx = context.WithValue(ctx,"role","admin")
	userinfo(ctx)
}