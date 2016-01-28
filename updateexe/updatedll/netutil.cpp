#include "stdafx.h"
#include "netutil.h"


//���HTML����
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

//���صĶ���С���ļ�,û����������
BOOLEAN downloadfile(char *url,char* localfile)
{
	CInternetSession session;
	CInternetFile* file = NULL;
	try
	{
		// �������ӵ�ָ��URL
		file = (CInternetFile*) session.OpenURL(url,1,INTERNET_FLAG_TRANSFER_BINARY|INTERNET_FLAG_DONT_CACHE); 
	}
	catch (CInternetException* m_pException)
	{
		// ����д���Ļ������ļ�Ϊ��
		file = NULL; 
		m_pException->Delete();
		return FALSE;
	}

	// ��dataStore�������ȡ���ļ�
	CStdioFile dataStore;

	if (file)
	{
		CString somecode;                            //Ҳ�ɲ���LPTSTR���ͣ�������ɾ���ı��е�\n�س���

		BOOL bIsOk = dataStore.Open(localfile,
			CFile::modeCreate 
			| CFile::modeWrite 
			| CFile::shareDenyWrite 
			| CFile::typeText);

		if (!bIsOk)
			return FALSE;

		// ��д�ļ���ֱ��Ϊ��
		while (file->ReadString(somecode) != NULL) //�������LPTSTR���ͣ���ȡ������nMax��0��ʹ�������ַ�ʱ����
		{
			dataStore.WriteString(somecode);
			dataStore.WriteString("\n");           //���somecode����LPTSTR����,�ɲ��ô˾�
		}

		file->Close();
		delete file;
	}
	else
	{
		dataStore.WriteString(_T("��ָ�������������ӽ���ʧ��..."));    
		return FALSE;
	}

	return TRUE;
}