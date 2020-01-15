package myon

var INT_MAX_ = 2147483647

func Add(_value1 int, _value2 int)int{
		return _value1 + _value2
}

func Sub(_value1 int, _value2 int)int{
		return _value1 - _value2
}

func Mul(_value1 int, _value2 int)int{
		return _value1 * _value2
}

func Div(_value1 int, _value2 int)int{
		if _value2 == 0{
				return INT_MAX_
		}
		return _value1 / _value2
}
