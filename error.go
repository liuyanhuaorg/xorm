// Copyright 2015 The Xorm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xorm

import (
	"errors"
	"fmt"
)

var (
	// ErrParamsType params error
	ErrParamsType   = errors.New("Params type error")
	ErrParamsFormat = errors.New("Params format error")
	// ErrTableNotFound table not found error
	ErrTableNotFound = errors.New("Table not found")
	// ErrUnSupportedType unsupported error
	ErrUnSupportedType = errors.New("Unsupported type error")
	// ErrNotExist record does not exist error
	ErrNotExist              = errors.New("Record does not exist")
	ErrNotInTransaction      = errors.New("Not in transaction.")
	ErrNestedTransaction     = errors.New("Nested transaction error.")
	ErrTransactionDefinition = errors.New("Transaction definition error.")
	// ErrCacheFailed cache failed error
	ErrCacheFailed = errors.New("Cache failed")
	// ErrNeedDeletedCond delete needs less one condition error
	ErrNeedDeletedCond = errors.New("Delete action needs at least one condition")
	// ErrNotImplemented not implemented
	ErrNotImplemented = errors.New("Not implemented")
	// ErrConditionType condition type unsupported
	ErrConditionType = errors.New("Unsupported condition type")
	// ErrNeedMoreArguments need more arguments
	ErrNeedMoreArguments = errors.New("Need more sql arguments")
)

// ErrFieldIsNotExist columns does not exist
type ErrFieldIsNotExist struct {
	FieldName string
	TableName string
}

func (e ErrFieldIsNotExist) Error() string {
	return fmt.Sprintf("field %s is not valid on table %s", e.FieldName, e.TableName)
}

// ErrFieldIsNotValid is not valid
type ErrFieldIsNotValid struct {
	FieldName string
	TableName string
}

func (e ErrFieldIsNotValid) Error() string {
	return fmt.Sprintf("field %s is not valid on table %s", e.FieldName, e.TableName)
}
