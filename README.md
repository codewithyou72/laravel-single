# laravel-single
一个100%兼容微服务框架go-zero、为PHP语言Laravel框架的PHPer们准备、支持GRPC的单体框架。

### 介绍
为了便于php的laravel开发者上手才基于go-zero开发的集群版本。laravel-single项目是其单机版本，其是该项目是100%遵循go-zero官方标准定制项目。修改内容如下:
1. 修改如下模版，并且修改了对应版本goctl代码：
   - api的handler和main；
   - rpc的main；
   - api和rpc以及model引入model；
   - 数据库配置Redis配置etc/index.yaml、app/internal/config/config.go、app/internal/svc/serviceContext.go

2. 所有生成代码都是基于goctl命令的--style=goZero风格,也可使用--style=go_zero.【强烈不推荐】使用默认--style=gozero风格；
3. 所有时间类型保持同laravel一致使用created_at和updated_at；
4. 路由规则尽量符合laravel的RestFul风格；

### 【开发准备一】编译指定版本goctl，代码地址是laravel-cli
```shell
go build -o laravel-cli.exe goctl.go
mv ./laravel-cli.exe  $GOPATH/bin/laravel-cli.exe

````

### 【开发准备二】laravel-single的mysql数据库，数据库密码见第1步，根据需要修改。
```mysql
DROP TABLE IF EXISTS `demo`;
CREATE TABLE `demo`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL DEFAULT 0 COMMENT '关联user表主键id',
  `picture_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '图片链接[image]',
  `video_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '视频链接[video]',
  `is_display` tinyint(4) NOT NULL DEFAULT 1 COMMENT '(用户端)是否显示[check]：1=不显示、2=显示、0=全部：默认=1',
  `sort` int(11) NOT NULL DEFAULT 0 COMMENT '(降序)排序',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_del` tinyint(4) NOT NULL DEFAULT 1 COMMENT '是否删除:1=否、2=是、0=全部：默认=1',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '创建表格demo' ROW_FORMAT = DYNAMIC;
```
【加索引的sql】
```mysql

CREATE TABLE `demo2` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '关联user表主键id',
  `picture_url` varchar(255) NOT NULL DEFAULT '' COMMENT '图片链接[image]',
  `video_url` varchar(255) NOT NULL DEFAULT '' COMMENT '视频链接[video]',
  `is_display` tinyint(4) NOT NULL DEFAULT '1' COMMENT '(用户端)是否显示[check]：1=不显示、2=显示、0=全部：默认=1',
  `sort` int(11) NOT NULL DEFAULT '0' COMMENT '(降序)排序',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_del` tinyint(4) NOT NULL DEFAULT '1' COMMENT '是否删除:1=否、2=是、0=全部：默认=1',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE COMMENT '用户id唯一索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='创建表格demo';

```

### 【开发准备三】准备redis

### 第1步：修改app/etc/index.yaml的mysql链接信息（加#号的根据需要自行修改）
```yaml
GormCon:
  Path: 127.0.0.1  #
  Port: 3306  #
  Dbname: laravel-single #
  Username: root #
  Password: "123456" #
  MaxIdleConns: 10
  MaxOpenConns: 10
  LogMode: dev
  LogZap: false
  Config: charset=utf8mb4&parseTime=True&loc=Local
  
# Redis缓存配置
RedisCache:
  - Host: 127.0.0.1:6379  #
    Pass: #

```

### 第2步：启动项目 在项目根目录执行这个命令
```shell
go run app/index.go -f app/etc/index.yaml #有效命令、根目录

go run index.go -f etc/index.yaml  #推荐使用有效命令、app目录

#[无效命令] go run index.go -f etc/index.yaml 执行这个命令会导致public目录静态文件无法访问
```
如上下面正常启动

### 第3步：创建新增api，重新生成新接口文件，根目录复制执行如下命令
```shell
laravel-cli api go -api app/routes/web.api -dir app --style=goZero
```
【说明】：
创建api步骤见官方https://go-zero.dev/docs/reference和https://go-zero.dev/docs/tutorials，也可以找博主获取数据库生成.api文件的sql2c工具


# 第4步：生成model命令
```shell
laravel-cli model mysql datasource -url="root:123456@tcp(127.0.0.1:3306)/laravel-single" -table="demo"  -dir="./"  --style=goZero
laravel-cli model mysql datasource -url="root:123456@tcp(127.0.0.1:3306)/laravel-single" -table="demo"  -dir="./"  --style=goZero
goctl1.5.5 model mysql datasource -url="root:123456@tcp(127.0.0.1:3306)/laravel-single" -table="demo"  -dir="./app/model/gormNoCache"  --style=goZero --home="./template/gorm/1.4.2"
goctl1.5.5 model mysql datasource -url="root:123456@tcp(127.0.0.1:3306)/laravel-single" -table="demo2"  -dir="./app/model/gormCache" -cache=true --style=goZero --home="./template/gorm/1.4.2"
goctl1.5.5 model mysql datasource -url="root:123456@tcp(127.0.0.1:3306)/laravel-single" -table="demo2"  -dir="./app/model/Cache" -cache=true --style=goZero
goctl1.5.5 model mysql datasource -url="root:123456@tcp(127.0.0.1:3306)/laravel-single" -table="demo"  -dir="./app/model/NoCache" --style=goZero
```



# 其他功能、在routes文件夹下面使用goctl-swagger插件执行以下命令生成swagger2.0.json文件
```shell
#依赖
#go install github.com/zeromicro/go-zero/tools/goctl@latest
#go install github.com/zeromicro/goctl-swagger@latest
laravel-cli api plugin -p goctl-swagger="swagger -filename swagger2.0.json" --api app/routes/web.api --dir ./public

```
