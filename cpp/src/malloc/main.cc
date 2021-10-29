#include <iostream>
#include <stdlib.h>
#include <vector>
#include <map>
#include <set>
#include <algorithm>

using namespace std;

int main() {
    int * p = (int*)malloc(1);
    *p = 5;
    printf("p=%p,*p=%d\n",p,*p);

    char * str = (char *)malloc(1);
    strcpy(str,"dflbkajw");
    printf("str=%s\n,&str=%p",str,&str);
    return 0;
}