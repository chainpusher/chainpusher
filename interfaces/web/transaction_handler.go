package web

import (
	"github.com/chainpusher/chainpusher/interfaces/facade/impl"
	"net/http"
)

func GetTransactionHandler(writer http.ResponseWriter, request *http.Request) {
	serviceFacade := impl.NewTransactionServiceFacadeImpl()

	_, err := serviceFacade.GetTransactions(nil)

	if err != nil {

	}

	_, err = writer.Write([]byte("Hello"))
	if err != nil {
		return
	}
}
