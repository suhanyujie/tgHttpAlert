# 监控服务器
监控服务器的应用进程是否存在，如果不存在，则发送消息到 tg 群中提醒

## build
- 编译到 linux: `GOOS=linux GOARCH=amd64 go build -o tgAlertSender.exe main.go`

## run
- eg: `nohup /home/app/webapps/scripts/tools/tgAlertSender.exe > /dev/null 2>&1 &`
- 相关的辅助 shell

```shell
#!/bin/bash
# 1019 服务监控，如果进程不存在，则发送告警消息到 tg 群
pCount=$(ps axu | grep "web" | grep -v grep | wc -l)
if [ "$pCount" -lt "4" ];then
    echo "$(date +'%Y-%m-%d %H:%M:%S'), web 进程不存在，发送告警消息到 tg 群"
    curl -X POST --location "http://127.0.0.1:9101/tg/alert1" \
        -H "Content-Type: application/json" \
        -d '{
              "env": "prod",
              "serverFlag": "2029",
              "app": "monitor1",
              "msg": "web 相关进程不存在，请检查"
            }'
else
    echo "$(date +'%Y-%m-%d %H:%M:%S') web 相关进程个数为 ${pCount}，相关进程存在，无需重启"
fi
```