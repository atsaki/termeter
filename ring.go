package termeter

type float64RingBuffer struct {
	buffer   []float64
	length   int
	capacity int
	tail     int
}

func newFloat64RingBuffer(capacity int) *float64RingBuffer {
	return &float64RingBuffer{
		buffer:   make([]float64, capacity, capacity),
		length:   0,
		capacity: capacity,
		tail:     0,
	}
}

func (r *float64RingBuffer) Len() int {
	return r.length
}

func (r *float64RingBuffer) Capacity() int {
	return r.capacity
}

func (r *float64RingBuffer) Add(v float64) {
	if r.length < r.capacity {
		r.length += 1
	}
	r.buffer[r.tail] = v
	r.tail = (r.tail + 1) % r.capacity
}

func (r *float64RingBuffer) Slice(i, j int) []float64 {
	if r.length < r.capacity {
		if r.length < j {
			j = r.length
		}
		return r.buffer[i:j]
	}
	s := append(r.buffer[r.tail:r.capacity], r.buffer[:r.tail]...)
	return s[i:j]
}

func (r *float64RingBuffer) Last(n int) []float64 {
	start := r.length - n
	if start < 0 {
		start = 0
	}
	return r.Slice(start, r.length)
}

type stringRingBuffer struct {
	buffer   []string
	length   int
	capacity int
	tail     int
}

func newStringRingBuffer(capacity int) *stringRingBuffer {
	return &stringRingBuffer{
		buffer:   make([]string, capacity, capacity),
		length:   0,
		capacity: capacity,
		tail:     0,
	}
}

func (r *stringRingBuffer) Len() int {
	return r.length
}

func (r *stringRingBuffer) Capacity() int {
	return r.capacity
}

func (r *stringRingBuffer) Add(v string) {
	if r.length < r.capacity {
		r.length += 1
	}
	r.buffer[r.tail] = v
	r.tail = (r.tail + 1) % r.capacity
}

func (r *stringRingBuffer) Slice(i, j int) []string {
	if r.length < r.capacity {
		if r.length < j {
			j = r.length
		}
		return r.buffer[i:j]
	}
	s := append(r.buffer[r.tail:r.capacity], r.buffer[:r.tail]...)
	return s[i:j]
}

func (r *stringRingBuffer) Last(n int) []string {
	start := r.length - n
	if start < 0 {
		start = 0
	}
	return r.Slice(start, r.length)
}
