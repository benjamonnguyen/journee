package types

// Pointer returns a pointer to the given object.
// Can be used for "nullable" objects.
func Pointer[T any](o T) *T {
	return &o
}

// type Optional[T any] struct {
// 	Value   *T
// 	Default T
// }

// func (o Optional[T]) HasValue() bool {
// 	return o.Value != nil
// }

// func (o Optional[T]) ValueOrDefault() T {
// 	if o.HasValue() {
// 		return *o.Value
// 	}

// 	return o.Default
// }
