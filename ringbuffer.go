package ringbuffer

type RingBuffer struct {
	buff          []float64
	head          int
	tail          int
	high          int
	low           int
	size          int
	recordHighLow bool
}

func NewBuffer(size int, recordHighLow bool) *RingBuffer {

	rb := &RingBuffer{}
	rb.recordHighLow = recordHighLow
	rb.size = size
	rb.buff = make([]float64, rb.size+1)
	rb.head = 0
	rb.tail = 1
	rb.high = 0
	rb.low = 0
	return rb
}

//Push adds a new element to the buffer
func (rb *RingBuffer) Push(value float64) {
	rb.buff[rb.head] = value

	if rb.recordHighLow == true {
		// --------------------
		// rb.high end of life,
		// --------------------
		if rb.high == rb.tail {
			rb.buff[rb.tail] = 0
			hVal := float64(0)
			for i, val := range rb.buff {
				if val > hVal {
					hVal = val
					rb.high = i
				}

			}
		}

		// --------------------
		// rb.low end of life
		// --------------------
		if rb.low == rb.tail {
			lVal := float64(0Xffffffff)
			rb.buff[rb.tail] = lVal
			for i, val := range rb.buff {
				if val < lVal {
					lVal = val
					rb.low = i
				}

			}
		}

		if value >= rb.buff[rb.high] {
			rb.high = rb.head
		}

		if value <= rb.buff[rb.low] {
			rb.low = rb.head
		}
	}
	rb.head++
	rb.tail++
	//rb.high++
	//rb.low++

	if rb.tail%(rb.size+1) == 0 {
		rb.tail = 0
	}

	if rb.head%(rb.size+1) == 0 {
		rb.head = 0
	}

}

//Tail returns the element at the buffer tail
func (rb *RingBuffer) Tail() float64 {
	return rb.buff[rb.tail]
}

//Head returns the element at the buffer tail
func (rb *RingBuffer) Head() float64 {
	return rb.buff[rb.head]
}

//MostRecent returns the element at the head - 1
func (rb *RingBuffer) MostRecent() float64 {
	//if rb.head == 0 {
	//	return rb.buff[rb.size-1]
	//}
	//return rb.buff[rb.head-1]
	return rb.buff[rb.head]

}

//Oldest returns the element at the head - 1
func (rb *RingBuffer) Oldest() float64 {
	//if rb.tail == 0 {
	//	return rb.buff[rb.size-1]
	//}
	return rb.buff[rb.tail]

}

//Tail returns the element at the buffer tail
func (rb *RingBuffer) High() float64 {
	return rb.buff[rb.high]
}

//Head returns the element at the buffer tail
func (rb *RingBuffer) Low() float64 {
	return rb.buff[rb.low]
}

func (rb *RingBuffer) tailNext() int {
	tailNext := rb.tail + 1

	if tailNext%(rb.size+1) == 0 {
		return 0
	}

	return tailNext

}
