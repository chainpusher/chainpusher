package gorm_test

import (
	"testing"
)

func TestChargeRepository_FindChargingByTransactions(t *testing.T) {
	//db := test.SetupTestDB()
	//
	//var a1 *account.Account
	//db.First(&a1)
	//
	//w := charge.Wallet{Text: "1"}
	//p := charge.WalletPool{
	//	Wallets: []charge.Wallet{w},
	//}
	//
	//txs := shared.Slice[*transaction.Transaction]{
	//	{
	//		Payee:  "1",
	//		Amount: 10000,
	//	},
	//}
	//
	//repo := gorm.NewChargeRepository(db)
	//
	//c := &charge.Charge{
	//	Amount: 10000,
	//	//Pool:   p,
	//	Status: charge.Unpaid,
	//}
	//
	//db.Create(c)
	//
	//c, err := repo.Find(c.ID)
	//
	//charges, err := repo.FindChargingByTransactions(txs)
	//
	//assert.NotNil(t, err)
	//assert.NotNil(t, charges)
}
