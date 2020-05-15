// Copyright (c) 2012-present The upper.io/db authors. All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package db

import (
	"reflect"
	"time"

	"github.com/upper/db/internal/adapter"
)

// Comparison represents relationships between values.
type Comparison = adapter.Comparison

// Gte is a comparison that means: is greater than or equal to value.
func Gte(value interface{}) *Comparison {
	return adapter.NewComparisonOperator(adapter.ComparisonOperatorGreaterThanOrEqualTo, value)
}

// Lte is a comparison that means: is less than or equal to value.
func Lte(value interface{}) *Comparison {
	return adapter.NewComparisonOperator(adapter.ComparisonOperatorLessThanOrEqualTo, value)
}

// Eq is a comparison that means: is equal to value.
func Eq(value interface{}) *Comparison {
	return adapter.NewComparisonOperator(adapter.ComparisonOperatorEqual, value)
}

// NotEq is a comparison that means: is not equal to value.
func NotEq(value interface{}) *Comparison {
	return adapter.NewComparisonOperator(adapter.ComparisonOperatorNotEqual, value)
}

// Gt is a comparison that means: is greater than value.
func Gt(value interface{}) *Comparison {
	return adapter.NewComparisonOperator(adapter.ComparisonOperatorGreaterThan, value)
}

// Lt is a comparison that means: is less than value.
func Lt(value interface{}) *Comparison {
	return adapter.NewComparisonOperator(adapter.ComparisonOperatorLessThan, value)
}

// In is a comparison that means: is any of the values.
func In(value ...interface{}) *Comparison {
	return adapter.NewComparisonOperator(adapter.ComparisonOperatorIn, toInterfaceArray(value))
}

// After is a comparison that means: is after the (time.Time) value.
func After(value time.Time) *Comparison {
	return adapter.NewComparisonOperator(adapter.ComparisonOperatorGreaterThan, value)
}

// Before is a comparison that means: is before the (time.Time) value.
func Before(value time.Time) *Comparison {
	return adapter.NewComparisonOperator(adapter.ComparisonOperatorLessThan, value)
}

// OnOrAfter is a comparison that means: is on or after the (time.Time) value.
func OnOrAfter(value time.Time) *Comparison {
	return adapter.NewComparisonOperator(adapter.ComparisonOperatorGreaterThanOrEqualTo, value)
}

// OnOrBefore is a comparison that means: is on or before the (time.Time) value.
func OnOrBefore(value time.Time) *Comparison {
	return adapter.NewComparisonOperator(adapter.ComparisonOperatorLessThanOrEqualTo, value)
}

// Between is a comparison that means: is between valueA and valueB.
func Between(valueA interface{}, valueB interface{}) *Comparison {
	return adapter.NewComparisonOperator(adapter.ComparisonOperatorBetween, []interface{}{valueA, valueB})
}

// NotBetween is a comparison that means: is not between valueA and valueB.
func NotBetween(valueA interface{}, valueB interface{}) *Comparison {
	return adapter.NewComparisonOperator(adapter.ComparisonOperatorNotBetween, []interface{}{valueA, valueB})
}

// Is is a comparison that means: is equivalent to nil, true or false.
func Is(value interface{}) *Comparison {
	return adapter.NewComparisonOperator(adapter.ComparisonOperatorIs, value)
}

// IsNot is a comparison that means: is not equivalent to nil, true nor false.
func IsNot(value interface{}) *Comparison {
	return adapter.NewComparisonOperator(adapter.ComparisonOperatorIsNot, value)
}

// IsNull is a comparison that means: is equivalent to nil.
func IsNull() *Comparison {
	return adapter.NewComparisonOperator(adapter.ComparisonOperatorIs, nil)
}

// IsNotNull is a comparison that means: is not equivalent to nil.
func IsNotNull() *Comparison {
	return adapter.NewComparisonOperator(adapter.ComparisonOperatorIsNot, nil)
}

// Like is a comparison that checks whether the reference matches the wildcard
// value.
func Like(value string) *Comparison {
	return adapter.NewComparisonOperator(adapter.ComparisonOperatorLike, value)
}

// NotLike is a comparison that checks whether the reference does not match the
// wildcard value.
func NotLike(value string) *Comparison {
	return adapter.NewComparisonOperator(adapter.ComparisonOperatorNotLike, value)
}

// RegExp is a comparison that checks whether the reference matches the regular
// expression.
func RegExp(value string) *Comparison {
	return adapter.NewComparisonOperator(adapter.ComparisonOperatorRegExp, value)
}

// NotRegExp is a comparison that checks whether the reference does not match
// the regular expression.
func NotRegExp(value string) *Comparison {
	return adapter.NewComparisonOperator(adapter.ComparisonOperatorNotRegExp, value)
}

// Op returns a custom comparison operator.
func Op(customOperator string, value interface{}) *Comparison {
	return adapter.NewCustomComparisonOperator(customOperator, value)
}

func toInterfaceArray(value interface{}) []interface{} {
	rv := reflect.ValueOf(value)
	switch rv.Type().Kind() {
	case reflect.Ptr:
		return toInterfaceArray(rv.Elem().Interface())
	case reflect.Slice:
		elems := rv.Len()
		args := make([]interface{}, elems)
		for i := 0; i < elems; i++ {
			args[i] = rv.Index(i).Interface()
		}
		return args
	}
	return []interface{}{value}
}
