package indicators

import (
	"math"

	"gitlab.com/softwarengineer2/gobot/utils"
)

//Chaikin A/D Oscillator
//Source : https://www.tradingview.com/wiki/Chaikin_Oscillator
type chaikin struct {
	count     uint
	period    uint
	prev      float64
	priceType string
	emaArr    [2]*ema
}

//NewCHAIKIN : To return chaikin struct instance
//Params
//short		 : short EMA period
//long		 : long EMA period
//priceType  : To use it as base price for calculations from OHLCV Buffer (open, high, low, close)
func NewCHAIKIN(short, long uint, priceType string) *chaikin {
	return &chaikin{
		period:    long,
		prev:      math.NaN(),
		priceType: priceType,
		emaArr:    [2]*ema{NewEMA(short, priceType), NewEMA(long, priceType)},
	}
}

//Calculate : method to Calculate chaikin and return results as float array
//Return	: Chaiking result (1 value in return array)
//Params
//newData	: OHLCV instance to use its values for price calculation
func (ins *chaikin) Calculate(newData utils.OHLCV) []float64 {
	newPrice := newData.GetByType(ins.priceType)

	if math.IsNaN(newPrice) {
		return []float64{ins.prev}
	}

	ins.count++

	acc := ((newData.Close - newData.Low) - (newData.High - newData.Close)) / (newData.High - newData.Low) * newData.Volume

	accData := utils.NewOHLCV(ins.priceType, acc)
	EMAShortResult := ins.emaArr[0].Calculate(accData)[0]
	EMALongResult := ins.emaArr[1].Calculate(accData)[0]

	osc := (EMAShortResult - EMALongResult)

	if ins.count < ins.period {
		return []float64{math.NaN()}
	}

	ins.prev = osc

	return []float64{ins.prev}
}
