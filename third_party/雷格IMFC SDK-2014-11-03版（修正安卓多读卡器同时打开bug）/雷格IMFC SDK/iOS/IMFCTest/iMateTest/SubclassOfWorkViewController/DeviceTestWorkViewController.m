//
//  DeviceTestWorkViewController.m
//  iMateTest
//
//  Created by hxsmart on 14-7-23.
//  Copyright (c) 2014年 hxsmart. All rights reserved.
//

#import "DeviceTestWorkViewController.h"
#import <HXiMateSDK/iMateAppFace.h>

@interface DeviceTestWorkViewController () <iMateAppFaceDelegate>

@property (strong, nonatomic) iMateAppFace *imateAppFace;

@end

@implementation DeviceTestWorkViewController

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
    
    self.titleLabel.text = @"部件检测接口测试";

    //获取iMateAppFace实例
    _imateAppFace = [iMateAppFace sharedController];
    _imateAppFace.delegate = self;

    // 部件检测。可检测的部件包括二代证模块，射频卡模块。（IMFC还包括指纹模块、SD模块）
    // componentsMask的bit来标识检测的部件：
    //      0x01 二代证模块
    //      0x02 射频模块
    //      0x40 IMFC 指纹模块（iMate不支持）
    //      0x80 IMFC SD卡模块（iMate不支持）
    //      0xFF 全部部件检测
    // 检测的结果通过delegate响应
    
    [self showLogMessage:@"正在部件测试..."];
    [_imateAppFace deviceTest:0xff];
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

- (void)iMateDelegateDeviceTest:(NSInteger)returnCode resultMask:(Byte)resultMask error:(NSString *)error
{
    NSString *str = [NSString stringWithFormat:@"deviceTest返回码:%d",(int)returnCode];
    [self showLogMessage:str];
    
    if ( !returnCode ) {
        str = [NSString stringWithFormat:@"resultMask:%02x", resultMask];
        [self showLogMessage:str];
        
        if (!resultMask) {
            [self showLogMessage:@"部件测试全部正常"];
        }
        
        if (resultMask & 0x01) {
           [self showLogMessage:@"二代证模块故障或不存在"];
        }
        if (resultMask & 0x02) {
            [self showLogMessage:@"射频卡模块故障或不存在"];
        }
        if (resultMask & 0x04) {
            [self showLogMessage:@"IMFC指纹模块故障或不存在"];
        }
        if (resultMask & 0x08) {
            [self showLogMessage:@"IMFC SD卡故障或不存在"];
        }
    }
    else {
        str = [NSString stringWithFormat:@"错误信息:%@",error];
        [self showLogMessage:str];
    }
}

@end
