package engine

type ParseResult struct {
	Requesrts []Request
	Items []Item
}



type Item struct{

	Url string
	Type string
	Id string
	Payload interface{}
}

type Request struct{
	Url string
	ParseFunc func([]byte) ParseResult
}


func NilParse([]byte) ParseResult{
	return ParseResult{}
}