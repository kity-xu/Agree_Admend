#ifndef NETUTIL_H
#define NETUTIL_H

void GetHostName(char *hostname,int len);

//获得HTML内容
DWORD GetHTMLContent(char *url,char* buf ,int len);

//下载文件
BOOLEAN downloadfile(char *fn,char* localfile);

#endif