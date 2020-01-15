package myon


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
				// mathパッケージをimportすれば見れるけどそれだと意味ないのでとりあえず決め打ち
				return 2147483647
		}
		return _value1 / _value2
}
