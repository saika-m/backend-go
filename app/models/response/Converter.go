package response

import (
	"encoding/json"

	"github.com/Takina-Space/backend-go/libraries/httpResponse"
)

func Converter[R any](data any) (*R, error) {
	var result R
	b, err := json.Marshal(&data)
	if err != nil {
		httpResponse.InternalServerError(err)
	}
	err = json.Unmarshal(b, &result)
	if err != nil {
		httpResponse.InternalServerError(err)
	}
	return &result, nil
}
