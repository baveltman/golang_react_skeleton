package main

import (
	"encoding/json"
	"os"
)

func NewJsonContext(context interface{}) (JsonContext, error) {
	json, err := json.Marshal(context)
	if err != nil {
		return JsonContext{}, err
	}

	jsonContext := JsonContext{
		Json: string(json),
	}

	if os.Getenv(ENV_STAGE) == VALUE_PRODUCTION {
		jsonContext.IsProduction = true
	}

	return jsonContext, nil
}
