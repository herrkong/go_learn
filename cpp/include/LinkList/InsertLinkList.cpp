#include "InsertLinkList.h"
#include <algorithm>

void InsertLinkList::insert(node * pNode){

}


void InsertLinkList::print(node * head){
     while (head){
         std::cout << head->data << "," ;
     }
}

// 尾插法建立单链表
void InsertLinkList::init(){
    node * head = new node;
    node * tail = new node;
    createlinklist(m_phead,m_ptail,NULL);
    createlinklist(head,tail,m_pflag);
}

void InsertLinkList::createlinklist(node * head,node * tail,node * pflag){
     head->next = NULL;
     tail->next = NULL;
     int a[5]= {1,2,3,4,5};
     //std::random_shuffle(a,a+5);
     for (int i = 0 ; i < 5 ; i++){
        node * p = new node;
        p->data = i;
        tail->next = p;
        tail = p;
        if (i == SetIndex){
            m_pflag = p;
        }
     }
     tail->next = NULL;
     print(head);
}



void InsertLinkList::TestInsertLinkList(){
    init();
    //insert();
}


