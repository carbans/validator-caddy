package validator

import (
	"fmt"
	"net/http"
)

// ValidateRequest envía la petición a un servicio de validación externo
func ValidateRequest(r *http.Request) bool {
	validationURL := "http://validator-service:9000/check"

	req, err := http.NewRequest("POST", validationURL, r.Body)
	if err != nil {
		fmt.Println("Error creando request:", err)
		return false
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error enviando request:", err)
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK
}
