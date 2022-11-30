package errors

import (
	"fmt"
	"net/http"
	"sync"
)

/**
CoderExtend 拓展,支持国际化
*/

var (
	unknownCoderExtend defaultCoderExtend = defaultCoderExtend{1,
		http.StatusInternalServerError,
		map[string]string{
			"__default__": "An internal server error occurred",
			"zh-cn":       "An internal server error occurred",
			"en":          "内部服务异常",
		},
	}
)

// Coder defines an interface for an error code detail information.
type CoderExtend interface {
	// HTTP status that should be used for the associated error code.
	HTTPStatus() int

	// Code returns the code of the coder
	Code() int

	// External (user) facing error text.
	Message(lang string) string
}

type defaultCoderExtend struct {
	// C refers to the integer code of the ErrCode.
	C int

	// HTTP status that should be used for the associated error code.
	HTTP int

	// client message en zn
	Messages map[string]string
}

// Code returns the integer code of the coder.
func (coder defaultCoderExtend) Code() int {
	return coder.C

}

// HTTPStatus returns the associated HTTP status code, if any. Otherwise,
// returns 200.
func (coder defaultCoderExtend) HTTPStatus() int {
	if coder.HTTP == 0 {
		return 500
	}

	return coder.HTTP
}

// Message returns the  zh-cn or en
func (coder defaultCoderExtend) Message(lang string) string {
	return coder.Messages[lang]
}

// codes contains a map of error codes to metadata.
var codesExtend = map[int]CoderExtend{}
var codeExtendMux = &sync.Mutex{}

// Register register a user define error code.
// It will overrid the exist code.
func RegisterExtend(coder CoderExtend) {
	if coder.Code() == 0 {
		panic("code `0` is reserved by `github.com/marmotedu/errors` as unknownCode error code")
	}

	codeMux.Lock()
	defer codeMux.Unlock()

	codesExtend[coder.Code()] = coder
}

// MustRegister register a user define error code.
// It will panic when the same Code already exist.
func MustRegisterExtend(coder CoderExtend) {
	if coder.Code() == 0 {
		panic("code '0' is reserved by 'github.com/marmotedu/errors' as ErrUnknown error code")
	}

	codeMux.Lock()
	defer codeMux.Unlock()

	if _, ok := codes[coder.Code()]; ok {
		panic(fmt.Sprintf("code: %d already exist", coder.Code()))
	}

	codesExtend[coder.Code()] = coder
}

// ParseCoder parse any error into *withCode.
// nil error will return nil direct.
// None withStack error will be parsed as ErrUnknown.
func ParseCoderExtend(err error) CoderExtend {
	if err == nil {
		return nil
	}

	if v, ok := err.(*withCode); ok {
		if coder, ok := codesExtend[v.code]; ok {
			return coder
		}
	}

	return unknownCoderExtend
}

func init() {
	codesExtend[unknownCoderExtend.Code()] = unknownCoderExtend
}
