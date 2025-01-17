# Techan Release notes

## 0.13.3
* Back KC from sma to ema

## 0.13.2
* Fix KC from ema to sma

## 0.13.1
* Fix KC

## 0.13.0
* Fix Indicator Squeeze Momentum completed big to float internally

## 0.12.8
* Fix Indicator Squeeze Momentum completed

## 0.12.7
* Add Indicator Squeeze Momentum completed

## 0.12.6
* Add Indicator Squeeze Momentum.

## 0.12.5
* Update KeltnerChannel back to EMA

## 0.12.4
* Update KeltnerChannel ATR windows/2 vs SMA window

## 0.12.3
* Update KeltnerChannel based on SMA (not EMA)

## 0.12.2
* Customize NewKeltnerChannelLowerIndicator & NewKeltnerChannelUpperIndicator with mul parameter

## 0.12.1
* Fixes EMA window calculation (thanks @danhenke and @joelnordell!)

## 0.12.0
* Add MaximumValue and MinimumValue Indicators
* Add [MaximumDrawdownIndicator](https://www.investopedia.com/terms/m/maximum-drawdown-mdd.asp).

## 0.11.0
* Add BollingerUpperBandIndicator and BollingerLowerBandIndicator (thanks @shkim!)

## 0.10.0
* Add TimePeriod#In to modify timezone information
* Add TimePeriod#UTC to set location to UTC

## 0.9.0
* Add [AroonIndicator](https://www.investopedia.com/terms/a/aroon.asp)
* Deprecate [Parse](https://godoc.org/github.com/sdcoffey/techan#Parse) for time periods, introduce `ParseTimePeriod`

## 0.8.0
* Add [MMAIndicator](https://en.wikipedia.org/wiki/Moving_average#Modified_moving_average)
* Add GainIndicator
* Add LossIndicator
* Fix bug in RSI calculation (fixes #13)

## 0.7.1
* Fix bug in trendline indicator to prevent low-end index-OOB errors

## 0.7.0
* Add Trendline indicator
* Update big to v0.4.1

## 0.6.1
* Merge #10, which fixes a bug in TotalProfitAnalysis not taking short positions into effect

## 0.6.0
* **BREAKING**: Standard Deviation Indicator and Variance indicator now use the NewXIndicator pattern used throughout the library. Any usages creating the struct directly will need to be udpated.
* Migrate to go module

## 0.5.0
* Add StandardDeviationIndicator
* Add VarianceIndicator

## 0.4.0
* Add DerivativeIndicator

## 0.3.0
* Rename talib4g to techan

## 0.2.0
* Remove NewOrder methods and prefer struct initializer
* Add missing test coverage
* Add godoc

## 0.1.1
* Update documentation

## 0.1.0
* Initial release of talib4g
* Support for basic indicators
* Support for timeseries
* Support for basic strategies
* Support for entry and exit rules
