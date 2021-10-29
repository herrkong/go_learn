

#### 多态的实现方法

1 函数重载 ：函数同名 但是参数个数和类型不同
2 覆盖： 父类为虚函数 子类重新定义父类的虚函数


#### 字符串类型作为传参
char const * 
std::string const &  == const std::string & 
std::string 




const char* // 指针可修改 指向的char元素不能修改

char const * // 指向mutable 字符的常量指针  


#### std::forward完美转发
传入的左值(右值) 给下一个函数传参也是左值(右值)

#### std::move : 移除左值属性

