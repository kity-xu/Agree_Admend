#ifndef UTIL_H
#define UTIL_H
#include <string.h>

using namespace std;

//设置文件的最后修改时间
//指：通过属性中看出的文件最后修改时间
void SetFileModifyTime(const char* filename,const char* modifytime);

//创建层级目录
bool CreateMultipleDirectroy(char*name,int len);

//把utf8码转成gb2312码
void UTF8ToGB(char* szOut);

//执行某个文件
void RunFile(const char *path);

//关闭执行的exe文件
void closerunningexe(const char *exepath);

//比较两个数字
int CompareNumber(const char *fir,const char* sec,char sep);

//获取文件修改时间
void GetFileModifyTime(const char *filename, string &modifytime);

//获得文件版本号
void GetFileVersion(const char *filename, string &version);

void shutdown();

void reboot();

void GetHostName(char* hostname);

void DeloginUser(const char* user);

#endif
