//
//  ReadIcCardWorkViewController.m
//  iMateTest
//
//  Created by hxsmart on 14-7-22.
//  Copyright (c) 2014年 hxsmart. All rights reserved.
//

#import "ReadPbocInfoWorkViewController.h"
#import <HXiMateSDK/iMateAppFace.h>
#import <HXiMateSDK/iMateAppFace+Pboc.h>

@interface ReadPbocInfoWorkViewController () <iMateAppFaceDelegate>

@property (strong, nonatomic) iMateAppFace *imateAppFace;

@end

@implementation ReadPbocInfoWorkViewController

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

    self.titleLabel.text = @"读取Pboc卡信息";
    
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
    NSString *str = [NSString stringWithFormat:@"icResetCard返回码 : %ld", (long)returnCode];
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
    [self showLogMessage:@"正在读Pboc卡信息..."];
    [self.imateAppFace pbocReadInfo];
}

- (void)iMateDelegatePbocReadInfo:(NSInteger)returnCode
                         cardInfo:(NSString *)cardInfo
                            error:(NSString *)error
{
    NSString *str = [NSString stringWithFormat:@"pbocReadInfo返回码:%ld", (long)returnCode];
    [self showLogMessage:str];
    
    if ( returnCode ) {
        str = [NSString stringWithFormat:@"错误信息:%@",error];
        [self showLogMessage:str];
    }
    else {
        str = [NSString stringWithFormat:@"卡信息:\n%@",cardInfo];
        [self showLogMessage:str];
    }
}


@end
