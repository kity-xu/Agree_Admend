//
//  PrinterWorkViewController.m
//  iMateTest
//
//  Created by hxsmart on 14-7-23.
//  Copyright (c) 2014年 hxsmart. All rights reserved.
//

#import "PrinterWorkViewController.h"
#import <HXiMateSDK/iMateAppFace.h>

@interface PrinterWorkViewController () <iMateAppFaceDelegate>

@property (strong, nonatomic) iMateAppFace *imateAppFace;

@end

@implementation PrinterWorkViewController

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

    self.titleLabel.text = @"打印机测试";
    
    //获取iMateAppFace实例
    _imateAppFace = [iMateAppFace sharedController];
    _imateAppFace.delegate = self;
    
    [self showLogMessage:@"目前SDK支持斯普瑞特的蓝牙打印机SP-T7BT针式打印机，为了方便开发，将打印机的SDK与iMate SDK组合在一起。\n"];

    [_imateAppFace printerStatus];
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

#pragma mark printer delegate
- (void)printerDelegateStatusResponse:(NSInteger)status
{
    switch (status) {
        case PRINTER_OK:
            [self showLogMessage:@"打印机状态正常!"];
            
            // 该方法成功后无返回结果
            [_imateAppFace print:@"深圳华信智能科技有限公司\n========================\nhxsmart.com\n2014.7.22\n\n\n"];
            break;
        case PRINTER_CONNECTED:
            [self showLogMessage:@"打印机已连接!"];
            break;
        case PRINTER_NOT_CONNECTED:
            [self showLogMessage:@"打印机未连接!"];
            break;
        case PRINTER_OUT_OF_PAPER:
            [self showLogMessage:@"打印机缺纸!"];
            break;
    }
}

@end
