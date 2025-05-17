#!/bin/bash

# 提示用户输入项目信息
read -p "请输入项目名称: " PROJECT_NAME
read -p "请输入模块名称: " MODULE_NAME

# 递归查找项目中的所有文件，并替换 your-app-name 为 $PROJECT_NAME
find . -type f -not -path './.git/*' -exec sed -i '' "s/your-app-name/$PROJECT_NAME/g" {} +

# 替换 go.mod 中的模块名
sed -i '' "s/your-module-name/$MODULE_NAME/g" go.mod

# 初始化 Go 模块
go mod tidy

echo "项目初始化完成！"