/**
 * 以太坊处理类
 * @author Bourne<liuweifu@digifinex.com>
 * 2018-09-21 16:20:22
 */

const { logstr, sendQywxMsg, sendJson } = require('../my_modules/common.js');//公共函数库
const common = require('../my_modules/common.js');//公共函数库
const bigNumber = require('bignumber.js');
const secret = require("../my_modules/secret.js");
const withdraw = require("../my_modules/withdraw.js");
let qianbaoDB = require('../my_modules/database/index.js');
const exec = require('child_process').exec;

const eth = class {
	//初始化
	constructor(initData) {
    	let keys = Object.keys(initData);
    	for(let [index, key] of keys.entries()){
    		this[key] = initData[key];
    	}

    	let options = {
          	host: this.ipc,
          	ipc:true,
          	personal: true,
          	admin: false,
          	debug: false
        };

        delete require.cache[require.resolve('web3_ipc')]
        const web3_extended = require('web3_ipc');

	    this.web3 = web3_extended.create(options);
  	}

  	/**
  	 * 查询指定地址余额
  	 * @Author   LiuWeiFu
  	 * @DateTime 2018-10-18T16:39:20+0800
  	 * @param    [key]/[type]/[desc]
  	 * @param    {[type]}                 arg      [description]
  	 * @param    {Function}               callback [description]
  	 * @return   {[type]}                          [description]
  	 */
  	query(arg, callback) {
  		let { address } = arg,
  			obj = {};
	    this.web3.eth.getBalance(address, (error, result) => {//这里必须使用箭头函数,否则下面的this会指向getBalance方法
	        logstr("Response received", arg.coin);

	        if(error) {
	            logstr(error, arg.coin);
	            obj.ret = "fail";
	            obj.msg = `web3 error: ${JSON.stringify(error)}`;
	        }else{
	            logstr(`result: ${result}`, arg.coin);
	            let accountBalance =  Number((new bigNumber(result)).dividedBy(this.token_digits).toString(10));
	            obj.ret = "succ";
	            obj.address = address;
	            obj.bal = accountBalance;
	        }
	        logstr(`eth query response: ${JSON.stringify(obj)}`, arg.coin);
            callback(obj);
	    });
  	}

  	/**
  	 * 钱包地址创建
  	 * @Author   LiuWeiFu
  	 * @DateTime 2018-10-18T16:39:51+0800
  	 * @param    [key]/[type]/[desc]
  	 * @param    {[type]}                 arg      [description]
  	 * @param    {Function}               callback [description]
  	 * @return   {[type]}                          [description]
  	 */
  	create(arg, callback) {
  	this.web3.personal.newAccount(secret.Data["etsc_password"]+arg.uid+"digifinex", (error,result) => {
            logstr(`eth create result: ${result}`, arg.coin);
            let obj = new Object();
            let checkObj = new Object();
            checkObj.address = result;
            checkObj.coin = arg.coin;
            let coin = arg.coin;
            let keystore_path = this.keystore_path;

            if(error){
                obj.ret = "fail";
                obj.msg = `eth create web3 error: ${JSON.stringify(error)}`;
                logstr(`eth create response: ${JSON.stringify(error)}`, arg.coin);
                callback(obj);
            }

            if(coin== "etc" || coin == "etsc" || coin=="eth"){
                let addressWithout0x = result.replace("0x","");
                let cmd = 'find '+keystore_path+' -name "*'+addressWithout0x+ '*"';
                logstr('exec before: ' + addressWithout0x, coin);
                exec(cmd, function(error, stdout, stderr) {
                    if (error) {
                        logstr('request kal create error: ' + cmd, coin);
                        obj.ret = "fail";
                        obj.msg = "error:" + JSON.stringify(error);
                        logstr(JSON.stringify(error), coin);
                        callback(obj);
                    } else {
                        logstr('exec after: 111', coin);
                        logstr(stdout, coin);
                        if(stdout){
                            obj.ret = "succ";
                            obj.uid = arg.uid;
                            obj.address = result;
                        }else{
                            obj.ret = "fail";
                            obj.msg = "create address not in the wallet";
                        }
                        callback(obj);
                    }
                });
            }else{
                this.check_address_in_wallet(checkObj,res=>{
                        let ifInRes = res.data;
                        if(ifInRes){
                            obj.ret = "succ";
                            obj.uid = arg.uid;
                            obj.address = result;
                        }else{
                            obj.ret = "fail";
                            obj.msg = "create address not in the wallet";
                        }
                        logstr(`eth create response: ${JSON.stringify(obj)}`, arg.coin);
                        callback(obj);
                    },
                    err=> {
                        obj.ret = "fail";
                        obj.msg = "something wrong when check address in wallet";
                        logstr(`eth create response: ${JSON.stringify(obj)}`, arg.coin);
                        callback(obj);
                    });
            }

        });	
  	}

    check_address_in_wallet(arg, succ, fail) {
  	    let obj = new Object();
        this.web3.eth.getAccounts(function(error, result) {
            logstr("Response received", arg.coin);
            if (error) {
                logstr(error, arg.coin);
                obj.ret = "fail";
                obj.msg = "error:" + JSON.stringify(error.message);
                logstr(JSON.stringify(error.message), arg.coin);
                fail(obj);
            } else {
                obj.ret = "fail";
                obj.data = false;
                for (var i = 0; i < result.length; i++){
                    if(arg.address == result[i]){
                        obj.ret = "succ";
                        obj.data = true;
                        break;
                    }
                }
                logstr(JSON.stringify(obj), arg.coin);
                succ(obj);
            }
        });
    }

  	/**
  	 * 执行真实转账接口
  	 * @Author   LiuWeiFu
  	 * @DateTime 2018-10-18T16:58:53+0800
  	 * @param    [key]/[type]/[desc]
  	 * @return   {[type]}                 [description]
  	 */
  	raw_transfer(transactionObject, succ, fail) {
  		logstr(`raw_transfer: ${JSON.stringify(transactionObject)}`, transactionObject.coin);
	    try {
	        if (transactionObject.coin_force == "1"){
	        	//...
	        }else{
	            transactionObject.to = this.g_address;
	        }
	        this.web3.eth.sendTransaction(transactionObject, (error, result) => {
	            logstr(`sendTransaction err: ${JSON.stringify(error)}`, transactionObject.coin);
	            logstr(`sendTransaction result: ${JSON.stringify(result)}`, transactionObject.coin);
	            if(!error)
	            {
	                logstr(`sendTransaction result in succ`, transactionObject.coin);
	                succ(result);
	            }else{
	                logstr('sendTransaction result in fail', transactionObject.coin);
	                fail(error);
	            }
	        });
	    } catch (err) {
	        logstr(`raw_transfer catch err: ${JSON.stringify(err)}`, transactionObject.coin);
	        fail(err);
	    }
  	}

  	/**
  	 * 转账统一调用接口
  	 * @Author   LiuWeiFu
  	 * @DateTime 2018-10-18T16:40:08+0800
  	 * @param    [key]/[type]/[desc]
  	 * @return   {[type]}                 [description]
  	 */
  	transfer(arg, callback) {
  		this.query(arg, (queryObj) => {
  		    if(!arg.gasPrice || !arg.gasLimit) {
  		        arg.gasPrice = this.gasPrice;
  		        arg.gasLimit = this.gasLimit;
            }
            let sender = arg.address,
            	receiver = arg.to_address,
            	coin_force = arg.coin_force,
            	before_amount = queryObj.bal,
            	amount = this.web3.toWei(before_amount, "ether"),
				transferAmount = (new bigNumber(amount)).minus(arg.gasPrice * arg.gasLimit).toString(10),
            	transactionObject = new Object(),
                passwd = secret.Data["eth_password"],
                obj = new Object();

            if(arg.is_old == "1") {
                passwd = secret.Data["eth_password"];
            } else {
                passwd = secret.Data["etsc_password"]+arg.uid+"digifinex";
            }

            if(arg.amount){
                transferAmount = this.web3.toWei(parseFloat(arg.amount), "ether");
            }

            logstr(`before_amount: ${before_amount}, amount: ${amount}, transferAmount: ${transferAmount}, gasLimit: ${arg.gasLimit}`, arg.coin);

            transactionObject.from = sender;
            transactionObject.to = receiver;
            transactionObject.value = transferAmount;
            transactionObject.gas = arg.gasLimit;
            transactionObject.gasPrice = arg.gasPrice;
            transactionObject.coin_force = coin_force;
            transactionObject.coin = arg.coin;

            logstr(`transactionObject: ${JSON.stringify(transactionObject)}`, arg.coin);
            if (amount > 0 ){
                this.raw_transfer(transactionObject,
                result => {
                    logstr(`succ, result: ${JSON.stringify(result)}`, arg.coin);
                    let obj = new Object();
                    obj.ret = "succ";
                    obj.uid = arg.uid;
                    obj.address = arg.address;
                    obj.to_address = arg.to_address;
                    obj.balance = this.web3.fromWei(queryObj.bal * this.token_digits, 'ether');
                    obj.amount = transferAmount / this.token_digits ;
                    obj.hashid = result;
                    logstr(`raw_transfer response: ${JSON.stringify(obj)}`, arg.coin);
                    callback(obj);
                },
                error => {
                    logstr(`err:${JSON.stringify(error)}, check need to unlock.`, arg.coin);
                    logstr(`error message: ${error.message}`, arg.coin);
                    if(error.message == "authentication needed: password or unlock" || error.message == "account is locked") {
                        if("0x021fc4ba776fb26cead2c5db6376affc33b95cc8" == transactionObject.from || "0x1905dd88042fb9e6ea5c3b02d47c93ad25865ce1" == transactionObject.from) {
                            passwd = "iwala.net";
                        }
                        this.web3.personal.unlockAccount(transactionObject.from, passwd, (err,result) => {
                            if(!err){
                                this.raw_transfer(transactionObject,
                                result => {
                                    logstr(result, arg.coin);
                                    obj.ret = "succ";
                                    obj.uid = arg.uid;
                                    obj.address = arg.address;
                                    obj.to_address = arg.to_address;
                                    obj.balance = this.web3.fromWei(queryObj.bal * this.token_digits, 'ether');
                                    obj.amount = transferAmount / this.token_digits;
                                    obj.hashid = result;
                                    logstr(obj, arg.coin);
                                    callback(obj);
                                },
                                error => {
                                    obj.ret = "fail";
                                    obj.msg = `second web3 error: ${JSON.stringify(error.message)}`;
                                    logstr(`second raw transfer err: ${JSON.stringify(error.message)}`, arg.coin);
                                    callback(obj);
                                });
                            }else{
                                obj.ret = "fail";
                                obj.msg = `unlock error: ${JSON.stringify(err.message)}`;
                                logstr(`unlock err: ${JSON.stringify(err.message)}`, arg.coin);
                                callback(obj);
                            }
                        });
                    }else{
                        obj.ret = "fail";
                        obj.msg = `web3 error: ${JSON.stringify(error)}`;
                        logstr(`raw transfer err: ${JSON.stringify(error)}`, arg.coin);
                        callback(obj);
                    }
                });
            }else{
                obj.ret = "fail";
                obj.msg = `balance of address: ${arg.address} is ${before_amount}`;
                callback(obj);
            }
        });
  	}

  	/**
  	 * 钱包余额查询
  	 * @Author   LiuWeiFu
  	 * @DateTime 2018-10-18T16:40:30+0800
  	 * @param    [key]/[type]/[desc]
  	 * @return   {[type]}                 [description]
  	 */
  	all_amount(arg, callback) {
  		let address = withdraw.Data[arg.coin],
  			obj = new Object();
        this.web3.eth.getBalance(address, (error, result) => {
            logstr(`eth all_amount response received`, arg.coin);

            if(error) {
                logstr(`eth all_amount error: ${JSON.stringify(error)}`, arg.coin);
                obj.ret = "fail";
                obj.msg = "web3 error:" + JSON.stringify(error);
            }else {
                logstr(`eth all_amount result: ${result}`, arg.coin);
                let accountBalance =  Number((new bigNumber(result)).dividedBy(this.token_digits).toString(10));

                obj.ret = "succ";
                obj.address = address;
                obj.coin = arg.coin;
                obj.bal = accountBalance;
                logstr(`eth all_amount response: ${JSON.stringify(obj)}`, arg.coin);
            }
            callback(obj);
        });
  	}

  	/**
  	 * 自动提币处理
  	 * @Author   LiuWeiFu
  	 * @DateTime 2018-10-18T17:54:16+0800
  	 * @param    [key]/[type]/[desc]
  	 * @param    {[type]}                 arg      [description]
  	 * @param    {Function}               callback [description]
  	 * @return   {[type]}                          [description]
  	 */
  	auto_draw(arg, callback) {
        let obj = new Object();
  		if ("hot_wallet" == arg.address ){
            let transfer_amount = parseFloat(arg.amount);
            if (!arg.amount || transfer_amount > withdraw.Data[arg.coin + "_max"]) {
                logstr(`auto draw, amount invalid, transfer_amount: ${transfer_amount}, g_auto_max_amount: ${withdraw.Data[arg.coin + "_max"]}`, arg.coin);
                obj.ret = "fail";
                obj.msg = "auto draw,amount invalid,transfer_amount:" + transfer_amount;
                callback(obj);
            }else{
                arg.address = withdraw.Data[arg.coin];
                arg.coin_force = "1";
                arg.uid="246926";
                logstr(`auto draw,para: ${JSON.stringify(arg)}`, arg.coin);
                arg.gasPrice = this.gasPrice;
                arg.gasLimit = this.gasLimit;
                this.transfer(arg, transferResp => {
                	callback(transferResp);
                });
            }
        }else{
            logstr(`auto draw,address invalid,address: ${arg.address}`, arg.coin);
            obj.ret = "fail";
            obj.msg = "auto draw,address invalid,address:" + arg.address;
            callback(obj);
        }
  	}

  	/**
  	 * 给指定地址添加汽油费
  	 * @Author   LiuWeiFu
  	 * @DateTime 2018-10-18T16:40:37+0800
  	 * @param    [key]/[type]/[desc]
  	 * @param    {[type]}                 arg [description]
  	 */
  	add_gas(arg, callback) {
        let gasPrice = this.gasPrice;
        let gasLimit = this.gasLimit;
        qianbaoDB.GetErc20Config(arg.real_coin, arg.real_link).then((data) => {
            if(data){
                if(data["merge_gas_price"]) {
                    gasPrice = data["merge_gas_price"];
                } else if(data["gas_price"]) {
                    gasPrice = data["gas_price"];
                }
                if(data["merge_gas_limit"]) {
                    gasLimit = data["merge_gas_limit"];
                } else if(data["gas_limit"]) {
                    gasLimit = data["gas_limit"];
                }
            }
            let obj = new Object(),
                passwd = secret.Data["eth_password"];
            //参数兼容一下，外面调用的时候填的参数是address
            if (!arg.to_address && arg.address) {
                arg.to_address = arg.address;
            }
            if (arg.to_address && arg.to_address != withdraw.Data[arg.coin]){
                //首先判断是不是我们的账户
                this.web3.personal.unlockAccount(withdraw.Data[arg.coin], passwd, (err, result) => {
                    logstr(`unlockAccount err: ${JSON.stringify(err)}`, arg.coin);
                    if(!err)
                    {
                        arg.address = withdraw.Data[arg.coin];
                        arg.coin_force = "1";
                        arg.amount = (gasPrice*gasLimit/this.token_digits*(1+0.2*Math.random())).toFixed(6);
                        arg.gasPrice = gasPrice;
                        arg.gasLimit = gasLimit;
                        logstr(`add_gas,para: ${JSON.stringify(arg)}`, arg.coin);
                        this.transfer(arg, transferResp => {
                            callback(transferResp);
                        });
                    }else{
                        obj.ret = "fail";
                        obj.msg = "unlock error:" + arg.to_address +",not our account";
                        logstr(`unlock err: ${JSON.stringify(err)}`, arg.coin);
                        callback(obj);
                    }
                });
            }else{
                obj.ret = "fail";
                obj.msg = `params error: ${JSON.stringify(arg)}`;
                callback(obj);
            }
        });
	}

  	trace(arg, callback) {
  		callback({ret: 'succ', msg: 'eth trace support in blockchain project.'});
  	}

    /**
     * [get_wallet_all_address 获取钱包所有地址]{"data":{"msg":"error: address=0x5c3e981e4544f8733d106f6e0711f082fca05f3f,result.to=0x5c3e981e4544f8733d106f6e0711f082fca05f3f,amount=11.95063900,result.value=11.950639,hash=0x02313462c9c8f9f1cc01a987fb63306d4c2dcef260e4e68339286f3312d97180,result.hash=0x02313462c9c8f9f1cc01a987fb63306d4c2dcef260e4e68339286f3312d97180","ret":"fail"},"errcode":0,"errmsg":"\u6210\u529f"}
     * @param  {[type]}   arg      [description]
     * @param  {Function} callback [description]
     * @return {[type]}            [description]
     */
    get_wallet_all_address(arg, callback) {
        var obj = new Object();

        logstr('get_wallet_all_address start:'+JSON.stringify(arg), arg.coin);

        this.web3.eth.getAccounts(function(error, result) {

            logstr("Response received", arg.coin);

            if (error) {
                logstr(error, arg.coin);
                obj.ret = "fail";
                obj.msg = "error:" + JSON.stringify(error);
                logstr(JSON.stringify(error), arg.coin);
                callback(obj);
            } else {
                //logstr(JSON.stringify(result.length), arg.coin);
                obj.ret = "succ";
                obj.data = result;
                // logstr(obj, arg.coin);
                callback(obj);
            }
        });
    }

    /**
     * [hash_query_info 通过hashid查找充提币信息]
     * @param  {[type]}   arg      [description]
     * @param  {Function} callback [description]
     * @return {[type]}            [description]
     */
    hash_query_info(arg, callback) {
        var obj = new Object();
        var hash = arg.hash ? arg.hash : '';
        var address = arg.address;
        var amount = arg.amount;
        var fee = 0;

        if (!hash || !address || !amount) {
            logstr('param error, please input hash、address、amount', arg.coin);
            obj.ret = "fail";
            obj.msg = "error: param error, please input hash、address、amount";
            callback(obj);
            return;
        }



        this.web3.eth.getTransaction(hash, (error, result) => {
            logstr(JSON.stringify(result), arg.coin);
            logstr("Response received", arg.coin);

            this.web3.eth.getTransactionReceipt(hash, (error1, result1) => {
                var gasUsed = 0;
                if (!error1) {
                    gasUsed = result1.gasUsed;
                }

                if (error) {
                    logstr(error, arg.coin);
                    obj.ret = "fail";
                    obj.fee = fee;
                    obj.msg = "error:" + JSON.stringify(error);
                    logstr(JSON.stringify(error), arg.coin);
                    callback(obj);
                } else {
                    var resultAddress = result.to;
                    var resultValue = this.web3.fromWei(result.value, 'ether');
                    var resultHash = result.hash;
                    fee = this.web3.fromWei(result.gasPrice * gasUsed, 'ether');

                    if (address == resultAddress && Math.abs(amount - resultValue) < 0.00000001 && hash == resultHash) {
                        logstr(JSON.stringify(result), arg.coin);
                        obj.ret = "succ";
                        obj.fee = fee;
                        obj.data = {
                            address: resultAddress,
                            amount: resultValue,
                            hash: resultHash
                        };
                        callback(obj);
                    } else {
                        obj.ret = "fail";
                        obj.fee = fee;
                        obj.msg = 'error: address='+address+',resultAddress='+resultAddress+',amount='+amount+',resultValue='+resultValue+',hash='+hash+',resultHash='+resultHash;
                        logstr(JSON.stringify(obj), arg.coin);
                        callback(obj);
                    }
                }
            });
        });
    }
}

exports.handler = eth;
