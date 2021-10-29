#include <iostream>
#include <stdlib.h>
#include <vector>
#include <map>
#include <set>
#include <algorithm>

using namespace std;

//线程安全的单例模式
// 还需要AutoRelease自动释放类 atexit()


class Singleton{
public:
    // 对外提供获取唯一实例的接口
    static Singleton * GetSingleton(){
        if (!m_pSingle){
            m_pSingle = new Singleton();
        }
        return m_pSingle;
    }
    void Clear(){
        if(m_pSingle){
            delete m_pSingle;
            m_pSingle = NULL;
            cout << "clear pSingle" << endl;
        } 
    }
private:
    //构造函数私有化 不可创建多个对象
    Singleton(){ cout<< "generate Singleton" << endl;}

private:
    static Singleton * m_pSingle;

};

Singleton * Singleton::m_pSingle = new Singleton; //初始化唯一的pSingle

int main() {
    //只构造了一次
    Singleton * pSingle = Singleton::GetSingleton();
    Singleton * pSecode = Singleton::GetSingleton();
    pSingle->Clear();
    pSecode->Clear();
    return 0;
}