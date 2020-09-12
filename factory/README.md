# 工厂设计模式



## 简单工厂模式 simple

比如案例中使用消息 Alert 作为接口，规定需要实现的方法 Send()  
- Mail/SMS/Wechat 实现 Alert 接口，编写自己 Send() 逻辑

通过 `NewAlert(type) Alert` 传入的 **参数不同** ，生产出不同的对象，以接口类型的方式返回（多态）

缺点：如果新增加产品，需要更改 `NewAlert(type) Alert` 中的代码，比如已经添加的 `SMS` ;违反了 设计模式六大原则中的`开闭原则`，

 `开闭原则` ：软件对象（类、模块、方法等）应该对于扩展是开放的，对修改是关闭的 



## 工厂模式 factory

工厂方法模式的实质是“定义一个创建对象的接口，但让 **实现这个接口的类来决定实例化哪个类** 。工厂方法让类的实例化推迟到子类中进行。”
比如案例中， `AlertFactory` `Alert` 两个接口  
`AlertFactory`接口定义 `Create() Alert` 方法，返回的是 Alert 接口类型  
产品则需要分别创建两个 struct，一个实现  `AlertFactory`  一个实现 `Alert` ，
实例的初始化交给 FactoryStruct 实现，具体实例的方法交给 AlertStruct实现  

如果需要新增实例，只需要新增文件并实现上述的两个 interface 即可  

调用处，只不过是 new(Factory1) 或者是 new(Factory2) 来进行控制，实例的创建下发给各个工厂struct去实现  

缺点：每次增加一个产品需要增具体struct，导致系统类数目过多，复杂性增加



## 抽象工厂模式 abstract 

通过工厂模式我们可以看出，工厂与产品都是一对一的模式  
抽象工厂模式则打破了壁垒，扩展了工厂的功能，使其可以 **生产多个大类**
如果只是生产一种大类，那么跟`工厂模式`区别不大, 但是如果对大类进行扩展，势必会修改抽象工厂与具体工厂的代码，违反了`开闭原则`



案例中：以 DAO 层对 `OrderMain` 和 `OrderDetail` 两个表的操作作为案例，对于 MYSQL Oracle 等数据库的操作作为扩展

其中的工厂与产品对应关系：

- `DAOFactory` 对应两个产品  `OrderMain` 与 `OrderDetail`
- `MySQLFactory`  负责工厂的具体对象实例化，通过两个函数分别返回**两个产品的实例**
- `MySQLOrderMainDAO` 和 `MySQLOrderDetailDAO` 分别对应具体产品的功能实现



可以看出，抽象模式工厂不仅可以生产 `OrderMain` 产品， 也可以生产 `OrderDetail` 产品，相对于上述工厂模式的案例中只能生产 `Alert` 类产品有了很大的扩展性。但是同样的，如果抽象工厂模式需要增加产品，比如 `OrderHistory` 则需要更改代码，违反了 `开闭原则`



## 总结

三种工厂模式最终目的都是为了得到指定的对象，按需使用，能达到解耦目的就是好模式
