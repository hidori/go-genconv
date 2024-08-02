package dsl

type CreateMapperDirective any

func CreateMapper[TSource any, TDestination any](name string, directives ...CreateMapperDirective) any {
	return nil
}

func ForMember[TSourceField any, TDestinationField any](dstFieldSelector TDestinationField, conv func(srcFieldValue TSourceField) (TDestinationField, error)) CreateMapperDirective {
	return nil
}

func Ignore[TDestinationField any](dstFieldSelector TDestinationField) CreateMapperDirective {
	return nil
}
