### 标签管理

#### 创建标签

**请求方式：** POST

**请求地址：** https://udcp.api.com/openapi/v1/tag?access_token=ACCESS_TOKEN

**请求包体：**

```
{
   "tag_name": "UI",
   "tag_id": 12
}
```

**参数说明：**

| 参数 	| 必选 | 类型   | 说明                                    |
| :--- | :--- | :----- | --------------------------------------- |
| access_token | 是 | string | 调用接口凭证                            |
| tag_name | 是 | string    | 标签名称，长度限制为32个字以内（汉字或英文字母），标签名不可与其他标签重名。 |
| tag_id | 否 | int | 标签id，非负整型，指定此参数时新增的标签会生成对应的标签id，不指定时则以目前最大的id自增。 |


**返回结果：**

```
{
   "code": 0,
   "message": "created",
   "data": {
       "tag_id": 12
   }
}
```

**参数说明：**

| 参数 			| 类型 			| 说明                         |
| :------- | :------- | :--------------------------- |
| code     | int      | 状态码(0成功，非0都是失败) |
| message  | string   | 提示信息                     |
| tag_id   | int      | 标签id                     |











#### 删除标签

**请求方式：** DELETE

**请求地址：** https://udcp.api.com/openapi/v1/tag?access_token=ACCESS_TOKEN

**请求包体：**

```
{
   "tag_id": 12
}
```

**参数说明：**

| 参数 	| 必选 | 类型   | 说明                                    |
| :--- | :--- | :----- | --------------------------------------- |
| access_token | 是 | string | 调用接口凭证                            |
| tag_id | 是 | int | 标签id，非负整型，指定此参数时新增的标签会生成对应的标签id，不指定时则以目前最大的id自增。 |


**返回结果：**

```
{
   "code": 0,
   "message": "ok",
}
```

**参数说明：**

| 参数 			| 类型 			| 说明                         |
| :------- | :------- | :--------------------------- |
| code     | int      | 状态码(0成功，非0都是失败) |
| message  | string   | 提示信息                     |















#### 部门执行策略展示

**请求方式：** GET

**请求地址：**[/api/v1/policy_view](/api/v1/policy_group)

**请求包体：**

```
{
   "dpt_id": "12",
}
```

**参数说明：**

| 参数   | 必选 | 类型 | 说明  |
| :----- | :--- | :--- | :---- |
| dpt_id | 是   | int  | 部门i |

**返回结果：**

```
{
   "code": 0,
   "message": "created",
   "data": {
       {
                  "number": "PG-000001",
                  "name": "策略组示例1",
                  "force_publish": true,
                  "disable": false,
                  "remark": "示例",
                  "created_at": "2021-01-01 20:00:00",
                  "updated_at": "2021-01-01 20:00:00",
                  "scope": {
                    "departments": [
                      {
                        "id": 10,
                        "policy_key": "研发一部"
                      },
                      {
                        "id": 11,
                        "policy_key": "研发二部"
                      },
                      {
                        "id": 13,
                        "policy_key": "研发四部"
                      }
                    ],
                    "pc_type": [
                      1
                    ],
                    "os_archs": [
                      "amd",
                      "arm",
                      "mips"
                    ],
                    "os_versions": [
                      "1020",
                      "1021",
                      "1022",
                      "1030"
                    ],
                    "exec_type": 2,
                    "pc_groups": [
                      {
                        "id": 100,
                        "policy_key": "testPC-100"
                      },
                      {
                        "id": 101,
                        "policy_key": "testPC-101"
                      },
                      {
                        "id": 103,
                        "policy_key": "testPC-103"
                      }
                    ]
                  },
                  "policys": [
                    {
                      "id": 100,
                      "policy_key": "screen_saver",
                      "description": "配置是否自动启用屏保，可设置自动启用时间",
                      "disable": true,
                      "lock": false,
                      "settings": [
                        {
                          "id": 100,
                          "setting_key": "time_out",
                          "setting_val": 1
                        }
                      ]
                    },
                    {
                      "id": 101,
                      "policy_key": "close_screen",
                      "description": "配置是否自动息屏，可设置自动息屏时间",
                      "disable": true,
                      "lock": false,
                      "settings": [
                        {
                          "id": 101,
                          "setting_key": "time_out",
                          "setting_val": 1
                        }
                      ]
                    },
                    {
                      "id": 102,
                      "policy_key": "wallpaper",
                      "description": "设置终端桌面壁纸，支持PNG、JPG格式图片，小于2M，推荐3:2比例",
                      "disable": true,
                      "lock": false,
                      "settings": [
                        {
                          "id": 102,
                          "setting_key": "wallpaper",
                          "setting_val": "https://example.com/udcp/policy/data/wallpaper/111.ipg"
                        }
                      ]
                    },
                    {
                      "id": 103,
                      "policy_key": "script",
                      "description": "上传终端可执行的脚本，终端接收后自动执行",
                      "disable": true,
                      "lock": false,
                      "settings": [
                        {
                          "id": 103,
                          "setting_key": "script",
                          "setting_val": "https://example.com/udcp/policy/data/script/update.sh"
                        },
                        {
                          "id": 104,
                          "setting_key": "script",
                          "setting_val": "https://example.com/udcp/policy/data/script/update2.sh"
                        },
                        {
                          "id": 105,
                          "setting_key": "script",
                          "setting_val": "https://example.com/udcp/policy/data/script/update5.sh"
                        }
                      ]
                    },
                    {
                      "id": 104,
                      "policy_key": "shortcut",
                      "description": "可配置是否允许使用系统快捷键",
                      "disable": true,
                      "lock": false,
                      "settings": [
                        {
                          "id": 106,
                          "setting_key": "shortcut",
                          "setting_val": "1"
                        }
                      ]
                    },
                    {
                      "id": 105,
                      "policy_key": "taskbar",
                      "description": "任务栏功能配置",
                      "disable": true,
                      "lock": false,
                      "settings": [
                        {
                          "id": 106,
                          "setting_key": "display_mode",
                          "setting_val": "1"
                        },
                        {
                          "id": 107,
                          "setting_key": "launcher",
                          "setting_val": "1"
                        },
                        {
                          "id": 108,
                          "setting_key": "right_menu",
                          "setting_val": "1"
                        },
                        {
                          "id": 109,
                          "setting_key": "dock_app",
                          "setting_val": "1"
                        }
                      ]
                    },
                    {
                      "id": 106,
                      "policy_key": "app",
                      "description": "任务栏右侧应用管理",
                      "disable": true,
                      "lock": false,
                      "settings": [
                        {
                          "id": 110,
                          "setting_key": "trash",
                          "setting_val": "1"
                        }
                      ]
                    },
                    {
                      "id": 107,
                      "policy_key": "tray_icon",
                      "description": "托盘区域功能管理",
                      "disable": true,
                      "lock": false,
                      "settings": [
                        {
                          "id": 111,
                          "setting_key": "network_tray_icon",
                          "setting_val": "1"
                        },
                        {
                          "id": 112,
                          "setting_key": "usb_tray_icon",
                          "setting_val": "1"
                        }
                      ]
                    },
                    {
                      "id": 109,
                      "policy_key": "control_center",
                      "description": "控制中心配置",
                      "disable": true,
                      "lock": false,
                      "settings": [
                        {
                          "id": 114,
                          "setting_key": "control_center_union_id",
                          "setting_val": "1"
                        },
                        {
                          "id": 115,
                          "setting_key": "control_center_display",
                          "setting_val": "1"
                        },
                        {
                          "id": 116,
                          "setting_key": "control_center_defapp",
                          "setting_val": "1"
                        },
                        {
                          "id": 117,
                          "setting_key": "control_center_personalization",
                          "setting_val": "1"
                        },
                        {
                          "id": 118,
                          "setting_key": "control_center_network",
                          "setting_val": "1"
                        },
                        {
                          "id": 119,
                          "setting_key": "control_center_sound",
                          "setting_val": "1"
                        },
                        {
                          "id": 120,
                          "setting_key": "control_center_bluetooth",
                          "setting_val": "1"
                        },
                        {
                          "id": 121,
                          "setting_key": "control_center_datetime",
                          "setting_val": "1"
                        },
                        {
                          "id": 122,
                          "setting_key": "control_center_notification",
                          "setting_val": "1"
                        },
                        {
                          "id": 123,
                          "setting_key": "control_center_power",
                          "setting_val": "1"
                        },
                        {
                          "id": 124,
                          "setting_key": "control_center_mouse",
                          "setting_val": "1"
                        },
                        {
                          "id": 125,
                          "setting_key": "control_center_wacom",
                          "setting_val": "1"
                        },
                        {
                          "id": 126,
                          "setting_key": "control_center_keyboard",
                          "setting_val": "1"
                        },
                        {
                          "id": 127,
                          "setting_key": "control_center_voice",
                          "setting_val": "1"
                        },
                        {
                          "id": 128,
                          "setting_key": "control_center_systeminfo",
                          "setting_val": "1"
                        },
                        {
                          "id": 129,
                          "setting_key": "control_center_commoninfo",
                          "setting_val": "1"
                        }
                      ]
                    },
                    {
                      "id": 110,
                      "policy_key": "firewall",
                      "description": "使用iptables配置终端防火墙策略",
                      "disable": true,
                      "lock": false,
                      "settings": [
                        {
                          "id": 130,
                          "setting_key": "firewall",
                          "setting_val": ""
                        }
                      ]
                    },
                    {
                      "id": 111,
                      "policy_key": "install_apps",
                      "description": "管理终端需要安装的应用清单",
                      "disable": true,
                      "lock": false,
                      "settings": [
                        {
                          "id": 131,
                          "setting_key": "firefox",
                          "setting_val": "{\"policy_key\":\"firefox\",\"package_policy_key\":\"firefox\",\"version\":\"1.0.0.1\",\"architecture\":\"arm\"}"
                        },
                        {
                          "id": 132,
                          "setting_key": "chrome",
                          "setting_val": "{\"policy_key\":\"chrome\",\"package_policy_key\":\"chrome\",\"version\":\"1.0.0.1\",\"architecture\":\"arm\"}"
                        }
                      ]
                    },
                    {
                      "id": 112,
                      "policy_key": "user_password_strength",
                      "description": "配置域管帐户的密码强度策略",
                      "disable": true,
                      "lock": false,
                      "settings": [
                        {
                          "id": 133,
                          "setting_key": "password_length",
                          "setting_val": "[8,32]"
                        }
                      ]
                    }
                  ]
                }
              }
   }
}
```

**参数说明：**

| 参数    | 类型   | 说明                              |
| :------ | :----- | :-------------------------------- |
| code    | int    | 状态码(0成功，非0都是失败)        |
| message | string | 提示信息                          |
| data    | int    | 策略详情   （具体细项同策略详情） |







####  删除策略（单个/多个)

**请求方式：** DELETE

**请求地址：**[/api/v1/policy_group]

**请求包体：**

```
{
    "pg_id": [1，3，5]
}
```

**参数说明：**

| 参数  | 必选 | 类型  | 说明                        |
| :---- | :--- | :---- | :-------------------------- |
| pg_id | 是   | int[] | 删除的策略组id （数组形式） |

**返回结果：**

```
{
    		"code": 0,
        "msg": "",
        "data": {}
}
```

**参数说明：**

| 参数 | 类型   | 说明                       |
| :--- | :----- | :------------------------- |
| code | int    | 状态码(0成功，非0都是失败) |
| msg  | string | 提示信息                   |
| data | int    | 空对象                     |











####  策略组排序

**请求方式：** POST

**请求地址：**[/api/v1/policy_sort]

**请求包体：**

```
[
  {
    "pg_id": 10,
    "sorting": 1
  },
   {
    "pg_id": 130,
    "sorting": 2
  },
   {
    "pg_id": 133,
    "sorting": 3
  },
]
```

**参数说明：**

| 参数    | 必选 | 类型 | 说明                                              |
| :------ | :--- | :--- | :------------------------------------------------ |
| pg_id   | 是   | int  | 策略组id                                          |
| sorting | 是   | int  | 前端排序的顺序号  严格按照1，2，3，4。。。的 顺序 |

**返回结果：**

```
{
    		"code": 0,
        "msg": "",
        "data": {}
}
```

**参数说明：**

| 参数 | 类型   | 说明                       |
| :--- | :----- | :------------------------- |
| code | int    | 状态码(0成功，非0都是失败) |
| msg  | string | 提示信息                   |
| data | int    | 空对象                     |







#### 部门策略视图  （关联 & 继承）

**请求方式：** GET

**请求地址：**[/api/v1/policy_relation]

**请求包体：**

```
 {
    "dpt_id": 10,
    "type":0
  }
```

**参数说明：**

| 参数   | 必选 | 类型 | 说明                        |
| :----- | :--- | :--- | :-------------------------- |
| dpt_id | 是   | int  | 部门id                      |
| type   | 是   | int  | 0(关联视图）  1（继承视图） |

**返回结果：**

```
{
 				"code": 0,
        "msg": "",
        "data": {
            [
              {
                "id": "1",
                "name": "NAME1",//策略组名称
                "create_time": 1557838797,
                "desc": "memo",//策略组备注
                "department": "研发中心",//执行部门
                "publisher":"aaa"//发布人
                "force": 0,  //强制   1是 0否
                "type": 1   //类型 0(关联视图）  1（继承视图）
              },
              {
                "id": "2",
                "name": "NAME1",//策略组名称
                "create_time": 1557838797,
                "desc": "memo",//策略组备注
                "department": "研发中心",//执行部门
                "publisher":"aaa"//发布人
                "force": 0,  //强制   1是 0否
                "type": 1   //类型 0(关联视图）  1（继承视图）
              }
            ]
        }
}
```

**参数说明：**

| 参数 | 类型   | 说明                       |
| :--- | :----- | :------------------------- |
| code | int    | 状态码(0成功，非0都是失败) |
| msg  | string | 提示信息                   |
| data | int    | 空对象                     |