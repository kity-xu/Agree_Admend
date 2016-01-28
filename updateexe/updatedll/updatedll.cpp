// updatedll.cpp : Defines the exported functions for the DLL application.
//

#include "stdafx.h"
#include "updatedll.h"

#include <string>
#include <list>
#include <ostream>
#include "tinyxml.h"
#include "util.h"
#include "logutil.h"
#include "netutil.h"

using namespace std;

TiXmlElement* FindDstLocationFile(TiXmlElement* te,const char* dstlocation)
{
	TiXmlElement* curElement = te;
	const char* location = NULL;
	char tmp[MAX_PATH];
	while(curElement)
	{
		location = curElement->Attribute("Dstlocation");
		if(location == NULL){
			return NULL;
		}
		memcpy(tmp,location,strlen(location)+1);
		UTF8ToGB(tmp);
		if((location != NULL) && (strcmp(tmp,dstlocation) == 0))
		{
			return curElement;
		}
		curElement = curElement->NextSiblingElement();
	}
	return NULL;
}

//���������нű�
void CreateSelfBat(char *srclocation,char* dstlocation){
	static FILE *fp ;
	fp = fopen ("selfupdate.bat","w");
	char com[500];
	sprintf_s(com,500,"copy \"%s\" \"%s\" ",srclocation,dstlocation);
	fwrite(com,500,500,fp);
	fclose(fp);	
}

UPDATEDLL_API void StartUpdate(char *url){
	LOG_PRINT("��ʼ������ҳ��http��ַ %s",url);
	char  xmlfilelocation[MAX_PATH];
	DWORD xmllen = GetHTMLContent(url,xmlfilelocation,MAX_PATH);
	if(xmllen==NULL){
		LOG_PRINT("��ȡ������ҳ����URL %s",url);
		return;
	}
	xmlfilelocation[xmllen] = 0;

	char *httptoken = strstr(url,"//");
	if (httptoken == NULL){
		LOG_PRINT("�����url schema��ʽ��URL %s",url);
		return;
	}
	char *endtoken = httptoken+3;
	while(endtoken < url+strlen(url) && *endtoken !='/' && *endtoken != ':'){
		endtoken++;
	}

	LOG_PRINT("%s %s",httptoken,endtoken);
	if(endtoken == url){
		LOG_PRINT("�����url schema��ʽ��URL %s",url);
		return;
	}
	string ip;
	string rawip;
	ip.append("config/");
	ip.append(httptoken+2,endtoken - httptoken-2);
	rawip.append(httptoken+2,endtoken - httptoken-2);
	ip.append("/desc.xml");
	CreateMultipleDirectroy((char*)ip.c_str(),ip.size());

	LOG_PRINT("��ʼ�����ļ���http��ַ %s�����ص�ַ:%s",xmlfilelocation,ip.c_str());

	BOOL uloadok = downloadfile(xmlfilelocation,(char*)ip.c_str());
	
	//û���������,�˳�
	if(uloadok == FALSE){
		LOG_PRINT("��ȡxml�ļ������ļ���%s");
		return;
	}

	TiXmlDocument *doc = new TiXmlDocument((char*)ip.c_str());
	bool loadOkay = doc->LoadFile();
	if(loadOkay == FALSE){
		LOG_PRINT("����update.xml����,%s",doc->ErrorDesc());
		delete doc;
		return;
	}

	TiXmlElement *RootElement = doc->RootElement();

	TiXmlElement *sectionElement = RootElement->FirstChildElement();
	TiXmlElement *curElement;

	string localfile;
	localfile.append("config/");
	localfile.append(rawip);
	localfile.append("/update.xml");
	TiXmlDocument *local_doc = new TiXmlDocument((char*)localfile.c_str());
	TiXmlElement  *update_rootelement = NULL;
	loadOkay = local_doc->LoadFile();
	if(loadOkay == FALSE) //�ļ�������,�½�
	{
		LOG_PRINT("�Ҳ��������ļ�,���´�����%s",localfile.c_str());
		local_doc = new TiXmlDocument();
		TiXmlDeclaration * decl = new TiXmlDeclaration( "1.0", "utf-8", "" ); 
		local_doc->LinkEndChild( decl ); 
		TiXmlElement te("root");
		local_doc->InsertEndChild(te);
	}
	update_rootelement = local_doc->RootElement();

	//��Ҫ�滻���ļ��汾
	const char* version = NULL;

	//��http�������ϵĵ�ַ
	const char* xsrcloc = NULL;

	//���ص�����Ŀ¼�ĵ�ַ
	const char* xdstloc = NULL;

	//���û���ļ��汾��,��ͨ���ļ��޸�ʱ�����ж�
	const char* modifytime = NULL;

	//���غ�Ҫִ�еķ���,Ĭ�Ϻ��ǽ��п���
	const char* runmethod = NULL;

	TiXmlElement *foundelement = NULL;

	bool needupdate = false;

	const char* cversion = NULL;
	const char* cmodifytime = NULL;

	char exepath[MAX_PATH];

	const char* utfexepath = NULL;
	const char* preupdate = NULL;
	const char* updateend = NULL;
	const char* userexit = NULL;

	bool o_exit = false;

	list<string> delayaction;

	char srclocation[MAX_PATH];
	char dstlocation[MAX_PATH];

	bool has_pre_exec = false;
	const char* selfupdate = false;
	//��һ������Ҫ�����
	while(sectionElement)
	{
		//ʧ��
		if(strcmp(sectionElement->Value(),"Section")!=0)
		{
			LOG_PRINT("��⵽���ϸ��xmlԪ�� %s",sectionElement->Value());
			sectionElement = sectionElement->NextSiblingElement();
			continue;
		}
		o_exit = false;
		memset(exepath,0,MAX_PATH);
		has_pre_exec = false;
		utfexepath = sectionElement->Attribute("Exe");
		if(utfexepath != NULL)
		{
			memcpy(exepath,utfexepath,strlen(utfexepath)+1);
			UTF8ToGB(exepath);
		}

		//ĳЩ��Ҫ�Ӻ�ִ�е�����,Ҳ�ᱻupdateend,������Ϊû���������,�����
		preupdate = sectionElement->Attribute("Preupdate");
		updateend = sectionElement->Attribute("Updateend");
		userexit = sectionElement->Attribute("Userexit");
		if((userexit != NULL) && strcmp(userexit,"true") == 0)	{
			o_exit = true;
		}

		curElement = sectionElement->FirstChildElement();
		while(curElement)
		{
			version = curElement->Attribute("Version");
			xsrcloc = curElement->Attribute("Srclocation");
			xdstloc = curElement->Attribute("Dstlocation");

			if((xsrcloc == NULL) || (xdstloc == NULL))	{
				LOG_PRINT("�����ļ���srclocation��dstlocation����Ϊ��,�Թ�");
				curElement = curElement->NextSiblingElement();
				continue;
			}

			memcpy(srclocation,xsrcloc,strlen(xsrcloc));
			memcpy(dstlocation,xdstloc,strlen(xdstloc));
			srclocation[strlen(xsrcloc)] = NULL;
			dstlocation[strlen(xdstloc)] = NULL;
			UTF8ToGB(srclocation);
			UTF8ToGB(dstlocation);

			modifytime = curElement->Attribute("Modifytime");
			runmethod = curElement->Attribute("Runmethod");
			selfupdate = curElement->Attribute("Type");


			//�ҵ��������ļ��е����Ԫ��,���û��,���½�
			foundelement = FindDstLocationFile(update_rootelement->FirstChildElement(),dstlocation);
			if(foundelement == NULL){
				TiXmlElement tmp("File");
				tmp.SetAttribute("Dstlocation",xdstloc);
				if(version != NULL)
				{
					string lversion;
					GetFileVersion(dstlocation,lversion);
					tmp.SetAttribute("Version",lversion.c_str());
				}
				if(modifytime != NULL)
				{
					string lmodify;
					GetFileModifyTime(dstlocation,lmodify);
					tmp.SetAttribute("Modifytime",lmodify.c_str());
				}

				foundelement = (TiXmlElement*)update_rootelement->InsertEndChild(tmp);

			}

			needupdate = false;
			//�������жϰ汾
			if(version != NULL)
			{
				cversion = foundelement->Attribute("Version");
				int res = CompareNumber(version,cversion,'.');
				if(res == -1){
					needupdate = true;
				}
			}
			if(modifytime != NULL)
			{
				cmodifytime = foundelement->Attribute("modifytime");

				int res = CompareNumber(modifytime,cmodifytime,' ');
				if(res == -1)
				{
					needupdate = true;
				}
			}

			//�����Ҫ����,��ִ�и���
			if(needupdate == true)
			{
				LOG_PRINT("��ʼ�����ļ�");

				string GenFile;
				GenFile.append(rawip);
				GenFile.append("/");
				char* lstocc = strrchr(srclocation,'/');
				if (lstocc == NULL){
					GenFile.append(srclocation);
				}else{
					GenFile.append(lstocc+1);
				}
				
				BOOL srcdownload = downloadfile(srclocation,(char*)GenFile.c_str());
				
				if(srcdownload == FALSE)//�����ļ�����,����ִ����һ���ڵ�
				{
					LOG_PRINT("�����ļ�����... http source %s ,local file : %s",srclocation,GenFile.c_str());
					curElement = curElement->NextSiblingElement();
					continue;
				}

				//ִ��Ԥ�ȵĶ���,һ�������ͷ�һ������,Ȼ���ڽ��п���
				if(!has_pre_exec)
				{
					if(preupdate == NULL)
					{

					}
					else if(strcmp(preupdate,"close") == 0)
					{
						closerunningexe(exepath);
					}
					else
					{
						LOG_PRINT("��֧�ֵ�Ԥ���嶯�� %s",preupdate);
					}

					//����delayaction��
					if((updateend != NULL) && (strcmp(updateend,"closeie") == 0))
					{
						string tt(updateend);
						delayaction.push_back(tt);
					}
					//����delayaction��
					if((updateend != NULL) && (strcmp(updateend,"reboot") == 0))
					{
						string tt(updateend);
						delayaction.push_back(tt);
					}

					has_pre_exec = true;
				}

				//�����Ը���
				if(selfupdate != NULL && strncmp(selfupdate,"true",4) ==0){
					CreateSelfBat(srclocation,dstlocation);
					continue;
				}
				CreateMultipleDirectroy(dstlocation,strlen(dstlocation)); //���ȴ���Ŀ¼
				//�����������и���

				int copyretry_count = 0;
retrycopy:
				int res = MoveFile((char *)GenFile.c_str(),dstlocation);
				
				if(res == 0)
				{
					char e[10];
					int i  = GetLastError();
					//�ļ����ʳ�ͻ
					if(i == ERROR_SHARING_VIOLATION)
					{
						LOG_PRINT("�ļ������ͻ");
						if(o_exit == true)
						{
							DeloginUser(userexit);
							Sleep(500);
							if(copyretry_count++ < 15 )
							{
								goto retrycopy;
							}
							else
							{
								curElement = curElement->NextSiblingElement();
								continue;
							}
						}
						else
						{
							curElement = curElement->NextSiblingElement();
							continue;
						}
					}
					itoa(i,e,10);

				}
				else
				{
					LOG_PRINT("�ļ�δ�ܿ����ɹ� %d",GetLastError());
				}
				DeleteFile((char*)localfile.c_str()); //ɾ���ļ�

				if(modifytime != NULL)
				{
					SetFileModifyTime(dstlocation,modifytime);
					foundelement->SetAttribute("modifytime",modifytime); //�����ļ��޸�ʱ��
				}
				if(version != NULL)
				{
					foundelement->SetAttribute("version",version); //�����ļ��汾
				}

				//����ȥִ��action,��ʱֻ֧��run
				if(runmethod != NULL)
				{
					//����bat֮���
					if(strcmp(runmethod,"run") == 0)
					{
						RunFile(dstlocation);
					}
					//ע��dll
					else if(strcmp(runmethod,"reg") == 0)
					{
						string re("regsvr32 /s \"");
						re.append(dstlocation);
						re.append("\"");
						system(re.c_str());
					}
					else{
						LOG_PRINT("δ�����runmethod����",runmethod);
					}
				}

			}

nextElement:
			curElement = curElement->NextSiblingElement();
		}

		if(has_pre_exec)
		{
			if(updateend == NULL)
			{

			}
			else if(strcmp(updateend,"open") == 0){
				RunFile(exepath);
			}
			else
			{
				OutputDebugString("unsupported updateend action");
			}
		}

		sectionElement = sectionElement->NextSiblingElement();
	}

	//�����ļ�������Ŀ¼
	local_doc->SaveFile((char*)localfile.c_str());

	if(doc != NULL){
		delete doc;
		doc = NULL;
	}
	if(local_doc != NULL)
	{
		delete local_doc;
		local_doc = NULL;
	}
	//�����Ӻ���,��Ҫ���⴦��
	list<string>::iterator it = delayaction.begin();
	while(it != delayaction.end())
	{
		if((*it).compare("closeie") == 0)
		{
			
		}
		else if((*it).compare("reboot") == 0) //����
		{
			reboot(); //ִ����������
		}
		it++;
	}
}