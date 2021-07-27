package JSON

import (
	"encoding/json"
	"reflect"
)

type Array struct {
	objects []interface{}
}

func NewArray() *Array {
	j := &Array{}
	return j
}

//解析jsonArray字符串
//`{"name":"json"}` 也可以传入JSON.ToJsonString(interface)
func ParseArray(s string) *Array {
	data := []byte(s)
	var final interface{}
	_ = json.Unmarshal(data, &final)
	array := &Array{}
	for _, v := range final.([]interface{}) {
		array.objects = append(array.objects, v)
	}
	return array
}

//获取jsonArray中的值
func (j *Array) Get(index int) interface{} {
	return j.objects[index]
}

//获取jsonArray中的jsonObject
func (j *Array) GetObject(index int) *Object {
	obj := &Object{}
	obj.v = j.objects[index]
	return obj
}

//获取jsonArray中的jsonArray
func (j *Array) GetArray(index int) *Array {
	array := &Array{}
	array.objects = j.objects[index].([]interface{})
	return array
}

//获取字符串
func (j *Array) GetString(index int) string {
	return j.Get(index).(string)
}

//获取bool类型
func (j *Array) GetBoolean(index int) bool {
	return j.Get(index).(bool)
}

//获取int64
func (j *Array) GetLong(index int) int64 {
	return int64(j.Get(index).(float64))
}

//获取int
func (j *Array) GetInt(index int) int {
	return int(j.Get(index).(float64))
}

//获取float64
func (j *Array) GetFloat(index int) float64 {
	return j.Get(index).(float64)
}

//向jsonArray中添加值
//可以是JSON.Array、Json.Object、interface
//原json对象中数组内容不变
func (j *Array) Put(v interface{}) *Array {
	t := reflect.TypeOf(v)
	if t.String() == "*JSON.Object" {
		item := v.(*Object).v
		j.objects = append(j.objects, item)
	} else if t.String() == "*JSON.Array" {
		item := v.(*Array).objects
		j.objects = append(j.objects, item)
	} else {
		j.objects = append(j.objects, v)
	}
	return j
}

//删除数组中元素
//原json对象中数组内容不变
func (j *Array) Remove(index int) *Array {
	j.objects = append(j.objects[:index], j.objects[index+1:]...)
	return j
}

//jsonArray长度
func (j *Array) Length() int {
	return len(j.objects)
}

//jsonArray转字符串
func (j *Array) ToString() string {
	result, _ := json.Marshal(j.objects)
	return string(result)
}

//jsonArray转格式化字符串
func (j *Array) ToStringIndent() string {
	result, _ := json.MarshalIndent(j.objects, "", "\t")
	return string(result)
}
