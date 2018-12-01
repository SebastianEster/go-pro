package FunctionalProgramming

type Pair struct {
	first Any
	second Any
}

type Mapper func(Any) Any
type Predicate func(Any) bool
type Accumulator func(Any, Any) Any

type Stream interface {
	Map(m Mapper) Stream
	Filter(p Predicate) Stream
	Reduce(a Accumulator) Any
}

type SliceStream struct {
	data []Any
}

func ToStream(slice []Any) *SliceStream {
	newStream := newStream()
	newStream.data = slice
	return newStream
}

func newStream() *SliceStream {
	return new(SliceStream)
}

func (stream *SliceStream) Map(m Mapper) Stream {
	retStream := newStream()
	retStream.data = make([]Any, len(stream.data))
	for index, element := range stream.data {
		retStream.data[index] = m(element)
	}
	return retStream
}

func (stream *SliceStream) Filter(p Predicate) Stream {
	retStream := newStream()
	for _, element := range stream.data {
		if p(element) {
			retStream.data = append(retStream.data, element)
		}
	}
	return retStream
}

func (stream *SliceStream) Reduce(a Accumulator) Any {
	var ret = stream.data[0]
	for i := 1; i < len(stream.data); i++ {
		ret = a(ret, stream.data[i])
	}
	return ret
}