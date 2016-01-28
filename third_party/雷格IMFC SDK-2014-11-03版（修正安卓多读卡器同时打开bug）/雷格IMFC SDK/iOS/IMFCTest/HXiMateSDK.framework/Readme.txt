2014.8.1 关于在多线程中调用syncomm有关的方法，不要使用主线程队列，否则会阻止蓝牙数据的发送。

2013.1.24 更新说明:
1、弃用方法 －initWithDelegate 和 ＋instanceWithDelegate；
2、新增方法如下：
    ＋sharedController	获取iMateAppFace的实例指针，在第一次调用时，实例被创建。
    -openSession    	建立与iMate的蓝牙连接session。该方法只在AppDelegate.m中使用。
    -closeSession   	关闭与iMate的蓝牙连接session，需要App进入后台时调用一次。
3、修改的方法：
    －connectingTest 	该方法用于判断蓝牙连接是否正常，可多次调用。上一个版本只允许调用一次。
    

HXiMateSDK使用说明：

1、初始化(必须）
    在AppDelegate中添加以下代码，用于初始化iMateFaceApp实例，并建立连接Session，改段代码需要添加在didFinishLaunchingWithOptions的顶部。
    - (BOOL)application:(UIApplication *)application didFinishLaunchingWithOptions:(NSDictionary *)launchOptions
    {
        iMateAppFace *iMateFace = [iMateAppFace sharedController];
        [iMateFace openSession];
    
        ...
    
    
2、进入后台时关闭Session(必须）
    在AppDelegate中添加以下代码，该段代码需要添加到applicationDidEnterBackground的底部位置。
    - (void)applicationDidEnterBackground:(UIApplication *)application
    {
        ...

        [[iMateAppFace sharedController] closeSession];
    
3、App从后台进入前台之前，重新建立Session(必须）
    在AppDelegate中添加以下代码，该段代码需要添加到applicationDidEnterBackground的顶部位置。
    - (void)applicationWillEnterForeground:(UIApplication *)application
    {
        [[iMateAppFace sharedController] openSession];
    
        ...
    
4、UI View中使用iMateFaceApp
    a) 在viewWillAppear中添加以下代码，来获取iMateAppFace的指针以及设置delegate。（必须）, 检查session是否正常。（可选）
        - (void)viewWillAppear:(BOOL)animated
        {
            [super viewWillAppear:animated];
            
            //获取iMateAppFace的实例, 该步骤也可以放在viewDidLoad中做
            _imateAppFace = [iMateAppFace sharedController];

            //如果使用delegate协议，需要将delegate设置为self，必须在进入view之前设置。 （重要!!）
            _imateAppFace.delegate = self;

            if ( ![_imateAppFace connectingTest] ) {
                NSLog(@"iMate未连接!");
                return;
            }
            NSLog(@"iMate已连接!");
            ...
    b) 在viewWillDisappear中添加以下代码，来取消尚未响应上次请求。（可选）
    - (void)viewWillDisappear:(BOOL)animated
    {
        [_imateAppFace cancel];

        [super viewWillDisappear:animated];        
        ...
    
    c) 在UI View中实现iMateAppFae的delegate，并调用相关接口。（略）
    
5、关于在App使用过程中重新关闭和打开iMate
    如果在App使用过程中，iMate关闭电源，iMateAppFace将通过iMateDelegateConnectStatus通知session的状态。
    重新打开iMate后，iMateAppFace又将通过iMateDelegateConnectStatus通知session的状态。
    以上两个过程中，App的用户代码不需要做任何操作，iMateAppFace将自动重新建立或关闭session。
