package dto

type JsonRpcResponseDto struct {
	Call   *JsonRpcDto
	Result interface{}
}
