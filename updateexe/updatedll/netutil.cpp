#include "stdafx.h"
#include "netutil.h"


//获得HTML内容
DWORD GetHTMLContent(char *url,char* buf ,int len){
	try
	{
		CInternetSession mySession(NULL, 0);       
		CString m_strHTML, myData;    
		CHttpFile *myHttpFile = NULL; 

		myHttpFile=(CHttpFile *)mySession.OpenURL(url);      
		while(myHttpFile->ReadString(myData))      
		{      
			m_strHTML= m_strHTML+ TEXT(" ") + myData;      
		}      
		myHttpFile->Close();      
		mySession.Close();
		m_strHTML.Trim();
		_tcscpy_s(buf, len, m_strHTML);
		return m_strHTML.GetLength();   
	}
	catch (CException* e)
	{
		e=NULL;
	}
	return NULL;
}

//下载的都是小型文件,没有做进度条
BOOLEAN downloadfile(char *url,char* localfile)
{
	CInternetSession session;
	CInternetFile* file = NULL;
	try
	{
		// 试着连接到指定URL
		file = (CInternetFile*) session.OpenURL(url,1,INTERNET_FLAG_TRANSFER_BINARY|INTERNET_FLAG_DONT_CACHE); 
	}
	catch (CInternetException* m_pException)
	{
		// 如果有错误的话，置文件为空
		file = NULL; 
		m_pException->Delete();
		return FALSE;
	}

	// 用dataStore来保存读取的文件
	CStdioFile dataStore;

	if (file)
	{
		CString somecode;                            //也可采用LPTSTR类型，将不会删除文本中的\n回车符

		BOOL bIsOk = dataStore.Open(localfile,
			CFile::modeCreate 
			| CFile::modeWrite 
			| CFile::shareDenyWrite 
			| CFile::typeText);

		if (!bIsOk)
			return FALSE;

		// 读写文件，直到为空
		while (file->ReadString(somecode) != NULL) //如果采用LPTSTR类型，读取最大个数nMax置0，使它遇空字符时结束
		{
			dataStore.WriteString(somecode);
			dataStore.WriteString("\n");           //如果somecode采用LPTSTR类型,可不用此句
		}

		file->Close();
		delete file;
	}
	else
	{
		dataStore.WriteString(_T("到指定服务器的连接建立失败..."));    
		return FALSE;
	}

	return TRUE;
}