##软件安装：
1. golang: https://golang.org/
2. MongoDB: https://www.mongodb.org/
3. golang的MongoDB驱动: https://labix.org/mgo
4. beego: http://beego.me/

##运行
1. 运行mongodb server程序
2. 导入数据：`mongoimport --db GraduateDesignDB --collection PlayersColl players.json`
3. 进入src目录，运行脚本：run.sh
4. 在新进入的bash里面执行：`bee run`
5. 浏览器访问：http://ip:8080/web/welcome.html

##运行模拟数据产生程序
1. `go run playerEmulator.go`

#设计部分
服务器开发使用的是GO语言，具体来说是GO语言的一个web开发框架beego。  
在server/src目录下的目录结构是标准的beego应用程序结构。  
主要开发目录在controllers和models目录下，其他目录的作用请参考beego文档。  

##Controllers
在beego中，controller是负责处理一个url请求的类，通过定义GET，POST等方法，即可响应和处理同名的HTTP方法。  
在目前的服务器代码中，包括5个controller：auth，players，rawtranrecord，trainhistory，trianremark。分别作用为：授权、运动员信息相关、提交原始传感器数据、查询已保存在数据库中的以往训练记录、查询本场训练处理过的数据指标。功能下面分别介绍，具体请求参数请看代码注释，及论文附录B。
+  auth  
目前设计了一个简单地授权方案，通过一开始POST管理员账号密码，如果检查通过，则在cookie中设置token。所有除此方法的调用外，都会检查cookie中是否有token，如果没有，将会重定向到auth接口上。但是目前检查cookie这一步禁用了。打开只需要把routers/router.go的init函数第一行去注释即可。
+  players  
GET方法获取运动员信息，POST方法修改/增加运动员，这里有一点需要改进，修改运动员信息应使用PATCH方法，才符合REST风格。
+  rawtranrecord  
POST方法提交原始数据，参数中有一个op参数，值为`append`和`flush`，flush操作应在结束训练时调用，此操作将把内存中的数据写到数据库中。
+  trainhistory  
GET方法获取历史训练记录（历史场次）。
+  trianremark  
GET方法获取本场训练的各项指标数据。

##Models
models目录下主要包括三部分：ORM，运行时内存数据维护类，原始数据处理类。下面将分别介绍其功能。
+  types  
定义了所有的数据类型。
+  dbhelper  
因为beego原生不支持MongoDB的ORM，所以我自己写了一个dbhelper。在models/dbhelper目录下。
+  storage  
内存数据维护类。包括：初始化时从数据库载入运动员信息，token，运行时保存原始数据处理之后的指标数据等。
+  processor  
原始数据处理模块，processor.go定义了processor的接口，通过实现该接口，然后在storage的AppendRawTrainData函数中，调用接口实现类的RawData2Record方法即可，目前实现了一个简单地NaiveProcessor版本，改进的实现只需要在models/processor目录下实现新的实例，然后在storage.go里面使用新的实例即可。

##其他
+  因为前端的现实需要使用JS渲染，所以并未使用beego的模板引擎。
+  tests目录下为对beego的测试代码，编写可以参照已有测例，运行方法：
  +  进入server/src目录
  +  执行脚本run.sh
  +  在新的bash shell里面执行：`go test -bench . -benchmem -parallel 10 tests/`，其中-bench是性能测试选项，其他更多选项请参考GO文档。
+  run.sh  
因为运行时需要先把一些目录添加到GOPATH环境变量中（GOPATH目录在安装golang的时候就会设置，请务必按照官网步骤设置），所以通过脚本先添加，然后开启一个bash子shell，在子shell中的GOPATH环境变量就会包含需要的目录，离开子shell后，这些目录不会保留在系统GOPATH环境变量中。

##前端
+  前端所有代码均在web目录下，此目录是beego配置的静态路径。在router.go的init函数中可以配置。
+  前端图表的绘制使用HighChart库，热力图目前使用的是百度地图提供的SDK。
+  屏幕适配主要是匹配ipad，其他尺寸需要在CSS样式中调整/增加适配。
