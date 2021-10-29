
#### all_amount
curl "http://127.0.0.1:9199/get/json?action=all_amount&coin=mrx"


curl "http://127.0.0.1:9197/get/json?action=all_amount&coin=mrx"

####  hash_query_info


curl "http://127.0.0.1:9199/get/json?action=hash_query_info&coin=mrx&hash=9448717b83fc1655b2d65a46a417676fdf33e42bd8c29f5381b1dd66f5e25ba0&address=MBJN7XA8PcAVAkLV8DGG6SCUZuUpxbze8c&amount=5000.00000000"

curl "http://127.0.0.1:9197/get/json?action=hash_query_info&coin=mrx&hash=9448717b83fc1655b2d65a46a417676fdf33e42bd8c29f5381b1dd66f5e25ba0&address=MBJN7XA8PcAVAkLV8DGG6SCUZuUpxbze8c&amount=5000.00000000"



 "gettransaction \"txid\" ( includeWatchonly )\n"
            "\nGet detailed information about in-wallet transaction <txid>\n"

            "\nArguments:\n"
            "1. \"txid\"    (string, required) The transaction id\n"
            "2. \"includeWatchonly\"    (bool, optional, default=false) Whether to include watchonly addresses in balance calculation and details[]\n"