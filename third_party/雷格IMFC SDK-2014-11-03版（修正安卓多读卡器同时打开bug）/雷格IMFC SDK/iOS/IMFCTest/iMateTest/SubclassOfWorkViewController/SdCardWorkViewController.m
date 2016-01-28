//
//  SdCardWorkViewController.m
//  iMateTest
//
//  Created by hxsmart on 14-7-22.
//  Copyright (c) 2014年 hxsmart. All rights reserved.
//

#import "SdCardWorkViewController.h"
#import <HXiMateSDK/iMateAppFace.h>
#import <HXiMateSDK/iMateDevice.h>

@interface SdCardWorkViewController () <iMateAppFaceDelegate>

@property (strong, nonatomic) iMateAppFace *imateAppFace;

@end

@implementation SdCardWorkViewController

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
    
    self.titleLabel.text = @"SD ICC 测试";
    
    //获取iMateAppFace实例
    _imateAppFace = [iMateAppFace sharedController];
    _imateAppFace.delegate = self;

    //执行SD ICC测试
    [self performSelectorInBackground:@selector(sdIccTestThread) withObject:nil];
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

- (void)sdIccTestThread
{
    uint ret;
    ret = uiSD_Init();
    if (ret == 99) {
        [self showLogTextInMainThread:@"不支持SD ICC"];
        return;
    }
    if (ret) {
        [self showLogTextInMainThread:@"接口初始化失败"];
        return;
    }
    [self showLogTextInMainThread:@"接口初始化成功"];
    
    if (uiSDSCConnectDev()) {
        [self showLogTextInMainThread:@"识别SD_ICC失败"];
        return;
    }
    [self showLogTextInMainThread:@"识别SD_ICC成功"];
    
    unsigned char sVer[50 + 1];
    unsigned int uiLen = 50;
    if (uiSDSCGetFirmwareVer(sVer, &uiLen)) {
        [self showLogTextInMainThread:@"获取SD_ICC固件版本号失败"];
        return;
    }
    
    [self showLogTextInMainThread:@"获取SD_ICC固件版本号成功"];
    uiLen = 50;
    if (uiSDSCGetSDKVersion((char*)sVer, &uiLen)) {
        [self showLogTextInMainThread:@"获取SD_ICC SDK版本号失败"];
        return;
    }
    [self showLogTextInMainThread:[NSString stringWithFormat:@"SDKVersion(%d):%s", uiLen, sVer]];
    
    uint uiSCIOType;
    if (uiSDSCGetSCIOType(&uiSCIOType)) {
        NSLog(@"获取获取SD_ICC IO类型失败");
        return;
    }
    [self showLogTextInMainThread:[NSString stringWithFormat:@"获取SD_ICC IO类型:%d", uiSCIOType]];
    
    [self showLogTextInMainThread:@"uiSDSCTransmit测试"];
    unsigned char apduBuff[512];
    unsigned int uiCosState;
    
    uiLen = sizeof(apduBuff);
    if (uiSDSCTransmit((unsigned char*)"\x00\x84\x00\x00\x08", 5, 0, apduBuff,  &uiLen, &uiCosState)) {
        [self showLogTextInMainThread:@"SD ICC APDU失败"];
        return;
    }
    [self showLogTextInMainThread:[NSString stringWithFormat:@"apdu成功, uiCosState = %04x, uiLen = %d", uiCosState, uiLen]];
    [self showLogTextInMainThread:[NSString stringWithFormat:@"随机数：%02x%02x%02x%02x%02x%02x%02x%02x", apduBuff[0], apduBuff[1],apduBuff[2],apduBuff[3],apduBuff[4],apduBuff[5],apduBuff[6],apduBuff[7]]];
    
    [self showLogTextInMainThread:@"uiSDSCTransmitEx测试"];
    uiLen = sizeof(apduBuff);
    if (uiSDSCTransmitEx((unsigned char*)"\x00\x84\x00\x00\x08", 5, 0, apduBuff,  &uiLen)) {
        [self showLogTextInMainThread:@"SD ICC APDU Ex失败"];
        return;
    }
    [self showLogTextInMainThread:[NSString stringWithFormat:@"CosState:%02x%02x", apduBuff[uiLen-2], apduBuff[uiLen-1]]];
    
    uiLen = sizeof(apduBuff);
    if(uiSDSCResetCard(apduBuff, &uiLen)) {
        [self showLogTextInMainThread:@"SD ICC热复位失败"];
        return;
    }
    if (uiSDSCDisconnectDev()) {
        [self showLogTextInMainThread:@"关闭SD_ICC失败"];
        return;
    }
    vSD_DeInit();
    [self showLogTextInMainThread:@"关闭SD卡电源"];
    
    [self showLogTextInMainThread:@"SD ICC测试成功"];
    
}

@end
