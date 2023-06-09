package techan

import (
	"github.com/sdcoffey/big"
)

const (
	SqzOn  = -1.0
	SqzOff = 1.0
	NoSqz  = 0.0
)

type squeezeMomentumTypeIndicator struct {
	lowerBB Indicator
	upperBB Indicator
	lowerKC Indicator
	upperKC Indicator
}

type squeezeMomentumValueIndicator struct {
	// PREREQUISITE FOR SQUEEZE MOMENTUM INDICATOR
	//lowestLowBidIndicator := techan.NewMinimumValueIndicator(LowBidIndicator, lengthBB)
	//highestHighBidIndicator := techan.NewMaximumValueIndicator(HighBidIndicator, lengthBB)
	//SMACloseBidIndicator := techan.NewSimpleMovingAverage(CloseBidIndicator, lengthBB)
}

func NewSqueezeMomentumTypeIndicator(lowerBBIndicator Indicator, lowerKCIndicator Indicator, upperBBIndicator Indicator, upperKCIndicator Indicator) Indicator {
	return squeezeMomentumTypeIndicator{
		lowerBB: lowerBBIndicator,
		upperBB: upperBBIndicator,
		lowerKC: lowerKCIndicator,
		upperKC: upperKCIndicator,
	}
}

func (smi squeezeMomentumTypeIndicator) Calculate(index int) big.Decimal {
	sqzon := (smi.lowerBB.Calculate(index).GT(smi.lowerKC.Calculate(index))) && (smi.upperBB.Calculate(index).LT(smi.upperKC.Calculate(index)))
	sqzoff := (smi.lowerBB.Calculate(index).LT(smi.lowerKC.Calculate(index))) && (smi.upperBB.Calculate(index).GT(smi.upperKC.Calculate(index)))
	nosqz := (!sqzon) && (!sqzoff)

	if nosqz {
		return big.NewDecimal(NoSqz)
	}
	if sqzon {
		return big.NewDecimal(SqzOn)
	}
	return big.NewDecimal(SqzOff)
}

func IsSqzOn(indicator big.Decimal) bool {
	return indicator.EQ(big.NewDecimal(SqzOn))
}

func IsSqzOff(indicator big.Decimal) bool {
	return indicator.EQ(big.NewDecimal(SqzOff))
}

func IsNoSqz(indicator big.Decimal) bool {
	return indicator.EQ(big.NewDecimal(NoSqz))
}
