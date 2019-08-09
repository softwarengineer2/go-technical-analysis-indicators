package indicators

import (
	"math"

	"gitlab.com/softwarengineer2/gobot/utils"
)

//Stochastic
//Source : https://www.tradingview.com/wiki/Stochastic_(STOCH)
type stoch struct {
	period    uint
	prev      float64
	priceType string
	buf       *utils.OHLCVBuffer
	smaArr    []*sma
}

//NewSTOCH 	: To return stoch struct instance
//Params
//short		: short sma period
//long		: long sma period
//priceType : To use it as base price for calculations from OHLCV Buffer (open, high, low, close)
func NewSTOCH(short, long uint, priceType string) *stoch {
	return &stoch{
		period:    long,
		prev:      math.NaN(),
		priceType: priceType,
		buf:       utils.NewOHLCVBuffer(long),
		smaArr:    []*sma{NewSMA(short, priceType), NewSMA(short, priceType)},
	}
}

//Calculate : method to Calculate stoch and return results as float array
//Return	: k and d value (2 values in return array)
//Params
//newData	: OHLCV instance to use its values for price calculation
func (ins *stoch) Calculate(newData utils.OHLCV) []float64 {
	newPrice := newData.GetByType(ins.priceType)

	if math.IsNaN(newPrice) {
		return []float64{ins.prev}
	}

	ins.buf.Add(newData)

	if ins.buf.Pushes < ins.buf.Size {
		return []float64{math.NaN()}
	}

	stochResult := (100.0 * (newPrice - ins.buf.Min().Low) / (ins.buf.Max().High - (ins.buf.Min().Low)))

	stochData := utils.NewOHLCV(ins.priceType, stochResult)

	stochKResult := ins.smaArr[0].Calculate(stochData)[0]

	stochKData := utils.NewOHLCV(ins.priceType, stochKResult)

	ins.prev = ins.smaArr[1].Calculate(stochKData)[0]

	return []float64{ins.prev, stochKResult}
}
