# 工厂设计模式

## 简单工厂模式 simple
比如案例中使用消息 Alert 作为接口，规定需要实现的方法 Send()  
- Mail/SMS/Wechat 实现 Alert 接口，编写自己 Send() 逻辑

通过 `NewAlert(type) Alert` 传入的 **参数不同** ，生产出不同的对象，以接口类型的方式返回（多态）

缺点：如果新增加产品，需要更改 `NewAlert(type) Alert` 中的代码，比如已经添加的 `SMS` ;违反了 设计模式六大原则中的`开闭原则`，
PS: `开闭原则` ：软件对象（类、模块、方法等）应该对于扩展是开放的，对修改是关闭的 

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


