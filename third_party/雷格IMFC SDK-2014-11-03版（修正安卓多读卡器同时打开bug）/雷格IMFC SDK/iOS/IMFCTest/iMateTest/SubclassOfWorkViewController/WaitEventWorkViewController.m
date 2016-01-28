//
//  WaitEventWorkViewController.m
//  iMateTest
//
//  Created by hxsmart on 14-7-23.
//  Copyright (c) 2014年 hxsmart. All rights reserved.
//

#import "WaitEventWorkViewController.h"
#import <HXiMateSDK/iMateAppFace.h>

@interface WaitEventWorkViewController () <iMateAppFaceDelegate>

@property (strong, nonatomic) iMateAppFace *imateAppFace;

@end

@implementation WaitEventWorkViewController

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
    
    self.titleLabel.text = @"等待事件测试";

    //获取iMateAppFace实例
    _imateAppFace = [iMateAppFace sharedController];
    _imateAppFace.delegate = self;
    
    [self showLogMessage:@"请刷卡、插卡或放置射频卡..."];

    // 等待事件，包括磁卡刷卡、Pboc IC插入、放置射频卡。timeout是最长等待时间(秒)
    // eventMask的bit来标识检测的部件：
    //      0x01    等待刷卡事件
    //      0x02    等待插卡事件
    //      0x04    等待射频事件
    //      0xFF    等待所有事件
    // 等待的结果通过delegate响应
    [_imateAppFace waitEvent:0xFF timeout:10];
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

- (void)iMateDelegateWaitEvent:(NSInteger)returnCode eventId:(NSInteger)eventId data:(NSData *)data error:(NSString *)error
{
    NSString *str = [NSString stringWithFormat:@"waitEvent返回码:%d",(int)returnCode];
    [self showLogMessage:str];
    
    if ( !returnCode ) {
        str = [NSString stringWithFormat:@"eventId : %ld",(long)eventId];
        [self showLogMessage:str];
        switch (eventId) {
            case 0x01:
                [self showLogMessage:@"检测到刷卡"];
                [self showLogMessage:[NSString stringWithFormat:@"二磁道数据 : %.37s\n三磁道数据 : %s", data.bytes, data.bytes + 37]];
                break;
            case 0x02:
                [self showLogMessage:@"检测到IC卡插卡"];
                [self showLogMessage:[NSString stringWithFormat:@"Data : %@", [iMateAppFace oneTwoData:data]]];
                break;
            case 0x04:
                [self showLogMessage:@"检测到放置射频卡"];
                [self showLogMessage:[NSString stringWithFormat:@"Data : %@", [iMateAppFace oneTwoData:data]]];
                break;
        }
    }
    else {
        str = [NSString stringWithFormat:@"错误信息:%@",error];
        [self showLogMessage:str];
    }
}

@end
