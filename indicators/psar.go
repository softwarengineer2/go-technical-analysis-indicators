package indicators

import (
	"math"

	"gitlab.com/softwarengineer2/gobot/utils"
)

//Parabolic SAR
//Source : https://www.tradingview.com/wiki/Parabolic_SAR_(SAR)
type psar struct {
	inc        float64
	maximum    float64
	start      float64
	prev       float64
	priceType  string
	count      int
	buf        *utils.OHLCVBuffer
	prevPrice  float64
	prevMaxMin float64
	prevAcc    float64
	prevPos    float64
}

//NewPSAR 	: To return psar struct instance
//Params
//start		: start value
//inc		: increment value
//maximum	: maximum value
//priceType : To use it as base price for calculations from OHLCV Buffer (open, high, low, close)
func NewPSAR(start, inc, maximum float64, priceType string) *psar {
	return &psar{
		inc:       inc,
		maximum:   maximum,
		start:     start,
		prev:      math.NaN(),
		priceType: priceType,
		buf:       utils.NewOHLCVBuffer(2),
	}
}

//Calculate : method to Calculate psar and return results as float array
//Return	: psar result (1 value in return array)
//Params
//newData	: OHLCV instance to use its values for price calculation
func (ins *psar) Calculate(newData utils.OHLCV) []float64 {
	newPrice := newData.GetByType(ins.priceType)

	if math.IsNaN(newPrice) {
		return []float64{ins.prev}
	}

	ins.buf.Add(newData)

	ins.count++

	if ins.count < 2 { //ins.buf.Size {
		return []float64{math.NaN()}
	}
	minTick := 1e-7
	out := math.NaN()
	pos := math.NaN()
	maxMin := math.NaN()
	acc := math.NaN()
	prev := ins.prev

	outSet := false
	if ins.count == 2 {
		if ins.count > 1 && newData.Close > ins.buf.Vals[ins.buf.Capacity-2].Close {
			pos = 1
			maxMin = math.Max(newData.High, ins.buf.Vals[ins.buf.Capacity-2].High)
			prev = math.Min(newData.Low, ins.buf.Vals[ins.buf.Capacity-2].Low)
		} else {
			pos = -1
			maxMin = math.Min(newData.Low, ins.buf.Vals[ins.buf.Capacity-2].Low)
			prev = math.Max(newData.High, ins.buf.Vals[ins.buf.Capacity-2].High)
		}
		acc = ins.start
	} else {
		pos = ins.prevPos
		acc = ins.prevAcc
		maxMin = ins.prevMaxMin
	}
	if pos == 1 {
		if newData.High > maxMin {
			maxMin = newData.High
			acc = math.Min(acc+ins.inc, ins.maximum)
		}
		if newData.Low <= prev {
			pos = -1
			out = maxMin
			maxMin = newData.Low
			acc = ins.start
			outSet = true
		}
	} else {
		if newData.Low < maxMin {
			maxMin = newData.Low
			acc = math.Min(acc+ins.inc, ins.maximum)
		}
		if newData.High >= prev {
			pos = 1
			out = maxMin
			maxMin = newData.High
			acc = ins.start
			outSet = true
		}
	}

	if outSet == false {
		out = prev + acc*(maxMin-prev)
		if pos == 1 {
			if out >= newData.Low {
				out = newData.Low - minTick
			}
		}
		if pos == -1 {
			if out <= newData.High {
				out = newData.High + minTick
			}
		}
	}
	ins.prevPos = pos
	ins.prevAcc = acc
	ins.prevMaxMin = maxMin
	ins.prev = out

	return []float64{ins.prev}
}
