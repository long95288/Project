#include<windows.h>
#include<Psapi.h>
#include<stdio.h>
#pragma comment(lib,"Psapi.lib")

void LogCurrentProcessMemoryInfo(){
  HANDLE handle = GetCurrentProcess();

  PROCESS_MEMORY_COUNTERS pmc;
  GetProcessMemoryInfo(handle,&pmc,sizeof(pmc));
  printf("%d Byte \r",pmc.WorkingSetSize);
}

int main(){
  LogCurrentProcessMemoryInfo();
  int i;
  scanf("%d",&i);
  return 0;
}