package models

// mapping from std.BaseAccount without pub key
type AccountInfo struct {
	Addr           string `json:"addr"`
	Balance        string `json:"balance"`
	AccountNumber  uint64 `json:"account_number"`
	SequenceNumber uint64 `json:"sequence"`
}
