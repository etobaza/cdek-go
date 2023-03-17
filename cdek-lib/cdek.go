package cdek_lib

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// CalculatorResponse represents tariff
type CalculatorResponse struct {
	TariffCodes []Tariff `json:"tariff_codes"`
}

// Calculate returns tariffs
func (c *Client) Calculate(fromLocation, toLocation Location, size Size) ([]Tariff, error) {
	requestBody := fmt.Sprintf(`
	{
		"from_location": {
			"code": "%s",
			"postal_code": "%s",
			"country_code": "%s",
			"city": "%s",
			"address": "%s"
		},
		"to_location": {
			"code": "%s",
			"postal_code": "%s",
			"country_code": "%s",
			"city": "%s",
			"address": "%s"
		},
		"packages": [{"weight": %d, "length": %d, "width": %d, "height": %d}]
	}
	`, fromLocation.Code, fromLocation.PostalCode, fromLocation.CountryCode, fromLocation.City, fromLocation.Address, toLocation.Code, toLocation.PostalCode, toLocation.CountryCode, toLocation.City, toLocation.Address, size.Weight, size.Length, size.Width, size.Height)

	req, err := http.NewRequest("POST", c.ApiURL+"/v2/calculator/tarifflist", strings.NewReader(requestBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.AccessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get a response from CDEK API: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var calculatorResponse CalculatorResponse
	err = json.Unmarshal(body, &calculatorResponse)
	if err != nil {
		return nil, err
	}

	return calculatorResponse.TariffCodes, nil
}

// GetAccessToken returns access token for CDEK API
func GetAccessToken(apiURL, account, securePassword string) (string, error) {
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", account)
	data.Set("client_secret", securePassword)

	req, err := http.NewRequest("POST", apiURL+"/v2/oauth/token", strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get response from CDEK API: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var tokenResponse struct {
		AccessToken string `json:"access_token"`
	}

	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		return "", err
	}

	return tokenResponse.AccessToken, nil
}
