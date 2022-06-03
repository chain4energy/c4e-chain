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
		BlocksPerYear: 10,
		NewCoinsMint:  20596877,
		MintDenom:     "uc4e",
	}
}
