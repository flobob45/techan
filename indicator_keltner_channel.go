package techan

import (
	"github.com/sdcoffey/big"
)

type keltnerChannelIndicator struct {
	ema Indicator
	//sma    Indicator
	atr    Indicator
	mul    big.Decimal
	window int
}

func NewKeltnerChannelUpperIndicator(series *TimeSeries, window int, mult float64) Indicator {
	return keltnerChannelIndicator{
		atr: NewAverageTrueRangeIndicator(series, window),
		ema: NewEMAIndicator(NewClosePriceIndicator(series), window),
		//sma:    NewSimpleMovingAverage(NewClosePriceIndicator(series), window),
		mul:    big.NewDecimal(mult),
		window: window,
	}
}

func NewKeltnerChannelLowerIndicator(series *TimeSeries, window int, mult float64) Indicator {
	return keltnerChannelIndicator{
		atr: NewAverageTrueRangeIndicator(series, window),
		ema: NewEMAIndicator(NewClosePriceIndicator(series), window),
		//sma:    NewSimpleMovingAverage(NewClosePriceIndicator(series), window),
		mul:    big.NewDecimal(mult).Neg(),
		window: window,
	}
}

func (kci keltnerChannelIndicator) Calculate(index int) big.Decimal {
	if index <= kci.window-1 {
		return big.ZERO
	}

	//return kci.sma.Calculate(index).Add(kci.atr.Calculate(index).Mul(kci.mul))
	return kci.ema.Calculate(index).Add(kci.atr.Calculate(index).Mul(kci.mul))
}
