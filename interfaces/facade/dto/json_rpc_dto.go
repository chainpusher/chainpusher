package dto

type JsonRpcDto struct {
	Method string `json:"method"`

	Params []interface{} `json:"params"`
}
