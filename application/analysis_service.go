package application

import "github.com/chainpusher/blockchain/model"

type AnalysisService interface {
	AnalyzeTrade(block *model.Block) error
}
