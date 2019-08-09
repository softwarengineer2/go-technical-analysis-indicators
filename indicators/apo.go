package indicators

import (
	"math"

	"gitlab.com/softwarengineer2/gobot/utils"
)

//Absolute Price Oscillator
//Source : https://www.fidelity.com/learning-center/trading-investing/technical-analysis/technical-indicator-guide/apo
type apo struct {
	count     uint
	period    uint
	prev      float64
	priceType string
	short     *ema
	long      *ema
}

//NewAPO 	: To return apo struct instance
//Params
//short		: long EMA period
//long		: short EMA period
//priceType : To use it as base price for calculations from OHLCV Buffer (open, high, low, close)
func NewAPO(short, long uint, priceType string) *apo {
	return &apo{
		period:    long,
		prev:      math.NaN(),
		priceType: priceType,
		short:     NewEMA(short, priceType),
		long:      NewEMA(long, priceType),
	}
}

//Calculate : method to Calculate APO and return results as float array
//Return	: APO Calculation Result  (1 value in return array)
//Params
//newData	: OHLCV instance to use its values for price calculation
func (ins *apo) Calculate(newData utils.OHLCV) []float64 {
	newPrice := newData.GetByType(ins.priceType)

	if math.IsNaN(newPrice) {
		return []float64{ins.prev}
	}

	ins.count++

	ins.prev = ins.short.Calculate(newData)[0] - ins.long.Calculate(newData)[0]

	if ins.count < ins.period {
		return []float64{math.NaN()}
	}

	return []float64{ins.prev}
}
