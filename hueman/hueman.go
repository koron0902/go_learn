package hueman

type Hueman struct{
		name_ string
		old_ int
		height_ int
		weight_ int
}

func Reserve(_name string) *Hueman {
		return &Hueman{name_ : _name}
}

func (_hueman Hueman) GetName() string {
		return _hueman.name_
}

func (_hueman *Hueman) SetName(_name string){
		_hueman.name_ = _name
}
