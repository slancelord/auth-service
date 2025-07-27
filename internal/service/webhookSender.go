package service

import (
	"auth-service/internal/config"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func SendWebhook(payload map[string]string) error {
	hookUrl := config.GetConfig().WebhookUrl
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	resp, err := http.Post(hookUrl, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("webhook returned status %d", resp.StatusCode)
	}

	return nil
}
