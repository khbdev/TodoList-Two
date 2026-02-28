package connection

import (
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
)



func Connection(serviceName string) *grpc.ClientConn{
	envKey := serviceName + "_URL"
	url := os.Getenv(envKey)
	if url == "" {
		log.Fatalf("Env %s is not found", serviceName)
	}

	var conn *grpc.ClientConn
	var err error

	for i := 1; i <= 3; i++ {
		ctx, cancel  := context.WithTimeout(context.Background(), 5*time.Second)
	  defer cancel()

	  conn, err := grpc.DialContext(ctx, url, grpc.WithInsecure(), grpc.WithBlock())
	}
}