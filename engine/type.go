package engine

import (
	"time"
)

const (
	FlagEnd  = 1
	VideoFlag  = 2

	TELEGRAM = "Telegram"

)
type ParseResult struct {
	Requesrts []*Request
	Items []interface{}
	RootRequesrt *RootRequest
	Flag  int32 		  //
	WebFlag string					  //
}

type Request struct{
	Url string
	ParseFunc func([]byte) ParseResult
	RootRequesrt *RootRequest
}

type RootRequest struct {
	Url string
	ParseFunc func([]byte) ParseResult
	WaitPeriod time.Duration
	ChildCount int64
	WebFlag  string
	DedupMap1 map[string]struct{}
	DedupMap2 map[string]struct{}
}

func NilParse([]byte) ParseResult{
	return ParseResult{}
}