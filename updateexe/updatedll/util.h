#ifndef UTIL_H
#define UTIL_H
#include <string.h>

using namespace std;

//�����ļ�������޸�ʱ��
//ָ��ͨ�������п������ļ�����޸�ʱ��
void SetFileModifyTime(const char* filename,const char* modifytime);

//�����㼶Ŀ¼
bool CreateMultipleDirectroy(char*name,int len);

//��utf8��ת��gb2312��
void UTF8ToGB(char* szOut);

//ִ��ĳ���ļ�
void RunFile(const char *path);

//�ر�ִ�е�exe�ļ�
void closerunningexe(const char *exepath);

//�Ƚ���������
int CompareNumber(const char *fir,const char* sec,char sep);

//��ȡ�ļ��޸�ʱ��
void GetFileModifyTime(const char *filename, string &modifytime);

//����ļ��汾��
void GetFileVersion(const char *filename, string &version);

void shutdown();

void reboot();

void GetHostName(char* hostname);

void DeloginUser(const char* user);

#endif
