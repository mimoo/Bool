package Bool

//
// option.Bool()
type Option struct {
	value bool
	err   error
}

func (b Option) getValue() interface{} {
	return b.value
}

func (b Option) isErr() bool {
	return b.err != nil
}

func (b Option) getErr() error {
	return b.err
}

func Match(option Option) interface{} {
	if option.isErr() {
		return option.getErr()
	}
	return option.getValue()
}

func Ok(value bool) Option {
	return Option{
		value: value,
		err:   nil,
	}
}

func Err(err error) Option {
	return Option{
		value: false,
		err:   err,
	}
}
