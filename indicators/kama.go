package indicators

import (
	"math"

	"gitlab.com/softwarengineer2/gobot/utils"
)

//Kaufman's Moving Average
//Source : https://stockcharts.com/school/doku.php?id=chart_school:technical_indicators:kaufman_s_adaptive_moving_average
type kama struct {
	period    uint
	prev      float64
	priceType string
	buf       *utils.OHLCVBuffer
}

//NewKAMA 	: To return kama struct instance
//Params
//short		: calculation period
//priceType : To use it as base price for calculations from OHLCV Buffer (open, high, low, close)
func NewKAMA(period uint, priceType string) *kama {
	return &kama{
		period:    period + 1,
		prev:      0.0,
		priceType: priceType,
		buf:       utils.NewOHLCVBuffer(period + 1),
	}
}

//Calculate : method to Calculate kama and return results as float array
//Return	: Kama result (1 value in return array)
//Params
//newData	: OHLCV instance to use its values for price calculation
func (ins *kama) Calculate(newData utils.OHLCV) []float64 {
	newPrice := newData.GetByType(ins.priceType)

	if math.IsNaN(newPrice) {
		return []float64{ins.prev}
	}

	ins.buf.Add(newData)

	change := math.Abs(newPrice - ins.buf.Vals[0].GetByType(ins.priceType))
	var volatility float64
	if ins.buf.Pushes >= ins.buf.Size {
		for i := ins.buf.Capacity - 1; i > 0; i-- {
			volatility += (math.Abs(ins.buf.Vals[i].GetByType(ins.priceType) - ins.buf.Vals[i-1].GetByType(ins.priceType)))
		}
	}

	ER := 0.0
	if volatility != 0 {
		ER = change / volatility
	}

	fastLength := 2.0
	slowLength := 30.0

	fastAlpha := (2.0 / (fastLength + 1.0))
	slowAlpha := (2.0 / (slowLength + 1.0))

	SC := math.Pow((ER*(fastAlpha-slowAlpha))+slowAlpha, 2.0)

	ins.prev = (SC * newPrice) + ((1 - SC) * ins.prev)

	//if ins.buf.Pushes < ins.buf.Size {
	//	return []float64{math.NaN()}
	//}

	return []float64{ins.prev}
}
