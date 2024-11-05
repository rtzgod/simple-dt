package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/rtzgod/simple-dt/internal/entity"
	"log/slog"
	"net/http"
)

type Client struct {
	log *slog.Logger
	Url string
}

func NewClient(log *slog.Logger, url string) *Client {
	return &Client{
		log: log,
		Url: url}
}

func (c *Client) SetData(value string) error {
	const op = "httpclient.SetData"

	c.log.With(slog.String("op", op), slog.String("url", c.Url), slog.String("value", value))

	c.log.Info("Setting data", slog.String("url", c.Url), slog.String("value", value))

	data := entity.Data{
		Value: value,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	resp, err := http.Post(c.Url+"/data/", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%s: %w", op, fmt.Errorf("unexpected status code: %d", resp.StatusCode))
	}

	return nil
}
