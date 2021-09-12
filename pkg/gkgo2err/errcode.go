package gkgo2err

type CodeType uint8

type Code struct {
	Type CodeType
	Category uint8
	SubCode int
	Msg string `json:"message"`
	Code int `json:"retcode"`
}

const (
	CodeTypeCommon = 1
	CodeTypeBusiness = 2
)

const (
	System = 1
	User = 2
)

type CodeCategoryFactory struct {
	Type CodeType
}

type CodeFactory struct {
	Type CodeType
	Category uint8
}

func Common() *CodeCategoryFactory {
	return &CodeCategoryFactory{
		Type: CodeTypeCommon,
	}
}

func Business() *CodeCategoryFactory {
	return &CodeCategoryFactory{
		Type: CodeTypeBusiness,
	}
}

func (f *CodeCategoryFactory) New(cat uint8) *CodeFactory {
	return &CodeFactory{
		Type:     f.Type,
		Category: cat,
	}
}

func (f *CodeFactory) New(subCode int, msg string) *Code {
	c := &Code{
		Type:     f.Type,
		Category: f.Category,
		SubCode:  subCode,
		Msg:      msg,
	}
	c.Code = (10*int(c.Type)+int(c.Category))*1000 + c.SubCode
	return c
}
