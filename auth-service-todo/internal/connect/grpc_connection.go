package connect

import (
	"auth-service/internal/descovry"
	"context"
	"fmt"

	"log"
	"time"

	"google.golang.org/grpc"
)


func ConnectService(serviceName string) (*grpc.ClientConn, error) {
	addr, err := descovry.GetServiceAddress(serviceName)
	if err != nil {
		return nil, fmt.Errorf("%s address topilmadi: %v", serviceName, err)
	}

	var conn *grpc.ClientConn
	var dialErr error

	for i := 1; i <= 3; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		conn, dialErr = grpc.DialContext(ctx, addr, grpc.WithInsecure(), grpc.WithBlock())
		if dialErr == nil {
			log.Printf("[gRPC]  %s bilan ulanish o‘rnatildi: %s", serviceName, addr)
			return conn, nil
		}

		log.Printf("[gRPC]  %s bilan ulanish muvaffaqiyatsiz (urinish %d/3): %v", serviceName, i, dialErr)
		time.Sleep(1 * time.Second)
	}

	return nil, fmt.Errorf("%s bilan 3 marta ulanish muvaffaqiyatsiz: %v", serviceName, dialErr)
}