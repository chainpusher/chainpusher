package dto

type JsonRpcEvent struct {
	Name string `json:"name"`

	Data interface{} `json:"data"`
}
