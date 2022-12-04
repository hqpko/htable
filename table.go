package htable

import (
	"fmt"
)

type TableItemType string

const (
	TableItemTypeInt     TableItemType = "int"
	TableItemTypeInt64                 = "int64"
	TableItemTypeFloat                 = "float"
	TableItemTypeFloat64               = "float64"
	TableItemTypeString                = "string"
	TableItemTypeCustom                = "custom" // 自定义类型
)

type Table struct {
	Name  string
	Path  string
	Items []*TableItem
}

type TableItem struct {
	Struct    *TableItemStruct
	Values    []*TableItemValue
	valuesMap map[string]*TableItemValue
}

func (ti *TableItem) get(value string) *TableItemValue {
	if (ti.valuesMap) == nil {
		ti.createMap()
	}
	return ti.valuesMap[value]
}

func (ti *TableItem) has(value string) bool {
	return ti.get(value) != nil
}

func (ti *TableItem) createMap() {
	ti.valuesMap = map[string]*TableItemValue{}
	for _, value := range ti.Values {
		if ti.Struct.IsArray {
			panic(fmt.Sprintf("table item key(%s) is array", ti.Struct.Name))
		} else if ti.Struct.Type == TableItemTypeCustom {
			panic(fmt.Sprintf("table item key(%s) is custom type", ti.Struct.Name))
		}
		ti.valuesMap[value.str(ti.Struct.Type)] = value
	}
}

type TableItemStruct struct {
	Name       string
	Type       TableItemType
	IsArray    bool
	CustomType []*TableItemStruct // type == custom 时的类型说明
}

type TableItemValue struct {
	Int         int
	Int64       int64
	Float       float32
	Float64     float64
	String      string
	CustomValue []*TableItemValue
	ArrayValues []*TableItemValue
}

func (tiv *TableItemValue) str(t TableItemType) string {
	if t == TableItemTypeInt {
		return fmt.Sprintf("%d", tiv.Int)
	} else if t == TableItemTypeInt64 {
		return fmt.Sprintf("%d", tiv.Int64)
	} else if t == TableItemTypeFloat {
		return fmt.Sprintf("%f", tiv.Float)
	} else if t == TableItemTypeFloat64 {
		return fmt.Sprintf("%f", tiv.Float64)
	} else if t == TableItemTypeString {
		return tiv.String
	}
	return tiv.String
}
