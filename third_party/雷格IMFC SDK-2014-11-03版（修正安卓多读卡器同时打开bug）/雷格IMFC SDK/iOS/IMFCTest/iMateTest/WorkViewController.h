//
//  WorkViewController.h
//  iMateTest
//
//  Created by hxsmart on 14-7-21.
//  Copyright (c) 2014年 hxsmart. All rights reserved.
//

#import <UIKit/UIKit.h>
#import <HXiMateSDK/iMateAppFace.h>

@protocol WorkViewControllerDelegate <NSObject>
@optional
- (void)workViewControllerFinish;
@end

@interface WorkViewController : UIViewController

@property (nonatomic) id<WorkViewControllerDelegate>delegate;

@property (strong, nonatomic) IBOutlet UILabel *titleLabel;
@property (strong, nonatomic) IBOutlet UIImageView *idImageView;

- (void)showLogMessage:(NSString *)logMessage;

- (void)playGifInImageView:(NSString *)gifImageFileName;
- (void)showLogTextInMainThread:(NSString*)logTxt;

- (void)hidenGifImageView;

@end
