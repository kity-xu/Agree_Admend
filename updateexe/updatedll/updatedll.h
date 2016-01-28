#ifndef UPDATE_DLL
#define UPDATE_DLL
#include <windows.h>

#ifdef UPDATEDLL_EXPORTS
#define UPDATEDLL_API __declspec(dllexport)
#else
#define UPDATEDLL_API __declspec(dllimport)
#endif

extern "C"{
UPDATEDLL_API void StartUpdate(char *url);
};


#endif