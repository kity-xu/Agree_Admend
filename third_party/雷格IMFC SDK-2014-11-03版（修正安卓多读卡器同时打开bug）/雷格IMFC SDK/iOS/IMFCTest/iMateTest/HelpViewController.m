//
//  HelpViewController.m
//  iMateTest
//
//  Created by hxsmart on 14-7-24.
//  Copyright (c) 2014年 hxsmart. All rights reserved.
//

#import "HelpViewController.h"

@interface HelpViewController ()  <UIWebViewDelegate>

@end

@implementation HelpViewController

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
    // Do any additional setup after loading the view from its nib.
    
    self.title = @"iMate快速入门指南";
    
    self.navigationController.navigationBarHidden = NO;
    self.navigationController.toolbarHidden = YES;
    
    self.view.backgroundColor = [UIColor whiteColor];
    
    CGRect rect = [[UIScreen mainScreen] bounds];
    CGSize screenSize = rect.size;
    UIWebView *webView = [[UIWebView alloc] initWithFrame:CGRectMake(0,0,screenSize.width,screenSize.height)];
    webView.delegate = self;
    
    NSString *path = [[NSBundle mainBundle] pathForResource:@"menu" ofType:@"pdf"];
    NSURL *targetURL = [NSURL fileURLWithPath:path];
    NSURLRequest *request = [NSURLRequest requestWithURL:targetURL];
    [webView loadRequest:request];
    
    [self.view addSubview:webView];
}

- (void)didReceiveMemoryWarning
{
    [super didReceiveMemoryWarning];
    // Dispose of any resources that can be recreated.
}

#pragma mark UIWebViewDelegate

- (void)webViewDidFinishLoad:(UIWebView *)webView
{
    webView.scalesPageToFit = YES;
    webView.contentMode = UIViewContentModeScaleAspectFill;
}

@end
