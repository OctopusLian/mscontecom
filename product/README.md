<!--
 * @Description: 
 * @Author: neozhang
 * @Date: 2022-02-10 13:11:06
 * @LastEditors: neozhang
 * @LastEditTime: 2022-03-07 22:17:18
-->
# Product service 
当前服务 名称为 Product 类型 service 

创建初始化模版请使用

```
sudo docker run --rm -v $(pwd): $(pwd) -w  $(pwd) -e cap-protoc -I ./ --micro_out=./ --go_out=./ ./proto/product/product.proto
```

## 快速开始

- [Product service](#product-service)
  - [快速开始](#快速开始)
  - [配置信息](#配置信息)
  - [使用](#使用)

## 配置信息

- 服务名称: go.micro.service.product
- 类型: service
- 简称: product

## 使用
根据 proto 自动生成
```
make proto
```

编译
```
make proto
```

构建镜像
```
make docker
```