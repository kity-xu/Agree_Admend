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

//创建自运行脚本
void CreateSelfBat(char *srclocation,char* dstlocation){
	static FILE *fp ;
	fp = fopen ("selfupdate.bat","w");
	char com[500];
	sprintf_s(com,500,"copy \"%s\" \"%s\" ",srclocation,dstlocation);
	fwrite(com,500,500,fp);
	fclose(fp);	
}

UPDATEDLL_API void StartUpdate(char *url){
	LOG_PRINT("开始访问网页。http地址 %s",url);
	char  xmlfilelocation[MAX_PATH];
	DWORD xmllen = GetHTMLContent(url,xmlfilelocation,MAX_PATH);
	if(xmllen==NULL){
		LOG_PRINT("获取更新网页错误。URL %s",url);
		return;
	}
	xmlfilelocation[xmllen] = 0;

	char *httptoken = strstr(url,"//");
	if (httptoken == NULL){
		LOG_PRINT("错误的url schema形式。URL %s",url);
		return;
	}
	char *endtoken = httptoken+3;
	while(endtoken < url+strlen(url) && *endtoken !='/' && *endtoken != ':'){
		endtoken++;
	}

	LOG_PRINT("%s %s",httptoken,endtoken);
	if(endtoken == url){
		LOG_PRINT("错误的url schema形式。URL %s",url);
		return;
	}
	string ip;
	string rawip;
	ip.append("config/");
	ip.append(httptoken+2,endtoken - httptoken-2);
	rawip.append(httptoken+2,endtoken - httptoken-2);
	ip.append("/desc.xml");
	CreateMultipleDirectroy((char*)ip.c_str(),ip.size());

	LOG_PRINT("开始下载文件。http地址 %s。本地地址:%s",xmlfilelocation,ip.c_str());

	BOOL uloadok = downloadfile(xmlfilelocation,(char*)ip.c_str());
	
	//没有下载完成,退出
	if(uloadok == FALSE){
		LOG_PRINT("获取xml文件错误。文件名%s");
		return;
	}

	TiXmlDocument *doc = new TiXmlDocument((char*)ip.c_str());
	bool loadOkay = doc->LoadFile();
	if(loadOkay == FALSE){
		LOG_PRINT("加载update.xml错误,%s",doc->ErrorDesc());
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
	if(loadOkay == FALSE) //文件不存在,新建
	{
		LOG_PRINT("找不到配置文件,重新创建。%s",localfile.c_str());
		local_doc = new TiXmlDocument();
		TiXmlDeclaration * decl = new TiXmlDeclaration( "1.0", "utf-8", "" ); 
		local_doc->LinkEndChild( decl ); 
		TiXmlElement te("root");
		local_doc->InsertEndChild(te);
	}
	update_rootelement = local_doc->RootElement();

	//需要替换的文件版本
	const char* version = NULL;

	//在http服务器上的地址
	const char* xsrcloc = NULL;

	//下载到本地目录的地址
	const char* xdstloc = NULL;

	//如果没有文件版本号,则通过文件修改时间来判断
	const char* modifytime = NULL;

	//下载后要执行的方法,默认后是进行拷贝
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
	//这一段是需要处理的
	while(sectionElement)
	{
		//失败
		if(strcmp(sectionElement->Value(),"Section")!=0)
		{
			LOG_PRINT("检测到不合格的xml元素 %s",sectionElement->Value());
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

		//某些需要延后执行的命令,也会被updateend,不过因为没有这个命令,会忽略
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
				LOG_PRINT("更新文件的srclocation和dstlocation不能为空,略过");
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


			//找到在配置文件中的这个元素,如果没有,则新建
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
			//接下来判断版本
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

			//如果需要更新,则执行更新
			if(needupdate == true)
			{
				LOG_PRINT("开始更新文件");

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
				
				if(srcdownload == FALSE)//下载文件错误,继续执行下一个节点
				{
					LOG_PRINT("下载文件错误... http source %s ,local file : %s",srclocation,GenFile.c_str());
					curElement = curElement->NextSiblingElement();
					continue;
				}

				//执行预先的动作,一般是先释放一个进程,然后在进行拷贝
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
						LOG_PRINT("不支持的预定义动作 %s",preupdate);
					}

					//放入delayaction内
					if((updateend != NULL) && (strcmp(updateend,"closeie") == 0))
					{
						string tt(updateend);
						delayaction.push_back(tt);
					}
					//放入delayaction内
					if((updateend != NULL) && (strcmp(updateend,"reboot") == 0))
					{
						string tt(updateend);
						delayaction.push_back(tt);
					}

					has_pre_exec = true;
				}

				//创建自更新
				if(selfupdate != NULL && strncmp(selfupdate,"true",4) ==0){
					CreateSelfBat(srclocation,dstlocation);
					continue;
				}
				CreateMultipleDirectroy(dstlocation,strlen(dstlocation)); //首先创建目录
				//如果存在则进行覆盖

				int copyretry_count = 0;
retrycopy:
				int res = MoveFile((char *)GenFile.c_str(),dstlocation);
				
				if(res == 0)
				{
					char e[10];
					int i  = GetLastError();
					//文件访问冲突
					if(i == ERROR_SHARING_VIOLATION)
					{
						LOG_PRINT("文件共享冲突");
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
					LOG_PRINT("文件未能拷贝成功 %d",GetLastError());
				}
				DeleteFile((char*)localfile.c_str()); //删除文件

				if(modifytime != NULL)
				{
					SetFileModifyTime(dstlocation,modifytime);
					foundelement->SetAttribute("modifytime",modifytime); //设置文件修改时间
				}
				if(version != NULL)
				{
					foundelement->SetAttribute("version",version); //设置文件版本
				}

				//接下去执行action,暂时只支持run
				if(runmethod != NULL)
				{
					//诸如bat之类的
					if(strcmp(runmethod,"run") == 0)
					{
						RunFile(dstlocation);
					}
					//注册dll
					else if(strcmp(runmethod,"reg") == 0)
					{
						string re("regsvr32 /s \"");
						re.append(dstlocation);
						re.append("\"");
						system(re.c_str());
					}
					else{
						LOG_PRINT("未定义的runmethod动作",runmethod);
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

	//保存文件到本地目录
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
	//最后的延后动作,需要特殊处理
	list<string>::iterator it = delayaction.begin();
	while(it != delayaction.end())
	{
		if((*it).compare("closeie") == 0)
		{
			
		}
		else if((*it).compare("reboot") == 0) //重启
		{
			reboot(); //执行重启操作
		}
		it++;
	}
}