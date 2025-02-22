package utils

import (
	"crypto/rand"
	"encoding/hex"

	"example.com/shorturl/db"
)

func GenerateHexID(length int) (string, error) {
	bytes := make([]byte, length/2)
	_, err := rand.Read(bytes)

	if err != nil {
		return "", err
	}

	hexID := hex.EncodeToString(bytes)

	return hexID, nil
}

func UniqueCheck(key string) (bool, error) {
	rows, err := db.DB.Query("SELECT uniqueKey FROM urllist")
	if err != nil {
		return false, err
	}
	defer rows.Close()

	allkeys := make(map[string]bool)

	for rows.Next() {
		var uniqueKey string
		err = rows.Scan(&uniqueKey)
		if err != nil {
			return false, err
		}
		allkeys[uniqueKey] = true
	}

	if _, exists := allkeys[key]; exists {
		return false, nil
	}
	return true, nil
}

func ProvideUniqueKey() (string, error) {
	keyLength := 6

	for {
		key, err := GenerateHexID(keyLength)
		if err != nil {
			return "", err
		}

		isUnique, err := UniqueCheck(key)
		if err != nil {
			return "", err
		}

		if isUnique {
			return key, nil
		}
	}
}
