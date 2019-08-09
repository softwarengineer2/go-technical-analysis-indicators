package utils

//Buffer is a data structure that holds float64 values
type Buffer struct {
	Size     uint
	Pushes   uint
	Index    uint
	Capacity uint
	Sum      float64
	Vals     []float64
}

//NewBuffer : To return new Buffer instance
//Params
//Size : Buffer size
func NewBuffer(Size uint) *Buffer {
	return &Buffer{
		Size: Size,
		Vals: make([]float64, Size),
	}
}

//Shift : Method to shift array elements
func (ins *Buffer) Shift() {
	for i := 1; i < len(ins.Vals); i++ {
		ins.Vals[i-1] = ins.Vals[i]
	}
}

//Add : Method to add new array element
func (ins *Buffer) Add(v float64) {
	if ins.Pushes >= ins.Size {
		ins.Sum -= ins.Vals[ins.Index]
		ins.Shift()
		ins.Index = ins.Size - 1
	} else {
		ins.Capacity++
	}

	ins.Sum += v
	ins.Vals[ins.Index] = v
	ins.Pushes++
	ins.Index++
	if ins.Index >= ins.Size {
		ins.Index = 0
	}
}

//Total : To return Total element of array
func (ins *Buffer) Total() float64 {
	return ins.Sum
}

//Max : To return Max element of array
func (ins *Buffer) Max() float64 {
	m := ins.Vals[0]

	for _, v := range ins.Vals[1:] {
		if v > m {
			m = v
		}
	}

	return m
}

//Min : To return Min element of array
func (ins *Buffer) Min() float64 {
	m := ins.Vals[0]

	for _, v := range ins.Vals[1:] {
		if v < m {
			m = v
		}
	}

	return m
}

//Get : To return specific element of array
//If it doesn't exist, return 0
//Params
//i : index of element
func (ins *Buffer) Get(i uint) float64 {
	oneData := 0.0
	if ins.Capacity > 0 && i < ins.Capacity && i >= 0 {
		oneData = ins.Vals[i]
	}
	return oneData
}

//GetR : To return specific element of array
//If it doesn't exist, return 0
//Params
//i : index of element
func (ins *Buffer) GetR(i uint) float64 {
	oneData := 0.0
	if ins.Capacity > 0 && ins.Capacity-i >= 0 && i >= 0 {
		oneData = ins.Vals[ins.Capacity-i]
	}
	return oneData
}

//Last : To return last element of array
func (ins *Buffer) Last() float64 {
	return ins.Get(ins.Capacity - 1)
}

//First : To return first element of array
func (ins *Buffer) First() float64 {
	return ins.Get(0)
}
