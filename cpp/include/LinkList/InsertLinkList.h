#include <stdlib.h>
#include <iostream>

//  p q 为单链表 头尾节点 将其插入另外一个链表s节点之后

#define TestLength 5
#define SetIndex   3

struct node{
    int data;
    node* next;
};

class InsertLinkList {

public:
    InsertLinkList():m_length(TestLength),m_phead(nullptr),m_ptail(nullptr),m_pflag(nullptr){}
    ~InsertLinkList(){
        delete m_phead;
        delete m_ptail;
        delete m_pflag;
    }

public:
    void TestInsertLinkList();
    //初始化两条链表
    void init();
    void createlinklist(node * head,node * tail,node * pflag);   
    void print(node * head);
    int  getlength(){ return m_length;}
    void insert(node * pNode);

private:
    int m_length;   // 维护两条链表，长度都为TestLength
    node * m_phead; // 待插入链表的头尾节点
    node * m_ptail;
    node * m_pflag; // 被插入链表的其中一个节点SetIndex
};