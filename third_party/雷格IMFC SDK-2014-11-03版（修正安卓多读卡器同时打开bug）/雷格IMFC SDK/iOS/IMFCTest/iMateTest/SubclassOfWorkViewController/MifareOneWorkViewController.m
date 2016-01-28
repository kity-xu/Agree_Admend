//
//  MifareOneWorkViewController.m
//  iMateTest
//
//  Created by hxsmart on 14-7-22.
//  Copyright (c) 2014年 hxsmart. All rights reserved.
//

#import "MifareOneWorkViewController.h"
#import <HXiMateSDK/iMateAppFace.h>
#import <HXiMateSDK/iMateMifCard.h>

@interface MifareOneWorkViewController () <iMateAppFaceDelegate>

@property (strong, nonatomic) iMateAppFace *imateAppFace;

@end

@implementation MifareOneWorkViewController

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
    
    self.titleLabel.text = @"M1卡测试";
    
    //获取iMateAppFace实例
    _imateAppFace = [iMateAppFace sharedController];
    _imateAppFace.delegate = self;
    
    //[self playGifInImageView:@"id_play"];
    
    dispatch_async(dispatch_get_global_queue(DISPATCH_QUEUE_PRIORITY_DEFAULT, 0), ^{
        @autoreleasepool {
            [self showLogTextInMainThread:@"请放置M1卡 ..."];
            
            double timeSeconds = [self currentTimeSeconds] + 5; //整体5秒等待时间
            unsigned char sSerialNo[10];
            unsigned int uiRet;
            while ([self currentTimeSeconds] < timeSeconds) {
                uiRet = uiMifCard(sSerialNo);
                if (uiRet > 0)
                    break;
            }
            //[self performSelectorOnMainThread:@selector(hidenGifImageView) withObject:nil waitUntilDone:NO];
            if ([self currentTimeSeconds] >= timeSeconds) {
                [self showLogTextInMainThread:@"等待M1卡超时"];
            }
            else {
                [self showLogTextInMainThread:[NSString stringWithFormat:@"检测到M1卡,UID:%02x%02x%02x%02x", sSerialNo[0],sSerialNo[1],sSerialNo[2],sSerialNo[3]]];
                
                [self M1TestThread];
            }
        }
    });
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

void mifInitMoney(long balence, unsigned char *buf)
{
    buf[0] = (balence & 0xff);
    buf[1] = ((balence>>8) & 0xff);
    buf[2] = ((balence>>16) & 0xff);
    buf[3] = ((balence>>24) & 0xff);
    buf[4] = ~buf[0];
    buf[5] = ~buf[1];
    buf[6] = ~buf[2];
    buf[7] = ~buf[3];
    buf[8] = buf[0];
    buf[9] = buf[1];
    buf[10] = buf[2];
    buf[11] = buf[3];
    
    buf[12] = 0x01;
    buf[13] = 0xfe;
    buf[14] = 0x01;
    buf[15] = 0xfe;
}


- (void)M1TestThread
{
    dispatch_async(dispatch_get_global_queue(DISPATCH_QUEUE_PRIORITY_DEFAULT, 0), ^{
        @autoreleasepool {
            unsigned int ret;
            /*
            double timeSeconds = [self currentTimeSeconds] + 20; //整体20秒等待时间
            unsigned char sSerialNo[10];
            while ([self currentTimeSeconds] < timeSeconds) {
                ret = uiMifCard(sSerialNo);
                if (ret > 0)
                    break;
            }
            if ([self currentTimeSeconds] >= timeSeconds) {
                [self showLogTextInMainThread:@"等待MIF卡超时"];
                return;
            }
            [self showLogTextInMainThread:[NSString stringWithFormat:@"检测到MIF卡,UID:%02x%02x%02x%02x", sSerialNo[0],sSerialNo[1],sSerialNo[2],sSerialNo[3]]];
            */
            
            //2、认证扇区
            ret = uiMifAuth(1, 0, (unsigned char*)"\xff\xff\xff\xff\xff\xff");
            if (ret) {
                [self showLogTextInMainThread:@"扇区1认证失败"];
                return;
            }
            [self showLogTextInMainThread:@"扇区1认证成功"];
            
            //3、读1扇区0块
            unsigned char buf[16];
            ret = uiMifReadBlock(1, 0, buf);
            if (ret) {
                [self showLogTextInMainThread:@"读1扇区0块失败"];
                return;
            }
            [self showLogTextInMainThread:[NSString stringWithFormat:@"1扇区0块读卡成功， Data:%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x", buf[0],buf[1],buf[2],buf[3],buf[4],buf[5],buf[6],buf[7],buf[8],buf[9],buf[10],buf[11],buf[12],buf[13],buf[14],buf[15]]];
            
            //4、写1扇区0块
            ret = uiMifWriteBlock(1, 0, (unsigned char*)"\x1\x2\x3\x4\x5\x6\x7\x8\x9\x0\xa\xb\xc\xd\xe\xf");
            if (ret) {
                [self showLogTextInMainThread:@"写1扇区0块失败"];
                return;
            }
            memset(buf, 0 , sizeof(buf));
            ret = uiMifReadBlock(1, 0, buf);
            if (ret) {
                [self showLogTextInMainThread:@"读1扇区0块失败"];
                return;
            }
            [self showLogTextInMainThread:[NSString stringWithFormat:@"1扇区0块写卡成功， Data:%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x", buf[0],buf[1],buf[2],buf[3],buf[4],buf[5],buf[6],buf[7],buf[8],buf[9],buf[10],buf[11],buf[12],buf[13],buf[14],buf[15]]];
            
            //恢复1扇区0块数据
            memset(buf, 0, sizeof(buf));
            ret = uiMifWriteBlock(1, 0, buf);
            if (ret) {
                [self showLogTextInMainThread:@"恢复1扇区0块数据失败"];
                return;
            }
            
            //6、初始化钱包（1扇区1块，100.00元）
            mifInitMoney(0,buf);
            ret = uiMifWriteBlock(1, 1, buf);
            if (ret) {
                [self showLogTextInMainThread:@"恢复1扇区0块数据失败"];
                return;
            }
            [self showLogTextInMainThread:@"1扇区1块写钱包初始化成功"];
            
            //7、钱包加值(10.00元）
            ret = uiMifIncrement(1, 1, 1000);
            if (ret) {
                [self showLogTextInMainThread:@"1扇区1块钱包加值失败"];
                return;
            }
            [self showLogTextInMainThread:@"1扇区1块钱包加值成功"];
            
            //8、读1扇区1块钱包（验证）
            ret = uiMifReadBlock(1, 1, buf);
            if (ret) {
                [self showLogTextInMainThread:@"读1扇区1块失败"];
                return;
            }
            [self showLogTextInMainThread:[NSString stringWithFormat:@"1扇区1块读卡成功， Data:%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x", buf[0],buf[1],buf[2],buf[3],buf[4],buf[5],buf[6],buf[7],buf[8],buf[9],buf[10],buf[11],buf[12],buf[13],buf[14],buf[15]]];
            //9、钱包减值(10.00元）
            ret = uiMifDecrement(1, 1, 1000);
            if (ret) {
                [self showLogTextInMainThread:@"1扇区1块钱包减值失败"];
                return;
            }
            [self showLogTextInMainThread:@"1扇区1块钱包减值成功"];
            
            
            //10、读1扇区1块钱包（验证）
            ret = uiMifReadBlock(1, 1, buf);
            if (ret) {
                [self showLogTextInMainThread:@"读1扇区1块失败"];
                return;
            }
            [self showLogTextInMainThread:[NSString stringWithFormat:@"1扇区1块读卡成功， Data:%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x", buf[0],buf[1],buf[2],buf[3],buf[4],buf[5],buf[6],buf[7],buf[8],buf[9],buf[10],buf[11],buf[12],buf[13],buf[14],buf[15]]];
            
            
            //块拷贝（1块拷贝到2块）
            ret = uiMifCopy(1, 1, 1, 2);
            if (ret) {
                [self showLogTextInMainThread:@"1扇区1块拷贝到1扇区2块失败"];
                return;
            }
            [self showLogTextInMainThread:@"1扇区1块拷贝到1扇区2块成功"];
            //读1扇区2块钱包（验证）
            ret = uiMifReadBlock(1, 2, buf);
            if (ret) {
                [self showLogTextInMainThread:@"读1扇区2块失败"];
                return;
            }
            [self showLogTextInMainThread:[NSString stringWithFormat:@"1扇区2块读卡成功， Data:%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x", buf[0],buf[1],buf[2],buf[3],buf[4],buf[5],buf[6],buf[7],buf[8],buf[9],buf[10],buf[11],buf[12],buf[13],buf[14],buf[15]]];
            
            [self showLogTextInMainThread:@"MIF卡测试完成"];
        }
    });
}


@end
