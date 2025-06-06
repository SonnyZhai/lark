package do

type KeysValues struct {
	Keys   []string `json:"keys"`
	Values []string `json:"values"`
}

type KeyMaps struct {
	Key  interface{}       `json:"key"`
	Maps map[string]string `json:"maps"`
	Ex   interface{}       `json:"ex"`
}

type KeyFieldValue struct {
	Key   interface{} `json:"key"`
	Field interface{} `json:"field"`
	Value interface{} `json:"value"`
}

type KeysFieldValues struct {
	Keys   []string `json:"keys"`
	Field  string   `json:"field"`
	Values []string `json:"values"`
}

type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
