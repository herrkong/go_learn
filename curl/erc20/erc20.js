/**
 * ERC20处理类
 * @author Bourne<liuweifu@digifinex.com>
 * 2018-09-21 16:20:22
 */

const web3_extended = require('web3_ipc');
const { logstr, sendQywxMsg, sendJson } = require('../my_modules/common.js');//公共函数库
const common = require('../my_modules/common.js');//公共函数库
const bigNumber = require('bignumber.js');
const secret = require("../my_modules/secret.js");
const withdraw = require("../my_modules/withdraw.js");
var web3, contract;
const erc20 = class {
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

		if (!web3) {
    		// web3.currentProvider.connection.destroy();
			this.web3 = web3_extended.create(options);
			web3 = this.web3;
		}else {
    		        this.web3 = web3;
			if(!web3.currentProvider.connection.writable){
				web3.currentProvider.connection.destroy();
				this.web3 = web3_extended.create(options);
				web3 = this.web3;
			}
		}
		this.contract = this.web3.eth.contract(this.abi).at(this.contract_address);


  	}

  	/**
  	 * ERC20查询地址余额
  	 * @Author   LiuWeiFu
  	 * @DateTime 2018-10-22T12:30:28+0800
  	 * @param    [key]/[type]/[desc]
  	 * @param    {[type]}                 arg      [description]
  	 * @param    {Function}               callback [description]
  	 * @return   {[type]}                          [description]
  	 */
  	query(arg, callback) {
  		let { address } = arg,
  			obj = {};

	    this.contract.balanceOf(address, (error, result) => {
	        logstr("Response received", arg.coin);

	        if(error) {
	            logstr(error, arg.coin);
	            obj.ret = "fail";
	            obj.msg = `web3 error: ${JSON.stringify(error)}`;
	        }else{
	            logstr(`result: ${result}`, arg.coin);
	            let accountBalance =  Number((new bigNumber(result.toString())).dividedBy(this.token_digits).toString(10));
	            //let accountBalance =  Number((new bigNumber(result.toString())).dividedBy(this.token_digits).toString(10));
	            obj.ret = "succ";
	            obj.address = address;
	            obj.bal = accountBalance;
				obj.strBal = (new bigNumber(result.toString())).dividedBy(this.token_digits).toString(10);
	        }
	        logstr(JSON.stringify(obj), arg.coin);
            callback(obj);
	    });
  	}

  	/**
  	 * ERC20地址创建
  	 * @Author   LiuWeiFu
  	 * @DateTime 2018-10-22T12:30:53+0800
  	 * @param    [key]/[type]/[desc]
  	 * @param    {[type]}                 arg      [description]
  	 * @param    {Function}               callback [description]
  	 * @return   {[type]}                          [description]
  	 */
  	create(arg, callback) {
        let obj = new Object();
  		try{
	        this.web3.personal.newAccount(secret.Data["erc20_password"],function(error,result){
	            if(!error){
	                logstr(`${arg.coin} create result: ${result}`, arg.coin);
	                obj.ret = "succ";
	                obj.uid = arg.uid;
	                obj.address = result;
	                logstr(obj, arg.coin);
	            }else{
	                obj.ret = "fail";
	                obj.msg = "web3 error:"+ JSON.stringify(error);
	                logstr(JSON.stringify(error), arg.coin);
	            }
                callback(obj);
	        });
	    }catch(e){
	        logstr(`req_handle catch, error ${JSON.stringify(e)}, arg: ${JSON.stringify(arg)}`, arg.coin);
	        obj.ret = "fail";
	        obj.msg = `${arg.coin} address create error`;
	        logstr(JSON.stringify(e), arg.coin);
            callback(obj);
	    }
  		return 'eth create';
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
	        if (transactionObject.coin_force == "1") {
	        	//...
	        }else{
	            transactionObject.to = this.g_address;
	        }
	        this.contract.transfer(transactionObject.to,transactionObject.amount,transactionObject, (error,result) => {
	            logstr(`sendTransaction err: ${JSON.stringify(error)}`, transactionObject.coin);
	            let  errStr = this.err_2_str(error);
	            logstr("sendTransaction err:" + errStr, transactionObject.coin);
	            logstr(`sendTransaction result: ${JSON.stringify(result)}`, transactionObject.coin);
	            if(!error)
	            {
	                logstr('sendTransaction result in succ', transactionObject.coin);
	                succ(result);
	            }
	            else
	            {
	                logstr('sendTransaction result in fail', transactionObject.coin);
	                fail(error);
	            }
	        });
	    } catch (err) {
	        logstr(`raw_transfer catch err: ${JSON.stringify(err)}`, transactionObject.coin);
	        fail(err);
	    }
	}
	/***
	 **  输出错误码
	***/
	err_2_str(error)
	{
	    return error?("name:" + error.name +",msg:"+error.message):" null";
	}

	/**
  	 * 转账统一调用接口
  	 * @Author   LiuWeiFu
  	 * @DateTime 2018-10-18T16:40:08+0800
  	 * @param    [key]/[type]/[desc]
  	 * @return   {[type]}                 [description]
  	 */
  	transfer(arg, callback) {
  		try{
	        this.query(arg, queryObj => {
	            let sender = arg.address,
	            	receiver = arg.to_address,
	            	coin_force = arg.coin_force,
                	transactionObject = new Object(),
                	obj = new Object();
				var before_amount = queryObj.strBal;
				var amount = (new bigNumber(before_amount.toString()).multipliedBy(this.token_digits).toString(10));

	            logstr(`before_amount: ${before_amount}, arg: ${JSON.stringify(arg)}, queryObj: ${JSON.stringify(queryObj)}, amount: ${amount}`, arg.coin);
	            if(arg.amount > 0)
	            {
	                let diff = new bigNumber(0.000001);
					let before_amount_big = new bigNumber(before_amount);
					let is_more_query = before_amount_big.isGreaterThan(new bigNumber(arg.amount));// 是否有更多的余额

					if (arg.isNotMerge) {
						before_amount = (arg.amount);
					} else {
						if (is_more_query) {// 如果够
							before_amount = (arg.amount);
						} else {
							let is_more_diff = diff.isGreaterThan(new bigNumber(arg.amount).minus(before_amount_big));// 是否超过精度
							if (!is_more_diff) {// 相差较多不归集
								let obj = new Object();
								obj.ret = "fail";
								obj.msg = "error: transfer amount large difference";
								logstr(obj.msg, arg.coin);
								callback(obj);
								return;
							}
						}
					}

	                // before_amount = (arg.amount);
	                //amount = Number(new bigNumber(before_amount.toString()).multipliedBy(this.token_digits).toString(10));
	                amount = (new bigNumber(before_amount.toString()).multipliedBy(this.token_digits).toString(10));
	                logstr(`22 before_amount: ${before_amount}, arg: ${JSON.stringify(arg)}, queryObj: ${JSON.stringify(queryObj)}, amount: ${amount}`, arg.coin);
	            } else {
	            	//避免将钱包余额全部转出！
                    obj.ret = "fail";
                    obj.msg = "error: transfer amount <= 0";
                    logstr(obj.msg, arg.coin);
                    callback(obj);
                    return;
				}

	            logstr(`after numsub before_amount: ${before_amount}, arg: ${JSON.stringify(arg)}, queryObj: ${JSON.stringify(queryObj)}, amount: ${amount}`, arg.coin);

                logstr(`amount: ${amount}, token_digits: ${this.token_digits}, gasLimit: ${this.gasLimit}, gasPrice: ${this.gasPrice}`, arg.coin);

                transactionObject.from = sender;
                transactionObject.to = receiver;
                transactionObject.gas = this.gasLimit;
                transactionObject.gasPrice = this.gasPrice;
                transactionObject.coin_force = coin_force;
                transactionObject.amount = (amount);
                transactionObject.coin = arg.coin;

                logstr(`transactionObject: ${JSON.stringify(transactionObject)}`, arg.coin);
                if (amount > 0 )
                {
                    this.raw_transfer(transactionObject,
                    result => {
                        logstr("succ,result:"+JSON.stringify(result), arg.coin);
                        obj.ret = "succ";
                        obj.uid = arg.uid;
                        obj.address = arg.address;
                        obj.to_address = arg.to_address;
                        obj.amount = Number(amount);
                        obj.hashid = result;
                        logstr(obj, arg.coin);
                        callback(obj);
                    },
                    error => {
                        logstr("err transactionObject step1 :"+ JSON.stringify(transactionObject), arg.coin);
                        logstr("err:"+ JSON.stringify(error)+",check need to unlock.", arg.coin);
                        if(error.message == "authentication needed: password or unlock")
                        {
		            var pwd = secret.Data["erc20_password"];
			    if(transactionObject.from=="0x13819886a3aa607d4d9eb93501b09bae75944158"){
				 pwd = secret.Data["erc202_password"];
		            } else if (transactionObject.from=="0xb2dd30f7aed2a92af29a8a85e0b7ce54fabe01de") {
				 pwd = secret.Data["erc203_password"];
		            } else if (transactionObject.from=="0x054afc88da199d5da42dd5df4305066a7bad0250") {
                                 pwd = secret.Data["usdte_password"];
                            }
                            this.web3.personal.unlockAccount(transactionObject.from, pwd, (err,result) => {
                                logstr("unlockAccount err:"+ JSON.stringify(err), arg.coin);
                                if(!err)
                                {
                                    transactionObject.to = arg.to_address;
                                    logstr("transactionObject step2 :"+ JSON.stringify(transactionObject), arg.coin);
                                    this.raw_transfer(transactionObject,
                                    result => {
                                        logstr(result, arg.coin);
                                        obj.ret = "succ";
                                        obj.uid = arg.uid;
                                        obj.address = arg.address;
                                        obj.to_address = arg.to_address;
                                        obj.amount = Number(amount);
                                        obj.hashid = result;
                                        logstr(obj, arg.coin);
                                        callback(obj);
                                    },
                                    error => {
                                            obj.ret = "fail";
                                            obj.msg = "second web3 error:" + this.err_2_str(error);
                                            logstr("second raw transfer err:"+this.err_2_str(error), arg.coin);
                                            callback(obj);
                                    });
                                }
                                else
                                {
                                    obj.ret = "fail";
                                    obj.msg = "unlock error:" + JSON.stringify(err);
                                    logstr("unlock err:"+JSON.stringify(err), arg.coin);
                                    callback(obj);
                                }
                            });
                        }
                        else
                        {
                            obj.ret = "fail";
                            obj.msg = "web3 error:" + JSON.stringify(error);
                            logstr("raw transfer err:"+JSON.stringify(error), arg.coin);
                            callback(obj);
                        }
                    });
                }
                else
                {
                    obj.ret = "fail";
                    obj.msg = "balance of address:"+arg.address +" is "+before_amount;
                    callback(obj);
                }
	        })
	    }
	    catch(e)
	    {
	        logstr('req_handle catch,name:'+ e.name +',msg:'+e.message+",arg:"+ JSON.stringify(arg), arg.coin);
	        let obj = new Object();
	        obj.ret = "fail";
	        obj.msg = `${arg.coin} transfer error: ${JSON.stringify(e)}`;
	        logstr(e, arg.coin);
	        callback(obj);
	    }
  	}

  	/**
  	 * 钱包余额查询
  	 * @Author   LiuWeiFu
  	 * @DateTime 2018-10-18T16:40:30+0800
  	 * @param    [key]/[type]/[desc]
  	 * @return   {[type]}                 [description]
  	 */
  	all_amount(arg, callback) {
  		logstr("arg:"+JSON.stringify(arg), arg.coin);
        let obj = new Object();
	    try
	    {
               let address = "";
            if(arg.link_type) {
                address = withdraw.Data[arg.coin + "_" + arg.link_type];
            } else if (withdraw.Data[arg.coin]) {
	        address = withdraw.Data[arg.coin];
			logstr("arg.coin address:" + withdraw.Data[arg.coin],arg.coin)
	    } else {
                address = withdraw.Data["erc20"];
				logstr("erc20 address:" + withdraw.Data["erc20"],arg.coin)
            }
			logstr("address:" + address, arg.coin);
			logstr("this.contract address:" + this.contract_address,arg.coin)
			logstr("this.abi:" + this.abi,arg.coin)
	        this.contract.balanceOf(address, (error,result) => {
	            logstr("Response received", arg.coin);
	            if(error) {
	                logstr(error, arg.coin);
	                obj.ret = "fail";
	                obj.msg = "web3 error:" + JSON.stringify(error);
	                logstr(JSON.stringify(error), arg.coin);
	                callback(obj);
	            }
	            else
	            {
	                logstr(JSON.stringify(error) +",query result:" + result, arg.coin);
	                //let accountBalance = result.toNumber()/this.token_digits;
	                let accountBalance =  Number((new bigNumber(result.toString())).dividedBy(this.token_digits).toString(10));
	                obj.ret = "succ";
	                obj.address = address;
	                obj.bal = accountBalance;
	                logstr(obj, arg.coin);
	                callback(obj);
	            }
	        });
	    }
	    catch(e)
	    {
	        logstr('req_handle catch,error '+ JSON.stringify(e) +",arg:"+ JSON.stringify(arg) + this.token_digits, arg.coin);
	        obj.ret = "fail";
	        obj.coin = arg.coin;
	        obj.msg = JSON.stringify(e);
	        logstr(JSON.stringify(e), arg.coin);
	        callback(obj);
	    }
  	}

  	/**
  	 * ERC20自动提币
  	 * @Author   LiuWeiFu
  	 * @DateTime 2018-10-22T17:04:18+0800
  	 * @param    [key]/[type]/[desc]
  	 * @return   {[type]}                 [description]
  	 */
  	auto_draw(arg, callback) {
  		let obj = new Object();
  		if ("hot_wallet" == arg.address ){
            let transfer_amount = parseFloat(arg.amount);
            if (!arg.amount || transfer_amount > this.g_auto_max_amount) {
                logstr(`auto draw, amount invalid, transfer_amount: ${transfer_amount}, g_auto_max_amount: ${this.g_auto_max_amount}`, arg.coin);
                obj.ret = "fail";
                obj.msg = "auto draw,amount invalid,transfer_amount:" + transfer_amount;
                callback(obj);
            }else{
                if(withdraw.Data[arg.coin]) {
                    arg.address = withdraw.Data[arg.coin];
				} else {
                    arg.address = withdraw.Data["erc20"];
				}
				arg.coin_force = "1";
				arg.isNotMerge = true;
                logstr(`auto draw,para: ${JSON.stringify(arg)}`, arg.coin);
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
                //logstr(JSON.stringify(result), arg.coin);
                obj.ret = "succ";
                obj.data = result;
                //logstr(obj, arg.coin);
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
            fail(obj);
            return;
        }

        this.web3.eth.getTransactionReceipt(hash, (error, result) => {

		if (!result) {
                obj.ret = "fail";
                obj.fee = 0;
                obj.msg = 'error: not result';
                logstr(JSON.stringify(obj), arg.coin);
                callback(obj);
                return;
            }
            logstr("Response received", arg.coin);

			this.web3.eth.getTransaction(hash, (error1, result1) => {
			if (!result1) {
                    obj.ret = "fail";
                    obj.fee = 0;
                    obj.msg = 'error: not result1';
                    logstr(JSON.stringify(obj), arg.coin);
                    callback(obj);
                    return;
                }
				var gasPrice = 0;
				if (!error1) {
					gasPrice = result1.gasPrice;
				}
				if (error) {
					logstr(error, arg.coin);
					obj.ret = "fail";
					obj.fee = fee;
					obj.msg = "error:" + JSON.stringify(error);
					logstr(JSON.stringify(error), arg.coin);
					callback(obj);
				} else {
					var logs = new Object();
					fee = this.web3.fromWei(result.gasUsed * gasPrice, 'ether');

					if (result.logs.length > 0) {
						for (var i = 0, len = result.logs.length; i < len; i++) {
							if (result.logs[i].transactionHash == hash) {
								logs = result.logs[i];
							}
						}
					}

					if (typeof logs.topics == 'undefined') {
						obj.ret = "fail";
						obj.fee = fee;
						obj.msg = "hashid request error:" + JSON.stringify(logs);
						logstr(JSON.stringify(logs), arg.coin);
						callback(obj);
						return;
					}

					var resultTo = logs.topics[2].replace(/^0x0+/, "");
					resultTo =  '0x' + Array(40 - resultTo.length + 1).join('0') + resultTo;
					var resultValue = Number((new bigNumber(parseInt(logs.data).toString())).dividedBy(this.token_digits).toString(10));

					if (address == resultTo && Math.abs(amount - resultValue) < 0.00000001 && hash == result.transactionHash) {
						logstr(JSON.stringify(result), arg.coin);
						obj.ret = "succ";
						obj.fee = fee;
						obj.data = {
							address: resultTo,
							amount: resultValue,
							hash: result.transactionHash
						};
						callback(obj);
					} else {
						obj.ret = "fail";
						obj.fee = fee;
						obj.msg = 'error: address='+address+',result.to='+resultTo+',amount='+amount+',result.value='+resultValue+',hash='+hash+',result.hash='+result.transactionHash;
						logstr(JSON.stringify(obj), arg.coin);
						callback(obj);
					}
				}
			});
        });
    }


  	trace(arg, callback) {
  		callback({ret: 'succ', msg: `${arg.coin} trace support in blockchain project.`});
  	}
}

exports.handler = erc20;
