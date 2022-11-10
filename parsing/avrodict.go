package parsing

import "encoding/json"

type AvroDict struct {
	Meta struct {
		FileName        string `json:"file_name"`
		FileDescription string `json:"file_description"`
		Package         string `json:"package"`
		License         string `json:"license"`
		Source          string `json:"source"`
		AdaptedBy       string `json:"adapted_by"`
		Updated         string `json:"updated"`
		Encoding        string `json:"encoding"`
	} `json:"meta"`
	Data struct {
		Patterns []struct {
			Find    string `json:"find"`
			Replace string `json:"replace"`
			Rules   []struct {
				Matches []struct {
					Type  string `json:"type"`
					Scope string `json:"scope"`
					Value string `json:"value"`
				} `json:"matches"`
				Replace string `json:"replace"`
			} `json:"rules,omitempty"`
		} `json:"patterns"`
		Vowel         string `json:"vowel"`
		Consonant     string `json:"consonant"`
		CaseSensitive string `json:"casesensitive"`
		Number        string `json:"number"`
	} `json:"data"`
}

func ConvertToJSON() (*AvroDict, error) {
	inter := &AvroDict{}
	err := json.Unmarshal(data, inter)
	return inter, err
}
