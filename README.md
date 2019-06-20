# golang + iris + go.mod

最开始没有使用go.mod

之后使用go.mod时，出现无法下载相关的扩展包，这需要设置代理，go官方代理：设置环境变量：GOPROXY="https://goproxy.io"

我出现无法启动的情况，之后切换不使用go.mod再切换使用go.mod后就可以了。

主要测试一些工具，比如redis，elastic等，编写对应的工具类
