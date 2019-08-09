package utils

//OHLCVBuffer is a data structure that holds float64 values
type OHLCVBuffer struct {
	Size     uint
	Pushes   uint
	Index    uint
	Capacity uint
	Vals     []OHLCV
	Sum      OHLCV
}

//NewOHLCVBuffer : To return new OHLCVBuffer instance
//Params
//Size : Buffer size
func NewOHLCVBuffer(Size uint) *OHLCVBuffer {
	return &OHLCVBuffer{
		Size: Size,
		Vals: make([]OHLCV, Size),
	}
}

//Shift : Method to shift array elements
func (ins *OHLCVBuffer) Shift() {
	for i := 1; i < len(ins.Vals); i++ {
		ins.Vals[i-1] = ins.Vals[i]
	}
}

//Add : Method to add new array element
func (ins *OHLCVBuffer) Add(newData OHLCV) {
	if ins.Pushes >= ins.Size {
		ins.Sum.Open -= ins.Vals[ins.Index].Open
		ins.Sum.High -= ins.Vals[ins.Index].High
		ins.Sum.Low -= ins.Vals[ins.Index].Low
		ins.Sum.Close -= ins.Vals[ins.Index].Close
		ins.Sum.Volume -= ins.Vals[ins.Index].Volume
		ins.Shift()
		ins.Index = ins.Size - 1
	} else {
		ins.Capacity++
	}

	ins.Sum.Open += newData.Open
	ins.Sum.High += newData.High
	ins.Sum.Low += newData.Low
	ins.Sum.Close += newData.Close
	ins.Sum.Volume += newData.Volume
	ins.Vals[ins.Index] = newData
	ins.Pushes++
	ins.Index++
	if ins.Index >= ins.Size {
		ins.Index = 0
	}
}

//Total : To return Total element of array
func (ins *OHLCVBuffer) Total() OHLCV {
	return ins.Sum
}

//Max : To return Max element of array
func (ins *OHLCVBuffer) Max() OHLCV {
	oneData := ins.Vals[0]

	for i := 1; i < int(ins.Capacity); i++ {
		newData := ins.Vals[i]
		if newData.Open > oneData.Open {
			oneData.Open = newData.Open
		}
		if newData.High > oneData.High {
			oneData.High = newData.High
		}
		if newData.Low > oneData.Low {
			oneData.Low = newData.Low
		}
		if newData.Close > oneData.Close {
			oneData.Close = newData.Close
		}
		if newData.Volume > oneData.Volume {
			oneData.Volume = newData.Volume
		}
	}

	return oneData
}

//Min : To return Min element of array
func (ins *OHLCVBuffer) Min() OHLCV {
	oneData := ins.Vals[0]

	for i := 1; i < int(ins.Capacity); i++ {
		newData := ins.Vals[i]
		if newData.Open < oneData.Open {
			oneData.Open = newData.Open
		}
		if newData.High < oneData.High {
			oneData.High = newData.High
		}
		if newData.Low < oneData.Low {
			oneData.Low = newData.Low
		}
		if newData.Close < oneData.Close {
			oneData.Close = newData.Close
		}
		if newData.Volume < oneData.Volume {
			oneData.Volume = newData.Volume
		}
	}

	return oneData
}

//MaxIndex : To return Max element index of array
func (ins *OHLCVBuffer) MaxIndex() OHLCV {
	oneData := ins.Vals[0]
	tempData := OHLCV{}

	for i := 1; i < int(ins.Capacity); i++ {
		newData := ins.Vals[i]
		if newData.Open > oneData.Open {
			oneData.Open = newData.Open
			tempData.Open = float64(i)
		}
		if newData.High > oneData.High {
			oneData.High = newData.High
			tempData.High = float64(i)
		}
		if newData.Low > oneData.Low {
			oneData.Low = newData.Low
			tempData.Low = float64(i)
		}
		if newData.Close > oneData.Close {
			oneData.Close = newData.Close
			tempData.Close = float64(i)
		}
		if newData.Volume > oneData.Volume {
			oneData.Volume = newData.Volume
			tempData.Volume = float64(i)
		}
	}
	return tempData
}

//MinIndex : To return Min element index of array
func (ins *OHLCVBuffer) MinIndex() OHLCV {
	oneData := ins.Vals[0]
	tempData := OHLCV{}

	for i := 1; i < int(ins.Capacity); i++ {
		newData := ins.Vals[i]
		if newData.Open < oneData.Open {
			oneData.Open = newData.Open
			tempData.Open = float64(i)
		}
		if newData.High < oneData.High {
			oneData.High = newData.High
			tempData.High = float64(i)
		}
		if newData.Low < oneData.Low {
			oneData.Low = newData.Low
			tempData.Low = float64(i)
		}
		if newData.Close < oneData.Close {
			oneData.Close = newData.Close
			tempData.Close = float64(i)
		}
		if newData.Volume < oneData.Volume {
			oneData.Volume = newData.Volume
			tempData.Volume = float64(i)
		}
	}

	return tempData
}

//Get : To return specific element of array
//If it doesn't exist, return 0
//Params
//i : index of element
func (ins *OHLCVBuffer) Get(i uint) OHLCV {
	return ins.Vals[(ins.Index+ins.Size-1-i)%ins.Size]
}

//GetR : To return specific element of array
//Params
//i : index of element
func (ins *OHLCVBuffer) GetR(i uint) OHLCV {
	return ins.Vals[ins.Capacity-i]
}

//Avg : To return average of array for OHLCV values
func (ins *OHLCVBuffer) Avg() OHLCV {
	oneData := OHLCV{}

	for i := 0; i < int(ins.Capacity); i++ {
		newData := ins.Vals[i]
		oneData.Open += newData.Open
		oneData.High += newData.High
		oneData.Low += newData.Low
		oneData.Close += newData.Close
		oneData.Volume += newData.Volume
	}

	oneData.Open /= float64(ins.Capacity)
	oneData.High /= float64(ins.Capacity)
	oneData.Low /= float64(ins.Capacity)
	oneData.Close /= float64(ins.Capacity)
	oneData.Volume /= float64(ins.Capacity)

	return oneData
}

//SpecialAvg : To return special average of array for OHLCV values
func (ins *OHLCVBuffer) SpecialAvg(start, end int) OHLCV {
	oneData := OHLCV{}

	if end < 0 {
		start = int(ins.Capacity) + end
		end = int(ins.Capacity)
	}
	if start < 0 {
		start = 0
	}

	for i := 0; i < int(ins.Capacity); i++ {
		newData := ins.Vals[i]
		oneData.Open += newData.Open
		oneData.High += newData.High
		oneData.Low += newData.Low
		oneData.Close += newData.Close
		oneData.Volume += newData.Volume
	}

	oneData.Open /= float64(ins.Capacity)
	oneData.High /= float64(ins.Capacity)
	oneData.Low /= float64(ins.Capacity)
	oneData.Close /= float64(ins.Capacity)
	oneData.Volume /= float64(ins.Capacity)

	return oneData
}

//Last : To return last element of array
func (ins *OHLCVBuffer) Last() OHLCV {
	return ins.Vals[ins.Capacity-1]
}

//First : To return first element of array
func (ins *OHLCVBuffer) First() OHLCV {
	return ins.Vals[0]
}
