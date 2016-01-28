// idcardreader.cpp : Defines the exported functions for the DLL application.
//

#include "stdafx.h"
#include "iccardreader.h"
#include <WinSCard.h>
#include <string>

using namespace std;

#define SMART_FORMAT_ERROR 2

char* errorfunc(long rc)
{
	char * errdesc = NULL;
	switch(rc)
	{
	case ERROR_BROKEN_PIPE:
		{
			errdesc = "ERROR_BROKEN_PIPE";
			break;
		}
	case SCARD_E_BAD_SEEK:
		{
			errdesc = "SCARD_E_BAD_SEEK";
			break;
		}
	case SCARD_E_CANCELLED:
		{
			errdesc = "SCARD_E_CANCELLED";
			break;
		}
	case SCARD_E_CANT_DISPOSE :
		{
			errdesc = "SCARD_E_CANT_DISPOSE";
			break;
		}
	case SCARD_E_CARD_UNSUPPORTED:
		{
			errdesc = "SCARD_E_CARD_UNSUPPORTED";
			break;
		}
	case SCARD_E_CERTIFICATE_UNAVAILABLE:
		{
			errdesc = "SCARD_E_CERTIFICATE_UNAVAILABLE";
			break;
		}
	case SCARD_E_COMM_DATA_LOST:
		{
			errdesc = "SCARD_E_COMM_DATA_LOST";
			break;
		}
	case SCARD_E_DIR_NOT_FOUND :
		{
			errdesc = "SCARD_E_DIR_NOT_FOUND";
			break;
		}
	case SCARD_E_DUPLICATE_READER:
		{
			errdesc = "SCARD_E_DUPLICATE_READER";
			break;
		}
	case SCARD_E_FILE_NOT_FOUND:
		{
			errdesc = "SCARD_E_FILE_NOT_FOUND";
			break;
		}
	case SCARD_E_ICC_CREATEORDER:
		{
			errdesc = "SCARD_E_ICC_CREATEORDER";
			break;
		}
	case SCARD_E_ICC_INSTALLATION:
		{
			errdesc = "SCARD_E_ICC_INSTALLATION";
			break;
		}
	case SCARD_E_INSUFFICIENT_BUFFER:
		{
			errdesc = "SCARD_E_INSUFFICIENT_BUFFER";
			break;
		}
	case SCARD_E_INVALID_ATR:
		{
			errdesc = "SCARD_E_INVALID_ATR";
			break;
		}
	case SCARD_E_INVALID_CHV :
		{
			errdesc = "SCARD_E_INVALID_CHV";
			break;
		}
	case SCARD_E_INVALID_HANDLE:
		{
			errdesc = "SCARD_E_INVALID_HANDLE";
			break;
		}
	case SCARD_E_INVALID_PARAMETER :
		{
			errdesc = "SCARD_E_INVALID_PARAMETER";
			break;
		}
	case SCARD_E_INVALID_TARGET:
		{
			errdesc = "SCARD_E_INVALID_TARGET";
			break;
		}
	case SCARD_E_INVALID_VALUE:
		{
			errdesc = "SCARD_E_INVALID_VALUE";
			break;
		}
	case SCARD_E_NO_ACCESS :
		{
			errdesc = "SCARD_E_NO_ACCESS";
			break;
		}
	case SCARD_E_NO_DIR :
		{
			errdesc = "SCARD_E_NO_DIR";
			break;
		}
	case SCARD_E_NO_FILE:
		{
			errdesc = "SCARD_E_NO_FILE";
			break;
		}

	case SCARD_E_NO_MEMORY:
		{
			errdesc = "SCARD_E_NO_MEMORY";
			break;
		}
	case SCARD_E_NO_READERS_AVAILABLE :
		{
			errdesc = "SCARD_E_NO_READERS_AVAILABLE";
			break;
		}
	case SCARD_E_NO_SERVICE :
		{
			errdesc = "SCARD_E_NO_SERVICE";
			break;
		}
	case SCARD_E_NO_SMARTCARD:
		{
			errdesc = "SCARD_E_NO_SMARTCARD";
			break;
		}
	case SCARD_E_NO_SUCH_CERTIFICATE :
		{
			errdesc = "SCARD_E_NO_SUCH_CERTIFICATE";
			break;
		}
	case SCARD_E_NOT_READY:
		{
			errdesc = "SCARD_E_NOT_READY";
			break;
		}
	case SCARD_E_NOT_TRANSACTED:
		{
			errdesc = "SCARD_E_NOT_TRANSACTED";
			break;
		}
	case SCARD_E_PCI_TOO_SMALL:
		{
			errdesc = "SCARD_E_PCI_TOO_SMALL";
			break;
		}
	case SCARD_E_PROTO_MISMATCH :
		{
			errdesc = "SCARD_E_PROTO_MISMATCH";
			break;
		}
	case SCARD_E_READER_UNAVAILABLE:
		{
			errdesc = "SCARD_E_READER_UNAVAILABLE";
			break;
		}
	case SCARD_E_READER_UNSUPPORTED:
		{
			errdesc = "SCARD_E_READER_UNSUPPORTED";
			break;
		}
	case SCARD_E_SERVICE_STOPPED:
		{
			errdesc = "SCARD_E_SERVICE_STOPPED";
			break;
		}
	case SCARD_E_SHARING_VIOLATION:
		{
			errdesc = "SCARD_E_SHARING_VIOLATION";
			break;
		}
	case SCARD_E_SYSTEM_CANCELLED:
		{
			errdesc = "SCARD_E_SYSTEM_CANCELLED";
			break;
		}
	case SCARD_E_TIMEOUT:
		{
			errdesc = "SCARD_E_TIMEOUT";
			break;
		}
	case SCARD_E_UNEXPECTED:
		{
			errdesc = "SCARD_E_UNEXPECTED";
			break;
		}
	case SCARD_E_UNKNOWN_CARD:
		{
			errdesc = "SCARD_E_UNKNOWN_CARD";
			break;
		}
	case SCARD_E_UNKNOWN_READER:
		{
			errdesc = "SCARD_E_UNKNOWN_READER";
			break;
		}
	case SCARD_E_UNKNOWN_RES_MNG:
		{
			errdesc = "SCARD_E_UNKNOWN_RES_MNG";
			break;
		}
	case SCARD_E_UNSUPPORTED_FEATURE:
		{
			errdesc = "SCARD_E_UNSUPPORTED_FEATURE";
			break;
		}
	case SCARD_E_WRITE_TOO_MANY :
		{
			errdesc = "SCARD_E_WRITE_TOO_MANY";
			break;
		}
	case SCARD_F_COMM_ERROR:
		{
			errdesc = "SCARD_F_COMM_ERROR";
			break;
		}
	case SCARD_F_INTERNAL_ERROR :
		{
			errdesc = "SCARD_F_INTERNAL_ERROR";
			break;
		}
	case SCARD_F_UNKNOWN_ERROR:
		{
			errdesc = "SCARD_F_UNKNOWN_ERROR";
			break;
		}
	case SCARD_F_WAITED_TOO_LONG:
		{
			errdesc = "SCARD_F_WAITED_TOO_LONG";
			break;
		}
	case SCARD_P_SHUTDOWN:
		{
			errdesc = "SCARD_P_SHUTDOWN";
			break;
		}
	case SCARD_S_SUCCESS:
		{
			errdesc = "SCARD_S_SUCCESS";
			break;
		}
	case SCARD_W_CANCELLED_BY_USER:
		{
			errdesc = "SCARD_W_CANCELLED_BY_USER";
			break;
		}
	case SCARD_W_CHV_BLOCKED:
		{
			errdesc = "SCARD_W_CHV_BLOCKED";
			break;
		}
	case SCARD_W_EOF :
		{
			errdesc = "SCARD_W_EOF";
			break;
		}
	case SCARD_W_REMOVED_CARD:
		{
			errdesc = "SCARD_W_REMOVED_CARD";
			break;
		}
	case SCARD_W_RESET_CARD:
		{
			errdesc = "SCARD_W_RESET_CARD";
			break;
		}
	case SCARD_W_SECURITY_VIOLATION:
		{
			errdesc = "SCARD_W_SECURITY_VIOLATION";
			break;
		}
	case SCARD_W_UNPOWERED_CARD:
		{
			errdesc = "SCARD_W_UNPOWERED_CARD";
			break;
		}
	case SCARD_W_UNRESPONSIVE_CARD:
		{
			errdesc = "SCARD_W_UNRESPONSIVE_CARD";
			break;
		}
	case SCARD_W_UNSUPPORTED_CARD:
		{
			errdesc = "SCARD_W_UNSUPPORTED_CARD";
			break;
		}
	case SCARD_W_WRONG_CHV:
		{
			errdesc = "SCARD_W_WRONG_CHV";
			break;
		}

	case 0x000000007a :{
		errdesc = "INSUFFICIENT BUFFER";
		break;
					   }

	default:
		{
			char tmp[10];
			itoa(rc,tmp,10);
			OutputDebugString(tmp);
			errdesc = "unknown error";
			break;
		}
	}
	return errdesc;
}

#define DEBUGARGS __LINE__
#define DEBUG_PRINT(rc,errmsg) (SetErrDesc(rc,errmsg,DEBUGARGS))

void SetErrDesc(long rc,char* errmsg,int line){
	sprintf(errmsg,"(%d) %s",line,errorfunc(rc));
}

LPCSCARD_IO_REQUEST getProtocol(DWORD ap){
	if(ap == SCARD_PROTOCOL_T0){
		return SCARD_PCI_T0;
	}else if(ap == SCARD_PROTOCOL_T1){
		return SCARD_PCI_T1;
	}
}

UPDATEDLL_API int ReadICCard(int timeout,char* bkcard,int *msglen,char* errmsg){
	char *err_desc = NULL;
	LPTSTR            szReaders, szRdr;
	DWORD             cchReaders = SCARD_AUTOALLOCATE;
	DWORD             dwI, dwRdrCount;
	SCARD_READERSTATE rgscState[MAXIMUM_SMARTCARD_READERS];
	TCHAR             szCard[MAX_PATH];
	int dwProtocol;
	DWORD dwAP;
	SCARDCONTEXT hSC; //智能卡句柄
	string t;

	SCARDHANDLE     hCardHandle;

	long IReturn = 0;
	IReturn=SCardEstablishContext(SCARD_SCOPE_USER,NULL,NULL,&hSC);
	if ( SCARD_S_SUCCESS != IReturn )
	{
		DEBUG_PRINT(IReturn,errmsg); 
		return IReturn;
	}
	else
	{
		IReturn = SCardListReaders(hSC,
			NULL,
			(LPTSTR)&szReaders,
			&cchReaders );

		if ( SCARD_S_SUCCESS != IReturn )
		{
			DEBUG_PRINT(IReturn,errmsg); 
			SCardReleaseContext(hSC);
			return IReturn;
		}

		// Place the readers into the state array.
		szRdr = szReaders;
		for ( dwI = 0; dwI < MAXIMUM_SMARTCARD_READERS; dwI++ )
		{
			if ( '\0' == *szRdr )
				break;
			rgscState[dwI].szReader = szRdr;
			rgscState[dwI].dwCurrentState = SCARD_STATE_UNAWARE;
			szRdr = szRdr + strlen(szRdr) + 1;
			OutputDebugString(szRdr);
		}
		dwRdrCount = dwI;

		// If any readers are available, proceed.
		if ( 0 != dwRdrCount )
		{
			for (;;)
			{ 

				// Card not found yet; wait until there is a change.
				IReturn = SCardGetStatusChange(hSC,
					INFINITE, // infinite wait
					rgscState,
					dwRdrCount );
				if ( SCARD_S_SUCCESS != IReturn )
				{
					DEBUG_PRINT(IReturn,errmsg); 
					SCardReleaseContext(hSC);
					return IReturn;
				}
				else
				{
					for ( int cur_reader = 0; cur_reader < dwRdrCount; cur_reader++ )
					{
						//已经获取了卡片,需要获取数据
						if((rgscState[cur_reader].dwEventState & SCARD_STATE_PRESENT) != 0)
						{
							//如果读卡器中没有卡则会返回错误
							IReturn = SCardConnect( hSC, 
								(LPCTSTR)rgscState[cur_reader].szReader,
								SCARD_SHARE_SHARED,
								SCARD_PROTOCOL_T1 | SCARD_PROTOCOL_T0, 
								&hCardHandle,
								&dwAP );
							if ( SCARD_S_SUCCESS != IReturn )
							{
								DEBUG_PRINT(IReturn,errmsg); 
								SCardReleaseContext(hSC);
								return IReturn;
							}

							//此处加入设置要发送的命令的代码．．． 
							BYTE pbyReceived[200];
							DWORD dwRecLength = 200;
							BYTE    select_mf[]={0x00, 0xA4, 0x04, 0x00, 0x0E, 0x31,0x50,0x41,0x59,0x2E,0x53,0x59,0x53,0x2E,0x44,0x44,0x46,0x30,0x31};
							IReturn = SCardTransmit(hCardHandle,getProtocol(dwAP),(LPCBYTE)select_mf,19,NULL,pbyReceived,&dwRecLength);
							if ( SCARD_S_SUCCESS != IReturn )
							{
								DEBUG_PRINT(IReturn,errmsg); 
								SCardReleaseContext(hSC);
								return IReturn;
							}

							BYTE sw1;
							BYTE sw2;
							sw1 = pbyReceived[dwRecLength-2];
							sw2 = pbyReceived[dwRecLength-1];

							if(sw1 == 144 && sw2 == 0)
							{
								BYTE    select_mf2[]={0x00, 0xB2, 0x01, 0x0C, 0x00};
								dwRecLength = 200;
								IReturn = SCardTransmit(hCardHandle,getProtocol(dwAP),(LPCBYTE)select_mf2,5,NULL,pbyReceived,&dwRecLength);
								if ( SCARD_S_SUCCESS != IReturn )
								{
									DEBUG_PRINT(IReturn,errmsg); 
									SCardReleaseContext(hSC);
									return IReturn;
								}
								char dbg[10];
								sw1 = pbyReceived[dwRecLength-2];
								sw2 = pbyReceived[dwRecLength-1];

								if(sw1 == 144 && sw2 == 0)
								{
									select_mf2[4] = sw2;
									dwRecLength = 200;
									IReturn = SCardTransmit(hCardHandle,getProtocol(dwAP),(LPCBYTE)select_mf2,5,NULL,pbyReceived,&dwRecLength);

									if ( SCARD_S_SUCCESS != IReturn )
									{
										DEBUG_PRINT(IReturn,errmsg); 
										SCardReleaseContext(hSC);
										return IReturn;
									}
									
									sw1 = pbyReceived[dwRecLength-2];
									sw2 = pbyReceived[dwRecLength-1];

									if(sw1 != 144 || sw2 != 0)
									{
										char *err = "reteive memory location error";
										memcpy(errmsg,err,strlen(err));
										errmsg[strlen(err)]=0;
										SCardReleaseContext(hSC);
										return SMART_FORMAT_ERROR;
									}

									dwRecLength = sizeof(pbyReceived);
									BYTE select_mf3[]={0x00, 0xA4, 0x04, 0x00, 0x08,0xA0, 0x00, 0x00,0x03,0x33,0x01,0x01,0x01};
									IReturn = SCardTransmit(hCardHandle,getProtocol(dwAP),(LPCBYTE)select_mf3,13,NULL,pbyReceived,&dwRecLength);
									if ( SCARD_S_SUCCESS != IReturn )
									{
										DEBUG_PRINT(IReturn,errmsg); 
										SCardReleaseContext(hSC);
										return IReturn;
									}

									sw1 = pbyReceived[dwRecLength-2];
									sw2 = pbyReceived[dwRecLength-1];

									if(sw1 != 144 || sw2 != 0)
									{
										char *err = "reteive application error";
										memcpy(errmsg,err,strlen(err));
										errmsg[strlen(err)]=0;
										return SMART_FORMAT_ERROR;
									}

									BYTE select_mf4[]={0x00, 0xB2, 0x01, 0x0C, 0x00};
									dwRecLength = sizeof(pbyReceived);
									IReturn = SCardTransmit(hCardHandle,getProtocol(dwAP),(LPCBYTE)select_mf4,5,NULL,pbyReceived,&dwRecLength);

									if ( SCARD_S_SUCCESS != IReturn )
									{
										DEBUG_PRINT(IReturn,errmsg); 
										SCardReleaseContext(hSC);
										return IReturn;
									}

									sw1 = pbyReceived[dwRecLength-2];
									sw2 = pbyReceived[dwRecLength-1];

									if(sw1 != 144 || sw2 != 0)
									{
										char *err = "reteive msf error";
										memcpy(errmsg,err,strlen(err));
										errmsg[strlen(err)]=0;
										return SMART_FORMAT_ERROR;
									}

									//读取最后的数据
									int  i = 0;
									dwRecLength = sizeof(pbyReceived);
									while(i < dwRecLength)
									{
										if(pbyReceived[i] == 0x57)
										{
											break;
										}
										i++;
									}
									if(i == dwRecLength)
									{
										char *err = "reteive card no error";
										memcpy(errmsg,err,strlen(err));
										errmsg[strlen(err)]=0;
										SCardReleaseContext(hSC);
										return SMART_FORMAT_ERROR;
									}

									byte len = pbyReceived[i+1];
									char res[30];
									memcpy(res,pbyReceived+ i+2,len);
									res[len] = 0;
									i = 0;
									t.clear();

									BYTE raw = (BYTE)res[i];
									BYTE high = raw >> 4;
									BYTE low = raw & 0x0F;
									if(high != 0)
									{
										t.append(1,high+48);
									}
									if((low !=0 )|| (high!=0 && low == 0))
									{
										t.append(1,low+48);
									}
									i++;
									while(i < len)
									{
										raw = (BYTE)res[i];
										high = raw >>4;
										low = raw & 0x0F;
										if(high == '='){
											break;
										}
										t.append(1,high+48);
										if(low == '='){
											break;
										}
										t.append(1,low+48);
										i++;
									}
									t.erase(len);
									memcpy(bkcard,t.c_str(),t.size());
									bkcard[t.size()] = 0;
									*msglen = t.size();

									SCardDisconnect(hCardHandle,SCARD_RESET_CARD);
									SCardReleaseContext(hSC);

									return 0 ; //成功

								}
								else
								{
									char *err = "reteive directory error";
									memcpy(errmsg,err,strlen(err));
									errmsg[strlen(err)]=0;
									SCardReleaseContext(hSC);
									return SMART_FORMAT_ERROR;
								}

							}
							else
							{
								char *err = "reteive 1pay.sys.ddf01 error";
								memcpy(errmsg,err,strlen(err));
								errmsg[strlen(err)]=0;
								SCardReleaseContext(hSC);
								return SMART_FORMAT_ERROR;
							}
						
						}
						rgscState[cur_reader].dwCurrentState = rgscState[cur_reader].dwEventState;
					}
				}
			}  
		}else{
			char *err = "no smart card reader found";
			memcpy(errmsg,err,strlen(err));
			errmsg[strlen(err)]=0;
			SCardReleaseContext(hSC);
			return SMART_FORMAT_ERROR;
		}
	}

}
