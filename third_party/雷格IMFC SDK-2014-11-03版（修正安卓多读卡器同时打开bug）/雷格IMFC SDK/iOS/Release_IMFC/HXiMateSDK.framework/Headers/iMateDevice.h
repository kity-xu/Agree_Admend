//
//  iMateDevice.h
//  HXiMateSDK
//
//  Created by hxsmart on 14-4-17.
//  Copyright (c) 2014年 hxsmart. All rights reserved.
//

#ifndef HXiMateSDK_iMateDevice_h
#define HXiMateSDK_iMateDevice_h

// SD ICC IO Type
#define SDSC_SCIO_UNKNOWN                    0         // The type of the SCIO is unknown.
#define SDSC_SCIO_7816                       1         // The type of the SCIO is 7816.
#define SDSC_SCIO_SPI_V1                     2         // The type of the SCIO is SPI V1.
#define SDSC_SCIO_SPI_V2                     3         // The type of the SCIO is SPI V2.
#define SDSC_SCIO_SPI_V3                     4         // The type of the SCIO is SPI V3

// 打开指纹模块电源，并检测是否连接成功。指纹模块的启动时间在2秒左右。
// Device Mode采用A模式
// 返回码： 0      :   成功
//          99     :   不支持该功能
//          其它   :   失败或未检测到指纹模块
extern unsigned int uiFingerprintOpen(void);

// 关闭指纹模块电源
extern void vFingerprintClose(void);

// 检测指纹模块是否在线。该函数在fingerprintOpen后可不用调用。
// 返回码：
//      0         : 成功
//      其它      : 失败
extern unsigned int uiFingerprintLink(void);

// 向指纹模块发送命令, 不包含前导串和验证码
// 输入参数：
//      psIn      :   发送的数据缓冲
//      uiInLen   :   发送数据长度
// 返回码：
//          0     :   成功
//         其它   :   失败
extern void vFingerprintSend(unsigned char *psIn, unsigned int uiInLen);

// 在超时时间内等待指纹模块返回的数据，接收到的数据不包含前导串和验证码。
// 输入参数：
//      puiOutLen   :   psOut缓冲区大小
//      ulTimeOutMs :   等待返回数据的超时时间（毫秒）
// 输出参数：
//      psOut       :   输出数据的缓冲区（不包括前导串和验证码等），需要预先分配空间，maxlength = 512
//      puiOutLen   :   接收到的数据长度（不包括前导串和验证码等）
// 返回码：
//      0           :   成功
//      其它        :   失败
extern unsigned int uiFingerprintRecv(unsigned char *psOut, unsigned int *puiOutLen, unsigned long ulTimeOutMs);

// TF ICC Functions

// 打开SD卡电源，接口初始化
// 返回码：
//      0       :   成功
//      99      :   不支持该功能
//      其它    :   失败
extern unsigned int uiSD_Init(void);

// 关闭SD卡电源
extern void vSD_DeInit(void);

// 识别SD_ICC，冷复位
// 返回码：
//      0       :   成功
//      其它    :   失败
extern unsigned int uiSDSCConnectDev(void);

// 关闭SD_ICC
// 返回码：
//      0       :   成功
//      其它    :   失败
extern unsigned int uiSDSCDisconnectDev(void);

// 获取SD_ICC固件版本号
// 输入参数 ：
//      puiFirmwareVerLen   :   psFirmwareVer缓冲区长度
// 输出参数 ：
//      psFirmwareVer       :   固件版本缓冲区(需要预先分配空间，maxlength = 20）
//      puiFirmwareVerLen   :   固件版本数据的长度
// 返回码：
//      0                   :   成功
//      其它                :   失败
extern unsigned int uiSDSCGetFirmwareVer(unsigned char *psFirmwareVer, unsigned int *puiFirmwareVerLen);

// SD_ICC热复位
// 输入参数 ：
//      puiAtrLen    :   psAtr缓冲区长度
// 输出参数 :
//      psAtr        :   Atr数据缓冲区（需要预先分配空间，maxlength = 80）
//      puiAtrLen    :   返回复位数据长度
// 返回码   ：
//      0            :   成功
//      其它         :   失败
extern unsigned int uiSDSCResetCard(unsigned char *psAtr, unsigned int *puiAtrLen);


// 转换SD_ICC电源模式
// 返回码：
//      0       :   成功
//      其它    :   失败
extern unsigned int uiSDSCResetController(unsigned int uiSCPowerMode);

// SD_ICC APDU
// 输入参数 :
//      psCommand       :   ICC Apdu命令串
//      uiCommandLen    :   命令串长度
//      uiTimeOutMode   :   超时模式，固定使用0.
//      puiOutDataLen   :   psOutData缓冲区长度
// 输出参数 :
//      psOutData       :   响应串缓冲区（不包括状态字）,需要预分配空间，maxlength = 300
//      puiOutDataLen   :   返回响应数据长度
//      puiCosState     :   卡片执行状态字
// 返回码：
//      0               :   成功
//      其它            :   失败
extern unsigned int uiSDSCTransmit(unsigned char *psCommand, unsigned int uiCommandLen, unsigned int uiTimeOutMode, unsigned char *psOutData, unsigned int *puiOutDataLen, unsigned int *puiCosState);

// SD_ICC APDU EX
// 输入参数 :
//      psCommand       :   ICC Apdu命令串
//      uiCommandLen    :   命令串长度
//      uiTimeOutMode   :   超时模式，固定使用0.
//      puiOutDataLen   :   psOutData缓冲区长度
// 输出参数 :
//      psOutData       :   响应串缓冲区（包括状态字）,需要预分配空间，maxlength = 300
//      puiOutDataLen   :   返回响应数据长度
// 返回码：
//      0               :   成功
//      其它            :   失败
extern unsigned int uiSDSCTransmitEx(unsigned char *psCommand, unsigned int uiCommandLen, unsigned int uiTimeOutMode, unsigned char *psOutData, unsigned int *puiOutDataLen);

// 获取SD_ICC SDK版本号
// 输出参数 ：
//      puiVersionLen :   version缓冲区大小
// 输出参数 ：
//      pszVersion    :   SDK版本号数据缓冲，需要预先分配空间，maxlength = 20
//      puiVersionLen :   SDK版本号数据长度
// 返回码：
//      0             :   成功
//      其它          :   失败
extern unsigned int uiSDSCGetSDKVersion(char *pszVersion, unsigned int *puiVersionLen);

// 获取SD_ICC IO类型
// 输出参数：
//      puiSCIOType     :   IO类型,需预先创建对象：Integer SCIOType = new Integer(0)
//                          请参考：SD ICC IO Type
// 返回码：
//      0               :   成功
//      其它            :   失败
extern unsigned int uiSDSCGetSCIOType(unsigned int *puiSCIOType);

#endif

