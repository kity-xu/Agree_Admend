//
//  ReadIdCardWorkViewController.m
//  iMateTest
//
//  Created by hxsmart on 14-7-22.
//  Copyright (c) 2014年 hxsmart. All rights reserved.
//

#import "ReadIdCardWorkViewController.h"

@interface ReadIdCardWorkViewController () <iMateAppFaceDelegate>

@property (strong, nonatomic) iMateAppFace *imateAppFace;

@end

@implementation ReadIdCardWorkViewController

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
    
    self.titleLabel.text = @"二代证读卡测试";
    
    //获取iMateAppFace实例
    _imateAppFace = [iMateAppFace sharedController];
    _imateAppFace.delegate = self;
    
    [self playGifInImageView:@"id_play"];
    
    [self showLogMessage:@"请放置二代证..."];
    [_imateAppFace idReadMessage:20];
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

- (void)iMateDelegateIDReadMessage:(NSInteger)returnCode information:(NSData *)information photo:(NSData*)photo error:(NSString *)error
{
    [self hidenGifImageView];
    
    NSString *str = [NSString stringWithFormat:@"idReadMessage返回码:%ld", (long)returnCode];
    [self showLogMessage:str];
    
    if ( !returnCode ) {
        str = [NSString stringWithFormat:@"information data:\n%@",[iMateAppFace oneTwoData:information]];
        NSLog(@"%@",str);
        
        NSString *str2 = [NSString stringWithFormat:@"photo data:\n%@",[iMateAppFace oneTwoData:photo]];
        NSLog(@"%@%@",str,str2);
        
        for(NSString *info in [iMateAppFace processIdCardInfo:information]) {
            [self showLogMessage:info];
        }
        if (self.idImageView) {
            self.idImageView.hidden = NO;
            self.idImageView.image = [iMateAppFace processIdCardPhoto:photo];
        }
    }
    else {
        str = [NSString stringWithFormat:@"错误信息:%@",error];
        [self showLogMessage:str];
    }
}

@end
