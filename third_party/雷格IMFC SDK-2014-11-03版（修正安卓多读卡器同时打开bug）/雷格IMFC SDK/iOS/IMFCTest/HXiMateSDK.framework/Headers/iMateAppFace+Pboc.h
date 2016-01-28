//
//  iMateAppFace+PbocIssue.h
//  HXiMateSDK
//
//  Created by hxsmart on 13-12-23.
//  Copyright (c) 2013年 hxsmart. All rights reserved.
//

#import <HXiMateSDK/iMateAppFace.h>

struct structIssData {
	int  iAppVer;				// 应用版本号, 0x20或0x30, 分别表示Pboc2.0卡或Pboc3.0卡
    
	int  iKmcIndex;
	int  iAppKeyBaseIndex;
	int  iIssuerRsaKeyIndex;
    
	char szPan[19+1];			// 12-19
	int  iPanSerialNo;			// 0-99, -1表示无
	char szExpireDate[6+1];		// YYMMDD
	char szHolderName[45+1];	// Len>=2
	int  iHolderIdType;			// 0:身份证 1:军官证 2:护照 3:入境证 4:临时身份证 5:其它
	char szHolderId[40+1];
	char szDefaultPin[12+1];	// 4-12
    
	char szAid[32+1];			// OneTwo之后的
	char szLabel[16+1];			// 1-16
	int  iCaIndex;				// 0-255
	int  iIcRsaKeyLen;			// 64-192
	long lIcRsaE;				// 3 or 65537
	int  iCountryCode;			// 1-999
    
	int  iCurrencyCode;			// 1-999
};

@protocol iMateAppFacePbocDelegate <iMateAppFaceDelegate>

@optional
// 操作状态信息(后台运行状态报告）
- (void)iMateDelegateRuningStatus:(NSString *)statusString;


// 读PBOC卡信息操作执行结束后，该方法被调用，返回结果。cardInfo中的数据用","间隔
- (void)iMateDelegatePbocReadInfo:(NSInteger)returnCode
                         cardInfo:(NSString *)cardInfo
                            error:(NSString *)error;

// 读PBOC卡日志操作执行结束后，该方法被调用，返回结果。cardLog为NSString数组的方式。
- (void)iMateDelegatePbocReadLog:(NSInteger)returnCode
                         cardLog:(NSArray *)cardLog
                           error:(NSString *)error;

// Pboc发卡操作执行结束后，该方法被调用，返回结果。
- (void)iMateDelegatePbocIssCard:(NSInteger)returnCode
                           error:(NSString *)error;
@end

@interface iMateAppFace (Pboc)

// 设置IC读卡器类型，cardReaderType: 0 芯片读卡器；1 射频读卡器
- (void)pbocIcCardReaderType:(int)cardReaderType;

// 读取PBOC卡信息的操作请求，该操作需执行之前成功执行了icResetCard操作
// 结果由iMateDelegatePbocReadInfo获得
// 获取卡号，持卡人姓名，持卡人证件号码，应用失效日期、卡序列号
- (void)pbocReadInfo;

// 读取PBOC卡信息的操作请求(扩展方法），该操作需执行之前成功执行了icResetCard操作
// 结果由iMateDelegatePbocReadInfo获得
// 以XML格式返回卡片的信息，包括卡号、姓名、证件类型、证件号、二磁道数据、一磁道数据、现金余额、余额上限、应用失效日期、卡序列号
// 格式参数： outType  0：XML格式，1：用逗号分割（按照以上说明的顺序）
- (void)pbocReadInfoEx:(int)outType;

// 读取PBOC卡日志的操作请求，该操作需执行之前成功执行了icResetCard操作
- (void)pbocReadLog;

// PBOC发卡操作请求，该操作需执行之前成功执行了icResetCard操作, issData内容参考structIssData
- (void)pbocIssCard:(NSData *)issData;

@end
