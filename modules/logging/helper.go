package logging

import (
	"encoding/json"
	"log"
)

func EncodePersonalInformation(input string) string {
	sensitiveKeys := []string{"password", "pw", "email", "token"}

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
		redactedInCycle := false
		for key, value := range v {
			redactedInCycle = false
			for _, sensitiveKey := range sensitiveKeys {
				if key == sensitiveKey {
					v[key] = "****"
					redactedInCycle = true
					break
				}
			}
			// this is so the redaction from above doesn't get overridden
			// I don't like this approach too much, but it does the job
			if !redactedInCycle {
				v[key] = redactSensitiveFields(value, sensitiveKeys)
			}
		}
		return v

	case []interface{}:
		for i, value := range v {
			v[i] = redactSensitiveFields(value, sensitiveKeys)
		}
		return v
	}

	return data
}
