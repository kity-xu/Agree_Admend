package com.example.imateinterfacetest;

import android.os.Bundle;
import android.os.Handler;
import android.os.Message;
import android.annotation.SuppressLint;
import android.app.Activity;
import android.bluetooth.BluetoothAdapter;
import android.text.method.ScrollingMovementMethod;
import android.view.Menu;
import android.view.View;
import android.widget.Button;
import android.widget.TextView;

import com.hxsmart.imateinterface.*;
import com.hxsmart.imateinterface.imatedevice.iMateDevice;
import com.hxsmart.imateinterface.mifcard.MifCard;
import com.hxsmart.imateinterface.pinpad.*;

import com.ivsign.android.IDCReader.IDCReaderSDK;

public class MainActivity extends Activity {
	static private boolean threadStarted = false;

	final static String wltlibDirectory = "/sdcard/wltlib";
	
	public static BluetoothThread bluetoothThread;
	public Pinpad pinpad;
	public MifCard mifCard;
		
	public Button button0;
	public Button button1;
	public Button button2;
	public Button button3;
	public Button button4;
	public Button button5;
	public Button button6;
	public Button button7;
	public Button button8;
	public Button button9;
	public Button button10;
	
	public TextView logView;
	private LogViewAppendHandler logViewAppendHandler;
	private static boolean isWorking = false;
	
	@Override
	protected void onCreate(Bundle savedInstanceState) {
		super.onCreate(savedInstanceState);
		setContentView(R.layout.activity_main);
		
		// Begin of iMate Bluetooth init Codes
		if (!threadStarted) {
			BluetoothAdapter bluetoothAdapter = BluetoothAdapter
					.getDefaultAdapter();
			bluetoothThread = new BluetoothThread(bluetoothAdapter);
			bluetoothThread.start();
			threadStarted = true;
		}
		pinpad = new Pinpad();
		mifCard = new MifCard();

		// End of iMate Bluetooth init Codes
		
		button0 = (Button)findViewById(R.id.button00);
		button1 = (Button)findViewById(R.id.button01);
		button2 = (Button)findViewById(R.id.button02);
		button3 = (Button)findViewById(R.id.button03);
		button4 = (Button)findViewById(R.id.button04);
		button5 = (Button)findViewById(R.id.button05);
		button6 = (Button)findViewById(R.id.button06);
		button7 = (Button)findViewById(R.id.button07);
		button8 = (Button)findViewById(R.id.button08);
		button9 = (Button)findViewById(R.id.button09);
		button10 = (Button)findViewById(R.id.button10);
		
		logView = (TextView)findViewById(R.id.textView1);
		logView.setMovementMethod(new ScrollingMovementMethod());
		
		logViewAppendHandler = new LogViewAppendHandler();
		
		button0.setOnClickListener(new View.OnClickListener() {
			public void onClick(View view) {
				if (isWorking)
					return;
				logView.setText("");
				
				if (bluetoothThread.deviceIsConnecting()) {
					logView.append("蓝牙已经连接成功\n");
					logView.append("iMate固件版本：" + bluetoothThread.deviceVersion());
				}
				else {
					logView.append("蓝牙未连接\n");
				}
			}
		});
		button1.setOnClickListener(new View.OnClickListener() {
			public void onClick(View view) {
				if (isWorking) {
					logView.setText("");
					bluetoothThread.cancel();
					return;
				}
				logView.setText("刷卡...\n");				
				
				new Thread(new Runnable() {
					@Override
					public void run() {
						isWorking = true;
						int retCode;
						MagCardData cardData = new MagCardData();
						String message;
						retCode = bluetoothThread.swipeCard(cardData, 20);
						switch (retCode) {
						case 0:
							message = "刷卡成功:\n\n" + "卡号:" + cardData.getCardNoString() + "\n"
									+ "二磁道数据:" + cardData.getTrack2String() + "\n"
									+ "三磁道数据:" + cardData.getTrack3String();				
							break;
						case 1:
							message = "通讯超时";						
							break;
						case 9:
							message = "设备未连接";						
							break;
						default:
							message = cardData.getErrorString();												
							break;
						}	
						writeLogFromThread(message);
						isWorking = false;
					}
				}).start();
			}
		});
		button2.setOnClickListener(new View.OnClickListener() {
			public void onClick(View view) {
				if (isWorking) {
					logView.setText("");
					bluetoothThread.cancel();
					return;
				}
				logView.setText("读IC卡...\n");
				
				new Thread(new Runnable() {
					@Override
					public void run() {
						isWorking = true;
						int retCode;
						IcCardData icCardData = new IcCardData();
						String message;
						retCode = bluetoothThread.readIcCard(icCardData, 20);
						switch (retCode) {
						case 0:
							message = "读IC卡成功:\n\n" + "卡号:\t\t\t\t\t" + icCardData.getCardNoString() + "\n"
									+ "序列号:\t\t\t\t" + icCardData.getPanSequenceNoString() + "\n"
									+ "持卡人姓名:\t\t\t" + icCardData.getHolderNameString() + "\n"
									+ "持卡人证件号码:\t" + icCardData.getHolderIdString() + "\n"
									+ "有效期:\t\t\t\t" + icCardData.getExpireDateString() + "\n" 
									+"二磁道等效数据:\t" + icCardData.getTrack2String();						
							break;
						case 1:
							message = "通讯超时";						
							break;
						case 9:
							message = "设备未连接";						
							break;
						default:
							message = icCardData.getErrorString();												
							break;
						}
						writeLogFromThread(message);
						isWorking = false;
					}
				}).start();
			}
		});
		button3.setOnClickListener(new View.OnClickListener() {
			public void onClick(View view) {
				if (isWorking) {
					logView.setText("");
					bluetoothThread.cancel();
					return;
				}
				logView.setText("读二代证...\n");
				
				new Thread(new Runnable() {
					@Override
					public void run() {
						isWorking = true;
						int retCode;
						IdInformationData idInformationDataData = new IdInformationData();
						String message;
						retCode = bluetoothThread.readIdInformation(idInformationDataData, 20);
						switch (retCode) {
						case 0:
							message = "读二代证成功:\n\n" + "姓名:" + idInformationDataData.getNameString() + "\n"
												+ "性别:" + idInformationDataData.getSexString() + "\n"
												+ "民族:" + idInformationDataData.getNationString() + "\n"
												+ "出生:" + idInformationDataData.getBirthdayYearString() + "年" + idInformationDataData.getBirthdayMonthString() + "月" + idInformationDataData.getBirthdayDayString() + "日" + "\n"
												+ "住址:" + idInformationDataData.getAddressString() + "\n"
												+ "身份号码:" + idInformationDataData.getIdNumberString() + "\n"
												+ "签发机关:" + idInformationDataData.getIssuerString() + "\n"
												+ "有效期限:" + idInformationDataData.getValidDateString();						
							break;
						case 1:
							message = "通讯超时";						
							break;
						case 9:
							message = "设备未连接";						
							break;
						default:
							message = idInformationDataData.getErrorString();												
							break;
						}
						writeLogFromThread(message);
						
						/*
						if (retCode == 0) {						
							retCode = IDCReaderSDK.decodingPictureData(wltlibDirectory, idInformationDataData.getPictureData());
							switch (retCode) {
							case 0:
								message = "照片解码成功";
								break;							
							case 1:
								message = "照片解码初始化失败，需要检查传入的wltlibDirectory以及base.dat文件";
								break;
							case 2:
								message = "授权文件license.lic错误";
								break;
							case 3:
								message = "照片解码失败，其它错误";
								break;
							}
							writeLogFromThread(message);					
						}						
						*/
						isWorking = false;
					}
				}).start();
			}
		});		
		button4.setOnClickListener(new View.OnClickListener() {
			public void onClick(View view) {
				if (isWorking) {
					logView.setText("");
					pinpad.cancel();
					return;
				}
				logView.setText("密码键盘测试\n");

				new Thread(new Runnable() {
					@Override
					public void run() {
						isWorking = true;
						
						// 打开密码键盘电源
						try {
							pinpad.powerOn();
						}catch (Exception e) {
							writeLogFromThread(e.getMessage());
							isWorking = false;
							return;
						}
						writeLogFromThread("打开密码键盘电源成功");
						
						// 等待2秒，密码键盘完成上电后再继续操作
						try {
			                Thread.sleep(2000);
			            }
			            catch (InterruptedException e) {
			            }

						writeLogFromThread("密码键盘复位自检...");
						try {
							pinpad.reset(false);
						}catch (Exception e) {
							writeLogFromThread(e.getMessage());
							isWorking = false;
							return;
						}
						writeLogFromThread("密码键盘复位自检成功");
						
						writeLogFromThread("下载主密钥...");
						byte[] masterKey = {0x00, 0x01,0x02,0x03,0x04,0x05,0x06,0x07,0x08,0x09,0x0a,0x0b,0x0c,0x0d,0x0e,0x0f};
						try {
							pinpad.downloadMasterKey(true, 1, masterKey, masterKey.length);
						}catch (Exception e) {
							writeLogFromThread(e.getMessage());
							isWorking = false;
							return;
						}
						writeLogFromThread("下载主密钥成功");
						
						writeLogFromThread("下载工作密钥...");
						byte[] workingKey = {0x12, 0x34,0x56,0x78, (byte)0x90, (byte)0xab,(byte)0xcd,(byte)0xef,0x01,0x02,0x03,0x04,0x05,0x06,0x07,0x08};
						try {
							pinpad.downloadWorkingKey(true, 1, 1,  workingKey, workingKey.length);
						}catch (Exception e) {
							writeLogFromThread(e.getMessage());
							isWorking = false;
							return;
						}
						writeLogFromThread("下载工作密钥成功");
						
						writeLogFromThread("输入PinBlock...");
						byte[] pinblock;
						try {
							pinblock = pinpad.inputPinblock(true, true, 1, 1, "1234567890123", 6, 20);
						}catch (Exception e) {
							writeLogFromThread(e.getMessage());
							isWorking = false;
							return;
						}
						String retString = "";
						for (int m=0; m<pinblock.length; m++) {
							retString += Integer.toHexString((pinblock[m]&0x000000ff)|0xffffff00).substring(6);
						}
						retString += "\n输入PinBlock成功";
						writeLogFromThread(retString);
						
						try {
							pinpad.powerOff();
						}catch (Exception e) {
							writeLogFromThread(e.getMessage());
							isWorking = false;
							return;
						}
						writeLogFromThread("Pinpad下电成功");
						isWorking = false;
					}
				}).start();
			}
		});
		button5.setOnClickListener(new View.OnClickListener() {
			public void onClick(View view) {
				if (isWorking) {
					logView.setText("");
					bluetoothThread.cancel();
					return;
				}
				logView.setText("Mifware one卡测试...\n");

				new Thread(new Runnable() {
					@Override
					public void run() {
						isWorking = true;
						byte[] uid;
							
						writeLogFromThread("\n请放置M1卡...");
						
						//1、等待卡片
						try {
							uid = mifCard.waitCard(20);
						}catch (Exception e) {
							writeLogFromThread(e.getMessage());
							isWorking = false;
							return;
						}
						String retString = "寻卡成功， UID：";
						for (int m=0; m<uid.length; m++) {
							retString += Integer.toHexString((uid[m]&0x000000ff)|0xffffff00).substring(6);
						}
						writeLogFromThread(retString);
						
						//2、认证扇区						
						byte[] key = {(byte)0xff,(byte)0xff,(byte)0xff,(byte)0xff,(byte)0xff,(byte)0xff};
						try {
							mifCard.mifareAuth(MifCard.Mifare_KeyA, 1, key);
						}catch (Exception e) {
							writeLogFromThread(e.getMessage());
							isWorking = false;
							return;
						}
						
						writeLogFromThread("扇区1认证成功");
						
						//3、读1扇区0块
						byte[] block;
						try {
							block = mifCard.mifareRead(1*4+0);
						}catch (Exception e) {
							writeLogFromThread(e.getMessage());
							isWorking = false;
							return;
						}
						retString = "1扇区0块读卡成功， Data：";
						for (int m=0; m<block.length; m++) {
							retString += Integer.toHexString((block[m]&0x000000ff)|0xffffff00).substring(6);
						}
						writeLogFromThread(retString);
						
						//4、写1扇区0块
						byte[] block2 = {1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16};
						try {
							mifCard.mifareWrite(1*4, block2);
						}catch (Exception e) {
							writeLogFromThread(e.getMessage());							
							isWorking = false;
							return;
						}
						retString = "1扇区0块写卡成功， Data：";
						for (int m=0; m<block2.length; m++) {
							retString += Integer.toHexString((block2[m]&0x000000ff)|0xffffff00).substring(6);
						}
						writeLogFromThread(retString);
						
						//5、读1扇区0块（对比）
						try {
							block = mifCard.mifareRead(1*4);
						}catch (Exception e) {
							writeLogFromThread(e.getMessage());							
							isWorking = false;
							return;
						}
						retString = "1扇区0块读卡成功， Data：";
						for (int m=0; m<block.length; m++) {
							retString += Integer.toHexString((block[m]&0x000000ff)|0xffffff00).substring(6);
						}
						writeLogFromThread(retString);
						
						//6、初始化钱包（1扇区1块，100.00元）
						byte[] moneyBloack = mifInitMoney(10000);
						try {
							mifCard.mifareWrite(1*4+1, moneyBloack);
						}catch (Exception e) {
							writeLogFromThread(e.getMessage());							
							isWorking = false;
							return;
						}
						retString = "1扇区1块写钱包初始化成功， Data：";
						for (int m=0; m<moneyBloack.length; m++) {
							retString += Integer.toHexString((moneyBloack[m]&0x000000ff)|0xffffff00).substring(6);
						}
						writeLogFromThread(retString);
						
						//7、钱包加值(10.00元）
						try {
							mifCard.mifareInc(1*4+1, 1000);
						}catch (Exception e) {
							writeLogFromThread(e.getMessage());							
							isWorking = false;
							return;
						}
						writeLogFromThread("1扇区1块钱包加值成功");
						
						
						//8、读1扇区1块钱包（验证）
						try {
							block = mifCard.mifareRead(1*4+1);
						}catch (Exception e) {
							writeLogFromThread(e.getMessage());							
							isWorking = false;
							return;
						}
						retString = "1扇区1块读卡成功， Data：";
						for (int m=0; m<block.length; m++) {
							retString += Integer.toHexString((block[m]&0x000000ff)|0xffffff00).substring(6);
						}
						writeLogFromThread(retString);
						
						//9、钱包减值(10.00元）
						try {
							mifCard.mifareDec(1*4+1, 1000);
						}catch (Exception e) {
							writeLogFromThread(e.getMessage());							
							isWorking = false;
							return;
						}
						writeLogFromThread("1扇区1块钱包减值成功");
						
						
						//10、读1扇区钱包块（验证）
						try {
							block = mifCard.mifareRead(1*4+1);
						}catch (Exception e) {
							writeLogFromThread(e.getMessage());							
							isWorking = false;
							return;
						}
						retString = "1扇区1块读卡成功， Data：";
						for (int m=0; m<block.length; m++) {
							retString += Integer.toHexString((block[m]&0x000000ff)|0xffffff00).substring(6);
						}
						writeLogFromThread(retString);
						
						//块拷贝（1块拷贝到2块）
						try {
							mifCard.mifareCopy(1*4+1, 1*4+2);
						}catch (Exception e) {
							writeLogFromThread(e.getMessage());							
							isWorking = false;
							return;
						}
						writeLogFromThread("1扇区1块拷贝到1扇区2块成功");
						
						//读1扇区2块
						try {
							block = mifCard.mifareRead(1*4+2);
						}catch (Exception e) {
							writeLogFromThread(e.getMessage());							
							isWorking = false;
							return;
						}
						retString = "1扇区2块读卡成功， Data：";
						for (int m=0; m<block.length; m++) {
							retString += Integer.toHexString((block[m]&0x000000ff)|0xffffff00).substring(6);
						}
						writeLogFromThread(retString);
												
						//恢复成全0状态						
						byte[] block3 = {0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0};
						try {
							mifCard.mifareWrite(1*4, block3);
							mifCard.mifareWrite(1*4+1, block3);
							mifCard.mifareWrite(1*4+2, block3);
						}catch (Exception e) {
							writeLogFromThread(e.getMessage());							
							isWorking = false;
							return;
						}
						writeLogFromThread("扇区1数据复原成功");						
						writeLogFromThread("\n请移除卡片...");
						
						//移除卡片
						Boolean removal;
						try {
							removal = mifCard.waitRemoval(10);
						}catch (Exception e) {
							writeLogFromThread(e.getMessage());							
							isWorking = false;
							return;
						}
						if (removal)
							writeLogFromThread("卡片移除成功");
						else 
							writeLogFromThread("卡片还未移除");
											
						isWorking = false;
						
						return;
					}
				}).start();
			}
		});
		button6.setOnClickListener(new View.OnClickListener() {
			public void onClick(View view) {
				if (isWorking) {
					logView.setText("");
					bluetoothThread.cancel();
					return;
				}
				logView.setText("射频CPU卡测试...\n");

				new Thread(new Runnable() {
					@Override
					public void run() {
						isWorking = true;
						byte[] uid;
						
						writeLogFromThread("\n请放置射频CPU卡...");
						
						try {
							uid = mifCard.waitCard(20);
						}catch (Exception e) {
							writeLogFromThread(e.getMessage());							
							isWorking = false;
							return;
						}
						String retString = "寻卡成功，UID：";
						for (int m=0; m<uid.length; m++) {
							retString += Integer.toHexString((uid[m]&0x000000ff)|0xffffff00).substring(6);
						}
						writeLogFromThread(retString);
						
						try {
							mifCard.activeCard();
						}catch (Exception e) {
							writeLogFromThread(e.getMessage());							
							isWorking = false;
							return;
						}
						writeLogFromThread("CPU卡激活成功");
						
						byte[] in = {0,(byte)0xa4,4,0,0x08,(byte)0xA0,0x00,0x00,0x03,0x33,0x01,0x01,0x01};
						byte[] out;

						try {
							out = mifCard.apdu(in, in.length);
						}catch (Exception e) {
							writeLogFromThread(e.getMessage());							
							isWorking = false;
							return;
						}
						retString = "\nAPDU返回:\n";
						for (int m=0; m < out.length; m++) {
							retString += Integer.toHexString((out[m]&0x000000ff)|0xffffff00).substring(6);
						}
						writeLogFromThread(retString);
						
						//移除卡片
						writeLogFromThread("\n请移除射频CPU卡...");
						
						Boolean removal;
						try {
							removal = mifCard.waitRemoval(10);
						}catch (Exception e) {
							writeLogFromThread(e.getMessage());							
							isWorking = false;
							return;
						}
						if (removal)
							writeLogFromThread("卡片移除成功");
						else 
							writeLogFromThread("卡片还未移除");
						
						isWorking = false;
					}
				}).start();
			}
		});

		button7.setOnClickListener(new View.OnClickListener() {
			public void onClick(View view) {
				if (isWorking) {
					logView.setText("");
					pinpad.cancel();
					return;
				}
				if (!bluetoothThread.deviceIsConnecting()) {
					logView.setText("蓝牙未连接");
					return;
				}
				logView.setText("");
				logView.append("指纹模块测试...\n");

				new Thread(new Runnable() {
					@Override
					public void run() {
						isWorking = true;
						iMateDevice imateDevice = new iMateDevice();
						writeLogFromThread("正在打开指纹模块，需要等待2~3秒");
					    int ret = imateDevice.fingerprintOpen();
					    if (ret == 99) {
					    	writeLogFromThread("不支持指纹模块!");
					    	isWorking = false;;
					        return;
					    }
					    if (ret != 0) {
					    	writeLogFromThread("打开指纹模块失败");
					    	isWorking = false;;
					        return;
					    }
					    writeLogFromThread("打开指纹模块成功");
					    
					    /* 在fingerprintOpen之后就不需要调用
					    ret = imateDevice.fingerprintLink();
					    if (ret != 0) {
					        writeLogFromThread("打开指纹模块未连接");
					        isWorking = false;;
					        return;
					    }
					     */
					    writeLogFromThread("开始采集指纹");
					    byte[] sendData = new byte[2];
					    sendData[0] = (byte)0x83;
					    sendData[1] = (byte)0x00;
					    imateDevice.fingerprintSend(sendData, 2);
					    
					    byte[] buff = new byte[512];
					    Integer len = new Integer(0);
					    boolean done = false;
					    long timeMillis = System.currentTimeMillis() + 20 * 1000L; //整体20秒等待时间
					    while (System.currentTimeMillis() < timeMillis) {
					        ret = imateDevice.fingerprintRecv(buff, len, 200); //200毫秒延时
					        if (ret == 0) {
					            if (buff[0] == (byte)0x83 && buff[1] == 0x00 ) {
					            	writeLogFromThread("采样正确");
					            	done = true;
					                break;
					            }
					            if (buff[0] == (byte)0x83 && buff[1] == 0x02 ) {
					            	writeLogFromThread("参数不符合定义");
					                break;
					            }
					            if (buff[0] == (byte)0x83 && buff[1] == 0x03 ) {
					            	writeLogFromThread("校验和错");
					            	sendData[0] = (byte)0x83;
								    sendData[1] = (byte)0x00;
								    imateDevice.fingerprintSend(sendData, 2);
					                continue;
					            }
					            if (buff[0] == (byte)0x83 && buff[1] == 0x33 ) {
					            	writeLogFromThread("采样错误");
					            	sendData[0] = (byte)0x83;
								    sendData[1] = (byte)0x00;
								    imateDevice.fingerprintSend(sendData, 2);
					                continue;
					            }
					            if (buff[0] == (byte)0x83 && buff[1] == 0x30 ) {
					            	writeLogFromThread("采样超时");
					            	sendData[0] = (byte)0x83;
								    sendData[1] = (byte)0x00;
								    imateDevice.fingerprintSend(sendData, 2);
					                continue;
					            }
					            if (buff[0] == (byte)0x84 && buff[1] == 0x31 ) {
					            	writeLogFromThread("请按下手指");
					            }
					            if (buff[0] == (byte)0x84 && buff[1] == 0x32 ) {
					            	writeLogFromThread("请抬起手指");
					            }
					        }
							try {
								Thread.sleep(10L);
							} catch (InterruptedException e) {}
					    }
					    if (done == true) {
							String fingerprintString = "";
							for (int m=0; m < len - 2; m++) {
								fingerprintString += Integer.toHexString((buff[m + 2]&0x000000ff)|0xffffff00).substring(6);
							}
						    // 返回指纹模板结构，详细参考《TS36EBG 指纹识别模块直接接口开发指南》
							len -= 2;
							writeLogFromThread("指纹特征特征模板结构数据(" + len + "):" + fingerprintString);
						    writeLogFromThread("指纹采集完成");
					    }
					    else {
						    writeLogFromThread("指纹采集等待超时");							
						}
					    imateDevice.fingerprintClose();
						isWorking = false;
					}
				}).start();
			}
		});
		button8.setOnClickListener(new View.OnClickListener() {
			public void onClick(View view) {
				if (isWorking) {
					logView.setText("");
					pinpad.cancel();
					return;
				}
				if (!bluetoothThread.deviceIsConnecting()) {
					logView.setText("蓝牙未连接");
					return;
				}
				logView.setText("");
				logView.append("SD_ICC测试...\n");

				new Thread(new Runnable() {
					@Override
					public void run() {
						isWorking = true;
						iMateDevice imateDevice = new iMateDevice();
					    int ret = imateDevice.SD_Init();
					    if (ret == 99) {
					    	writeLogFromThread("不支持SD ICC");
					    	isWorking = false;
					        return;
					    }
					    if (ret != 0) {
					    	writeLogFromThread("接口初始化失败");
					    	isWorking = false;
					        return;
					    }
					    writeLogFromThread("接口初始化成功");
					    
					    if (imateDevice.SDSCConnectDev() != 0) {
					    	writeLogFromThread("识别SD_ICC失败");
					    	isWorking = false;
					        return;
					    }
					    writeLogFromThread("识别SD_ICC成功");
					    
					    byte[] ver = new byte[100];
					    Integer len = new Integer(0);
					    if (imateDevice.SDSCGetFirmwareVer(ver, len) != 0) {
					    	writeLogFromThread("获取SD_ICC固件版本号失败");
					    	isWorking = false;
					        return;
					    }
						String logString = "";
						for (int m=0; m < len; m++) {
							logString += Integer.toHexString((ver[m]&0x000000ff)|0xffffff00).substring(6);
						}
						writeLogFromThread("固件版本号：" + logString);
					    
					    if (imateDevice.SDSCGetSDKVersion(ver, len) != 0) {
					    	writeLogFromThread("获取SD_ICC SDK版本号失败");
					    	isWorking = false;
					        return;
					    }
					    logString = new String(ver);
						writeLogFromThread("SDK版本号：" + logString);
					     
					    Integer SCIOType = new Integer(0);
					    if (imateDevice.SDSCGetSCIOType(SCIOType) != 0) {
					    	writeLogFromThread("获取获取SD_ICC IO类型失败");
					        isWorking = false;
					        return;
					    }
					    writeLogFromThread("获取SD_ICC IO类型:" + SCIOType);
					    
					    writeLogFromThread("uiSDSCTransmit测试");
					    
					    byte[] apduBuff = new byte[300];
					    byte[] cosState = new byte[2];	
					    byte[] command = new byte[5];
					    command[0] = 0x00;
					    command[1] = (byte)0x84;
					    command[2] = 0x00;
					    command[3] = 0x00;
					    command[4] = 0x08;					    
					    if (imateDevice.SDSCTransmit(command, 5, 0, apduBuff, len, cosState) != 0) {
					    	writeLogFromThread("SD ICC APDU失败");
					    	isWorking = false;
					        return;
					    }
					    logString = "随机数:";
						for (int m=0; m < len; m++) {
							logString += Integer.toHexString((apduBuff[m]&0x000000ff)|0xffffff00).substring(6);
						}
						logString += ",CosState:";
						for (int m=0; m < 2; m++) {
							logString += Integer.toHexString((cosState[m]&0x000000ff)|0xffffff00).substring(6);
						}
					    writeLogFromThread("apdu成功" + logString);
					    
					    writeLogFromThread("SDSCTransmitEx测试");
					    command[0] = 0x00;
					    command[1] = (byte)0x84;
					    command[2] = 0x00;
					    command[3] = 0x00;
					    command[4] = 0x08;
					    
					    if (imateDevice.SDSCTransmitEx(command, 5, 0, apduBuff,  len) != 0) {
					    	writeLogFromThread("SD ICC APDU Ex失败");
					    	isWorking = false;
					        return;
					    }
					    logString = "随机数+CosState:";
						for (int m=0; m < len; m++) {
							logString += Integer.toHexString((apduBuff[m]&0x000000ff)|0xffffff00).substring(6);
						}
						writeLogFromThread("apduEx成功," + logString);
					    
						byte[] atr = new byte[100];
					    if(imateDevice.SDSCResetCard(atr, len) != 0) {
					    	writeLogFromThread("SD ICC热复位失败");
					    	isWorking = false;
					        return;
					    }
					    if (imateDevice.SDSCDisconnectDev() != 0) {
					    	writeLogFromThread("关闭SD_ICC失败");
					    	isWorking = false;
					        return;
					    }
					    imateDevice.SD_DeInit();
					    writeLogFromThread("关闭SD卡电源");
					    
					    writeLogFromThread("SD ICC测试成功");
						isWorking = false;
					}
				}).start();
			}
		});
		button9.setOnClickListener(new View.OnClickListener() {
			public void onClick(View view) {
				if (isWorking) {
					logView.setText("");
					pinpad.cancel();
					return;
				}
				if (!bluetoothThread.deviceIsConnecting()) {
					logView.setText("蓝牙未连接");
					return;
				}
				logView.setText("");
				logView.append("等待事件测试...\n\n");
				logView.append("请刷卡、插入PBOC IC卡或放置射频卡...\n");
				new Thread(new Runnable() {
					@Override
					public void run() {
						isWorking = true;
						byte[] retData;
						try {
							//等待所有事件，刷卡、插卡、射频
							retData = bluetoothThread.waitEvent(0xff, 20);
						}catch (Exception e) {
							writeLogFromThread(e.getMessage());
							isWorking = false;
							return;
						}
					    String logString = "";
						switch (retData[0]) {
						case 0x01:
							byte[] track2 = new byte[37];
							byte[] track3 = new byte[104];
							
							for (int i = 0; i < 37; i++)
								track2[i] = retData[i+1];
							for (int i = 0; i < 104; i++)
								track3[i] = retData[i+38];
							
							writeLogFromThread("检测到刷卡事件:\nTrack2:" + new String(track2) + "\nTrack3:" + new String(track3));
							break;
						case 0x02:
							for (int i = 0;  i < retData.length - 1; i++) {
								logString += Integer.toHexString((retData[i+1]&0x000000ff)|0xffffff00).substring(6);
							}
							writeLogFromThread("检测到PBOC IC卡插卡事件:\nIC卡复位数据:" + logString);							
							break;
						case 0x04:
							for (int i = 0;  i < retData.length - 1; i++) {
								logString += Integer.toHexString((retData[i+1]&0x000000ff)|0xffffff00).substring(6);
							}
							writeLogFromThread("检测到射频卡事件:\n卡片系列号:" + logString);							
							break;
						}
					    writeLogFromThread("检测事件测试成功");
						isWorking = false;
					}
				}).start();
			}
		});
		button10.setOnClickListener(new View.OnClickListener() {
			public void onClick(View view) {
				if (isWorking) {
					logView.setText("");
					pinpad.cancel();
					return;
				}
				if (!bluetoothThread.deviceIsConnecting()) {
					logView.setText("蓝牙未连接");
					return;
				}
				logView.setText("");
				logView.append("部件测试...\n");
				new Thread(new Runnable() {
					@Override
					public void run() {
						isWorking = true;
						int retCode;
						try {
							//检测所有的部件
							retCode = bluetoothThread.deviceTest(0xff);
						}catch (Exception e) {
							writeLogFromThread(e.getMessage());
							isWorking = false;
							return;
						}
					    String logString = "";
					    if (retCode == 0) {
					    	writeLogFromThread("设备所有模块测试通过");
					    }
					    else {
					    	if ((retCode & 0x01) != 0) {
						    	writeLogFromThread("检测二代证模块失败");
					    	}
					    	if ((retCode & 0x02) != 0) {
						    	writeLogFromThread("检测射频模块失败");						    		
					    	}
					    	if ((retCode & 0x04) != 0) {
						    	writeLogFromThread("检测指纹模块失败");						    								    		
					    	}
					    	if ((retCode & 0x08) != 0) {
						    	writeLogFromThread("检测 SD卡模块失败");						    								    		
					    	}
					    }
					    writeLogFromThread("部件测试完成");
						isWorking = false;
					}
				}).start();
			}
		});
	}
	private void writeLogFromThread(String logString)
	{
		Message message = new Message();
		message.obj = logString;
		logViewAppendHandler.sendMessage(message);
	}
	
	@SuppressLint("HandlerLeak")
	class  LogViewAppendHandler extends Handler {
		@Override
	    public void handleMessage(Message message) {
			super.handleMessage(message);
			logView.append((String)message.obj + "\n");
	    }
	}
	
	@Override
	public void onResume() {
		super.onResume();

		// App进入前台，建立蓝牙连接
		if (bluetoothThread != null)
			bluetoothThread.resumeThread();
		/*
		if (!threadStarted) {
			BluetoothAdapter bluetoothAdapter = BluetoothAdapter
					.getDefaultAdapter();
			bluetoothThread = new BluetoothThread(bluetoothAdapter);
			bluetoothThread.start();
			threadStarted = true;
			System.out.println(bluetoothThread);
		}
		*/
	}

	@Override
	public void onPause() {
		super.onPause();
		// App进入后台，关闭蓝牙连接，释放资源
		if (bluetoothThread != null)
			bluetoothThread.pauseThread();
		/*
		if (threadStarted) {
			bluetoothThread.exitThread();
			bluetoothThread = null;
			threadStarted = false;
		}
		*/
	}
	
	@Override
	public void onDestroy() {
		super.onDestroy();
		// App退出后，关闭蓝牙连接，释放资源
		if (bluetoothThread != null)
			bluetoothThread.pauseThread();
		/*
		if (threadStarted) {
			bluetoothThread.exitThread();
			bluetoothThread = null;
			threadStarted = false;
		}
		*/
	}

	@Override
	public boolean onCreateOptionsMenu(Menu menu) {
		// Inflate the menu; this adds items to the action bar if it is present.
		getMenuInflater().inflate(R.menu.main, menu);
		return true;
	}
	public byte[] mifInitMoney(long balence)
	{

		byte[] sTmp = new byte[16];

		sTmp[0] = (byte)(balence & 0xff);
		sTmp[1] = (byte)((balence>>8) & 0xff);
		sTmp[2] = (byte)((balence>>16) & 0xff);
		sTmp[3] = (byte)((balence>>24) & 0xff);
		sTmp[4] = (byte)~sTmp[0];
		sTmp[5] = (byte)~sTmp[1];
		sTmp[6] = (byte)~sTmp[2];
		sTmp[7] = (byte)~sTmp[3];
		sTmp[8] = sTmp[0];
		sTmp[9] = sTmp[1];
		sTmp[10] = sTmp[2];
		sTmp[11] = sTmp[3];
		    
		sTmp[12] = 0x01;
		sTmp[13] = (byte)0xfe;
		sTmp[14] = 0x01;
		sTmp[15] = (byte)0xfe;
		
		return sTmp;
	}
}


