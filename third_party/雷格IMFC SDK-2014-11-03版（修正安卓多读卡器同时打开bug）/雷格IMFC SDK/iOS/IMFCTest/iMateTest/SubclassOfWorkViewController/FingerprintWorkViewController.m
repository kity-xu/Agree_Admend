//
//  FingerprintWorkViewController.m
//  iMateTest
//
//  Created by hxsmart on 14-7-22.
//  Copyright (c) 2014年 hxsmart. All rights reserved.
//

#import "FingerprintWorkViewController.h"
#import <HXiMateSDK/iMateAppFace.h>
#import <HXiMateSDK/iMateDevice.h>

@interface FingerprintWorkViewController () <iMateAppFaceDelegate>

@property (strong, nonatomic) iMateAppFace *imateAppFace;

@end

@implementation FingerprintWorkViewController

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
    
    self.titleLabel.text = @"IMFC指纹模块测试";
    
    //获取iMateAppFace实例
    _imateAppFace = [iMateAppFace sharedController];
    _imateAppFace.delegate = self;
    
    
    [self performSelectorInBackground:@selector(fingerprintTestThread) withObject:nil];
}

#pragma mark iMateAppFaceDelegate

- (void)iMateDelegateConnectStatus:(BOOL)isConnecting
{
    [self hidenGifImageView];
    
    if ( isConnecting ) {
        [self showLogMessage:@"iMateDelegateConnectStatus : iMate连接成功!"];
    }
    else {
        [self showLogMessage:@"iMateDelegateConnectStatus : iMate连接断开!"];
    }
}

- (void)iMateDelegateNoResponse:(NSString *)error
{
    [self hidenGifImageView];
    
    [self showLogMessage:[NSString stringWithFormat:@"iMateDelegateNoResponse : %@", error]];
}

- (void)iMateDelegateResponsePackError
{
    [self hidenGifImageView];
    
    [self showLogMessage:@"iMateDelegateResponsePackError : 应答报文错误"];
}

- (double)currentTimeSeconds
{
    NSTimeInterval time= [[NSDate date] timeIntervalSince1970];
    return (double)time;
}

- (void)fingerprintTestThread
{
    [self showLogTextInMainThread:@"正在打开指纹模块，需要等待2~3秒"];
    uint ret = uiFingerprintOpen();
    if (ret == 99) {
        [self showLogTextInMainThread:@"不支持指纹模块!"];
        return;
    }
    if (ret) {
        [self showLogTextInMainThread:@"打开指纹模块失败"];
        return;
    }
    [self showLogTextInMainThread:@"打开指纹模块成功"];
    
    /* 在uiFingerprintOpen之后就不需要调用
     ret = uiFingerprintLink();
     if (ret) {
     NSLog(@"打开指纹模块未连接");
     return;
     }
     */
    [self showLogTextInMainThread:@"开始采集指纹"];
    vFingerprintSend((unsigned char*)"\x83\x00", 2);
    
    unsigned char buff[512];
    unsigned int len = 0;
    BOOL done = NO;
    double timeSeconds = [self currentTimeSeconds] + 20; //整体20秒等待时间
    while ([self currentTimeSeconds] < timeSeconds) {
        ret = uiFingerprintRecv(buff, &len, 200); //200毫秒延时
        if (ret == 0) {
            if (memcmp(buff, "\x83\x00",2) == 0) {
                [self showLogTextInMainThread:@"采样正确"];
                done = YES;
                break;
            }
            if (memcmp(buff, "\x83\x02",2) == 0) {
                [self showLogTextInMainThread:@"参数不符合定义"];
                break;
            }
            if (memcmp(buff, "\x83\x03",2) == 0) {
                [self showLogTextInMainThread:@"校验和错"];
                vFingerprintSend((unsigned char*)"\x83\x00", 2);
                continue;
            }
            if (memcmp(buff, "\x83\x33",2) == 0) {
                [self showLogTextInMainThread:@"采样错误"];
                vFingerprintSend((unsigned char*)"\x83\x00", 2);
                continue;
            }
            if (memcmp(buff, "\x83\x30",2) == 0) {
                [self showLogTextInMainThread:@"采样超时"];
                vFingerprintSend((unsigned char*)"\x83\x00", 2);
                continue;
            }
            if (memcmp(buff, "\x84\x31",2) == 0) {
                [self showLogTextInMainThread:@"请按下手指"];
            }
            if (memcmp(buff, "\x84\x32",2) == 0) {
                [self showLogTextInMainThread:@"请抬起手指"];
            }
        }
        usleep(10000);
    }
    if (done && len) {
        NSData *fingerprintData = [NSData dataWithBytes:buff + 2 length:len - 2];
        // 返回指纹模板结构，详细参考《TS36EBG 指纹识别模块直接接口开发指南》
        [self showLogTextInMainThread:[NSString stringWithFormat:@"指纹特征特征模板结构数据(%d):%@", len - 2, fingerprintData]];
        [self showLogTextInMainThread:@"指纹采集完成"];
    }
    else
        [self showLogTextInMainThread:@"指纹采集等待超时"];

    vFingerprintClose();
}

@end
