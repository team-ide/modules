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


### 登录 

* 基本信息 
  - **请求URL**: `/login`
  - **请求方式**: `POST`
  - **Content-Type**: `application/json`
  - **需要auth**: `否`
  - **说明**: 登录


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


### 加载登录信息 

* 基本信息 
  - **请求URL**: `/load/login`
  - **请求方式**: `POST`
  - **Content-Type**: `application/json`
  - **需要auth**: `否`
  - **说明**: 加载登录信息


## /manage/log/ 

### /manage/log/manage/log/page 

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


### /manage/log/manage/log/delete 

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


## /manage/login/ 

### /manage/login/manage/login/page 

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


### /manage/login/manage/login/delete 

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


## /manage/permission/ 

### /manage/permission/manage/permission/add 

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


### /manage/permission/manage/permission/query 

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


### /manage/permission/manage/permission/query/by/ids 

* 基本信息 
  - **请求URL**: `/manage/permission/query/by/ids`
  - **请求方式**: `POST`
  - **Content-Type**: `application/json`
  - **需要auth**: `是`
  - **说明**: 

* Headers 

| 参数名 | 类型 | 必填 | 描述 |
| - | - | - | - |
| token | string | 是 | 用于 auth 认证，通过 /login 接口获取 |


### /manage/permission/manage/permission/query/by/role/ids 

* 基本信息 
  - **请求URL**: `/manage/permission/query/by/role/ids`
  - **请求方式**: `POST`
  - **Content-Type**: `application/json`
  - **需要auth**: `是`
  - **说明**: 

* Headers 

| 参数名 | 类型 | 必填 | 描述 |
| - | - | - | - |
| token | string | 是 | 用于 auth 认证，通过 /login 接口获取 |


### /manage/permission/manage/permission/query/by/user/ids 

* 基本信息 
  - **请求URL**: `/manage/permission/query/by/user/ids`
  - **请求方式**: `POST`
  - **Content-Type**: `application/json`
  - **需要auth**: `是`
  - **说明**: 

* Headers 

| 参数名 | 类型 | 必填 | 描述 |
| - | - | - | - |
| token | string | 是 | 用于 auth 认证，通过 /login 接口获取 |


### /manage/permission/manage/permission/delete 

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


## /manage/role/ 

### /manage/role/manage/role/add 

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


### /manage/role/manage/role/list 

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


### /manage/role/manage/role/delete 

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


### /manage/role/manage/role/get/user/roles 

* 基本信息 
  - **请求URL**: `/manage/role/get/user/roles`
  - **请求方式**: `POST`
  - **Content-Type**: `application/json`
  - **需要auth**: `是`
  - **说明**: 

* Headers 

| 参数名 | 类型 | 必填 | 描述 |
| - | - | - | - |
| token | string | 是 | 用于 auth 认证，通过 /login 接口获取 |


### /manage/role/manage/role/add/role/users 

* 基本信息 
  - **请求URL**: `/manage/role/add/role/users`
  - **请求方式**: `POST`
  - **Content-Type**: `application/json`
  - **需要auth**: `是`
  - **说明**: 

* Headers 

| 参数名 | 类型 | 必填 | 描述 |
| - | - | - | - |
| token | string | 是 | 用于 auth 认证，通过 /login 接口获取 |


### /manage/role/manage/role/add/role/user 

* 基本信息 
  - **请求URL**: `/manage/role/add/role/user`
  - **请求方式**: `POST`
  - **Content-Type**: `application/json`
  - **需要auth**: `是`
  - **说明**: 

* Headers 

| 参数名 | 类型 | 必填 | 描述 |
| - | - | - | - |
| token | string | 是 | 用于 auth 认证，通过 /login 接口获取 |


## /manage/user/ 

### /manage/user/manage/user/add 

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


### /manage/user/manage/user/delete 

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


### /manage/user/manage/user/remove 

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


### /manage/user/manage/user/a 

* 基本信息 
  - **请求URL**: `/manage/user/a`
  - **请求方式**: `POST`
  - **Content-Type**: `application/json`
  - **需要auth**: `是`
  - **说明**: 

* Headers 

| 参数名 | 类型 | 必填 | 描述 |
| - | - | - | - |
| token | string | 是 | 用于 auth 认证，通过 /login 接口获取 |


