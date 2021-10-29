#include <stdlib.h>
#include <iostream>
#include <vector>


//   std::string operator()(const PKHash& id) const
//     {
//         std::vector<unsigned char> data = m_params.Base58Prefix(CChainParams::PUBKEY_ADDRESS);
//         data.insert(data.end(), id.begin(), id.end());
//         return EncodeBase58Check(data);
//     }

//std::vector<unsigned char>(1,65); 

void testInsertVector(){
    //[]byte serailazed byte pubkey 
    //std::vector<unsigned char> id();
    std::vector<unsigned char> data(1,65);
}

void print(std::vector<unsigned char> vec){
    for(int i = 0; i < vec.size() ; i++){
        std::cout << vec[i] << std::endl;
    }
}

void testInsertVector2(){
    std::vector<unsigned char> data(1,65);
    std::vector<unsigned char> id(1,66);
   //data.insert(data.end(), id.begin(), id.end());
    print(data);
    print(id);
}


int main(){

    testInsertVector2();
    return 0;
}