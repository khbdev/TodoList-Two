package connection

import (
	"log"
	"os"

	"google.golang.org/grpc"
)



func Connect(name string) *grpc.ClientConn{
	envKey := name + "_URL"
	url := os.Getenv(envKey)
	if url == "" {
		log.Fatalf("Env %d is not found", envKey)
	}
}