package indicators

import (
	"math"

	"gitlab.com/softwarengineer2/gobot/utils"
)

//Wilder's Moving Average
//Source : https://www.incrediblecharts.com/indicators/wilder_moving_average.php
type wilders struct {
	count     uint
	period    uint
	prev      float64
	priceType string
	sma1      *sma
}

//NewWILDERS : To return dema struct instance
//Params
//period	 : calculation period
//priceType  : To use it as base price for calculations from OHLCV Buffer (open, high, low, close)
func NewWILDERS(period uint, priceType string) *wilders {
	return &wilders{
		prev:      0,
		period:    period,
		priceType: priceType,
		sma1:      NewSMA(period, priceType),
	}
}

//Calculate : method to Calculate wilders and return results as float array
//Return	: wilders result (1 value in return array)
//Params
//newData	: OHLCV instance to use its values for price calculation
func (ins *wilders) Calculate(newData utils.OHLCV) []float64 {
	newPrice := newData.GetByType(ins.priceType)

	if math.IsNaN(newPrice) {
		return []float64{ins.prev}
	}

	ins.count++

	if ins.count <= ins.period {
		ins.prev = ins.sma1.Calculate(newData)[0]
	}
	if ins.count > ins.period {
		ins.prev = (newPrice-ins.prev)*(1./(float64(ins.period))) + ins.prev
	}

	if ins.count < ins.period {
		return []float64{math.NaN()}
	}

	return []float64{ins.prev}
}
