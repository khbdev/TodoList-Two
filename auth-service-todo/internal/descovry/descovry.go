package descovry

import (
	"auth-service/pkg/env"
	"fmt"
	"strings"
)




func GetServiceAddress(serviceName string) (string, error) {
	envKey := strings.ToUpper(strings.ReplaceAll(serviceName, "-", "_"))

	addr := env.LoadEnvKEY(envKey)
	if addr == "" {
		return "", fmt.Errorf("%s env topilmadi", envKey)
	}
	return addr, nil
}