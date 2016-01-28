#include "stdafx.h"
#include "util.h"


//�����ļ�������޸�ʱ��
void SetFileModifyTime(const char* filename,const char* modifytime)
{
	HANDLE hFind = CreateFile( filename, GENERIC_WRITE, 0, NULL, OPEN_EXISTING, FILE_ATTRIBUTE_NORMAL, NULL );
	if(hFind == INVALID_HANDLE_VALUE ) //������
	{
		return;
	}

	FILETIME ftWrite;

	SYSTEMTIME stLocal,stUTC;
	memset(&stLocal,0,sizeof(SYSTEMTIME));

	char tmp[10];
	const char* tt = strchr(modifytime,' ');
	const char* pre = modifytime;

	memcpy(tmp,pre,tt-pre);
	tmp[tt-pre] = NULL;
	stLocal.wYear = atoi(tmp);

	pre = tt+1;
	tt = strchr(tt+1,' ');
	memcpy(tmp,pre,tt-pre);
	tmp[tt-pre] = NULL;
	stLocal.wMonth = atoi(tmp);

	pre = tt+1;
	tt = strchr(tt+1,' ');
	memcpy(tmp,pre,tt-pre);
	tmp[tt-pre] = NULL;
	stLocal.wDay = atoi(tmp);

	pre = tt+1;
	tt = strchr(tt+1,' ');
	memcpy(tmp,pre,tt-pre);
	tmp[tt-pre] = NULL;
	stLocal.wHour = atoi(tmp);

	pre = tt+1;
	tt = modifytime + strlen(modifytime);
	memcpy(tmp,pre,tt-pre);
	tmp[tt-pre] = NULL;
	stLocal.wMinute = atoi(tmp);

	TzSpecificLocalTimeToSystemTime (NULL, &stLocal, &stUTC);  
	SystemTimeToFileTime(&stUTC, &ftWrite);
	SetFileTime(hFind,NULL,NULL,&ftWrite);
	CloseHandle(hFind);
}

//����Ŀ¼
bool CreateMultipleDirectroy(char*name,int len)
{
	int count = len-1;
	char tmp;
	WIN32_FIND_DATA FindFileData;
	HANDLE hFind;

	//����Ŀ¼
	count= 0;
	for(;count<len;count++)
	{
		if(name[count] == '/' || name[count] == '\\')
		{
			tmp = name[count];
			name[count] = '\0';
			hFind = FindFirstFile(name,&FindFileData);
			if(hFind == INVALID_HANDLE_VALUE) //û���ҵ�
			{
				CreateDirectory(name,NULL);
			}
			else
			{
				FindClose(hFind);
			}
			name[count] = tmp;
		}
	}
	return true;
}

void UTF8ToGB(char* szOut)
{
	WCHAR *strSrc;
	char *szRes;
	//����
	int i = MultiByteToWideChar(CP_UTF8, 0, szOut, -1, NULL, 0);
	strSrc = new WCHAR[i+1];
	MultiByteToWideChar(CP_UTF8, 0, szOut, -1, strSrc, i);

	//����
	i = WideCharToMultiByte(CP_ACP, 0, strSrc, -1, NULL, 0, NULL, NULL);
	szRes = new TCHAR[i+1];
	WideCharToMultiByte(CP_ACP, 0, strSrc, -1, szRes, i, NULL, NULL);

	memcpy(szOut,szRes,strlen(szRes)+1);

	delete []strSrc;
	delete []szRes;
}

//ִ��ĳ���ļ�
void RunFile(const char *path)
{
	if(path == NULL)
	{
		return;
	}
	PROCESS_INFORMATION pi;
	STARTUPINFOA si;

	si.dwFlags|=STARTF_USESHOWWINDOW;
	si.wShowWindow=SW_SHOWNORMAL;

	memset(&si,0,sizeof(si));
	si.cb= sizeof(si);

	const char *pch = strrchr(path,'\\');
	if(pch == NULL)
	{
		pch = strrchr(path,'/');
	}


	char curEnviron[MAX_PATH];
	memset(curEnviron,0,MAX_PATH*sizeof(char));
	if(pch != NULL)
	{
		memcpy(curEnviron,path,pch-path);
		curEnviron[pch-path] = NULL;
	}

	//1:����exeû��Ӱ�� 2 ����bat֮���dos�������޴�������
	if(!CreateProcessA(NULL, (LPSTR)path, NULL, NULL, false, CREATE_NO_WINDOW, NULL,curEnviron,&si,&pi)) {

	}
}

//�ر�ִ�е�exe�ļ�
void closerunningexe(const char *exepath)
{
	if(exepath == NULL)
	{
		return;
	}
	const char *spos = exepath;

	const char *tmp = strrchr(exepath,'/');
	if(tmp == NULL)
	{
		tmp = strrchr(exepath,'\\');
		if(tmp != NULL)
		{
			spos = ++tmp;
		}
	}
	else
	{
		spos = ++tmp;
	}

	HANDLE hProcessSnap = NULL;
	PROCESSENTRY32 process32;
	hProcessSnap = CreateToolhelp32Snapshot(TH32CS_SNAPPROCESS, 0);
	process32.dwSize = sizeof(PROCESSENTRY32); 
	BOOL b = Process32First(hProcessSnap, &process32);
	HANDLE hProcess;
	while (b)
	{
		if (strcmp(spos,process32.szExeFile) == 0)
		{
			DWORD dwId = process32.th32ProcessID;

			hProcess = OpenProcess(PROCESS_TERMINATE, FALSE, dwId);
			if ( NULL == hProcess )
			{
				return;
			}

			TerminateProcess(hProcess, 0);
			CloseHandle(hProcess);
		}
		b = Process32Next(hProcessSnap,&process32);
	}
	return;
}

//�Ƚ������ַ����Ĵ�С,Ҫ������ת��Ϊ����
//����ֵ -1:ǰ��Ĵ�
//0:���
//1:����Ĵ�
int CompareNumber(const char *fir,const char* sec,char sep)
{
	if((fir == NULL) && (sec == NULL))
	{
		return 0;
	}
	if(fir == NULL)
	{
		return 1;
	}
	if(sec == NULL)
	{
		return -1;
	}

	char f_v[10];
	char s_v[10];
	const char* f_i = fir;
	const char* s_i = sec;
	const char* l_f_i = f_i;
	const char* s_f_i = s_i;

	int f_com = 0;
	int s_com = 0;
	int end = 5;

	do{
		f_i = strchr(f_i,sep);
		s_i = strchr(s_i,sep);

		//���Ҫ�������һ���ж�,������߲����,��ô�������return����
		//����ȽϽ���Ծ������,��ô����ֵ�������end��ʾ
		if((f_i == NULL) && (s_i == NULL))
		{
			end = 0;
			f_i = l_f_i + strlen(l_f_i); //��λ�����һ������
			s_i = s_f_i + strlen(s_f_i);
		}
		else if(f_i == NULL)
		{
			f_i = l_f_i + strlen(l_f_i); //��λ�����һ������
			end = 1;
		}
		else if(s_i == NULL)
		{
			s_i = s_f_i + strlen(s_f_i);
			end =-1;
		}

		memcpy(f_v,l_f_i,f_i-l_f_i+1);
		f_v[f_i-l_f_i] = NULL;
		memcpy(s_v,s_f_i,s_i-s_f_i+1);
		s_v[s_i-s_f_i] = NULL;

		f_com = atoi(f_v);
		s_com = atoi(s_v);
		if(f_com > s_com)
		{
			return -1;
		}
		else if(f_com < s_com)
		{
			return 1;
		}
		l_f_i = f_i+1;
		s_f_i = s_i+1;
		f_i++;
		s_i++;
	}while(end == 5);
	return end;
}

//��ȡ�ļ��޸�ʱ��
void GetFileModifyTime(const char *filename, string &modifytime)
{
	modifytime.clear();
	WIN32_FIND_DATA ffd ;  
	HANDLE hFind = FindFirstFile(filename,&ffd);  
	if(hFind == INVALID_HANDLE_VALUE ) //������
	{
		modifytime.append("0000 00 00 00 00"); //������,��Ϊ��Сֵ
		return;
	}
	SYSTEMTIME stUTC, stLocal;  
	FileTimeToSystemTime(&(ffd.ftLastWriteTime), &stUTC);  
	SystemTimeToTzSpecificLocalTime(NULL, &stUTC, &stLocal);  

	char tmp[5];
	itoa(stLocal.wYear,tmp,10);
	modifytime.append(tmp);
	modifytime.append(" ");

	itoa(stLocal.wMonth,tmp,10);
	modifytime.append(tmp);
	modifytime.append(" ");

	itoa(stLocal.wDay,tmp,10);
	modifytime.append(tmp);
	modifytime.append(" ");

	itoa(stLocal.wHour,tmp,10);
	modifytime.append(tmp);
	modifytime.append(" ");

	itoa(stLocal.wMinute,tmp,10);
	modifytime.append(tmp);
}

//����ļ��汾��
void GetFileVersion(const char *filename, string &version)
{
	version.empty();

	int   iVerInfoSize;  
	char   *pBuf;  
	VS_FIXEDFILEINFO   *pVsInfo;  
	unsigned   int   iFileInfoSize   =   sizeof(   VS_FIXEDFILEINFO   );  


	iVerInfoSize   =   GetFileVersionInfoSize(filename,NULL);   

	if(iVerInfoSize!= 0)  
	{     
		pBuf   =   new   char[iVerInfoSize];  
		if(GetFileVersionInfo(filename,0,   iVerInfoSize,   pBuf   )   )     
		{     
			if(VerQueryValue(pBuf,"\\",(void   **)&pVsInfo,&iFileInfoSize))     
			{
				char tmp[5];
				itoa(HIWORD(pVsInfo->dwFileVersionMS),tmp,10);
				version.append(tmp);
				version.append(",");

				itoa(LOWORD(pVsInfo->dwFileVersionMS),tmp,10);
				version.append(tmp);
				version.append(",");

				itoa(HIWORD(pVsInfo->dwFileVersionLS),tmp,10);
				version.append(tmp);
				version.append(",");

				itoa(LOWORD(pVsInfo->dwFileVersionLS),tmp,10);
				version.append(tmp);
				version.append(",");

			}     
		}     
		delete   pBuf;     
	}
	else
	{
		int error = GetLastError();
		char e[10];
		itoa(error,e,10);
		version.append("0.0.0.0"); //����ļ�������,��汾��Ϊ0.0.0.0
	}
}

//�ر�ϵͳ
void shut_down(DWORD verInfo,int type)
{
	try
	{
		//�����ȡ����ϵͳdwMajorVersionֵ���ڵ���5,��ʾΪNT���ϲ���ϵͳ,����������Ȩ��
		if(verInfo>=5)
		{
			HANDLE ToHandle;
			TOKEN_PRIVILEGES tkp;
			//�򿪱����̷�������
			if(OpenProcessToken(GetCurrentProcess(),TOKEN_ADJUST_PRIVILEGES|TOKEN_QUERY,&ToHandle))
			{
				//�޸ı�����Ȩ��
				LookupPrivilegeValue(NULL,SE_SHUTDOWN_NAME,&tkp.Privileges[0].Luid);
				tkp.PrivilegeCount=1;
				tkp.Privileges[0].Attributes=SE_PRIVILEGE_ENABLED;
				//֪ͨϵͳ���޸�
				AdjustTokenPrivileges(ToHandle,FALSE,&tkp,0,(PTOKEN_PRIVILEGES)NULL,0);
			}
		}

		if(type == 1)
		{
			ExitWindowsEx(EWX_SHUTDOWN|EWX_FORCE,0);
		}
		else if(type ==2)
		{
			ExitWindowsEx(EWX_REBOOT|EWX_FORCE,0);
		}
		else if(type == 3)
		{
			ExitWindowsEx(EWX_LOGOFF|EWX_FORCE,0);
		}

	}
	catch(...)
	{
		return;
	}
}

DWORD GetVerInfo()
{
	OSVERSIONINFO osver={sizeof(OSVERSIONINFO)};
	GetVersionEx(&osver);
	return osver.dwMajorVersion;
}


void shutdown()
{
	// TODO: �ڴ����ʵ�ִ���

	shut_down(GetVerInfo(),1);
}


void reboot()
{
	// TODO: �ڴ����ʵ�ִ���
	shut_down(GetVerInfo(),2);
}


//ע��һ���û�
void DeloginUser(const char* user){
	WTS_SESSION_INFO  *wsi = NULL;
	DWORD count = 0;
	BOOL RetVal ;

	//���Դ���
	int retrycount = 0;

retry:

	RetVal = WTSEnumerateSessions(WTS_CURRENT_SERVER_HANDLE,0,1,&wsi,&count);
	if(RetVal == TRUE){
		for(int i=0;i<count;i++){
			LPTSTR	pBuffer = NULL;
			DWORD	dwBufferLen;

			BOOL bRes = WTSQuerySessionInformation(WTS_CURRENT_SERVER_HANDLE, wsi[i].SessionId, WTSUserName, &pBuffer, &dwBufferLen);

			if (bRes == FALSE)
				continue;

			if(strncmp(user,pBuffer,4) == 0){
				WTSFreeMemory(pBuffer);
				WTSFreeMemory(wsi);
				bool res = WTSLogoffSession(WTS_CURRENT_SERVER_HANDLE,wsi[i].SessionId,TRUE);
				if(res == 0){
					char err[10];
					DWORD e = GetLastError();
					itoa(e,err,10);
					
				}
				else{
					
				}
				return;
			}
			WTSFreeMemory(pBuffer);

		}

		if(retrycount ++ < 10){
			Sleep(1000);
			goto retry;	
		}
		else{
			
		}
	}
}

