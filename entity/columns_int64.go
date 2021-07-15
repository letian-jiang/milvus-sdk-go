// Code generated by go generate; DO NOT EDIT
// This file is generated by go generate at 2021-07-15 15:55:19.25274366 +0800 CST m=+0.002663389

//Package entity defines entities used in sdk
package entity 

import (
	"errors"

	"github.com/milvus-io/milvus-sdk-go/v2/internal/proto/schema"
)

// ColumnInt64 generated columns type for Int64
type ColumnInt64 struct {
	name   string
	values []int64
}

// Name returns column name
func (c *ColumnInt64) Name() string {
	return c.name
}

// Type returns column FieldType
func (c *ColumnInt64) Type() FieldType {
	return FieldTypeInt64
}

// Len returns column values length
func (c *ColumnInt64) Len() int {
	return len(c.values)
}

// FieldData return column data mapped to schema.FieldData
func (c *ColumnInt64) FieldData() *schema.FieldData {
	fd := &schema.FieldData{
		Type: schema.DataType_Int64,
		FieldName: c.name,
	}
	data := make([]int64, 0, c.Len())
	for i := 0 ;i < c.Len(); i++ {
		data = append(data, int64(c.values[i]))
	}
	fd.Field = &schema.FieldData_Scalars{
		Scalars: &schema.ScalarField{
			Data: &schema.ScalarField_LongData{
				LongData: &schema.LongArray{
					Data: data,
				},
			},
		},
	}
	return fd
}

// ValueByIdx returns value of the provided index
// error occurs when index out of range
func (c *ColumnInt64) ValueByIdx(idx int) (int64, error) {
	var r int64 // use default value
	if idx < 0 || idx >= c.Len() {
		return r, errors.New("index out of range")
	}
	return c.values[idx], nil
}

// NewColumnInt64 auto generated constructor
func NewColumnInt64(name string, values []int64) *ColumnInt64 {
	return &ColumnInt64 {
		name: name,
		values: values,
	}
}
