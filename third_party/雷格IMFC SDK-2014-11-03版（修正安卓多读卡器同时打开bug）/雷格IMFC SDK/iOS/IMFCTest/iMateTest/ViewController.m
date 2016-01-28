//
//  ViewController.m
//  iMateTest
//
//  Created by hxsmart on 14-7-17.
//  Copyright (c) 2014年 hxsmart. All rights reserved.
//

#import "ViewController.h"
#import <HXiMateSDK/iMateAppFace.h>
#import <HXiMateSDK/iMateAppFace+Pinpad.h>
#import "WorkViewController.h"
#import "UIViewController+MJPopupViewController.h"
#import "HelpViewController.h"

#import "DeviceTestWorkViewController.h"
#import "WaitEventWorkViewController.h"
#import "SwipeCardWorkViewController.h"
#import "ReadIdCardWorkViewController.h"
#import "XMemWorkViewController.h"
#import "PrinterWorkViewController.h"
#import "OtherWorkViewController.h"

#import "ReadPbocInfoWorkViewController.h"
#import "ReadPbocInfoExWorkViewController.h"
#import "ReadPbocLogWorkViewController.h"
#import "PbocIssueWorkViewController.h"

#import "MifareOneWorkViewController.h"

#import "PinpadKMYWorkViewController.h"
#import "PinpadXYDWorkViewController.h"

#import "FingerprintWorkViewController.h"
#import "SdCardWorkViewController.h"

@interface ViewController () <iMateAppFaceDelegate,WorkViewControllerDelegate>

@property (nonatomic, strong) NSArray *testSections;
@property (nonatomic, strong) NSArray *testLists;
@property (strong, nonatomic) iMateAppFace *imateAppFace;
@property (strong, nonatomic) WorkViewController *workViewController;
@property (strong, nonatomic) UILabel *statusLabel;
@end

@implementation ViewController

- (void)viewDidLoad
{
    [super viewDidLoad];
    
    /*
    UIBarButtonItem *helpButton;
    if ([[[[UIDevice currentDevice] systemVersion] substringToIndex:1] intValue] >= 7)
        helpButton = [[UIBarButtonItem alloc] initWithImage:[UIImage imageNamed:@"help"] style:UIBarButtonItemStylePlain target:self action:@selector(helpButtonPressed)];
    else
        helpButton = [[UIBarButtonItem alloc] initWithBarButtonSystemItem:UIBarButtonSystemItemBookmarks target:self action:@selector(helpButtonPressed)];
    
    self.navigationItem.rightBarButtonItem = helpButton;
     */
    
    _testSections = @[@"一般测试", @"Pboc IC卡接口测试", @"射频卡测试", @"指纹模块", @"SD ICC 卡", @"密码键盘"];
    _testLists = @[
            @[@"部件测试", @"等待事件", @"刷卡测试", @"二代证读卡测试", @"扩展内存测试", @"打印机测试", @"其它测试"],
            @[@"读取Pboc卡信息", @"读取Pboc卡信息(扩展)", @"读取Pboc卡日志", @"Pboc个人化演示"],
            @[@"M1卡测试"],
            @[@"指纹模块"],
            @[@"SD ICC 测试"],
            @[@"凯明扬密码键盘", @"信雅达密码键盘"]];
    [self.tableView setSeparatorStyle:UITableViewCellSeparatorStyleSingleLine];
    
    self.statusLabel = [[UILabel alloc] initWithFrame:CGRectMake(0, 0, self.view.bounds.size.width, self.navigationController.toolbar.bounds.size.height)];
    _statusLabel.textAlignment = NSTextAlignmentCenter;
    _statusLabel.textColor = [UIColor grayColor];
    
    if (!IS_IPAD)
        _statusLabel.font = [UIFont systemFontOfSize:12];
    
    [self.navigationController.toolbar addSubview:_statusLabel];
}

- (void)viewWillAppear:(BOOL)animated
{
    [super viewWillAppear:animated];
    
    self.navigationController.navigationBarHidden = NO;
    self.navigationController.toolbarHidden = NO;
    self.title = @"IMFC接口测试";
    
    //获取iMateAppFace实例
    self.imateAppFace = [iMateAppFace sharedController];
    //设置delegate对象
    self.imateAppFace.delegate = self;
    
    //检查iMate的连接情况
    if (![_imateAppFace connectingTest])
        [self showStatus:@"iMate未连接!"];
    else {
        [self showStatus:@"iMate已连接!"];
        //取消密码键盘输入模式
        [_imateAppFace pinPadCancel];
        //取消iMate的工作模式
        [_imateAppFace cancel];
    }
}

- (void)didReceiveMemoryWarning
{
    [super didReceiveMemoryWarning];
}

- (CGFloat)tableView:(UITableView *)atableView heightForHeaderInSection:(NSInteger)section
{
    if (IS_IPAD)
        return 50;
    return 30;
}

- (NSInteger)numberOfSectionsInTableView:(UITableView*) tableView
{
    if (IS_IPAD)
        return [_testSections count];
    return [_testSections count] - 1;
}

- (NSInteger)tableView:(UITableView *)tableView numberOfRowsInSection:(NSInteger)section
{
    return [[_testLists objectAtIndex:section] count];
}

- (CGFloat)tableView:(UITableView *)atableView heightForRowAtIndexPath:(NSIndexPath *)indexPath
{
    if (IS_IPAD)
        return 52.0;
    return 40;
}

- (NSString *)tableView:(UITableView *)tableView titleForHeaderInSection:(NSInteger)section
{
    return [_testSections objectAtIndex:section];
}

- (UITableViewCell *)tableView:(UITableView *)tableView cellForRowAtIndexPath:(NSIndexPath *)indexPath
{
    static NSString *CellIdentifier = @"TestTableViewCell";
    
    UITableViewCell *cell = [tableView dequeueReusableCellWithIdentifier:CellIdentifier];
    if (cell == nil) {
        cell = [[UITableViewCell alloc] initWithStyle:UITableViewCellStyleSubtitle reuseIdentifier:CellIdentifier];
    }

    NSArray *sectionOfLists = [_testLists objectAtIndex:indexPath.section];
    cell.textLabel.text = [sectionOfLists objectAtIndex:indexPath.row];
    cell.textLabel.textColor = [UIColor colorWithRed:26/255.0 green:125/255.0 blue:119/255.0 alpha:1.0];
    
    if (!IS_IPAD)
        cell.textLabel.font = [UIFont systemFontOfSize:14];
    return cell;
}

- (void)showStatus:(NSString *)status
{
    _statusLabel.text = status;
}

#pragma mark - Table view delegate

- (void)tableView:(UITableView *)tableView didSelectRowAtIndexPath:(NSIndexPath *)indexPath
{
    /*
    if (indexPath.section != 5 && ![_imateAppFace connectingTest]) {
        UIAlertView *alert = [[UIAlertView alloc]
                              initWithTitle:nil
                              message:@"iMate未连接，请检查iMate是否开机或是否已经成功对码。"
                              delegate:nil
                              cancelButtonTitle:@"确定"
                              otherButtonTitles:nil];
        [alert show];
        return;
    }
     */
    _workViewController = nil;
    
    if (indexPath.section == 0) {
        switch (indexPath.row) {
            case 0:
                _workViewController = [[DeviceTestWorkViewController alloc] init];
                break;
            case 1:
                _workViewController = [[WaitEventWorkViewController alloc] init];
                break;
            case 2:
                _workViewController = [[SwipeCardWorkViewController alloc] init];
                break;
            case 3:
                _workViewController = [[ReadIdCardWorkViewController alloc] init];
                break;
            case 4:
                _workViewController = [[XMemWorkViewController alloc] init];
                break;
            case 5:
                _workViewController = [[PrinterWorkViewController alloc] init];
                break;
            case 6:
                _workViewController = [[OtherWorkViewController alloc] init];
        }
    }
    if (indexPath.section == 1) {
        switch (indexPath.row) {
            case 0:
                _workViewController = [[ReadPbocInfoWorkViewController alloc] init];
                break;
            case 1:
                _workViewController = [[ReadPbocInfoExWorkViewController alloc] init];
                break;
            case 2:
                _workViewController = [[ReadPbocLogWorkViewController alloc] init];
                break;
            case 3:
                _workViewController = [[PbocIssueWorkViewController alloc] init];
                break;
        }
    }
    if (indexPath.section == 2) {
        switch (indexPath.row) {
            case 0:
                _workViewController = [[MifareOneWorkViewController alloc] init];
                break;
        }
    }
    if (indexPath.section == 3) {
        switch (indexPath.row) {
            case 0:
                _workViewController = [[FingerprintWorkViewController alloc] init];
                break;
        }
    }
    if (indexPath.section == 4) {
        switch (indexPath.row) {
            case 0:
                _workViewController = [[SdCardWorkViewController alloc] init];
                break;
        }
    }
    if (indexPath.section == 5) {
        switch (indexPath.row) {
            case 0:
                _workViewController = [[PinpadKMYWorkViewController alloc] init];
                break;
            case 1:
                _workViewController = [[PinpadXYDWorkViewController alloc] init];
                break;
        }
    }

    if (_workViewController == nil)
        return;
    
    UIBarButtonItem *back = [[UIBarButtonItem alloc] init];
    self.navigationItem.backBarButtonItem = back;
    if ([[[[UIDevice currentDevice] systemVersion] substringToIndex:1] intValue] >= 7)
        back.title = @"";
    else
        back.title = @"返回";
    
    _workViewController.delegate = self;
    if (IS_IPAD)
        [self presentPopupViewController:_workViewController animationType:MJPopupViewAnimationFade];
    else
        [self.navigationController pushViewController:_workViewController animated:YES];
}

- (void)helpButtonPressed
{
    HelpViewController *helpViewController = [[HelpViewController alloc] init];
    
    UIBarButtonItem *back = [[UIBarButtonItem alloc] init];
    self.navigationItem.backBarButtonItem = back;
    if ([[[[UIDevice currentDevice] systemVersion] substringToIndex:1] intValue] >= 7)
        back.title = @"";
    else
        back.title = @"返回";
    [self.navigationController pushViewController:helpViewController animated:YES];
}

#pragma mark iMateAppFaceDelegate

- (void)iMateDelegateConnectStatus:(BOOL)isConnecting
{
    if ( isConnecting )
        [self showStatus:@"iMate连接成功!"];
    else
        [self showStatus:@"iMate连接断开!"];
}

#pragma mark - WorkViewControllerDelegate
- (void)workViewControllerFinish
{
    [self dismissPopupViewControllerWithanimationType:MJPopupViewAnimationFade];

    //取消iMate的工作模式
    [_imateAppFace cancel];
    
    _workViewController = nil;
    
    //设置delegate对象
    self.imateAppFace.delegate = self;
    
    //检查iMate的连接情况
    if (![_imateAppFace connectingTest])
        [self showStatus:@"iMate未连接!"];
    else {
        [self showStatus:@"iMate已连接!"];
        //取消密码键盘输入模式
        [_imateAppFace pinPadCancel];
        //取消iMate的工作模式
        [_imateAppFace cancel];
    }
}

@end
