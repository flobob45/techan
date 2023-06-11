package techan

import (
	"github.com/sdcoffey/big"
)

const (
	SqzOn  = -1.0
	SqzOff = 1.0
	NoSqz  = 0.0
	SQZON  = "sqzon"
	SQZOFF = "sqzoff"
	NOSQZ  = "nosqz"
)

type squeezeMomentumTypeIndicator struct {
	lengthBB int
	lowerBB  Indicator
	upperBB  Indicator
	lowerKC  Indicator
	upperKC  Indicator
}

type squeezeMomentumValueIndicator struct {
	// PREREQUISITE FOR SQUEEZE MOMENTUM INDICATOR
	lowBid   Indicator
	highBid  Indicator
	closeBid Indicator
	lengthBB int
}

func NewSqueezeMomentumValueIndicator(lowBidIndicator Indicator, highBidIndicator Indicator, closeBidIndicator Indicator, lengthBB int) Indicator {
	return squeezeMomentumValueIndicator{
		lowBid:   lowBidIndicator,
		highBid:  highBidIndicator,
		closeBid: closeBidIndicator,
		lengthBB: lengthBB,
	}
}

func (smvi squeezeMomentumValueIndicator) Calculate(index int) big.Decimal {
	if index < smvi.lengthBB {
		return big.ZERO
	}
	xData := make([]big.Decimal, 0)
	yData := make([]big.Decimal, 0)
	for i := index - smvi.lengthBB + 1; i <= smvi.lengthBB; i++ {
		xData = append(xData, big.NewFromInt(i))
		yData = append(yData, smvi.intermediateSqueezeMomentumValueIndicator(i))
	}
	slope, _ := LeastSquaresMethod(xData, yData)
	return slope
}

func (smvi squeezeMomentumValueIndicator) intermediateSqueezeMomentumValueIndicator(index int) big.Decimal {
	//close - Moyenne (
	//	Moyenne ( Max(high, sur_longueur_20)  ,  Min(low, sur_longueur_20)  )
	//SimpleMovingAverage(close, 20)
	//)
	lowestLowBidIndicator := NewMinimumValueIndicator(smvi.lowBid, smvi.lengthBB)
	highestHighBidIndicator := NewMaximumValueIndicator(smvi.highBid, smvi.lengthBB)
	smaCloseBidIndicator := NewSimpleMovingAverage(smvi.closeBid, smvi.lengthBB)

	avgMaxMin := (highestHighBidIndicator.Calculate(index).Add(lowestLowBidIndicator.Calculate(index))).Div(big.NewDecimal(2.0))
	smaClose := smaCloseBidIndicator.Calculate(index)
	avgMaxMinClose := (avgMaxMin.Add(smaClose)).Div(big.NewDecimal(2.0))
	return smvi.closeBid.Calculate(index).Sub(avgMaxMinClose)
}

func NewSqueezeMomentumTypeIndicator(lowerBBIndicator Indicator, lowerKCIndicator Indicator, upperBBIndicator Indicator, upperKCIndicator Indicator, lengthBB int) Indicator {
	return squeezeMomentumTypeIndicator{
		lowerBB:  lowerBBIndicator,
		upperBB:  upperBBIndicator,
		lowerKC:  lowerKCIndicator,
		upperKC:  upperKCIndicator,
		lengthBB: lengthBB,
	}
}

func (smi squeezeMomentumTypeIndicator) Calculate(index int) big.Decimal {
	if index < smi.lengthBB {
		return big.NewDecimal(NoSqz)
	}
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

func SqzTypeString(sqzType big.Decimal) string {
	if sqzType.EQ(big.NewDecimal(SqzOn)) {
		return SQZON
	}
	if sqzType.EQ(big.NewDecimal(SqzOff)) {
		return SQZOFF
	}
	return NOSQZ
}

func LeastSquaresMethod(xData []big.Decimal, yData []big.Decimal) (slope big.Decimal, intercept big.Decimal) {
	type Point struct {
		x big.Decimal
		y big.Decimal
	}

	points := make([]Point, 0)
	for i, _ := range xData {
		points = append(points, Point{
			x: xData[i],
			y: yData[i],
		})
	}

	n := big.NewFromInt(len(points))

	sumX := big.ZERO
	sumY := big.ZERO
	sumXY := big.ZERO
	sumXX := big.ZERO

	for _, p := range points {
		sumX = sumX.Add(p.x)
		sumY = sumY.Add(p.y)
		sumXY = sumXY.Add(p.x.Mul(p.y))
		sumXX = sumXX.Add(p.x.Mul(p.x))
	}

	//base := (n*sumXX - sumX*sumX)
	//a := (n*sumXY - sumX*sumY) / base
	//b := (sumXX*sumY - sumXY*sumX) / base

	base := (n.Mul(sumXX)).Sub(sumX.Mul(sumX))
	a := ((n.Mul(sumXY)).Sub(sumX.Mul(sumY))).Div(base)
	b := (sumXX.Mul(sumY).Sub(sumXY.Mul(sumX))).Div(base)

	return a, b
}
