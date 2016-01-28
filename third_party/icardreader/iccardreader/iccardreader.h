#ifndef ICCARDREADER_DLL
#define ICCARDREADER_DLL
#include <windows.h>

#ifdef ICCARDDLL_EXPORTS
#define UPDATEDLL_API  extern "C" __declspec(dllexport)
#else
#define UPDATEDLL_API  extern "C" __declspec(dllimport)
#endif

UPDATEDLL_API int ReadICCard(int timeout,char* bkcard,int *msglen,char* errmsg);


#endif