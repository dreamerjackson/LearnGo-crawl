package engine

type ParseResult struct {
	Requesrts []Request
	Items []Item
}


type Parser interface {
	Parse(contents []byte,url string) ParseResult
	Serialize()(name string,args interface{})
}


type Item struct{

	Url string
	Type string
	Id string
	Payload interface{}
}

type Request struct{
	Url string
	Parse Parser
}

type Nilparse struct {

}

func (Nilparse) Parse(contents []byte, url string) ParseResult {
	return ParseResult{}
}

func (Nilparse) Serialize() (name string, args interface{}) {
	return "Nilparse",nil
}

type ParseFunc func(contents []byte,url string) ParseResult

type FuncParser struct{
	parseer ParseFunc
	name string
}

func ( f FuncParser) Parse(contents []byte, url string) ParseResult {
	return f.parseer(contents,url)
}

func (f FuncParser) Serialize() (name string, args interface{}) {
	return f.name,nil
}


func NewFuncparse(p ParseFunc,name string) *FuncParser{
	return &FuncParser{
		parseer:p,
		name:name,
	}
}