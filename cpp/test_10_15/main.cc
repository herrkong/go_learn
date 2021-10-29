#include <iostream>

using namespace std;


class Counter {
public:
    int A() {x_ += 1; return x_;}
    int B() {x_ += 1; return x_;}
    int C() {x_ += 1; return x_;}
    Counter():x_(0) {}

private:
    int x_;
};




int main(){
//    float  a = 1.0f;
//    cout << (int)a << endl;
//    cout << (int&)a << endl;
//    cout << boolalpha << ((int)a == (int&)a) << endl;


//    float  b = 0.0f;
//    cout << (int)b << endl;
//    cout << (int&)b << endl;
//    cout << boolalpha << ((int)b == (int&)b) << endl;

 
    //cout << (true ? 1 : "1") << endl;
    // Counter _;
    // int  u = _.A() * 100 + _.B() * 10 + _.C();
    // cout << "u=" << u << endl;

    const int i = 3;
    int const j = 4;
    

    return 0;
}