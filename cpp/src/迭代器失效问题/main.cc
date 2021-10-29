#include <iostream>
#include <vector>
#include <map>
#include <iterator>

void print(std::string str , int len){
     std::cout << "str=" << str << ",len=" << len << std:: endl;
}

void print(std::string str){
     std::cout << "str=" << str << std:: endl;
}

void print_vector(std::vector<std::string> vec){
    std::cout << "print my_vec\n" << std::endl;
     for (std::vector<std::string>::iterator iter = vec.begin() ; iter != vec.end() ; iter++ ){
        print(*iter);
    }
}

void print_map(std::map<std::string,int> my_map){
    std::cout << "print my_map\n" << std::endl;
    for (std::map<std::string,int>::iterator iter = my_map.begin() ; iter != my_map.end() ; iter++ ){
        print(iter->first,iter->second);
    }
}


// 迭代器失效的原因 是 vector dequeue是连续内存 erase返回的是迭代器下一个位置 所有的元素的内存地址都往前挪了 当前位置到末尾的迭代器全部失效

// 关联容器 map  set multimap  multiset 等 只是被删除的元素的迭代器失效 

// map 底层红黑树

int main(){

    std::vector<std::string> my_vec = {"ross","rachel","chandler","joey","monica","pheobe"};
    std::map<std::string,int> my_map;
    for (int i = 0 ; i < my_vec.size(); i++){
        my_map[my_vec[i]] = my_vec[i].size();
        //print(my_vec[i],my_vec[i].size());
    }

     print_map(my_map);

    std::cout << "----test---map--" << std::endl;

    for (std::map<std::string,int>::iterator iter = my_map.begin() ; iter != my_map.end() ; ){
        if (iter->first == "chandler"){
            my_map.erase(iter++);
            //迭代器失效
            //my_map.erase(iter);
        }else{
            iter++;
        }
    }

    print_map(my_map);

    print_vector(my_vec);

    std::cout << "----test-----vec------" << std::endl;

    // std::erase
    for (std::vector<std::string>::iterator iter = my_vec.begin() ; iter != my_vec.end() ;  ){
        if ((*iter).size() > 4){
            iter = my_vec.erase(iter);
            //迭代器失效
            //my_vec.erase(iter);
        }else{
            iter++;
        }
    }
    print_vector(my_vec);

    return 0;
}

