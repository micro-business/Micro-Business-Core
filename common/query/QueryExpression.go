package query

type RequestQuery struct {
	Key        string            `json:Key`
	Children   []RequestQuery    `json:Children`
	Attributes map[string]string `json:Attributes`
}

type ResponseQuery struct {
	Key      string          `json:Key`
	Value    string          `json:Value`
	Children []ResponseQuery `json:Children`
}
