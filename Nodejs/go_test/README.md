## 交易风险监控

#### 介绍

交易过程中的实时监控风险是自动化交易中非常重要的组成部分. 你需要设计一个当风险累积达到阈值时发出警报的程序. 我们会提供一些初始的框架代码, 包含了一系列交易和风险参数(累计数量和累积增量)

代码设计需要保证可读, 可扩展和可维护的, 当然, 正确和高效是必要条件. 也许你无法在指定的时间内完全解决此问题, 在这种情况下, 应该尝试给出最低限度可行的代码方案

提交的解决方案, 请描述已实现的功能以及任何需要的功能和尚未实现的功能


#### 问题描述

##### 你需要监控两类的风险指标, 累计成交量和累计Delta

累计成交量是指在一定时间内累计成交的某个合约的量, 它是严格递增的. 举例说明, 当你BUY 1 BTCUSDT然后SELL 1 BTCUSDT, 那么累计成交量为2

累计Delta是指在一定时间内累计成交的某个合约的金额, BUY和SELL产生的金额会相互抵消


##### 三个主要的实体Symbol, Group和Exchange有如下规定

- 一个Exchange可以包含多个Symbol，如exchangeA包含BTCUSDT与ETHUSDT
- 一个Symbol可以出现在不同Exchange，如BTCUSDT同时出现在exchangeA与exchangeB
- 特定Exchange的Symbol至多出现在一个Group内


##### 交易回报是指在BUY和SELL交易产生后, 交易所返回的信息, 我们简化为下列结构

- timestamp：时间毫秒数，严格递增
- exchange：交易所名称，exchangeA、exchangeB...
- symbol：合约名字，BTCUSDT、ETHUSDT...
- side：交易方向，BUY / SELL
- price：价格，严格大于0
- quantity：数量，严格大于0


##### 监控结构描述

- id：监控的唯一标识
- type：监控的类型，symbol / exchange / group
- name：监控的名称，BTCUSDT、exchangeA、groupA...
- quantityInterval：监控累计成交量的时间间隔，毫秒数
- quantityLimit：累计成交量的告警阈值
- deltaInterval：监控累计成交额Delta的时间间隔，毫秒数
- deltaLimit：累计成交额Delta的告警阈值

如以下监控参数：

```json
{
    type: 'symbol',
    name: 'BTCUSDT',
    quantityInterval: 3600000, // 1h
    quantityLimit: 1000,
    ...
}
```

表示监控过去1小时的累计成交量大于等于1000时需发出告警，**但连续满足同一告警条件不重复发出告警**。

> - 若依次接收3笔交易回报t1、t2、t3，如果t1第一次满足告警条件则发出告警，t2依旧满足告警条件但由于与t1重复故不发出告警，t3若满足条件同理不发出告警
>
> - 若依次接收3笔交易回报t1、t2、t3，如果t1第一次满足告警条件则发出告警，t2不满足告警条件不发出告警，t3满足告警条件需要发出告警
>
> - 若依次接收2笔交易回报t1、t2，如果t1满足累计成交量告警条件，t2满足累计成交额Delta告警调整，t2不算连续满足同一条件


#### 项目运行与测试

安装Node.js环境，在代码目录运行以下命令:
``` shell
npm install
npm test
```


