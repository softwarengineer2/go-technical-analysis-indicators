package indicators

import (
	"math"

	"gitlab.com/softwarengineer2/gobot/utils"
)

//T3 Moving Average
//Source : http://www.binarytribune.com/forex-trading-indicators/t3-moving-average-indicator/
type t3ma struct {
	count     uint
	period    uint
	prev      float64
	priceType string
	gdemaArr  [3]*gdema
}

//NewT3MA 	: To return t3ma struct instance
//Params
//period	: calculation period
//vFactor	: volume factor
//priceType : To use it as base price for calculations from OHLCV Buffer (open, high, low, close)
func NewT3MA(period uint, vFactor float64, priceType string) *t3ma {
	return &t3ma{
		period:    period,
		prev:      math.NaN(),
		priceType: priceType,
		gdemaArr:  [3]*gdema{NewGDEMA(period, vFactor, priceType), NewGDEMA(period, vFactor, priceType), NewGDEMA(period, vFactor, priceType)},
	}
}

//Calculate : method to Calculate t3ma and return results as float array
//Return	: t3ma result (1 value in return array)
//Params
//newData	: OHLCV instance to use its values for price calculation
func (ins *t3ma) Calculate(newData utils.OHLCV) []float64 {
	newPrice := newData.GetByType(ins.priceType)

	if math.IsNaN(newPrice) {
		return []float64{ins.prev}
	}

	ins.count++

	GDEMAResult1 := ins.gdemaArr[2].Calculate(newData)[0]
	gdema1 := utils.NewOHLCV(ins.priceType, GDEMAResult1)
	GDEMAResult2 := ins.gdemaArr[1].Calculate(gdema1)[0]
	gdema2 := utils.NewOHLCV(ins.priceType, GDEMAResult2)

	if !math.IsNaN(GDEMAResult1) && !math.IsNaN(GDEMAResult2) {
		ins.prev = ins.gdemaArr[0].Calculate(gdema2)[0]
	} else {
		return []float64{math.NaN()}
	}

	if ins.count < ins.period {
		return []float64{math.NaN()}
	}

	return []float64{ins.prev}
}
