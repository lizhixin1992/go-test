# golang + iris

在使用go.mod时，出现无法下载相关的扩展包，这需要设置代理，go官方代理：设置环境变量：GOPROXY="https://goproxy.io"

我出现无法启动的情况，之后切换不使用go.mod再切换使用go.mod后就可以了。
