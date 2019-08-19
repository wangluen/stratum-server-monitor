package models

import "encoding/json"

type StratumMsg struct {
	Method string      `json:"method"`
	Params []string    `json:"params"`
	ID     interface{} `json:"id"`
}

type BasicReply struct {
	ID     interface{} `json:"id"`
	Error  StratumErr  `json:"error,omitempty"`
	Result bool        `json:"result"`
}

// StratumRsp is the basic response type from stratum.
type StratumRsp struct {
	Method string           `json:"method"`
	ID     interface{}      `json:"id"`
	Error  StratumErr       `json:"error,omitempty"`
	Result *json.RawMessage `json:"result,omitempty"`
}

// NotifyRes models the json from a mining.notify message.
type NotifyRes struct {
	JobID          string
	Hash           string
	CoinbaseTX1    string
	CoinbaseTX2    string
	MerkleBranches []string
	BlockVersion   string
	Nbits          string
	Ntime          string
	CleanJobs      bool
}

type StratumErr struct {
	ErrNum uint64
	ErrStr string
	Result *json.RawMessage `json:"result,omitempty"`
}

// SubscribeReply models the server response to a subscribe message.
type SubscribeReply struct {
	SubscribeID       string
	ExtraNonce1       string
	ExtraNonce2Length float64
}