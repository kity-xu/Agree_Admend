// updateexe.cpp : Defines the entry point for the console application.
//

#include "stdafx.h"

#include <windows.h>
#include <string>
#include <list>
#include "tinyxml.h"

using namespace std;

char *config_file = "config/updateurl.xml";

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

int _tmain(int argc, _TCHAR* argv[])
{
	HINSTANCE	dllhinst;
	typedef VOID (CALLBACK* LPFNDLLFUNC1)(char*);
	LPFNDLLFUNC1 lpfnDllFunc1;

	dllhinst=LoadLibrary("updatedll.dll");
	if(dllhinst == NULL)
	{
		MessageBoxW(GetDesktopWindow(),L"����update.dllʧ��",L"���ش���", MB_OK);
		return -1;
	}

	if (dllhinst!=NULL)
	{
		lpfnDllFunc1=(LPFNDLLFUNC1)GetProcAddress(dllhinst, "StartUpdate");
		if (!lpfnDllFunc1)
		{
			FreeLibrary(dllhinst);
			MessageBoxW(GetDesktopWindow(),L"����StartUpdate����ʧ��",L"���ش���", MB_OK);
			return -1;
		}
	}

	//���ȶ�ȡ�����ļ��е�����,��ȡ���·������б�
	list<const char*> updateserver;
	TiXmlDocument *doc = new TiXmlDocument(config_file);
	bool uploadok = doc->LoadFile();
	if(uploadok == TRUE){
		TiXmlElement *RElement = doc->RootElement();
		TiXmlElement *server = RElement->FirstChildElement();
		while(server){
			updateserver.push_back(server->Attribute("UpdateUrl"));
			server = server->NextSiblingElement();
		}
	}else{
		MessageBoxW(GetDesktopWindow(),L"����config/updateurl.xml�����ļ�����",L"���ش���", MB_OK);
		return FALSE;
	}

	list<const char*>::iterator uit = updateserver.begin();
	while(uit != updateserver.end()){
		lpfnDllFunc1((char*)*uit);
		uit++;
	}

	//����Ƿ����Ը��½ű�
	FreeLibrary(dllhinst);
	RunFile("selfupdate.bat");
	DeleteFile("selfupdate.bat");
	exit(0);
}

