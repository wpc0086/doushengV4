# 轻松拿下对不队！

# Goland 设置

打开settings, Go Modules->Enable Go modules integration

# 环境设置

修改consts中的**Addr地址
## 安装minio
vi docker-compose.yml  
```javascript
version: '3'
services:
minio:
image: minio/minio:RELEASE.2022-09-07T22-25-02Z
container_name: minio
ports:
- 9000:9000
- 9001:9001
volumes:
- /var/minio/data:/data
- /var/minio/config:/root/.minio
environment:
MINIO_ACCESS_KEY: "minioadmin"
MINIO_SECRET_KEY: "minioadmin"
command: server /data --console-address ":9001" -address ":9000"
restart: always
```
执行：docker-compose up -d  

http://yourIP:9001/login 进行访问登录 账号密码都是minioadmin，创建video桶后记得public
## ffmpeg安装
https://www.gyan.dev/ffmpeg/builds/

输入ffmpeg -version验证（需要重启）
# 启动

## pulish微服务

cd cmd/publish

sh build.sh

sh output/bootstrap.sh

## interact微服务

cd cmd/interact

sh build.sh

sh output/bootstrap.sh

## user微服务

cd cmd/user

sh build.sh

sh output/bootstrap.sh

## api启动

cd cmd/api

go run main.go