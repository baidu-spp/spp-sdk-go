# GO SDK 文档

# 概述

本文档主要介绍百度慧推Go语言版的开发者工具包（SDK），方便用户更迅速的搭建推送开发环境进行调试。

# 安装SDK工具包

## 运行环境

GO SDK可以在go1.3及以上环境下运行。

## 安装SDK

**直接从github下载**

使用`go get`工具从github进行下载：

```shell
go get github.com/xxx/push-sdk
```

# 使用步骤

## 示例


```go
package main

import (
	"fmt"
	"push"
)

func main() {
    // 初始化sdk
	client := push.NewClient($appkey, $masterSecret)
	
	broadcastMsg := &push.Message{
		// 通知栏消息类型
		MessageType: 0,
		// 通知栏消息
		Notification: &push.Notification{
			Title:   "特殊字符来袭@！",
			Content: "( • ̀ω•́ )✧ 酷炫(╯°Д°)╯︵ ┻━┻ 再次掀桌ԅ(¯﹃¯ԅ)─=≡Σ((( つ•̀ω•́)つ( •̥́ ˍ •̀ू )",
			// 点击后续动作
			Action: &push.Action{
				// 打开应用
				ActionType: 6,
				//自定义参数
				Param: make(map[string]string),
			},
		},
		// 筛选条件,可选条件见openapi文档
		Condition: []*push.Condition{
			&push.Condition{
				Key:     "age",
				Values:  []string{"未成年", "老年人"},
				Operate: "or",
			}},
		Option: &push.Option{
			// 消息保存十个小时
			Expire: 36000,
		},
	}
	if response, err := client.Broadcast(broadcastMsg); err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Printf("%#v\n", response)
	}

	unicastMsg := &push.Message{
		// 透传消息类型
		MessageType: 2,
		// 透传消息
		Transmission: &push.Transmission{
			Title:   "透传消息不会展示",
			Content: "content需客户端解析",
		},
		Option: &push.Option{
			Expire: 36000,
		},
	}

	if response, err := client.Unicast(unicastMsg, "zp_test_1"); err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Printf("%#v\n", response)
	}
    
	muticastMsg := &push.Message{
		// 通知栏消息类型
		MessageType: 0,
		// 通知栏消息
		Notification: &push.Notification{
			Title:   "打开链接",
			Content: "打开链接",
			// 点击后续动作
			Action: &push.Action{
				// 打开连接
				ActionType: 9,
				//自定义参数
				Param: map[string]string{
					"user_custom_key": "user_custom_value",
				},
				// 打开的链接
				URL:       "http://www.baidu.com",
				ClassName: "com.baidu.push.demo",
			},
		},
		Option: &push.Option{
			Expire: 36000,
		},
		PushTime: 1575518400,
	}
    // 发送多播
	if response, err := client.Muticast(muticastMsg, []string{"uuid_1", "uuid_2"}); err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Printf("%#v\n", response)
	}
	// 删除消息
    if response, err := client.DelMsg("6150"); err != nil {
        fmt.Printf("%v\n", err)
    } else {
        fmt.Printf("%#v\n", response)
    }

}

```
