package main

import (
	"cdek-project/cdek-lib"
	"fmt"
)

func main() {
	// Credentials
	account := "EMscd6r9JnFiQ3bLoyjJY6eM78JrJceI"
	securePassword := "PjLZkKBHEiLK3YsjtNrt3TGNG0ahs3kG"
	apiURL := "https://api.edu.cdek.ru"
	accessToken, err := cdek_lib.GetAccessToken(apiURL, account, securePassword)
	if err != nil {
		fmt.Println("Error getting access token:", err)
		return
	}

	// Initialize client
	client := cdek_lib.Client{
		ApiURL:      apiURL,
		AccessToken: accessToken,
	}

	// Origin
	fromLocation := cdek_lib.Location{
		CountryCode: "RU",
		City:        "Москва",
		Address:     "Cлавянский бульвар д.1",
	}

	// Destination
	toLocation := cdek_lib.Location{
		CountryCode: "RU",
		City:        "Воронеж",
		Address:     "ул. Ленина д.43",
	}

	// Package size
	size := cdek_lib.Size{
		Weight: 1000,
		Length: 20,
		Width:  20,
		Height: 20,
	}

	// Calculate and get tariffs
	tariffs, err := client.Calculate(fromLocation, toLocation, size)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Tariff Codes:")
	for _, tariff := range tariffs {
		fmt.Printf("Code:%d | Name:%s | Description:%s | DeliveryMode:%d | Delivery Sum:%.2f | Period Min:%d | Period Max:%d\n",
			tariff.TariffCode, tariff.TariffName, tariff.TariffDescription, tariff.DeliveryMode, tariff.DeliverySum, tariff.PeriodMin, tariff.PeriodMax)
	}
}
