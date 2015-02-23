package fragment

import (
	"github.com/durgeshm/goprismic/fragment/link"
)

type Interface interface {
	Decode(string, interface{}) error
	AsText() string
	AsHtml() string
	ResolveLinks(link.Resolver)
}
