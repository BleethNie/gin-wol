# go-wol

## 使用文档

### 获取子网所有设备信息


#### 接口状态
> 开发中

#### 接口URL
> /api/queryDeviceList?subnet=192.168.2.0/24

#### 请求方式
> GET

#### Content-Type
> form-data

#### 请求Query参数

参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
subnet | 192.168.2.0/24 | String | 是 | -

#### 成功响应示例

```json
{
	"code": 200,
	"message": "请求成功",
	"data": [
		{
			"ip": "192.168.2.252",
			"mac": "00-f4-8d-b7-18-a7",
			"nick_name": "",
			"host_name": "hadoop102"
		}
	]
}
```

### 更新设备信息


#### 接口状态
> 开发中

#### 接口URL
> /api/updateDeviceInfo

#### 请求方式
> POST

#### Content-Type
> json

#### 请求Body参数

```json
{
    "ip": "192.168.2.17",
    "mac": "18-56-80-95-30-d0",
    "nick_name": "",
    "host_name": "DESKTOP-N21BJ6M"
}
```

#### 成功响应示例
```json
{
	"ok": "更新成功"
}
```


## 发送wol

#### 接口URL
> /api/wol

#### 请求方式
> POST

#### Content-Type
> form-data

#### 请求Body参数
```json
{
  "mac": "00-f4-8d-b7-18-a7"
}
```

#### 成功响应示例

```json
{
	"code": 200,
	"message": "发送WOL成功",
	"data": ""
}
```

参数名 | 示例值 | 参数类型 | 参数描述
--- | --- | --- | ---
code | 200 | Integer |
message | 发送WOL成功 | String |
data | - | String |

## 清空数据库并更新

#### 接口状态
> 开发中

#### 接口URL
> /api/clearDbAndSave?subnet=192.168.2.0/24

#### 请求方式
> POST

#### Content-Type
> json

#### 请求Query参数

参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
subnet | 192.168.2.0/24 | String | 是 | -

#### 请求Body参数
```json

```

#### 成功响应示例

```json
{
	"code": 200,
	"message": "请求成功",
	"data": [
		
		{
			"ip": "192.168.2.252",
			"mac": "00-f4-8d-b7-18-a7",
			"nick_name": "",
			"host_name": "hadoop102"
		}
	]
}
```



