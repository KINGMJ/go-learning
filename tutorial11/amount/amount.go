package amount

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var (
	regAmount = regexp.MustCompile(`^\d+(\.\d{1,2})?$`)
)

// Amount 金额 单位：分
type Amount int64

// String xx.xx
func (a Amount) String() (s string) {
	var negative bool
	if a < 0 {
		negative = true
		a = -a
	}
	defer func() {
		if negative {
			s = "-" + s
		}
	}()
	if a < 100 {
		return fmt.Sprintf("0.%02d", a)
	}
	s = strconv.FormatInt(int64(a), 10)
	l := len(s)
	return s[0:l-2] + "." + s[l-2:l]
}

// CalcTotalPrice doc
func (a Amount) CalcTotalPrice(count int64) Amount {
	return Amount(int64(a) * count)
}

func (a Amount) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.String())
}

func (a *Amount) UnmarshalJSON(data []byte) error {
	var amountStr string
	if err := json.Unmarshal(data, &amountStr); err != nil {
		return fmt.Errorf("无法解析金额: %w", err)
	}
	amountCents, err := Parse(amountStr)
	if err != nil {
		return err
	}
	*a = amountCents
	return nil
}

func (a *Amount) UnmarshalText(data []byte) error {
	amountStr := string(data)
	amountCents, err := Parse(amountStr)
	if err != nil {
		return err
	}
	*a = amountCents
	return nil
}

// Parse 解析金额字符串
func Parse(amountStr string) (Amount, error) {
	// 移除金额字符串中的逗号
	amountStr = strings.ReplaceAll(amountStr, ",", "")
	// 检查金额格式
	if !regAmount.MatchString(amountStr) {
		return 0, fmt.Errorf("金额格式错误: %s", amountStr)
	}
	// 将金额字符串转换为浮点数
	amountFloat, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		return 0, fmt.Errorf("无法解析金额: %w", err)
	}

	// 将浮点数转换为以分为单位的整型数值
	return Amount(amountFloat * 100), nil
}

// UnmarshalJSONWithInt64 JSON反序列化结构体，将 int64 转换为 Amount
func UnmarshalJSONWithInt64(data []byte, v any) error {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Pointer {
		if val.IsNil() {
			return errors.New("nil pointer")
		}
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return fmt.Errorf("expected a struct")
	}

	srcType := val.Type()
	newType := replaceAmountWithInt64(srcType)

	newStruct := reflect.New(newType).Interface()
	if err := json.Unmarshal(data, newStruct); err != nil {
		return err
	}

	ns := reflect.ValueOf(newStruct).Elem()
	for i := 0; i < val.NumField(); i++ {
		field := srcType.Field(i)
		newVal := ns.Field(i)
		if field.Type == reflect.TypeOf((*Amount)(nil)).Elem() {
			val.Field(i).Set(reflect.ValueOf(Amount(newVal.Int())))
		} else {
			val.Field(i).Set(newVal)
		}
	}

	return nil
}

// MarshalWithAmount JSON序列化结构体，将 Amount 转换为 int64
func MarshalWithInt64(v any) ([]byte, error) {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Pointer {
		if val.IsNil() {
			return nil, errors.New("nil pointer")
		}
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected a struct")
	}

	srcType := val.Type()
	newType := replaceAmountWithInt64(srcType)

	newStruct := reflect.New(newType).Elem()
	for i := 0; i < val.NumField(); i++ {
		field := srcType.Field(i)
		newVal := newStruct.Field(i)
		if field.Type == reflect.TypeOf((*Amount)(nil)).Elem() {
			newVal.SetInt(int64(val.Field(i).Int()))
		} else {
			newVal.Set(val.Field(i))
		}
	}

	return json.Marshal(newStruct.Interface())
}

func replaceAmountWithInt64(srcType reflect.Type) reflect.Type {
	newFields := make([]reflect.StructField, srcType.NumField())
	for i := 0; i < srcType.NumField(); i++ {
		field := srcType.Field(i)
		if field.Type == reflect.TypeOf((*Amount)(nil)).Elem() {
			newFields[i] = reflect.StructField{
				Name:      field.Name,
				Type:      reflect.TypeOf(int64(0)),
				Tag:       field.Tag,
				Anonymous: field.Anonymous,
			}
		} else {
			newFields[i] = field
		}
	}
	return reflect.StructOf(newFields)
}
