//
//  PbocIssueWorkViewController.m
//  iMateTest
//
//  Created by hxsmart on 14-7-22.
//  Copyright (c) 2014年 hxsmart. All rights reserved.
//

#import "PbocIssueWorkViewController.h"
#import <HXiMateSDK/iMateAppFace.h>
#import <HXiMateSDK/iMateAppFace+Pboc.h>

@interface PbocIssueWorkViewController () <iMateAppFaceDelegate>

@property (strong, nonatomic) iMateAppFace *imateAppFace;

@end

@implementation PbocIssueWorkViewController

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
    
    self.titleLabel.text = @"Pboc个人化演示";
    
    //获取iMateAppFace实例
    _imateAppFace = [iMateAppFace sharedController];
    _imateAppFace.delegate = self;
    
    [self playGifInImageView:@"ic_play"];
    
    [self showLogMessage:@"请插卡..."];
    [_imateAppFace icResetCard:0 tag:0 timeout:20];
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

- (void)iMateDelegateICResetCard:(NSInteger)returnCode resetData:(NSData *)resetData tag:(NSInteger)tag error:(NSString *)error
{
    [self hidenGifImageView];
    //打印返回码
    NSString *str = [NSString stringWithFormat:@"icResetCard返回码 : %d",(int)returnCode];
    [self showLogMessage:str];
    
    //如果复位失败，不再继续
    if ( returnCode ) {
        str = [NSString stringWithFormat:@"错误信息 : %@",error];
        [self showLogMessage:str];
        return;
    }
    else {
        //打印卡片复位返回的数据
        str = [NSString stringWithFormat:@"复位数据:\n%@\n",[iMateAppFace oneTwoData:resetData]];
        [self showLogMessage:str];
    }
    
    struct structIssData IssData;
    memset(&IssData, 0, sizeof(IssData));
    IssData.iAppVer = 0x20;
    IssData.iKmcIndex = 1;
    IssData.iAppKeyBaseIndex = 3;
    IssData.iIssuerRsaKeyIndex = 2;
    strcpy(IssData.szPan, "8888800000123456"); //卡号必须88888开头
    IssData.iPanSerialNo = 12;
    strcpy(IssData.szExpireDate, "491231");
    strcpy(IssData.szHolderName, "HxSmart");
    IssData.iHolderIdType = 0;	// 0:身份证 1:军官证 2:护照 3:入境证 4:临时身份证 5:其它
    strcpy(IssData.szHolderId, "220104200001010001");
    strcpy(IssData.szDefaultPin, "000000");
    strcpy(IssData.szAid, "A000000333010101");
    strcpy(IssData.szLabel, "Pboc20 d/c");
    IssData.iCaIndex = 1;
    IssData.iIcRsaKeyLen = 128;
    IssData.lIcRsaE = 3;
    IssData.iCountryCode = 156;
    IssData.iCurrencyCode = 156;
    
    NSData *issData = [NSData dataWithBytes:&IssData length:sizeof(IssData)];
    
    [self.imateAppFace pbocIssCard:issData];
}

- (void)iMateDelegatePbocIssCard:(NSInteger)returnCode error:(NSString *)error
{
    NSString *str = [NSString stringWithFormat:@"iMateDelegatePbocIssCard返回码:%d",(int)returnCode];
    [self showLogMessage:str];
    
    if ( returnCode ) {
        str = [NSString stringWithFormat:@"错误信息:%@",error];
        [self showLogMessage:str];
    }
    else {
        str = [NSString stringWithFormat:@"发卡成功"];
        [self showLogMessage:str];
    }
}

- (void)iMateDelegateRuningStatus:(NSString *)statusString
{
    [self showLogMessage:statusString];
}

@end
