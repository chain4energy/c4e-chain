package types

//func NewMinter(newCpo, annualProvisions sdk.Dec) HalvingMinter {
//	//return HalvingMinter{
//	//	Inflation:        inflation,
//	//	AnnualProvisions: annualProvisions,
//	//}
//
//}

func (m HalvingMinter) NextCointCount(blockHeight int64) int64 {

	if blockHeight%m.BlocksPerYear == 0 {
		m.NewCoinsMint = m.NewCoinsMint / 2
	}
	return m.NewCoinsMint
}

func InitialHalvingMinter() HalvingMinter {
	return HalvingMinter{
		BlocksPerYear: 2000,
		NewCoinsMint:  100,
		MintDenom:     "C4E",
	}
}
