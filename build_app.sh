#!/bin/sh
# 功能
# 1. 复制一个ucm配置文件
# 2. 替换项目名
# 使用方法
# 新项目下 sh build_app.sh go_app1
app_name=$1
echo "app_name=${app_name}"
if [[ $app_name == "" ]];then
    echo "ERR: app_name不能为空"
    exit 1
fi
if [[ -d "../${app_name}" ]];then
    echo "ERR:项目${app_name}已经存在"
    exit 1
fi

str="mkdir ${app_name}"
echo "exec $str"
$str

str="cp -rf ./ ../${app_name}"
echo "exec $str"
$str

str="cd ${app_name}"
echo "exec $str"
$str

str="cp /root/.bgf/cache/ucm_config_${app_name} /root/.bgf/cache/ucm_config_${app_name}"
echo "exec $str"
$str


sed -i 's/${app_name}/${app_name}/g' `find ./ -type f`

echo "over"
