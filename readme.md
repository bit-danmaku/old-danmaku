# 使用方法

```bash
cd server

# 构建二进制
make build

# 运行二进制
./old-danmaku
```

使用 HTTPie 请求接口

```
http POST :8080/channel/12 text=hello time:=12.0 author=tes
```
