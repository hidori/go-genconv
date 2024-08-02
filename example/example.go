package example

import (
	"github.com/hidori/go-genmapper/converter"
	"github.com/hidori/go-genmapper/dsl"
)

type Source struct {
	Value1 int
	Value2 int
}

type Destination struct {
	Value1 int
	Value2 string
	Value3 int
}

var _ = dsl.CreateMapper[Source, Destination]("MapSourceToDestination", func(src Source, dst Destination) {
	dsl.ForMember(dst.Value2, converter.IntToString)
	dsl.Ignore(dst.Value3)
})
