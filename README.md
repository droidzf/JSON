# JSON
JSON解析

```go
testJson := JSON.ParseObj(`{
        "status": 0,
        "msg":"success",
        "success":true,
		"data": {
			"count": 4,
			"array":[[{"a":1},{"a":2}],[{"a":3},{"a":4}]],
			"rows": [
				{
					"id": 598,
					"name": "李白",
					"org": "唐朝"
				},
				{
					"id": 597,
					"user_name": "李黑"
				}
			]
        }   
	}`)
println(testJson.ToString())
```

####获取object
```go
testObj := result.GetObject("data")
println(testObj.ToString())
```
####获取array
```go
testarray :=testObj.GetArray("array").GetArray(0)
println(testarray.ToString())//[{"a":1},{"a":2}]
```
####获取值
```go
status := testJson.Get("status").(int)   //0
sint   := testJson.GetInt("status")      //0
suc    := testJson.GetString("msg")      //success
boo    := testJson.GetBoolean("success") //true
result.GetObject("data").GetArray("rows").GetObject(0).GetString("name")//李白
```
####API

|  方法名        |  说明  |
|  ----         | ----  |
| GetObject     | 获取jsonObject|
| GetArray      | 获取jsonArray|
| ToJsonString  | struct/map[string]interface{}转json字符串|
| ToString      | 转字符串 |
| ToStringIndent| 格式化输出字符串 |
| Get           | 获取值返回interface|
| GetString     | 获取字符串|
| GetBoolean    | 获取bool类型|
| GetLong       | 获取int64|
| GetInt        | 获取int|
| GetFloat      | 获取float64|

Object方法

|  方法名       | 说明  |
|  ----        | ----  |
| NewObject    | 生成新Object |
| ParseObject  | 解析json字符串为Object |
| Remove       | 删除<br>原json对象中数组内容会改变|
| Set          | 更改或添加<br>可以是JSON.Array、Json.Object、interface<br>原json对象中数组内容会改变|
| Keys         | 获取key和value数组<br>keys:[]string<br>values: JSON.Array|

Array 方法

|  方法名     | 说明  |
|  ----      | ----  |
| NewArray   | 生成新Array |
| ParseArray | 解析json字符串为Array |
| Remove     | 删除<br>原json对象中数组内容不变|
| Put        | 添加<br>可以是JSON.Array、Json.Object、interface<br>原json对象中数组内容不变|
| Length     | Array长度|