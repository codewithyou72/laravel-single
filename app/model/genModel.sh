#!/usr/bin/env bash

# 使用方法：
# 在[项目根目录]执行如下命令
# ./app/model/genModel.sh 库名 表名 model路径 缓存
# ./app/model/genModel.sh laravel-single user user false  可以简写成./genModel.sh laravel-cli user user  默认不生成缓存信息

#指定 goctl 自定义版本
CC=laravel-cli

#文件风格
style=goZero

#指定数据库名
dbname=$1
#生成的表名
tables=$2

#设置是否生成缓存,默认不生成
defaultCacheBool=false
inputCacheBool=$4
cacheBool=${inputCacheBool:-$defaultCacheBool}

#表生成的 model 目录
#modelDir==$3
modelDir="app/model/mysql"

# 数据库配置
host=127.0.0.1
port=3306
username=root
passwd=123456

echo "开始创建 $dbname 数据库的 $2 表："
mkdir -p "${modelDir}"
echo "下面执行的命令是:"
echo "${CC} model mysql datasource -url=${username}:${passwd}@tcp(${host}:${port})/${dbname} -table=${tables}  -dir=${modelDir} -cache=${cacheBool} --ignore-columns=created_at --ignore-columns=updated_at --style=${style}"
"${CC}" model mysql datasource -url="${username}:${passwd}@tcp(${host}:${port})/${dbname}" -table="${tables}" -dir="${modelDir}" -cache="${cacheBool}"  --ignore-columns="created_at" --ignore-columns="updated_at"  --style="${style}"
