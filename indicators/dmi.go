package indicators

import (
	"math"

	"gitlab.com/softwarengineer2/gobot/utils"
)

//Directional Movement Index
//Source : https://www.tradingview.com/wiki/Directional_Movement_(DMI)
type dmi struct {
	count     uint
	prev      float64
	period    uint
	signal    uint
	priceType string
	rmaArr    []*wilders
	prevPrice utils.OHLCV
}

//NewDMI 	: To return dmi struct instance
//Params
//period	: calculation period
//signal	: signal period
//priceType : To use it as base price for calculations from OHLCV Buffer (open, high, low, close)
func NewDMI(period, signal uint, priceType string) *dmi {
	return &dmi{
		prev:      math.NaN(),
		period:    period,
		signal:    signal,
		priceType: priceType,
		rmaArr:    []*wilders{NewWILDERS(period, priceType), NewWILDERS(period, priceType), NewWILDERS(period, priceType), NewWILDERS(signal, priceType)},
		prevPrice: utils.OHLCV{},
	}
}

//Calculate : method to Calculate dmi and return results as float array
//Return	: Adx, plus and minus values (3 values in return array)
//Params
//newData	: OHLCV instance to use its values for price calculation
func (ins *dmi) Calculate(newData utils.OHLCV) []float64 {
	newPrice := newData.GetByType(ins.priceType)

	if math.IsNaN(newPrice) {
		return []float64{ins.prev}
	}

	ins.count++

	up := 0.0
	down := 0.0
	tr := math.NaN()

	if ins.count != 1 {
		up = (newData.High - ins.prevPrice.High)
		down = -(newData.Low - ins.prevPrice.Low)
		tr = math.Max(math.Max(newData.High-newData.Low, math.Abs(newData.High-ins.prevPrice.Close)), math.Abs(newData.Low-ins.prevPrice.Close))
	}

	trData := utils.NewOHLCV(ins.priceType, tr)
	trur := ins.rmaArr[0].Calculate(trData)[0]

	plusVal := 0.0
	if up > down && up > 0.0 {
		plusVal = up
	}
	plusData := utils.NewOHLCV(ins.priceType, plusVal)
	plVal := ins.rmaArr[1].Calculate(plusData)[0]
	plus := (100.0 * plVal / trur)

	minusVal := 0.0
	if down > up && down > 0.0 {
		minusVal = down
	}
	minusData := utils.NewOHLCV(ins.priceType, minusVal)
	mnVal := ins.rmaArr[2].Calculate(minusData)[0]
	minus := (100.0 * mnVal / trur)

	sum := plus + minus

	adxVal := math.Abs(plus - minus)
	if sum != 0 {
		adxVal /= sum
	}
	adxData := utils.NewOHLCV(ins.priceType, adxVal)
	adx := 100.0 * ins.rmaArr[3].Calculate(adxData)[0]

	ins.prev = adx

	ins.prevPrice = newData

	if ins.count < ins.period*2 {
		return []float64{math.NaN()}
	}

	return []float64{ins.prev}
}
