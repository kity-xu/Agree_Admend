//
//  WorkViewController.m
//  iMateTest
//
//  Created by hxsmart on 14-7-21.
//  Copyright (c) 2014年 hxsmart. All rights reserved.
//

#import "WorkViewController.h"

@interface WorkViewController ()

@property (strong, nonatomic) IBOutlet UITextView *logMessage;
@property (strong, nonatomic) IBOutlet UIWebView *playWebView;

@end

@implementation WorkViewController

- (id)initWithNibName:(NSString *)nibNameOrNil bundle:(NSBundle *)nibBundleOrNil
{
    if (IS_IPAD)
        self = [super initWithNibName:@"WorkViewController~iPad" bundle:nibBundleOrNil];
    else
        self = [super initWithNibName:@"WorkViewController~iPhone" bundle:nibBundleOrNil];
    
    if (self) {
    }
    return self;
}

- (void)viewDidLoad
{
    [super viewDidLoad];

    if (IS_IPAD) {
        self.navigationController.navigationBarHidden = YES;
        self.navigationController.toolbarHidden = YES;
    }
    else {
        self.navigationController.navigationBarHidden = NO;
        self.navigationController.toolbarHidden = YES;
    }

    _titleLabel.text = @"";
    _logMessage.text = @"";
}

- (void)viewWillAppear:(BOOL)animated
{
    self.title = _titleLabel.text;
    [super viewWillAppear:animated];
}

- (void)playGifInImageView:(NSString *)gifImageFileName
{
    _playWebView.hidden = NO;
    NSData *gif = [NSData dataWithContentsOfFile: [[NSBundle mainBundle] pathForResource:gifImageFileName ofType:@"gif"]];
    
    _playWebView.userInteractionEnabled = NO;
    [_playWebView loadData:gif MIMEType:@"image/gif" textEncodingName:nil baseURL:nil];
}

- (void)hidenGifImageView
{
    _playWebView.hidden = YES;
}

- (IBAction)buttonPressed:(id)sender
{
    [_delegate workViewControllerFinish];
}

- (void)showLogMessage:(NSString *)logMessage
{
    _logMessage.text = [NSString stringWithFormat:@"%@%@\n", _logMessage.text, logMessage];
}

- (void)showLogTextInMainThread:(NSString *)logTxt
{
    [self performSelectorOnMainThread:@selector(showLogMessage:) withObject:logTxt waitUntilDone:NO];
}

@end
