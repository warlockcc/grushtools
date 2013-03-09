package aggregate

import (
	"strconv"
)

type Aggregator interface {
	Aggregate(value []byte)
	String() string
}

// Sums
type Adder float64

func (a *Adder) Aggregate(value []byte) {
	var f, _ = strconv.ParseFloat(string(value), 64)
	*a += Adder(f)
}

func (a Adder) String() string {
	return strconv.FormatFloat(float64(a), 'g', -1, 64)
}

// Counters
type Counter uint64

func (c *Counter) Aggregate(value []byte) {
	*c += Counter(1)
}

func (c *Counter) String() string {
	return strconv.FormatUint(uint64(*c), 10)
}

// Minimums
type Miner float64

func (m *Miner) Aggregate(value []byte) {
	var f, _ = strconv.ParseFloat(string(value), 64)
	if Miner(f) < *m {
		*m = Miner(f)
	}
}

func (m Miner) String() string {
	return strconv.FormatFloat(float64(m), 'g', -1, 64)
}

// Maximums
type Maxer float64

func (m *Maxer) Aggregate(value []byte) {
	var f, _ = strconv.ParseFloat(string(value), 64)
	if Maxer(f) > *m {
		*m = Maxer(f)
	}
}

func (m Maxer) String() string {
	return strconv.FormatFloat(float64(m), 'g', -1, 64)
}

// Averages
type Averager struct {
	sum float64
	num uint64
}

func (a *Averager) Aggregate(value []byte) {
	var f, _ = strconv.ParseFloat(string(value), 64)
	a.sum += f
	a.num++
}

func (a Averager) String() string {
	if a.num > 0 {
		return strconv.FormatFloat(a.sum/float64(a.num), 'g', -1, 64)
	}
	return "0"
}

// Firsts
type Firster []byte

func (f *Firster) Aggregate(value []byte) {
	if *f == nil {
		*f = make([]byte, len(value))
		copy(*f, value)
	}
}

func (f Firster) String() string {
	return string(f)
}

// Lasts
type Laster []byte

func (l *Laster) Aggregate(value []byte) {
	if *l == nil || len(*l) < len(value) {
		*l = make([]byte, len(value))
	}
	copy(*l, value)
}

func (l Laster) String() string {
	return string(l)
}

// Concatenation
type Concater struct {
	buffer []byte
	Delim  []byte
}

func (c *Concater) Aggregate(value []byte) {
	if c.buffer != nil {
		c.buffer = append(c.buffer, c.Delim...)
	}
	c.buffer = append(c.buffer, value...)
}

func (c Concater) String() string {
	return string(c.buffer)
}