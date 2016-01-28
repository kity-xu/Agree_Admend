//
//  OtherWorkViewController.m
//  iMateTest
//
//  Created by hxsmart on 14-7-23.
//  Copyright (c) 2014年 hxsmart. All rights reserved.
//

#import "OtherWorkViewController.h"
#import <HXiMateSDK/iMateAppFace.h>

@interface OtherWorkViewController () <iMateAppFaceDelegate>

@property (strong, nonatomic) iMateAppFace *imateAppFace;

@end

@implementation OtherWorkViewController

- (id)initWithNibName:(NSString *)nibNameOrNil bundle:(NSBundle *)nibBundleOrNil
{
    self = [super initWithNibName:nibNameOrNil bundle:nibBundleOrNil];
    if (self) {
        // Custom initialization
    }
    return self;
}

- (void)viewDidLoad
{
    [super viewDidLoad];
    
    self.titleLabel.text = @"其它测试";

    //获取iMateAppFace实例
    _imateAppFace = [iMateAppFace sharedController];
    _imateAppFace.delegate = self;
    
    [self showLogMessage:@"其它测试项目："];
    [self showLogMessage:@"1、iMate固件版本号"];
    [self showLogMessage:@"2、iMate设备系列号"];
    [self showLogMessage:@"3、取iMate电池电量\n"];
    
    [self showLogMessage:[NSString stringWithFormat:@"iMate固件版本:%@",[self.imateAppFace deviceVersion]]];
    [self showLogMessage:[NSString stringWithFormat:@"iMate设备系列号:%@",[self.imateAppFace deviceSerialNumber]]];

    // 电池电量测试
    [_imateAppFace batteryLevel];
}

#pragma mark iMateAppFaceDelegate

- (void)iMateDelegateConnectStatus:(BOOL)isConnecting
{
    if ( isConnecting ) {
        [self showLogMessage:@"iMateDelegateConnectStatus : iMate连接成功!"];
    }
    else {
        [self showLogMessage:@"iMateDelegateConnectStatus : iMate连接断开!"];
    }
}

- (void)iMateDelegateNoResponse:(NSString *)error
{
    [self showLogMessage:[NSString stringWithFormat:@"iMateDelegateNoResponse : %@", error]];
}

- (void)iMateDelegateResponsePackError
{
    [self showLogMessage:@"iMateDelegateResponsePackError : 应答报文错误"];
}

- (void)iMateDelegateBatteryLevel:(NSInteger)returnCode level:(NSInteger)level error:(NSString *)error
{
    NSString *str = [NSString stringWithFormat:@"\nbatteryLevel返回码:%d",(int)returnCode];
    [self showLogMessage:str];
    
    if ( !returnCode ) {
        str = [NSString stringWithFormat:@"level : %ld%%",(long)level];
        [self showLogMessage:str];
    }
    else {
        str = [NSString stringWithFormat:@"错误信息:%@",error];
        [self showLogMessage:str];
    }
}

@end
