const { logstr, sendQywxMsg, sendJson } = require('../my_modules/common.js');//公共函数库
const common = require('../my_modules/common.js');//公共函数库
const bigNumber = require('bignumber.js');
const { request } = require('http');
const { exec } = require('child_process');
const secret = require("../my_modules/secret.js");
const withdraw = require("../my_modules/withdraw.js");

const apl = class {
    //初始化
    constructor(initData) {
        let keys = Object.keys(initData);
        for(let [index, key] of keys.entries()){
            this[key] = initData[key];
        }
    }

    myRequest(path="",isPost = false, port=7876){
        let that=this
        return new Promise(function(resolve, reject)
        {
            if(isPost){
                var cmd = `curl -X POST "https://${that.host}:${that.port}/apl?${path}"`;
            }else{
                var cmd = `curl "https://${that.host}:${that.port}/apl?${path}"`;
            }
    
            exec(cmd, (error, stdout, stderr) => {
                if (error) {
                    logstr(`exec error: ${JSON.stringify(error)}`);
                    reject(error);
                }
                resolve(stdout);
            });
        });
    }
    
    async asyncHttpReq(path,isPost) {
        let ret = await this.myRequest(path,isPost);
        return JSON.parse(ret);
    }


    async query(arg, callback) {
        let obj = new Object(),
            address = arg.address;
        let path =  `requestType=getBalance&account=${address}`;
        let result = await this.asyncHttpReq(path,false);
        let accountBalance = result.balanceATM / this.token_digits;
        obj.ret = "succ";
        obj.address = arg.address;
        obj.bal = accountBalance;
        callback(obj);
    }


    async create(arg, callback) {
        try{
            let obj = new Object();
            let uid = arg.uid;
            let pass = parseInt(uid) + 888888;
            let path = `requestType=getAccountId&secretPhrase=${pass}`;

            let ret = await this.asyncHttpReq(path,false);
            logstr(ret,arg.coin);
            obj.ret = "succ";
            obj.uid = arg.uid;
            obj.address = ret.accountRS;
            callback(obj);
        }catch (e) {
            logstr('apl create error: '+ JSON.stringify(e.message) + ",arg: "+ JSON.stringify(arg),arg.coin);
            var obj = new Object();
            obj.ret = "fail";
            obj.msg = "apl create address error.";
            logstr(JSON.stringify(e.message),arg.coin);
            callback(obj);
        }

    }


    async raw_transfer(transactionObject, succ, fail) {
        logstr(`raw_transfer: ${JSON.stringify(transactionObject)}`, transactionObject.coin);
        try {
            if (transactionObject.coin_force == "1"){

            }else{
                transactionObject.to = this.g_address;
            }

            if(transactionObject.from == withdraw.Data['apl']){
                var secretWithdraw = secret.Data["apl_g_auto_out_address_password"];
		//var secretWithdraw = "1881digifinexpassnew";	
            }else{
                var secretWithdraw = parseInt(transactionObject.uid) + 888888;
            }
            
	   let totalAmount = transactionObject.amount*this.token_digits;
            let stringAmount = new Number(totalAmount);
            let stringAmount1 = stringAmount.toLocaleString();
            var amountAtm = 0;
            if(stringAmount1.indexOf(".") != -1){
                let stringAmount1Array = stringAmount1.split(".");
                let amountAtmActual = stringAmount1Array[0];
                amountAtm = amountAtmActual.replace(/,/g,'');
            }else{
                amountAtm = stringAmount1.replace(/,/g,'');
            } 
 
	    //let amountAtm = transactionObject.amount * this.token_digits;
            //http://localhost:7876/apl?requestType=sendMoney&secretPhrase=123456&recipient=APL-CVGQ-L6NM-6LAU-G4TH3&amountATM=1000000000&feeATM=100000000&deadline=600
            let path = `requestType=sendMoney&secretPhrase=${secretWithdraw}&recipient=${transactionObject.to}&amountATM=${amountAtm}&feeATM=100000000&deadline=600`;

            let transferRet = await this.asyncHttpReq(path,true);
            let error = transferRet.errorDescription ? transferRet.errorDescription : '';
	    logstr(`sendTransaction result: ${JSON.stringify(transferRet)}`,transactionObject.coin);
	    if(!error)
            {
                logstr("sendTransaction result in succ");
                succ(transferRet);
            }
            else
            {
                logstr("sendTransaction result in fail");
                fail(error);
            }


        } catch (err) {
            logstr(`raw_transfer catch err: ${JSON.stringify(err.Message)}`, transactionObject.coin);
            fail(err);
        }
    }


    transfer(arg, callback) {
        //如果amount只能是整型
        try
        {
            this.query(arg, queryObj => {
                let sender = arg.address,
                    receiver = arg.to_address,
                    amount = Number(arg.amount),
                    coin_force = arg.coin_force,
                    addressBal = Number(queryObj.bal),
                    uid = arg.uid;

                logstr(`sender: ${sender}; receiver: ${receiver}; amount: ${amount}; addressBal: ${addressBal}`, arg.coin);

                let transactionObject = new Object();
                transactionObject.from = sender;
                transactionObject.to = receiver;
                transactionObject.coin_force = coin_force;
                transactionObject.amount = amount;
                transactionObject.coin = arg.coin;
                transactionObject.uid = arg.uid;
                transactionObject.tag = arg.tag ? arg.tag : '';
                logstr("transactionObject:"+ JSON.stringify(transactionObject), arg.coin);
                if (amount <= addressBal)
                {
                    this.raw_transfer(transactionObject, tranRet => {
                            let transResult = tranRet;
                            logstr(transResult, arg.coin);
                            let obj = new Object();
                            obj.ret = "succ";
                            obj.uid = arg.uid;
                            obj.address = arg.address;
                            obj.to_address = arg.to_address;
                            obj.amount = amount;
                            obj.hashid = transResult.fullHash;
                            logstr(obj, arg.coin);
                            callback(obj);
                        },
                        error => {
                            let obj = new Object();
                            obj.ret = "fail";
                            obj.msg = `apl raw_transfer error: ${JSON.stringify(error)}`;
                            logstr(`raw transfer error: ${JSON.stringify(error)}`, arg.coin);
                            callback(obj);
                        });
                }else{
                    let obj = new Object();
                    obj.ret = "fail";
                    obj.msg = `balance of address: ${arg.address} is ${addressBal}`;
                    callback(obj);
                }

            })
        }catch(e){
            logstr('req_handle catch,name:'+ e.name +',msg:'+e.message+",arg:"+ JSON.stringify(arg), arg.coin);
            let obj = new Object();
            obj.ret = "fail";
            obj.msg = "catch error when theta transfer";
            logstr(e, arg.coin);
            callback(obj);
        }
    }


    async all_amount(arg, callback) {
        arg.address = withdraw.Data[arg.coin];
        this.query(arg, queryObj => {

            let address = arg.address;
            let addressBal = Number(queryObj.bal);
            let obj = new Object();

            obj.ret = "succ";
            obj.address = address;
            obj.coin = arg.coin;
            obj.bal = addressBal;
            logstr(`bar all_amount response: ${JSON.stringify(obj)}`, arg.coin);

            callback(obj);

        });
    }


    auto_draw(arg, callback) {
        //如果amount只能是整型
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

    trace(arg, callback) {
        callback({ret: 'succ', msg: `${arg.coin} trace support in blockchain project.`});
    }

    //hash检查
    async hash_query_info(arg, callback){
        var obj = new Object();
        var hash = arg.hash ? arg.hash : '';
        var address = arg.address;
        var amount = arg.amount;

        //  http://localhost:7876/apl?requestType=getTransaction&fullHash=75cc7772eba77341317dc215c8deea96a14551d3842a8bd2024f514604f5e866
        let path = `requestType=getTransaction&fullHash=${hash}`;

        let ret = await this.asyncHttpReq(path);
        logstr(ret, arg.coin);

        var chainAddress = "";
        var chainAmount =  "";

        if(ret.error != null){
            obj.ret = "fail";
            obj.msg = 'error:'+ JSON.stringify(ret.error);
            logstr(JSON.stringify(obj), arg.coin);
            callback(obj);
        }else{
            chainAddress = ret.recipientRS;
            chainAmount = ret.amountATM / this.token_digits;
        }

        if(chainAddress == address && Math.abs(chainAmount - amount) < 0.00001){
            obj.ret = "succ";
            obj.data = {
                address: address,
                amount: chainAmount,
                hash: hash
            };
            callback(obj);
        }else{
            obj.ret = "fail";
            obj.msg = 'error: address='+address+',resultAddress='+chainAddress+',amount='+amount+',resultValue='+chainAmount;
            logstr(JSON.stringify(obj), arg.coin);
            callback(obj);
        }
    }

}

exports.handler = apl;
