package itvmsq

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func SendCode(mobileProvider, phoneNumber string, code int) error {
	message := fmt.Sprintf("hoopla kod: %d %s", code, "qKGewJ132BX")

	if mobileProvider == "uzmobile" {
		message = fmt.Sprintf("hoopla sizning ro'yhatdan o'tish uchun kodingiz: %d %s", code, "qKGewJ132BX")
	}

	payload := map[string]interface{}{
		"credential_provider": mobileProvider,
		"credential_type":     "phone_number",
		"credential_value":    phoneNumber,
		"message":             message,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Post("https://msq.itv.uz/api/v1/send", "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return fmt.Errorf("failed to send POST request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("result: %d", resp.StatusCode)
	}

	return nil
}
