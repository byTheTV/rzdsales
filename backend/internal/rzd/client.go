package rzd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"rzd-sales/backend/internal/config"
	"rzd-sales/backend/internal/models"
)

// Client представляет клиент для работы с API РЖД
type Client struct {
	httpClient *http.Client
	config     *config.RZDConfig
}

// NewClient создает новый экземпляр клиента
func NewClient(cfg *config.RZDConfig) *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: cfg.Timeout,
		},
		config: cfg,
	}
}

// SearchStations выполняет поиск станций
func (c *Client) SearchStations(query string) ([]models.Station, error) {
	params := url.Values{}
	params.Add("layer_id", "5827")
	params.Add("stationNamePart", query)
	params.Add("compactMode", "y")

	url := fmt.Sprintf("%s?%s", c.config.BaseURL, params.Encode())
	log.Printf("Searching stations with URL: %s", url)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error searching stations: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
	}

	var result struct {
		Stations []models.Station `json:"stations"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result.Stations, nil
}

// SearchTrains выполняет поиск поездов
func (c *Client) SearchTrains(fromCode, toCode string, date time.Time) ([]models.Train, error) {
	params := url.Values{}
	params.Add("layer_id", "5827")
	params.Add("dir", "0")
	params.Add("tfl", "3")
	params.Add("checkSeats", "1")
	params.Add("code0", fromCode)
	params.Add("code1", toCode)
	params.Add("dt0", date.Format("02.01.2006"))

	url := fmt.Sprintf("%s/train_routes.php?%s", c.config.BaseURL, params.Encode())
	log.Printf("Searching trains with URL: %s", url)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error searching trains: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	log.Printf("Received response: %s", string(body))

	var result []models.Train
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
