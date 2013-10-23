package fragment

import (
	"fmt"
	"reflect"

	"github.com/SoCloz/goprismic/fragment/link"
)

// A number
type Number float64

func (n *Number) Decode(_ string, enc interface{}) error {
	dec, ok := enc.(float64)
	if !ok {
		return fmt.Errorf("unable to decode number fragment : %+v is a %s, not a number", enc, reflect.TypeOf(enc))
	}
	*n = Number(dec)
	return nil
}

func (n *Number) AsText() string {
	return fmt.Sprintf("%f", *n)
}

func (n *Number) AsHtml() string {
	return fmt.Sprintf("<span class=\"number\">%f</span>", *n)
}

func (n *Number) ResolveLinks(_ link.Resolver) {}