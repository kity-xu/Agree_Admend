//
//  SwipeCardWorkViewController.m
//  iMateTest
//
//  Created by hxsmart on 14-7-22.
//  Copyright (c) 2014年 hxsmart. All rights reserved.
//

#import "SwipeCardWorkViewController.h"
#import <HXiMateSDK/iMateAppFace.h>

@interface SwipeCardWorkViewController () <iMateAppFaceDelegate>

@property (strong, nonatomic) iMateAppFace *imateAppFace;

@end

@implementation SwipeCardWorkViewController

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

    self.titleLabel.text = @"刷卡测试";
    
    //获取iMateAppFace实例
    _imateAppFace = [iMateAppFace sharedController];
    _imateAppFace.delegate = self;
    
    [self playGifInImageView:@"mag_play"];
    
    [self showLogMessage:@"请刷卡..."];
    [_imateAppFace swipeCard:20];
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

- (void)iMateDelegateSwipeCard:(NSInteger)returnCode track2:(NSString*)track2 track3:(NSString *)track3 error:(NSString *)error
{
    [self hidenGifImageView];

    NSString *str = [NSString stringWithFormat:@"swipeCard返回码 : %ld", (long)returnCode];
    [self showLogMessage:str];
    
    if ( !returnCode ) {
        str = [NSString stringWithFormat:@"二磁道数据 : \n%@",track2];
        [self showLogMessage:str];
        str = [NSString stringWithFormat:@"三磁道数据 : \n%@",track3];
        [self showLogMessage:str];
    }
    else {
        str = [NSString stringWithFormat:@"错误信息 : %@",error];
        [self showLogMessage:str];
    }
}

@end
