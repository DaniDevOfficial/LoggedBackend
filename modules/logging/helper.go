package logging

import (
	"encoding/json"
	"log"
)

func EncodePersonalInformation(input string) string {
	sensitiveKeys := []string{"password", "pw", "email"}

	var data interface{}
	if err := json.Unmarshal([]byte(input), &data); err != nil {
		return input
	}

	data = redactSensitiveFields(data, sensitiveKeys)

	jsonString, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
		return input
	}

	return string(jsonString)
}

func redactSensitiveFields(data interface{}, sensitiveKeys []string) interface{} {
	switch v := data.(type) {
	case map[string]interface{}:
		for key, value := range v {
			for _, sensitiveKey := range sensitiveKeys {
				if key == sensitiveKey {
					v[key] = "****"
				}
			}
			v[key] = redactSensitiveFields(value, sensitiveKeys)
		}
	case []interface{}:
		for i, value := range v {
			v[i] = redactSensitiveFields(value, sensitiveKeys)
		}
	}
	return data
}
