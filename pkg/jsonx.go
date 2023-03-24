package pkg

import jsoniter "github.com/json-iterator/go"

var (
	json = jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		UseNumber:              true, // 防止int转成float64
	}.Froze()
)

// ToJSON skip error
func ToJSON(data any) (res string) {
	res, _ = json.MarshalToString(data)
	return res
}

// MustToJSON No-skip error
func MustToJSON(data any) (string, error) {
	return json.MarshalToString(data)
}
