# Web API

[TOC]

## 基础说明 

### 响应格式 

- **Content-Type**: `application/json`
- 报文

| 参数名 | 类型 | 必填 | 描述 |
| -  | -    | - | - |
|code|string|是  |“0” 为成功，其它异常|
|msg |string|否  |错误信息|
|data|any   |否  |返回数据|

### auth 认证 

* 参数配置 
  - **参数位置**: `header`
  - **参数名称**: `token`
  - **参数来源**: `接口 /login`


## 基础接口 

### 获取图形验证码 

* 基本信息 
  - **请求URL**: `/getSecurityCode`
  - **请求方式**: `POST`
  - **Content-Type**: `application/json`
  - **需要auth**: `否`
  - **说明**: 获取图形验证码

* Body (请求体) 

| 参数名 | 类型 | 必填 | 描述 |
| - | - | - | - |
| deviceId | string |  | 浏览器 初始化一个 设备ID 可以是 UUID 等，生成同一个 |
| width | int |  | 验证码 宽度 |
| height | int |  | 验证码 高度 |


### 登录 

* 基本信息 
  - **请求URL**: `/login`
  - **请求方式**: `POST`
  - **Content-Type**: `application/json`
  - **需要auth**: `否`
  - **说明**: 登录

* Body (请求体) 

| 参数名 | 类型 | 必填 | 描述 |
| - | - | - | - |
| account | string |  | 登录账号 |
| password | string |  | 登录密码 需要对密码 MD5 传入接口，输入的密码为 MD5(password) |
| deviceId | string |  | 浏览器 初始化一个 设备ID 可以是 UUID 等，生成同一个 |
| securityCodeKey | string |  | 验证码的 临时 key 验证的时候需要传入 |
| securityCode | string |  | 验证码 |
| sourceType | int |  | 登录 来源 1 为 管理页面 |
| sourceInfo | string |  | 登录 来源 信息 如 浏览器信息、客户端ip等 使用 json 格式 |


### 获取登录用户信息 

* 基本信息 
  - **请求URL**: `/session`
  - **请求方式**: `POST`
  - **Content-Type**: `application/json`
  - **需要auth**: `是`
  - **说明**: 获取登录用户信息

* Headers 

| 参数名 | 类型 | 必填 | 描述 |
| - | - | - | - |
| token | string | 是 | 用于 auth 认证，通过 /login 接口获取 |


### 登出 

* 基本信息 
  - **请求URL**: `/logout`
  - **请求方式**: `POST`
  - **Content-Type**: `application/json`
  - **需要auth**: `否`
  - **说明**: 登出

* Headers 

| 参数名 | 类型 | 必填 | 描述 |
| - | - | - | - |
| token | string |  |  |


## /manage/log/ 

### /manage/log/page 

* 基本信息 
  - **请求URL**: `/manage/log/page`
  - **请求方式**: `POST`
  - **Content-Type**: `application/json`
  - **需要auth**: `是`
  - **说明**: 

* Headers 

| 参数名 | 类型 | 必填 | 描述 |
| - | - | - | - |
| token | string | 是 | 用于 auth 认证，通过 /login 接口获取 |

* Body (请求体) 

| 参数名 | 类型 | 必填 | 描述 |
| - | - | - | - |
| logId | int64 |  | 日志id |
| loginId | int64 |  | 登录id |
| userId | int64 |  | 用户id |
| userName | string |  | 用户名称 |
| userAccount | string |  | 用户账号 |
| ip | string |  | ip |
| path | string |  | 请求 |
| comment | string |  | 备注 |
| pageNo | int64 |  | 页码 |
| pageSize | int64 |  | 每页数量 |


### /manage/log/delete 

* 基本信息 
  - **请求URL**: `/manage/log/delete`
  - **请求方式**: `POST`
  - **Content-Type**: `application/json`
  - **需要auth**: `是`
  - **说明**: 

* Headers 

| 参数名 | 类型 | 必填 | 描述 |
| - | - | - | - |
| token | string | 是 | 用于 auth 认证，通过 /login 接口获取 |

* Body (请求体) 

| 参数名 | 类型 | 必填 | 描述 |
| - | - | - | - |
| logIds | []int64 |  |  |


## /manage/login/ 

### /manage/login/page 

* 基本信息 
  - **请求URL**: `/manage/login/page`
  - **请求方式**: `POST`
  - **Content-Type**: `application/json`
  - **需要auth**: `是`
  - **说明**: 

* Headers 

| 参数名 | 类型 | 必填 | 描述 |
| - | - | - | - |
| token | string | 是 | 用于 auth 认证，通过 /login 接口获取 |

* Body (请求体) 

| 参数名 | 类型 | 必填 | 描述 |
| - | - | - | - |
| loginId | int64 |  | 登录id |
| userId | int64 |  | 用户ID |
| userAccount | string |  | 用户 账号 |
| userName | string |  | 用户 名称 |
| loginIp | string |  | 登录 ip |
| sourceType | int |  | 来源 类型 |
| pageNo | int64 |  | 页码 |
| pageSize | int64 |  | 每页数量 |


### /manage/login/delete 

* 基本信息 
  - **请求URL**: `/manage/login/delete`
  - **请求方式**: `POST`
  - **Content-Type**: `application/json`
  - **需要auth**: `是`
  - **说明**: 

* Headers 

| 参数名 | 类型 | 必填 | 描述 |
| - | - | - | - |
| token | string | 是 | 用于 auth 认证，通过 /login 接口获取 |

* Body (请求体) 

| 参数名 | 类型 | 必填 | 描述 |
| - | - | - | - |
| loginIds | []int64 |  |  |


## /manage/permission/ 

### /manage/permission/add 

* 基本信息 
  - **请求URL**: `/manage/permission/add`
  - **请求方式**: `POST`
  - **Content-Type**: `application/json`
  - **需要auth**: `是`
  - **说明**: 

* Headers 

| 参数名 | 类型 | 必填 | 描述 |
| - | - | - | - |
| token | string | 是 | 用于 auth 认证，通过 /login 接口获取 |

* Body (请求体) 

| 参数名 | 类型 | 必填 | 描述 |
| - | - | - | - |
| roleId | int64 |  | 角色 ID 给角色授权 |
| userId | int64 |  | 用户 ID 给用户授权 |
| permissionType | string |  | 权限类型 path: 路径 button: 按钮 tag: 标签 page: 页面 |
| permission | string |  | 权限 根据类型设置 如：/admin/user/add |
| authorizable | int |  | 可授权 2:不可授权 1:可以授权 |
| startAt | int64 |  | 权限开始 时间戳 毫秒 |
| expiredAt | int64 |  | 过期 时间戳 毫秒 |
| expiredDuration | int64 |  | 过期时间 分钟 |
| updateAt | int64 |  | 修改 时间戳 毫秒 |


### /manage/permission/query 

* 基本信息 
  - **请求URL**: `/manage/permission/query`
  - **请求方式**: `POST`
  - **Content-Type**: `application/json`
  - **需要auth**: `是`
  - **说明**: 

* Headers 

| 参数名 | 类型 | 必填 | 描述 |
| - | - | - | - |
| token | string | 是 | 用于 auth 认证，通过 /login 接口获取 |

* Body (请求体) 

| 参数名 | 类型 | 必填 | 描述 |
| - | - | - | - |
| permissionIds | []int64 |  |  |
| roleIds | []int64 |  |  |
| userIds | []int64 |  |  |


### /manage/permission/queryByIds 

* 基本信息 
  - **请求URL**: `/manage/permission/queryByIds`
  - **请求方式**: `POST`
  - **Content-Type**: `application/json`
  - **需要auth**: `是`
  - **说明**: 

* Headers 

| 参数名 | 类型 | 必填 | 描述 |
| - | - | - | - |
| token | string | 是 | 用于 auth 认证，通过 /login 接口获取 |

* Body (请求体) 

| 参数名 | 类型 | 必填 | 描述 |
| - | - | - | - |
| ids | []int64 |  |  |


### /manage/permission/queryByRoleIds 

* 基本信息 
  - **请求URL**: `/manage/permission/queryByRoleIds`
  - **请求方式**: `POST`
  - **Content-Type**: `application/json`
  - **需要auth**: `是`
  - **说明**: 

* Headers 

| 参数名 | 类型 | 必填 | 描述 |
| - | - | - | - |
| token | string | 是 | 用于 auth 认证，通过 /login 接口获取 |

* Body (请求体) 

| 参数名 | 类型 | 必填 | 描述 |
| - | - | - | - |
| roleIds | []int64 |  |  |


### /manage/permission/queryByUserIds 

* 基本信息 
  - **请求URL**: `/manage/permission/queryByUserIds`
  - **请求方式**: `POST`
  - **Content-Type**: `application/json`
  - **需要auth**: `是`
  - **说明**: 

* Headers 

| 参数名 | 类型 | 必填 | 描述 |
| - | - | - | - |
| token | string | 是 | 用于 auth 认证，通过 /login 接口获取 |

* Body (请求体) 

| 参数名 | 类型 | 必填 | 描述 |
| - | - | - | - |
| userIds | []int64 |  |  |


### /manage/permission/delete 

* 基本信息 
  - **请求URL**: `/manage/permission/delete`
  - **请求方式**: `POST`
  - **Content-Type**: `application/json`
  - **需要auth**: `是`
  - **说明**: 

* Headers 

| 参数名 | 类型 | 必填 | 描述 |
| - | - | - | - |
| token | string | 是 | 用于 auth 认证，通过 /login 接口获取 |

* Body (请求体) 

| 参数名 | 类型 | 必填 | 描述 |
| - | - | - | - |
| permissionIds | []int64 |  |  |
| roleIds | []int64 |  |  |
| userIds | []int64 |  |  |


## /manage/role/ 

### /manage/role/add 

* 基本信息 
  - **请求URL**: `/manage/role/add`
  - **请求方式**: `POST`
  - **Content-Type**: `application/json`
  - **需要auth**: `是`
  - **说明**: 

* Headers 

| 参数名 | 类型 | 必填 | 描述 |
| - | - | - | - |
| token | string | 是 | 用于 auth 认证，通过 /login 接口获取 |

* Body (请求体) 

| 参数名 | 类型 | 必填 | 描述 |
| - | - | - | - |
| name | string |  | 名称 |
| isSuper | int8 |  | 角色 是否 是超管 如果是 则拥有所有权限 1是 2否 |
| updateAt | int64 |  | 修改 时间戳 毫秒 |


### /manage/role/list 

* 基本信息 
  - **请求URL**: `/manage/role/list`
  - **请求方式**: `POST`
  - **Content-Type**: `application/json`
  - **需要auth**: `是`
  - **说明**: 

* Headers 

| 参数名 | 类型 | 必填 | 描述 |
| - | - | - | - |
| token | string | 是 | 用于 auth 认证，通过 /login 接口获取 |

* Body (请求体) 

| 参数名 | 类型 | 必填 | 描述 |
| - | - | - | - |
| roleId | int64 |  | 角色 ID |
| name | string |  | 名称 |
| isSuper | int8 |  | 角色 是否 是超管 如果是 则拥有所有权限 1是 2否 |
| createAt | int64 |  | 创建 时间戳 毫秒 |
| updateAt | int64 |  | 修改 时间戳 毫秒 |


### /manage/role/delete 

* 基本信息 
  - **请求URL**: `/manage/role/delete`
  - **请求方式**: `POST`
  - **Content-Type**: `application/json`
  - **需要auth**: `是`
  - **说明**: 

* Headers 

| 参数名 | 类型 | 必填 | 描述 |
| - | - | - | - |
| token | string | 是 | 用于 auth 认证，通过 /login 接口获取 |

* Body (请求体) 

| 参数名 | 类型 | 必填 | 描述 |
| - | - | - | - |
| roleIds | []int64 |  |  |


### /manage/role/getUserRoles 

* 基本信息 
  - **请求URL**: `/manage/role/getUserRoles`
  - **请求方式**: `POST`
  - **Content-Type**: `application/json`
  - **需要auth**: `是`
  - **说明**: 

* Headers 

| 参数名 | 类型 | 必填 | 描述 |
| - | - | - | - |
| token | string | 是 | 用于 auth 认证，通过 /login 接口获取 |

* Body (请求体) 

| 参数名 | 类型 | 必填 | 描述 |
| - | - | - | - |
| userId | int64 |  |  |


### /manage/role/addRoleUsers 

* 基本信息 
  - **请求URL**: `/manage/role/addRoleUsers`
  - **请求方式**: `POST`
  - **Content-Type**: `application/json`
  - **需要auth**: `是`
  - **说明**: 

* Headers 

| 参数名 | 类型 | 必填 | 描述 |
| - | - | - | - |
| token | string | 是 | 用于 auth 认证，通过 /login 接口获取 |

* Body (请求体) 

| 参数名 | 类型 | 必填 | 描述 |
| - | - | - | - |
| roleId | int64 |  |  |
| userIds | []int64 |  |  |


### /manage/role/addRoleUser 

* 基本信息 
  - **请求URL**: `/manage/role/addRoleUser`
  - **请求方式**: `POST`
  - **Content-Type**: `application/json`
  - **需要auth**: `是`
  - **说明**: 

* Headers 

| 参数名 | 类型 | 必填 | 描述 |
| - | - | - | - |
| token | string | 是 | 用于 auth 认证，通过 /login 接口获取 |

* Body (请求体) 

| 参数名 | 类型 | 必填 | 描述 |
| - | - | - | - |
| roleId | int64 |  |  |
| userId | int64 |  |  |


## /manage/user/ 

### /manage/user/add 

* 基本信息 
  - **请求URL**: `/manage/user/add`
  - **请求方式**: `POST`
  - **Content-Type**: `application/json`
  - **需要auth**: `是`
  - **说明**: 

* Headers 

| 参数名 | 类型 | 必填 | 描述 |
| - | - | - | - |
| token | string | 是 | 用于 auth 认证，通过 /login 接口获取 |

* Body (请求体) 

| 参数名 | 类型 | 必填 | 描述 |
| - | - | - | - |
| name | string |  | 名称 |
| account | string |  | 登录账号 |
| password | string |  |  |


### /manage/user/list 

* 基本信息 
  - **请求URL**: `/manage/user/list`
  - **请求方式**: `POST`
  - **Content-Type**: `application/json`
  - **需要auth**: `是`
  - **说明**: 

* Headers 

| 参数名 | 类型 | 必填 | 描述 |
| - | - | - | - |
| token | string | 是 | 用于 auth 认证，通过 /login 接口获取 |

* Body (请求体) 

| 参数名 | 类型 | 必填 | 描述 |
| - | - | - | - |
| userId | int64 |  | 用户 ID |
| name | string |  | 名称 |
| account | string |  | 登录账号 |
| status | int |  | 状态 1：正常 2：禁用 9：删除 |
| createAt | int64 |  | 创建 时间戳 毫秒 |
| updateAt | int64 |  | 修改 时间戳 毫秒 |
| deleteAt | int64 |  | 删除 时间戳 毫秒 |
| disableAt | int64 |  | 禁用 时间戳 毫秒 |


### /manage/user/page 

* 基本信息 
  - **请求URL**: `/manage/user/page`
  - **请求方式**: `POST`
  - **Content-Type**: `application/json`
  - **需要auth**: `是`
  - **说明**: 

* Headers 

| 参数名 | 类型 | 必填 | 描述 |
| - | - | - | - |
| token | string | 是 | 用于 auth 认证，通过 /login 接口获取 |

* Body (请求体) 

| 参数名 | 类型 | 必填 | 描述 |
| - | - | - | - |
| userId | int64 |  | 用户 ID |
| name | string |  | 名称 |
| account | string |  | 登录账号 |
| status | int |  | 状态 1：正常 2：禁用 9：删除 |
| createAt | int64 |  | 创建 时间戳 毫秒 |
| updateAt | int64 |  | 修改 时间戳 毫秒 |
| deleteAt | int64 |  | 删除 时间戳 毫秒 |
| disableAt | int64 |  | 禁用 时间戳 毫秒 |
| pageNo | int64 |  | 页码 |
| pageSize | int64 |  | 每页数量 |


### /manage/user/delete 

* 基本信息 
  - **请求URL**: `/manage/user/delete`
  - **请求方式**: `POST`
  - **Content-Type**: `application/json`
  - **需要auth**: `是`
  - **说明**: 

* Headers 

| 参数名 | 类型 | 必填 | 描述 |
| - | - | - | - |
| token | string | 是 | 用于 auth 认证，通过 /login 接口获取 |

* Body (请求体) 

| 参数名 | 类型 | 必填 | 描述 |
| - | - | - | - |
| userIds | []int64 |  |  |


### /manage/user/remove 

* 基本信息 
  - **请求URL**: `/manage/user/remove`
  - **请求方式**: `POST`
  - **Content-Type**: `application/json`
  - **需要auth**: `是`
  - **说明**: 

* Headers 

| 参数名 | 类型 | 必填 | 描述 |
| - | - | - | - |
| token | string | 是 | 用于 auth 认证，通过 /login 接口获取 |

* Body (请求体) 

| 参数名 | 类型 | 必填 | 描述 |
| - | - | - | - |
| userIds | []int64 |  |  |


