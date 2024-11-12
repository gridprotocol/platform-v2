# platform

## 合约数据同步
合约数据同步代码位于`platform/dumper`目录下。如需要同步合约中数据，需要先实例化dumper并另起一协程执行dumper.SubscribeGRID(ctx)，dumper会每隔10s拉取GRID合约中的事件，解析事件后保存在本地数据库中，同步代码如下：
```go
    chain := "dev"
    registryAddress := common.Address{}
    marketAddress := common.Address{}
    dumper, err := dumper.NewGRIDDumper(chain, registryAddress, marketAddress)
    if err != nil {
        return err
    }

    err = dumper.DumpGRID()
    if err != nil {
        return err
    }

    go dumper.SubscribeGRID(ctx)
```
注：实例化dumper时需要输入三个参数，分别为GRID合约所在的链chain(dev/product/test),registry合约地址以及market合约地址。这三个参数可以直接在代码中写死，也可以从配置文件中读取，也可以在启动platform程序时在命令行中输入。

### 读取本地数据库
本地数据库代码位于`platform/database`目录下。database下实现了一些基本的查询功能。若需要实现更复杂的查询，可以通过gorm库封装好的函数，或者使用`GlobalDataBase.DB()`获取数据库并通过SQL语言查询数据库。