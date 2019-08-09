package indicators

import (
	"math"

	"gitlab.com/softwarengineer2/gobot/utils"
)

//MESA Adaptive Moving Average
//Source: https://www.tradingview.com/script/foQxLbU3-Ehlers-MESA-Adaptive-Moving-Average-LazyBear/
type mesa struct {
	period    uint
	prev      float64
	priceType string
	buf       *utils.OHLCVBuffer
	spArr     *utils.Buffer
	pArr      *utils.Buffer
	dtArr     *utils.Buffer
	i1Arr     *utils.Buffer
	i2Arr     *utils.Buffer
	q1Arr     *utils.Buffer
	q2Arr     *utils.Buffer
	imArr     *utils.Buffer
	p1Arr     *utils.Buffer
	p3Arr     *utils.Buffer
	reArr     *utils.Buffer
	phaseArr  *utils.Buffer
	sppArr    *utils.Buffer
	mamaArr   *utils.Buffer
	famaArr   *utils.Buffer
}

//NewMESA 	: To return chaikin struct instance
//Params
//period	: calculation period
//priceType : To use it as base price for calculations from OHLCV Buffer (open, high, low, close)
func NewMESA(period uint, priceType string) *mesa {
	return &mesa{
		period:    period,
		prev:      math.NaN(),
		priceType: priceType,
		buf:       utils.NewOHLCVBuffer(period),
		spArr:     utils.NewBuffer(6),
		pArr:      utils.NewBuffer(1),
		dtArr:     utils.NewBuffer(6),
		q1Arr:     utils.NewBuffer(6),
		q2Arr:     utils.NewBuffer(1),
		i1Arr:     utils.NewBuffer(6),
		i2Arr:     utils.NewBuffer(1),
		imArr:     utils.NewBuffer(1),
		p1Arr:     utils.NewBuffer(1),
		p3Arr:     utils.NewBuffer(1),
		reArr:     utils.NewBuffer(1),
		phaseArr:  utils.NewBuffer(1),
		sppArr:    utils.NewBuffer(1),
		mamaArr:   utils.NewBuffer(1),
		famaArr:   utils.NewBuffer(1),
	}
}

//Calculate : method to Calculate mesa and return results as float array
//Return	: mama, fama results (2 values in return array)
//Params
//newData	: OHLCV instance to use its values for price calculation
func (ins *mesa) Calculate(newData utils.OHLCV) []float64 {
	newPrice := newData.GetByType(ins.priceType)

	if math.IsNaN(newPrice) {
		return []float64{ins.prev}
	}

	ins.buf.Add(newData)

	fl := 0.2
	sl := 0.02

	sp := 0.0
	if ins.buf.Pushes >= 4 {
		sp = (4*newData.HL2() + 3*ins.buf.Vals[ins.buf.Capacity-2].HL2() + 2*ins.buf.Vals[ins.buf.Capacity-3].HL2() + ins.buf.Vals[ins.buf.Capacity-4].HL2()) / 10.0
	}

	dt := (0.0962*sp + 0.5769*(ins.spArr.Get(ins.spArr.Capacity-2)) - 0.5769*(ins.spArr.Get(ins.spArr.Capacity-4)) - 0.0962*(ins.spArr.Get(ins.spArr.Capacity-6))) * (0.075*(ins.pArr.Last()) + 0.54)
	q1 := (0.0962*dt + 0.5769*(ins.dtArr.Get(ins.dtArr.Capacity-2)) - 0.5769*(ins.dtArr.Get(ins.dtArr.Capacity-4)) - 0.0962*(ins.dtArr.Get(ins.dtArr.Capacity-6))) * (0.075*(ins.pArr.Last()) + 0.54)

	i1 := (ins.dtArr.Get(ins.dtArr.Capacity - 3))
	jI := (0.0962*i1 + 0.5769*(ins.i1Arr.Get(ins.i1Arr.Capacity-2)) - 0.5769*(ins.i1Arr.Get(ins.i1Arr.Capacity-4)) - 0.0962*(ins.i1Arr.Get(ins.i1Arr.Capacity-6))) * (0.075*(ins.pArr.Last()) + 0.54)

	jq := (0.0962*q1 + 0.5769*(ins.q1Arr.Get(ins.q1Arr.Capacity-2)) - 0.5769*(ins.q1Arr.Get(ins.q1Arr.Capacity-4)) - 0.0962*(ins.q1Arr.Get(ins.q1Arr.Capacity-6))) * (0.075*(ins.pArr.Last()) + 0.54)
	i2_ := i1 - jq
	q2_ := q1 + jI

	i2 := 0.2*i2_ + 0.8*(ins.i2Arr.Last())
	q2 := 0.2*q2_ + 0.8*(ins.q2Arr.Last())

	re_ := i2*(ins.i2Arr.Last()) + q2*(ins.q2Arr.Last())
	im_ := i2*(ins.q2Arr.Last()) - q2*(ins.i2Arr.Last())
	re := 0.2*re_ + 0.8*(ins.reArr.Last())
	im := 0.2*im_ + 0.8*(ins.imArr.Last())

	p1 := 0.0
	if math.Abs(im) > 0.00000001 && math.Abs(re) > 0.00000001 {
		p1 = 360.0 / math.Atan(im/re)
	} else {
		p1 = (ins.pArr.Last())
	}

	p2 := 0.0
	if p1 > 1.5*(ins.p1Arr.Last()) {
		p2 = 1.5 * (ins.p1Arr.Last())
	} else {
		if p1 < 0.67*(ins.p1Arr.Last()) {
			p2 = 0.67 * (ins.p1Arr.Last())
		} else {
			p2 = p1
		}
	}

	p3 := 0.0
	if p2 < 6 {
		p3 = 6.0
	} else {
		if p2 > 50 {
			p3 = 50
		} else {
			p3 = p2
		}
	}
	p := 0.2*p3 + 0.8*(ins.p3Arr.Last())

	spp := 0.33*p + 0.67*(ins.sppArr.Last())
	phase := math.Atan(q1 / i1)
	dphase_ := (ins.phaseArr.Last()) - phase

	dphase := 0.0
	if dphase_ < 1 {
		dphase = 1
	} else {
		dphase = dphase_
	}

	alpha_ := fl / dphase
	alpha := 0.0
	if alpha_ < sl {
		alpha = sl
	} else {
		if alpha_ > fl {
			alpha = fl
		} else {
			alpha = alpha_
		}
	}

	mama := alpha*newData.HL2() + (1-alpha)*(ins.mamaArr.Last())
	fama := 0.5*alpha*mama + (1.0-0.5*alpha)*(ins.famaArr.Last())

	if !math.IsNaN(sp) && ins.buf.Pushes != 1 {
		ins.spArr.Add(sp)
	}
	if !math.IsNaN(p) {
		ins.pArr.Add(p)
	}
	if !math.IsNaN(dt) {
		ins.dtArr.Add(dt)
	}
	if !math.IsNaN(i1) {
		ins.i1Arr.Add(i1)
	}
	if !math.IsNaN(i2) {
		ins.i2Arr.Add(i2)
	}
	if !math.IsNaN(q1) {
		ins.q1Arr.Add(q1)
	}
	if !math.IsNaN(q2) {
		ins.q2Arr.Add(q2)
	}
	if !math.IsNaN(im) {
		ins.imArr.Add(im)
	}
	if !math.IsNaN(re) {
		ins.reArr.Add(re)
	}
	if !math.IsNaN(p1) {
		ins.p1Arr.Add(p1)
	}
	if !math.IsNaN(p3) {
		ins.p3Arr.Add(p3)
	}
	if !math.IsNaN(phase) {
		ins.phaseArr.Add(phase)
	}
	if !math.IsNaN(fama) {
		ins.famaArr.Add(fama)
	}
	if !math.IsNaN(mama) {
		ins.mamaArr.Add(mama)
	}
	if !math.IsNaN(spp) {
		ins.sppArr.Add(spp)
	}

	ins.prev = fama
	return []float64{mama, fama}
}
