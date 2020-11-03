package domain

// Dollar - represents a dollar bill
type Dollar struct {
	Amount int64
}

// NewDollar - constructor
func NewDollar(amount int64) Dollar {
	d := new(Dollar)
	d.Amount = amount
	return *d
}

// Times - multiply amount by function parameter
func (d *Dollar) Times(value int64) Dollar {
	return Dollar{
		Amount: d.Amount * value,
	}
}
