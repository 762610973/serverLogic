# serverLogic

> Go语言实现原神服务器的业务逻辑模块

## 基础模块

1. 需求分析结构搭建：实现玩家模块，玩家模块目前包含两个模块，基础模块和头像模块
2. 实现头像框修改和测试流程：判断用户有没有这个Icon，没有报错，有就修改
3. 架构微调和名片设置功能：将对外接口都移动到了玩家的主体下，同时增加名片设置功能，和头像设置功能的逻辑一样
4. 名字和签名系统（初遇公共管理模块：名字和签名的更新，对违禁词汇的简单处理，处理分为内部自带的违禁词库和不断更新的外界接口
5. main函数主线程设计前瞻：主要是利用协程
6. 定时器和违禁词库的启用:增加定时器，定时更新词库，这个更新词库的功能应该在服务器启动并且加载配置之后启动
7. 配置表：暂时手动写死一些违禁词汇的配置表，增加`icon,card,banword`，同步到其他函数中，并进行简单测试，梳理了整个测试的运行流程，函数调用情况
8. 增加人物阅历（经验）和等级换算:通过读取等级配置文件，增加等级经验模块，通过经验增加等级,并完成简单的测试
9. 增加任务模块和突破任务的判断；增加唯一突破任务模块，在固定的等级处需要突破相应的任务才能进行升级，理清了整个流程，进行了简单的测试
10. 游戏业务与Go基础的交锋--肆无忌惮使用map,在使用map时一定要注意是否产生读写冲突。玩家在打开背包时，进行了读，此时为了防止发生冲突，线程崩溃，会通过邮件的形式发送物品。还有比如一个排行榜，他是通过自己加锁，让玩家给排行榜模块发送数据，而不是排行榜去读取每个玩家的数据，排行榜读取每个玩家，那么每个玩家都要上锁，但是玩家给排行榜发送数据，只需要排行榜上锁
11. 世界等级降低和还原功能并测试（能灵活配置的切忌不要写成固定逻辑，防止后续重新打包编译），增加了世界等级降级的相关规则判定
12. 生日设置与生日判定
13. 展示名片功能
14. 展示阵容功能


## Tips 
- 对内接口：内部逻辑调用的接口
- 对外接口：和客户端打交道的接口