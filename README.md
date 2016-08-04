# json2struct
json to struct


> 将json文件 转换成go对应的结构体


> InputFileContent

```
{
    "userId": 1001,
    "nickname": "nickname",
    "nickname": "nickname",
    "loginToken": "test@facebook.com",
    "avatar": "/static/pic/myavatar.png",
    "registerCountryCode": "CN",
    "registerLocation": "China",
    "userType": 0,
    "registerTime": 1361416163
}
```

> OutPutContent

```
package output

type UserInfo struct {
	userId              float64 `json:"userId"`
	nickname            string  `json:"nickname"`
	loginToken          string  `json:"loginToken"`
	avatar              string  `json:"avatar"`
	registerCountryCode string  `json:"registerCountryCode"`
	registerLocation    string  `json:"registerLocation"`
	userType            float64 `json:"userType"`
	registerTime        float64 `json:"registerTime"`
}

```



