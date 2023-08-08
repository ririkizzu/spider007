# 蜘蛛007-后端

## 项目介绍

蜘蛛007微信小程序，一键查询你的手机号码在哪些App和网站平台上注册过账号

后端服务基于微信云托管部署，好处是部署、开发和维护变得更加便捷，坏处是成本高，更多微信云托管介绍可参考[官方文档](https://developers.weixin.qq.com/miniprogram/dev/wxcloudrun/src/basic/intro.html)

## 部署步骤

### 数据库初始化

1. 新建数据库 `d_spider007`
2. 导入sql到数据库 `mysql -uroot -p d_spider007 < d_spider007.sql`

### 项目初始化

修改配置文件 `conf/config.yaml`

```yaml
Server:
  HttpPort: 80
  ReadTimeout: 60
  WriteTimeout: 60
Database:
  DBType: mysql
  User: root
  Password:
  Host: 127.0.0.1:3306
  DBName: d_spider007
```

### 运行测试

`go run main.go`

1. ping测试
   ![ping](example/ping.jpg)

2. 查看支持查询的平台清单
   ![platform](example/platform.jpg)

3. 一键查询接口
   ![scan](example/scan.jpg)

### 部署到微信云托管

将项目文件压缩打包，部署发布
![deploy](example/deploy.jpg)