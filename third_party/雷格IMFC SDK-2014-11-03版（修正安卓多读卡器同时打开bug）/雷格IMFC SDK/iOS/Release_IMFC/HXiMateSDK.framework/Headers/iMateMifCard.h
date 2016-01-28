//
//  iMateMifCard.h
//  HXiMateSDK
//
//  Created by hxsmart on 14-4-18.
//  Copyright (c) 2014年 hxsmart. All rights reserved.
//

#ifndef HXiMateSDK_iMateMifCard_h
#define HXiMateSDK_iMateMifCard_h

// 检测射频卡
// 输出参数：psSerialNo : 返回卡片系列号
// 返    回：>0         : 成功, 卡片系列号字节数
//           0          : 失败
unsigned int uiMifCard(unsigned char *psSerialNo);

// MIF CPU卡激活
// 返    回：0          : 成功
//           其它       : 失败
unsigned int uiMifActive(void);

// 关闭射频信号
// 返    回：0			；成功
//   		 其它		：失败
unsigned int uiMifClose(void);

// MIF移除
// 返    回：0          : 移除
//           其它       : 未移除
unsigned int uiMifRemoved(void);

// M1卡扇区认证
// 输入参数：  ucSecNo	：扇区号
//			 ucKeyAB	：密钥类型，0x00：A密码，0x04: B密码
//			 psKey		: 6字节的密钥
// 返    回：0          : 成功
//           其它       : 失败
unsigned int uiMifAuth(unsigned char ucSecNo, unsigned char ucKeyAB, unsigned char *psKey);

// M1卡读数据块
// 			 ucSecNo	：扇区号
//			 ucBlock	: 块号
// 输出参数：psData		：16字节的数据
// 返    回：0          : 成功
//           其它       : 失败
unsigned int uiMifReadBlock(unsigned char ucSecNo, unsigned char ucBlock, unsigned char *psData);

// M1卡写数据块
// 输入参数：  ucSecNo	：扇区号
//			 ucBlock	: 块号
//			 psData		：16字节的数据
// 返    回：0          : 成功
//           其它       : 失败
unsigned int uiMifWriteBlock(unsigned char ucSecNo, unsigned char ucBlock, unsigned char *psData);


// M1钱包加值
// 输入参数：  ucSecNo	：扇区号
//			 ucBlock	: 块号
//			 ulValue	：值
// 返    回：0          : 成功
//           其它       : 失败
unsigned int uiMifIncrement(unsigned char ucSecNo,unsigned char ucBlock,unsigned long ulValue);

// M1钱包减值
// 输入参数：  ucSecNo	：扇区号
//			 ucBlock	: 块号
//			 ulValue	：值
// 返    回：0          : 成功
//           其它       : 失败
unsigned int uiMifDecrement(unsigned char ucSecNo,unsigned char ucBlock,unsigned long ulValue);

// M1卡块拷贝
// 输入参数： ucSrcSecNo	：源扇区号
//			 ucSrcBlock	: 源块号
//			 ucDesSecNo	: 目的扇区号
//			 ucDesBlock	: 目的块号
// 返    回：0          : 成功
//           其它       : 失败
unsigned int uiMifCopy(unsigned char ucSrcSecNo, unsigned char ucSrcBlock, unsigned char ucDesSecNo, unsigned char ucDesBlock);

// MIF CPU 卡 APDU
// 输入参数：psApduIn	：apdu命令串
//			 uiInLen	: apdu命令串长度
//			 psApduOut	: apdu返回串
//			 puiOutLen	: apdu返回串长度
// 返    回：0          : 成功
//           其它       : 失败
unsigned int uiMifApdu(unsigned char *psApduIn, unsigned int uiInLen, unsigned char *psApduOut, unsigned int *puiOutLen);

#endif
