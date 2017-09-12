package ringbuffer

type RingBuffer struct {
	buff []float64
	head int
	tail int
	size int
}

func NewBuffer(size int) *RingBuffer {

	rb := &RingBuffer{}
	rb.size = size
	rb.buff = make([]float64, rb.size+1)
	rb.head = 0
	rb.tail = 1
	return rb
}

//Push adds a new element to the buffer
func (rb *RingBuffer) Push(value float64) {
	rb.buff[rb.head] = value
	rb.head++
	rb.tail++

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
	if rb.head == 0 {
		return rb.buff[rb.size-1]
	}
	return rb.buff[rb.head-1]

}

//Oldest returns the element at the head - 1
func (rb *RingBuffer) Oldest() float64 {
	if rb.tail == 0 {
		return rb.buff[rb.size-1]
	}
	return rb.buff[rb.tail-1]

}
