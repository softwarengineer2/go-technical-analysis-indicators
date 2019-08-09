package indicators

import (
	"math"

	"gitlab.com/softwarengineer2/gobot/utils"
)

//Stochastic Relative Strength Index
//Source : https://www.tradingview.com/wiki/Stochastic_RSI_(STOCH_RSI)
type srsi struct {
	period    uint
	prev      float64
	priceType string
	buf       *utils.OHLCVBuffer
	smaArr    []*sma
	rsi1      *rsi
}

//NewSRSI 	: To return srsi struct instance
//Params
//short		: short sma period
//long		: long sma period
//priceType : To use it as base price for calculations from OHLCV Buffer (open, high, low, close)
func NewSRSI(short uint, long uint, priceType string) *srsi {
	return &srsi{
		period:    long,
		prev:      math.NaN(),
		priceType: priceType,
		buf:       utils.NewOHLCVBuffer(long),
		smaArr:    []*sma{NewSMA(short, priceType), NewSMA(short, priceType)},
		rsi1:      NewRSI(long, priceType),
	}
}

//Calculate : method to Calculate srsi and return results as float array
//Return	: k and d values (2 values in return array)
//Params
//newData	: OHLCV instance to use its values for price calculation
func (ins *srsi) Calculate(newData utils.OHLCV) []float64 {
	newPrice := newData.GetByType(ins.priceType)

	if math.IsNaN(newPrice) {
		return []float64{ins.prev}
	}

	rsiRes := ins.rsi1.Calculate(newData)[0]
	data := utils.NewOHLCV(ins.priceType, rsiRes)

	srsiKResult := math.NaN()
	if !math.IsNaN(rsiRes) {
		ins.buf.Add(data)

		srsiResult := (100.0 * (rsiRes - ins.buf.Min().Close) / (ins.buf.Max().Close - (ins.buf.Min().Close)))

		srsiData := utils.NewOHLCV(ins.priceType, srsiResult)

		srsiKResult = ins.smaArr[0].Calculate(srsiData)[0]

		srsiKData := utils.NewOHLCV(ins.priceType, srsiKResult)

		ins.prev = ins.smaArr[1].Calculate(srsiKData)[0]
	}

	return []float64{ins.prev, srsiKResult}
}
