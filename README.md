# 猫猫超市后端
前端：[github.com/SuperH-0630/cat-shop-front](https://github.com/SuperH-0630/cat-shop-front)

## 配置文件
```yaml
mysql:
  username: # mysql用户名
  password: # mysql用户密码
  address: # mysql地址
  post: # mysql端口号
  dbname: # mysql数据库名称

file:
  localpath: # 文件上传时保存的地理位置（为一个文件夹地址，可不存在）

http:
  address: # http监听地址 默认：localhost:2689
  debugmsg: # api是否返回调试信息
  baseapi: # api前缀 默认：/api （为空则使用默认）
  testapi: # 是否开启测试api

jwt:
  secretpath: # jwt密钥保存地址（为一个文件地址，可不存在）
  hour: 5 # jwt令牌有效期（小时）
  resetmin: 30  # jwt令牌重置倒计时，当令牌距离过期时间短于此设定时将自动更新令牌（分钟）
```

## 运行
参数：
```
-help 查看帮助详情（打印帮助信息，服务不会运行）
-config 配置文件信息（默认：config.yaml）
```

### 测试运行
可以通过`go run`直接运行项目。

```shell
# 显示把昂住信息
go run github.com/SuperH-0630/cat-shop-back/src/cmd/v1 -help

# 运行服务并指定配置文件
go run github.com/SuperH-0630/cat-shop-back/src/cmd/v1 -config ./etc/config.yaml
```

### 实际运行
构建成可执行程序后可实际运行。构建请参考下文。若可执行文件为`./shop.exe`，则运行方式为：

```shell
# 显示把昂住信息
./shop.exe -help

# 运行服务并指定配置文件
./shop.exe -config ./etc/config.yaml
```

## 构建
可以使用`go build`进行构建。

```shell
go build github.com/SuperH-0630/cat-shop-back/src/cmd/v1
```
