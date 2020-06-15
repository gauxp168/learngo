package engine

type ParseResult struct {
	Requests []Requset
	Items []interface{}
}
type Requset struct {
	Url string
	ParseFunc func([]byte) ParseResult
}

func NilParse([]byte) ParseResult {
	return ParseResult{}
}