#include<stdio.h>

int szTestData[1024*1024*10] = {1};
int main(){
	// int * piTestData = new int[1024*1024*1024];
	printf("big Memory");
	for(int i=0;i<100;i++){
		int szTestData1[1024*1024*10] = {1000};
	}
	printf("size: %d k\n",sizeof(szTestData)/1024);
	return 0;
}