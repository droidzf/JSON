package JSON

import (
	"encoding/json"
	"reflect"
)

type Object struct {
	v interface{}
}

func NewObject() *Object {
	j := &Object{}
	return j
}

//struct转json字符串
func ToJsonString(p interface{}) string {
	result, _ := json.Marshal(p)
	return string(result)
}

//解析jsonObject字符串
func ParseObject(s string) *Object {
	data := []byte(s)
	var final map[string]interface{}
	err := json.Unmarshal(data, &final)
	if err != nil {
		panic(err)
	}
	j := &Object{}
	j.v = final
	return j
}

//获取jsonObject
func (j *Object) GetObject(key string) *Object {
	res := &Object{}
	res.v = j.v.(map[string]interface{})[key]
	return res
}

//获取jsonArray
func (j *Object) GetArray(key string) *Array {
	array := &Array{}
	for _, v := range j.v.(map[string]interface{})[key].([]interface{}) {
		array.objects = append(array.objects, v)
	}
	return array
}

//获取jsonObject中的值
func (j *Object) Get(key string) interface{} {
	return j.v.(map[string]interface{})[key]
}

//获取字符串
func (j *Object) GetString(key string) string {
	return j.Get(key).(string)
}

//获取bool类型
func (j *Object) GetBoolean(key string) bool {
	return j.Get(key).(bool)
}

//获取int64
func (j *Object) GetLong(key string) int64 {
	return int64(j.Get(key).(float64))
}

//获取int
func (j *Object) GetInt(key string) int {
	return int(j.Get(key).(float64))
}

//获取float64
func (j *Object) GetFloat(key string) float64 {
	return j.Get(key).(float64)
}

//更改或添加
//可以是JSON.Array、Json.Object、interface
//原json对象中数组内容会改变
func (j *Object) Set(key string, value interface{}) *Object {
	t := reflect.TypeOf(value)
	if t.String() == "*JSON.Object" {
		value = value.(*Object).v
	} else if t.String() == "*JSON.Array" {
		value = value.(*Array).objects
	}
	j.v.(map[string]interface{})[key] = value
	return j
}

//删除
//原json对象中数组内容会改变
func (j *Object) Remove(key string) *Object {
	delete(j.v.(map[string]interface{}), key)
	return j
}

//获取key和value数组
//keys:[]string
//values: JSON.Array
func (j *Object) Keys() (keys []string, values *Array) {
	values = &Array{}
	keys = make([]string, len(j.v.(map[string]interface{})))
	for k, v := range j.v.(map[string]interface{}) {
		keys = append(keys, k)
		values.objects = append(values.objects, v)
	}
	return
}

//转字符串
func (j *Object) ToString() string {
	result, _ := json.Marshal(j.v)
	return string(result)
}

//格式化字符串
func (j *Object) ToStringIndent() string {
	result, _ := json.MarshalIndent(j.v, "", "\t")
	return string(result)
}
