#ifndef NETUTIL_H
#define NETUTIL_H

void GetHostName(char *hostname,int len);

//���HTML����
DWORD GetHTMLContent(char *url,char* buf ,int len);

//�����ļ�
BOOLEAN downloadfile(char *fn,char* localfile);

#endif