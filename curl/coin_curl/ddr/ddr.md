
#### all_amount
curl "http://127.0.0.1:9199/get/json?action=all_amount&coin=ddr"



####  hash_query_info


curl "http://127.0.0.1:9199/get/json?action=hash_query_info&coin=ddr&hash=b3cda776782afaab33f4a6f0e47e43ddb7da0ffa0ee2900a1efee976361589e4&address=DSw97BdqSeFQrviQ5553AJNqaBneuYiSbC&amount=1.00000000"



 "gettransaction \"txid\" ( includeWatchonly )\n"
            "\nGet detailed information about in-wallet transaction <txid>\n"

            "\nArguments:\n"
            "1. \"txid\"    (string, required) The transaction id\n"
            "2. \"includeWatchonly\"    (bool, optional, default=false) Whether to include watchonly addresses in balance calculation and details[]\n"



    
    curl "http://127.0.0.1:9197/get/json?action=hash_query_info&coin=ddr&hash=b3cda776782afaab33f4a6f0e47e43ddb7da0ffa0ee2900a1efee976361589e4&address=DSw97BdqSeFQrviQ5553AJNqaBneuYiSbC&amount=1.00000000"
