package main

import "fmt"

type sum struct {
	intVal     int64
	floatVal   float64
	complexVal complex128
}

func (s *sum) New(nums ...interface{}) *sum {
	var intCurVal int64
	var floatCurVal float64
	var complexCurVal complex128
	isIntCurVal := true
	var isFloatCurVal, isComplexCurVal bool

	for _, val := range nums {
		var isFloatVal, isComplexVal bool
		switch val.(type) {
		case float64:
			isFloatVal = true
		case complex128:
			isComplexVal = true
		}

		if isIntCurVal && isComplexVal {
			floatCurVal = float64(intCurVal)
		}

		if !isComplexCurVal && isComplexVal {
			isComplexCurVal = true
			isFloatCurVal = false
			isIntCurVal = false
			complexVal := val.(complex128)
			complexCurVal = complexVal + complex(floatCurVal, 0)
			continue
		}

		if isComplexCurVal && isComplexVal {
			var complexVal complex128
			complexVal = val.(complex128)
			complexCurVal += complexVal
			continue
		}

		if isComplexCurVal && isFloatVal {
			var floatVal float64
			floatVal = val.(float64)
			var complexVal complex128
			complexVal = complex(floatVal, 0)
			complexCurVal += complexVal
			continue
		}

		if isComplexCurVal {
			var intVal int64
			intVal = val.(int64)
			var floatVal float64
			floatVal = float64(intVal)
			var complexVal complex128
			complexVal = complex(floatVal, 0)
			complexCurVal += complexVal
			continue
		}

		if isFloatCurVal && isFloatVal {
			var floatVal float64
			floatVal = val.(float64)
			floatCurVal += floatVal
			continue
		}

		if isFloatCurVal {
			var intVal int64
			intVal = val.(int64)
			var floatVal float64
			floatVal = float64(intVal)
			floatCurVal += floatVal
			continue
		}

		if isIntCurVal && isFloatVal {
			isIntCurVal = false
			isFloatCurVal = true
			var floatVal float64
			floatVal = val.(float64)
			floatCurVal = float64(intCurVal) + floatVal
			continue
		}

		var intVal int64
		intVal = val.(int64)
		intCurVal += intVal
	}

	if isIntCurVal {
		s.intVal = intCurVal
	}

	if isFloatCurVal {
		s.floatVal = floatCurVal
	}

	if isComplexCurVal {
		s.complexVal = complexCurVal
	}

	return s
}

func (s *sum) Val() interface{} {
	if s.complexVal != 0 {
		return s.complexVal
	}

	if s.floatVal != 0 {
		return s.floatVal
	}

	return s.intVal
}

func (s *sum) Add(val int64) {
	if s.complexVal != 0 {
		floatVal := float64(val)
		s.complexVal += complex(floatVal, 0)
		return
	}

	if s.floatVal != 0 {
		floatVal := float64(val)
		s.floatVal += floatVal
		return
	}

	s.intVal += val
}

func main() {
	a := int64(1)
	b := int64(2)
	c := 2.5
	d := complex(2, 3)
	e := int64(2)
	f := 3.5
	g := complex(0, 3)
	h := complex(0, 3.5)
	var s sum
	sInst := s.New(a, b, c, d, e, f, g, h)
	fmt.Println(sInst.Val())
	sInst.Add(3)
	fmt.Println(sInst.Val())
}
