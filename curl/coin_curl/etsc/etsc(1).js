
//http://52.80.67.251:9196/get/json?action=trace&uid=22&address=0x984Fd5d34E29AEC7cE5fDe06Bd2Fc27da3d9c35E&coin=wlszb
//52.80.67.251:9196/get/json?action=create&uid=22&coin=etz
//52.80.67.251:9196/get/json?action=create&uid=22&coin=etz
//http://52.80.67.251:9196/get/json?action=query&uid=22&address=0x984Fd5d34E29AEC7cE5fDe06Bd2Fc27da3d9c35E&coin=etz
//http://52.80.67.251:9196/get/json?action=transfer&uid=22&address=0x021fc4ba776fb26cead2c5db6376affc33b95cc8&to_address=0x1905dd88042fb9e6ea5c3b02d47c93ad25865ce1&coin=etz
/*
{
    "ret": "succ",
    "uid": "22",
    "address": "0x021fc4ba776fb26cead2c5db6376affc33b95cc8",
    "to_address": "0x1905dd88042fb9e6ea5c3b02d47c93ad25865ce1",
    "amount": "0.4589776432558277",
    "hashid": "0x182ed40aa8c516cde75c689afadc63c8b81213d89c91016e3b59163812bbd7ca"
}

// var options = {fromBlock: 4884146, toBlock: 'latest', 'topics':['0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef',null,'0x0000000000000000000000009f87a7ede00311d232be212e69df3b32615733e6']};
undefined
// var filter = web3.eth.filter(options);
undefined
// filter.get();

*/
//etz.accounts.forEach(function(a) {console.log(a + " "+web3.fromWei(etz.getBalance(a), "ether"));})
//admin.peers.forEach(function(p) {console.log(p.network.remoteAddress);})
var fs = require("fs");
var mylog = require('../my_log.js');
//var express_handle = require('./express_handle.js');
var my_ajax = require('../my_ajax.js');
var web3_extended = require('web3_ipc');
var utility = require('../utility.js');
var BigNumber = require('bignumber.js');
var async = require("async");
var web3 = null;
//var nonce = 1000;

///////////////////////////这些参数需要修改

var gasPrice = 50000000000;
var gasLimit = 60000;

//var gasPrice = 60000000000;
//var gasLimit = 46000;
//var gasLimit = 1001000;
var token_digits = 1000000000000000000;

//自动转币最大值
var g_auto_max_amount = 1000000;
var g_auto_out_address = '0xf7ec0f4ba89b1645a6026f0577fd411e4baf4406';
var g_auto_add_gas_amount = 0.003;

//自动汇总最小额度
var g_auto_collect_num = 0.01

//var my_ajax = require
var  req_option = {
            host: 'api.etherscan.io',
            port: 80,
            path: '/api?module=account&action=txlist&startblock=4039429&endblock=99999999&sort=desc&apikey=ZF9QM77BSM3XKM95VG64AXN8Q8ZWQWD7PU&address=',
            times: 3,
            metzod: 'GET'
};

////////////////////////////////end
//自动归集地址
//var g_address = "0xB37640f5F7ef7b0fDCce2c0C053DB4f976945647";
var g_address = "0xf7ec0f4ba89b1645a6026f0577fd411e4baf4406";

//记录存储位置
var g_trace_prefix =  './etz_trace/';

function init(ipc,index,callback)
{
    var options = {
          host: ipc,
          ipc:true,
          personal: true, 
          admin: false,
          debug: false
        };

    web3 = web3_extended.create(options);
    mylog.log_str(JSON.stringify(options) + ",etz init ok!!");
    callback(index);
}

function handle(arg,postData,succ,fail,cache)
{
    mylog.log_str(" handle COME IN");
    if (arg.action == "create") 
    {
        return create(arg,postData,succ,fail);
    }
    else if(arg.action == "query") 
    {
        return query(arg,postData,succ,fail);
    }
    else if(arg.action == "all_amount") 
    {
        return all_amount(arg,postData,succ,fail);
    }
    else if(arg.action == "transfer")
    {
        
        return transfer(arg,postData,succ,fail);
        
       
    }
    else if (arg.action == "auto_draw") 
    {
        if ("hot_wallet" == arg.address )
        {
            var transfer_amount = parseFloat(arg.amount);
            if (!arg.amount || transfer_amount > g_auto_max_amount) 
            {
                mylog.log_str("auto draw,amount invalid,transfer_amount:" + transfer_amount + ",g_auto_max_amount:"+g_auto_max_amount);
                var obj = new Object();
                obj.ret = "fail";
                obj.msg = "auto draw,amount invalid,transfer_amount:" + transfer_amount;
                //mylog.log_str(JSON.stringify(error));
                succ(obj);
            }
            else
            {
                arg.address = g_auto_out_address;
                arg.coin_force = "1";
                mylog.log_str("auto draw,para:" + JSON.stringify(arg));
                return transfer(arg,postData,succ,fail);
            }
            
        }
        else
        {
            mylog.log_str("auto draw,address invalid,address:" + arg.address);
            var obj = new Object();
            obj.ret = "fail";
            obj.msg = "auto draw,address invalid,address:" + arg.address;
            //mylog.log_str(JSON.stringify(error));
            succ(obj);
        }
    }
    else if(arg.action == "trace")
    {
        return trace(arg,postData,succ,fail);
    }
    else if(arg.action == "add_gas")
    {
        return add_gas(arg,postData,succ,fail);
    }
    else
    {
        mylog.log_str("arg:"+ JSON.stringify(arg));
        //var address = arg.address;
        var obj = new Object();
        obj.ret = "fail";
        //obj.uid = arg.uid;
        obj.msg =  "etz metzod:"+arg.action+" not implement!!";
        //obj.address = address;
        //obj.bal = 0;
        mylog.log_str(obj);
        succ(obj);
    }
    mylog.log_str(" handle COME out");
     
}
//转手续费到指定账号
function add_gas(arg,postData,succ,fail)
{
    //return error?("name:" + error.name +",msg:"+error.message):" null";

    var passwd = "19840902";
    //首先判断是不是我们的账户
    web3.personal.unlockAccount(arg.to_address, passwd,function(err,result)
    {
        mylog.log_str("unlockAccount err:"+ err_2_str(err));
        if(!err)
        {
            arg.address = g_auto_out_address;
            arg.coin_force = "1";
            arg.amount = g_auto_add_gas_amount;
            mylog.log_str("add_gas,para:" + JSON.stringify(arg));
            return transfer(arg,postData,succ,fail);
        }
        else
        {
            var obj = new Object();
            obj.ret = "fail";
            obj.msg = "unlock error:" + arg.to_address +",not our account";
            mylog.log_str("unlock err:"+err_2_str(err));
            succ(obj);
        }
    });

    
}
function query(arg,postData,succ,fail)
{
    var address = arg.address;
    web3.eth.getBalance(address, function(error, result) {
        // This is executed once I got the response
        mylog.log_str("Response received");

        if(error) {
            //console.error(error);
            //return;

            mylog.log_str(error);
            var obj = new Object();
            obj.ret = "fail";
            obj.msg = "web3 error:" + JSON.stringify(error);
            mylog.log_str(JSON.stringify(error));
            succ(obj);
        }
        else
        {
            mylog.log_str(err_2_str(error) + ",result:" + result);
            //var accountBalance = result.toNumber()/token_digits;
            //var accountBalance =  web3.fromWei(result, "ether");
           //var amount = numberSize(Number(result),8);
            //var accountBalance = utility.div(result,token_digits);
            //var accountBalance =  NumLastSub(accountBalance,1);
            //var accountBalance = utility.div(result.toNumber(),token_digits);
            //console.log(accountBalance);
            //console.log(result);
            //console.log(result.toString(10));
            var tmpBalance = new BigNumber(result);
            var accountBalance =  tmpBalance.dividedBy(token_digits).toString(10);
            var obj = new Object();
            obj.ret = "succ";
            //obj.uid = arg.uid;
            obj.address = address;
            obj.bal = accountBalance;
            mylog.log_str(obj);

            succ(obj);
        }
        
    });

}

function create(arg,postData,succ,fail)
{
    try
    {
        //if( arg.action  == "response")
        //{
            ////直接回复，按照fail处理，实际上成功了
        //my_ajax.response_ajax(arg,postData,fail);

        web3.personal.newAccount("19840902",function(error,result){
            if(!error){
                mylog.log_str(result);
                var obj = new Object();
                obj.ret = "succ";
                obj.uid = arg.uid;
                obj.address = result;
                mylog.log_str(obj);
                succ(obj);
            }
            else
            {
                var obj = new Object();
                obj.ret = "fail";
                obj.msg = "web3 error:"+ err_2_str(error);
                mylog.log_str(err_2_str(error));
                succ(obj);
            }
        });
        
    }
    catch(e)
    {
        mylog.log_str('req_handle catch,error '+ err_2_str(error) +",arg:"+ JSON.stringify(arg));
        //fail("内部错误！！");
        var obj = new Object();
        obj.ret = "fail";
        obj.msg = "代码异常";
        mylog.log_str(error);
        succ(obj);
    }
}
function err_2_str(error)
{
    return error?("name:" + error.name +",msg:"+error.message):" null";
}
function raw_transfer(transactionObject,succ,fail) {
    mylog.log_str("raw_transfer:"+JSON.stringify(transactionObject));
    try {
        if (transactionObject.coin_force == "1") 
        {

        }
        else
        {
            transactionObject.to = g_address;
        }
        web3.eth.sendTransaction(transactionObject,function(error,result)
        {
            mylog.log_str("sendTransaction err:"+ err_2_str(error));
            mylog.log_str("sendTransaction result:" + JSON.stringify(result));
            if(!error)
            {
                mylog.log_str("sendTransaction result in succ");
                succ(result);
            }
            else
            {
                mylog.log_str("sendTransaction result in fail");
                fail(error);
            }
            
        });
        //return false;
    } catch (err) {
        mylog.log_str("raw_transfer catch err " + err_2_str(err));
        fail(err);
    }
}

function transfer(arg,postData,succ,fail)
{
    try
    {
        query(arg,null
        ,function(queryObj)//succ
        {
            var sender = arg.address;
            var receiver = arg.to_address;
            var coin_force = arg.coin_force;
            //var amount = web3.toWei(queryObj.bal, "ether");
            //var before_amount = fomatFloat((queryObj.bal),8);
            //var before_amount = numberSize(parseFloat(queryObj.bal),8);
            var before_amount = queryObj.bal;
            var amount = web3.toWei(before_amount, "ether");
            //var amount = parseFloat(queryObj.bal)* token_digits;
            

            //web3.eth.getGasPrice(function(gasErr,gasPrice)
            var gasErr = null;
            {

                if(!gasErr)
                {

                    //mylog.log_str("gasPrice:" + JSON.stringify(gasPrice)+",raw gasPrice:" +gasPrice);
                    mylog.log_str("gasPrice:" + JSON.stringify(gasPrice)+",raw gasPrice:" +gasPrice);
                    //BigNumber(0.3).minus(0.1) 
                    var tmpBalance = new BigNumber(amount);
                    var transferAmount = tmpBalance.minus(gasPrice*gasLimit).toString(10);
                    //transferAmount = fomatFloat(parseFloat(transferAmount),8);
                    if(arg.amount)
                    {
                        //transferAmount = parseFloat(arg.amount) * token_digits;
                        transferAmount = web3.toWei(parseFloat(arg.amount), "ether");
                    }

                    mylog.log_str("before_amount:" + before_amount +",amount:" + amount + ",transferAmount:" + transferAmount +",gasLimit:" + gasLimit);

                    var transactionObject = new Object();
                    transactionObject.from = sender;
                    transactionObject.to = receiver;
                    transactionObject.value = transferAmount;
                    transactionObject.gas = gasLimit;
                    transactionObject.gasPrice = gasPrice;
		    
		    
                    transactionObject.coin_force = coin_force;
		    
		    
                    //transactionObject.nonce = nonce;
                    //transactionObject.password = "iwala.net";
                    mylog.log_str("transactionObject:"+ JSON.stringify(transactionObject));
                    if (amount > 0 ) 
                    {
                        raw_transfer(transactionObject,
                        function(result)//succ
                        {
                            mylog.log_str("succ,result:"+JSON.stringify(result));
                            var obj = new Object();
                            obj.ret = "succ";
                            obj.uid = arg.uid;
                            obj.address = arg.address;
                            obj.to_address = arg.to_address;
                            obj.balance = web3.fromWei(queryObj.bal*token_digits, 'ether');
                            obj.amount = transferAmount / token_digits ;
                            obj.hashid = result;
                            mylog.log_str(obj);
                            succ(obj);
                        },
                        function(error)//fail
                        {
                            mylog.log_str("err:"+ err_2_str(error)+",check need to unlock.");
                            if(error.message == "authentication needed: password or unlock")
                            {
                                var passwd = "19840902";
                                if ("0x021fc4ba776fb26cead2c5db6376affc33b95cc8" == transactionObject.from || "0x1905dd88042fb9e6ea5c3b02d47c93ad25865ce1" == transactionObject.from)
                                {
                                    passwd = "iwala.net";
                                }
                                web3.personal.unlockAccount(transactionObject.from, passwd,function(err,result)
                                {
                                    mylog.log_str("unlockAccount err:"+ err_2_str(err));
                                    if(!err)
                                    {
                                        raw_transfer(transactionObject,
                                        function(result)//succ
                                        {
                                            mylog.log_str(result);
                                            var obj = new Object();
                                            obj.ret = "succ";
                                            obj.uid = arg.uid;
                                            obj.address = arg.address;
                                            obj.to_address = arg.to_address;
                                            obj.balance = web3.fromWei(queryObj.bal*token_digits, 'ether');
                                            obj.amount = transferAmount / token_digits;
                                            obj.hashid = result;
                                            mylog.log_str(obj);
                                            succ(obj);
                                        },
                                        function(error)//fail
                                        {
                                                var obj = new Object();
                                                obj.ret = "fail";
                                                obj.msg = "second web3 error:" + err_2_str(error);
                                                mylog.log_str("second raw transfer err:"+err_2_str(error));
                                                succ(obj);
                                        });
                                    }
                                    else
                                    {
                                        var obj = new Object();
                                        obj.ret = "fail";
                                        obj.msg = "unlock error:" + err_2_str(err);
                                        mylog.log_str("unlock err:"+err_2_str(err));
                                        succ(obj);
                                    }
                                });
                            }
                            else
                            {
                                var obj = new Object();
                                obj.ret = "fail";
                                obj.msg = "web3 error:" + err_2_str(error);
                                mylog.log_str("raw transfer err:"+err_2_str(error));
                                succ(obj);
                            }
                        });
                    }
                    else
                    {
                        var obj = new Object();
                        obj.ret = "fail";
                        obj.msg = "balance of address:"+arg.address +" is "+amount;
                        succ(obj);
                    }
                }
                else
                {
                    var obj = new Object();
                    obj.ret = "fail";
                    obj.msg = "get gas price error:"+err_2_str(gasErr);
                }
                

            }
            
        }
        ,function()//fail
        {
            var obj = new Object();
            obj.ret = "fail";
            obj.msg = "web3 query error";
            succ(obj);
        })
        
    }
    catch(e)
    {
        mylog.log_str('req_handle catch,name:'+ e.name +',msg:'+e.message+",arg:"+ JSON.stringify(arg));
        //fail("内部错误！！");
        var obj = new Object();
        obj.ret = "fail";
        obj.msg = "代码异常";
        mylog.log_str(error);
        succ(obj);
    }
}


function trace(arg,postData,succ,fail) {
    mylog.log_str("arg:"+JSON.stringify(arg));
    try {
        fs.exists(g_trace_prefix, function(exists) {  
            //console.log(exists ? "创建成功" : "创建失败");  
            mylog.log_str(g_trace_prefix+" is exists:"+exists);
            if (exists)  //文件存在直接读取文件
            {
                 var  file_name = g_trace_prefix+arg.address.toLowerCase();
                 fs.exists(file_name, function(file_exsit)
                 {
                    if (file_exsit) 
                    {
                        var traceData = fs.readFileSync(file_name);

                         var obj = new Object();
                         obj.ret = "succ";
                         obj.address = arg.address;
                                //obj.result = traceData.result;
                         obj.result = utility.parseHtml(traceData);
                         mylog.log_str(obj);
                         succ(obj);
                    }
                    else
                    {
                        var obj = new Object();
                         obj.ret = "succ";
                         obj.address = arg.address;
                                //obj.result = traceData.result;
                         obj.result = new Array();
                         //mylog.log_str(obj);
                         succ(obj);
                    }
                 });
                 

            }
            else //转发查询
            {

                mylog.log_str(" err ,etz trace err,no :" + g_trace_prefix);
                //callback();
                var obj = new Object();
                obj.ret = "fail";
                obj.msg = " err ,etz trace err,no :" + g_trace_prefix + " directory";
                //obj.address = address;
                //obj.bal = 0;
                mylog.log_str(obj);
                succ(obj);
            }
        }); 
        //return false;
    } catch (err) {
        mylog.log_str("trace catch err " + err_2_str(err));
        fail(err);
    }
}


function all_amount(arg,postData,succ,fail)
{
    mylog.log_str("arg:"+JSON.stringify(arg));
    try
    {
        var address = g_address;
        web3.eth.getBalance(address, function(error, result) {
            // This is executed once I got the response
            mylog.log_str("Response received");

            if(error) {
                //console.error(error);
                //return;

                mylog.log_str(error);
                var obj = new Object();
                obj.ret = "fail";
                obj.msg = "web3 error:" + JSON.stringify(error);
                mylog.log_str(JSON.stringify(error));
                succ(obj);
            }
            else
            {
                mylog.log_str(err_2_str(error) + ",result:" + result);
                //var accountBalance = result.toNumber()/token_digits;
                //var accountBalance =  web3.fromWei(result, "ether");
                //var accountBalance = utility.div(result.toNumber(),token_digits);
            //console.log(accountBalance);
                var tmpBalance = new BigNumber(result);
                var accountBalance =  tmpBalance.dividedBy(token_digits).toString(10);
                var obj = new Object();
                obj.ret = "succ";
                //obj.uid = arg.uid;
                obj.address = address;
                obj.coin = arg.coin;
                obj.bal = accountBalance;
                mylog.log_str(obj);

                succ(obj);
            }
            
        });
        
    }
    catch(e)
    {
        mylog.log_str('req_handle catch,error '+ err_2_str(e) +",arg:"+ JSON.stringify(arg));
        //fail("内部错误！！");
        var obj = new Object();
        obj.ret = "fail";
        obj.coin = arg.coin;
        obj.msg = err_2_str(e);
        mylog.log_str(err_2_str(e));
        succ(obj);
    }
}


function trace_external(arg,postData,succ,fail) {
    mylog.log_str("arg:"+JSON.stringify(arg));
    try {
        var options = JSON.parse(JSON.stringify(req_option));
        options.path = options.path + arg.address;
        return my_ajax.direct_ajax(options,function(traceDataStr,headDataStr)
        {
            //mylog.log_str("traceDataStr:"+traceDataStr);
            var traceData = utility.parseHtml(traceDataStr);
            mylog.log_str("traceData:"+JSON.stringify(traceData));
            mylog.log_str("message:"+(traceData.message));
            //mylog.log_str("result:"+(traceData.result));
            if (traceData && traceData.message == "OK" && traceData.result) 
            {
                var obj = new Object();
                obj.ret = "succ";
                obj.address = arg.address;
                //obj.result = traceData.result;
                obj.result = new Array();
                for (var i = 0; i < traceData.result.length; i++) {
                    if (traceData.result[i]["to"] == arg.address) 
                    {
                        traceData.result[i]["value"] = parseInt(traceData.result[i]["value"])/token_digits;
                        obj.result.push(traceData.result[i]);
                    }
                    /*
                    var item = new Object();
                    if (to == ) {}

                    item.hash = traceData.data.txs[i].txid;

                    item.timeStamp = traceData.data.txs[i].time;
                    item.isError = "0";
                    item.from = "";
                    item.to = arg.address;
                    item.confirmations = traceData.data.txs[i].confirmations;
                    item.value = traceData.data.txs[i].value;
                    //mylog.log_str(obj);
                    obj.result.push(item);
                    */
                }
                mylog.log_str(obj);
                succ(obj);
            }
            else
            {
                var obj = new Object();
                obj.ret = "fail";
                obj.msg = "请求接口,返回的数据没有记录，请联系管理员";
                mylog.log_str("data not valid!!");
                succ(obj);
            }
            
        }
        ,function()
        {
            var obj = new Object();
            obj.ret = "fail";
            obj.msg = "请求接口失败，请联系管理员";
            //mylog.log_str(error);
            succ(obj);
        });
        //return false;
    } catch (err) {
        mylog.log_str("trace catch err " + err_2_str(err));
        fail(err);
    }
}


function get_token()
{
    var tokenInfo = new Object();
    tokenInfo.token_digits = token_digits;
    tokenInfo.auto_collect_num = g_auto_collect_num;
    tokenInfo.add_gas_amount = g_auto_add_gas_amount;
    tokenInfo.collet_address = g_address.toLowerCase() ;
    tokenInfo.auto_address = g_auto_out_address.toLowerCase() ;
    return tokenInfo;
}

exports.add_gas = add_gas;
exports.get_token = get_token;
exports.handle = handle;
exports.init = init;
