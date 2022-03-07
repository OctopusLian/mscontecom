<!--
 * @Description: 
 * @Author: neozhang
 * @Date: 2022-02-10 12:57:01
 * @LastEditors: neozhang
 * @LastEditTime: 2022-03-07 22:12:23
-->

# Category service  
Go微服务定制容器
名称为 Category 类型 service 

## Getting Started

- [Category service](#category-service)
  - [Getting Started](#getting-started)
  - [Configuration](#configuration)
  - [Dependencies](#dependencies)
  - [构建](#构建)

## Configuration

- FQDN: go.micro.service.category
- Type: service
- Alias: category

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.


## 构建

```
make build
```

Run the service
```
./category-service
```

Build a docker image
```
make docker
```