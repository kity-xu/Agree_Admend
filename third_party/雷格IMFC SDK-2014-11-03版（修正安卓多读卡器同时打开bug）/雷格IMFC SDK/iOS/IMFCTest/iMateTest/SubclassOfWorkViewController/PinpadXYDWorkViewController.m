//
//  PinpadXYDWorkViewController.m
//  iMateTest
//
//  Created by hxsmart on 14-7-22.
//  Copyright (c) 2014年 hxsmart. All rights reserved.
//

#import "PinpadXYDWorkViewController.h"

#import <HXiMateSDK/iMateAppFace.h>
#import <HXiMateSDK/iMateAppFace+Pinpad.h>

static PinPadRequestType doListArray[] = {
    PinPadRequestTypePowerOn,
    PinPadRequestTypeReset,
    PinPadRequestTypeVersion,
    PinPadRequestTypeDownloadMasterKey,
    PinPadRequestTypeDownloadWorkingKey,
    PinPadRequestTypeInputPinBlock,
    PinPadRequestTypeEncrypt,
    PinPadRequestTypeMac,
    PinPadRequestTypePowerOff
};

static Byte masterKey[16] = {0x00, 0x01,0x02,0x03,0x04,0x05,0x06,0x07,0x08,0x09,0x0a,0x0b,0x0c,0x0d,0x0e,0x0f};
static Byte workingKey[16] = {0x12, 0x34,0x56,0x78, 0x90, 0xab,0xcd,0xef,0x01,0x02,0x03,0x04,0x05,0x06,0x07,0x08};
static Byte dataEncrypt[8] = {0,0,0,0,0,0,0,0};
static Byte dataMac[8] = {1,2,3,4,5,6,7,8};

@interface PinpadXYDWorkViewController () <iMateAppFaceDelegate>

@property (strong, nonatomic) iMateAppFace *imateAppFace;

@end

@implementation PinpadXYDWorkViewController

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

    self.titleLabel.text = @"信雅达密码键盘";
    
    //获取iMateAppFace实例
    _imateAppFace = [iMateAppFace sharedController];
    _imateAppFace.delegate = self;
    
    //设置Pinpad型号
    [_imateAppFace pinpadSetModel:PINPAD_MODEL_XYD];
    
    [self showLogMessage:@"信雅达Pinpad测试项目："];
    [self showLogMessage:@"1、上电"];
    [self showLogMessage:@"2、复位"];
    [self showLogMessage:@"3、取固件版本号"];
    [self showLogMessage:@"4、下装Masterkey"];
    [self showLogMessage:@"5、下装Workingkey"];
    [self showLogMessage:@"6、输入PinBlock"];
    [self showLogMessage:@"7、加密"];
    [self showLogMessage:@"8、计算MAC"];
    [self showLogMessage:@"9、下电\n"];

    [self showLogMessage:@"Pinpad正在上电..."];
    [_imateAppFace pinPadPowerOn];
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

#pragma mark pinpad delegate
- (void)pinPadDelegateResponse:(NSInteger)returnCode  requestType:(PinPadRequestType)type responseData:(NSData *)responseData error:(NSString *)error
{
    NSString *str;
    NSString *requestStr;
    
    switch (type) {
        case PinPadRequestTypePowerOn:
            requestStr = @"PinpadRequestTypePowerOn";
            sleep(2); //等待2秒
            break;
        case PinPadRequestTypePowerOff:
            requestStr = @"PinpadRequestTypePowerOff";
            break;
        case PinPadRequestTypeReset:
            requestStr = @"PinpadRequestTypeReset";
            break;
        case PinPadRequestTypeVersion:
            requestStr = @"PinpadRequestTypeVersion";
            break;
        case PinPadRequestTypeDownloadMasterKey:
            requestStr = @"PinpadRequestTypeDownloadMasterKey";
            break;
        case PinPadRequestTypeDownloadWorkingKey:
            requestStr = @"PinpadRequestTypeDownloadWorkingKey";
            break;
        case PinPadRequestTypeInputPinBlock:
            requestStr = @"PinpadRequestTypeInputPinBlock";
            break;
        case PinPadRequestTypeEncrypt:
            requestStr = @"PinpadRequestTypeEncrypt";
            break;
        case PinPadRequestTypeMac:
            requestStr = @"PinpadRequestTypeMac";
            break;
    }
    if ( returnCode ) {
        str = [NSString stringWithFormat:@"%@,返回码:%d,错误信息:%@",requestStr, (int)returnCode, error];
        [self showLogMessage:str];
    }
    else {
        if (responseData) {
            if (type == PinPadRequestTypeVersion)
                [self showLogMessage:[NSString stringWithFormat:@"data:%@",[NSString stringWithFormat:@"%s",responseData.bytes]]];
            else
                [self showLogMessage:[NSString stringWithFormat:@"data:%@",[iMateAppFace oneTwoData:responseData]]];
        }
        str = [NSString stringWithFormat:@"%@,Pinpad操作成功",requestStr];
        [self showLogMessage:str];
        
        // do next step
        int nextStepType = -1;
        for (int i = 0; i < sizeof(doListArray)/sizeof(PinPadRequestType); i++) {
            if (doListArray[i] == type && doListArray[i] != PinPadRequestTypePowerOff) {
                nextStepType = doListArray[i + 1];
                break;
            }
        }
        if (nextStepType > PinPadRequestTypePowerOn) {
            switch (nextStepType) {
                case PinPadRequestTypeReset:
                    [self.imateAppFace pinPadReset:NO];
                    break;
                case PinPadRequestTypeVersion:
                    [self.imateAppFace pinPadVersion];
                    break;
                case PinPadRequestTypeDownloadMasterKey:
                    [self.imateAppFace pinPadDownloadMasterKey:YES index:1 masterKey:masterKey keyLength:16];
                    break;
                case PinPadRequestTypeDownloadWorkingKey:
                    [self.imateAppFace pinPadDownloadWorkingKey:YES masterIndex:1 workingIndex:1 workingKey:workingKey keyLength:20];
                    break;
                case PinPadRequestTypeInputPinBlock:
                    [self.imateAppFace pinPadInputPinblock:YES isAutoReturn:YES masterIndex:1 workingIndex:1 cardNo:@"8881234567890123" pinLength:6 timeout:20];
                    break;
                case PinPadRequestTypeEncrypt:
                    [self.imateAppFace pinPadEncrypt:YES algo:ALGO_ENCRYPT masterIndex:1 workingIndex:1 data:dataEncrypt dataLength:8];
                    break;
                case PinPadRequestTypeMac:
                    [self.imateAppFace pinPadMac:YES masterIndex:1 workingIndex:1 data:dataMac dataLength:8];
                    break;
                case PinPadRequestTypePowerOff:
                    [self.imateAppFace pinPadPowerOff];
                    break;
            }
        }
    }
}

@end
