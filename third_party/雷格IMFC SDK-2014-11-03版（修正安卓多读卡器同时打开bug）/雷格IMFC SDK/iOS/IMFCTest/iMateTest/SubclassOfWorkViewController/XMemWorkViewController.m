//
//  XMemWorkViewController.m
//  iMateTest
//
//  Created by hxsmart on 14-7-23.
//  Copyright (c) 2014年 hxsmart. All rights reserved.
//

#import "XMemWorkViewController.h"
#import <HXiMateSDK/iMateAppFace.h>

static Byte *testBytes = (Byte*)"\x1\x2\x3\x4\x5\x6\x7\x8\x9\x0\x1\x2\x3\x4\x5\x6";

@interface XMemWorkViewController () <iMateAppFaceDelegate>

@property (strong, nonatomic) iMateAppFace *imateAppFace;

@end

@implementation XMemWorkViewController

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
    
    self.titleLabel.text = @"扩展内存测试";

    //获取iMateAppFace实例
    _imateAppFace = [iMateAppFace sharedController];
    _imateAppFace.delegate = self;
    
    [_imateAppFace xmemWrite:0 data:[NSData dataWithBytes:testBytes length:16]];
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

- (void)iMateDelegateXmemWrite:(NSInteger)returnCode error:(NSString *)error
{
    NSString *str = [NSString stringWithFormat:@"xmemWrite返回码:%ld", (long)returnCode];
    [self showLogMessage:str];
    
    if ( returnCode ) {
        str = [NSString stringWithFormat:@"错误信息:%@",error];
        [self showLogMessage:str];
    }
    else {
        [_imateAppFace xmemRead:0 length:16];
    }
}

- (void)iMateDelegateXmemRead:(NSInteger)returnCode data:(NSData *)data error:(NSString *)error
{
    NSString *str = [NSString stringWithFormat:@"xmemRead返回码:%ld", (long)returnCode];
    [self showLogMessage:str];
    
    if ( !returnCode ) {
        str = [NSString stringWithFormat:@"read data:%@",[iMateAppFace oneTwoData:data]];
        [self showLogMessage:str];
        
        if (memcmp(data.bytes, testBytes, 16) == 0)
            [self showLogMessage:@"数据比较正确"];
        else
            [self showLogMessage:@"数据比较错误"];

    }
    else {
        str = [NSString stringWithFormat:@"错误信息:%@",error];
        [self showLogMessage:str];
    }
}

@end
