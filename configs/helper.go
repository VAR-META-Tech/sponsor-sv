package configs

import (
	"encoding/json"
	"log"
)

func PrettyPrint(T any) {

	// Marshal the struct into indented JSON bytes
	jsonData, err := json.MarshalIndent(T, "", "  ") // Use desired indentation string
	if err != nil {
		log.Println("Error marshalling JSON:", err)
		return
	}

	// Print the pretty-printed JSON string
	log.Println(string(jsonData))
}
