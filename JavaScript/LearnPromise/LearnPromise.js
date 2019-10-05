console.log('Hello World');
// new Promise(function(){});
function log(text){
  console.log(text);
}
// 承诺函数体
function test(resolve,reject){
    var timeOut = Math.random()*2;
    log(`延时${timeOut}秒`);
    setTimeout(()=>{
        if(timeOut <1){
            log(`调用resolve函数`);
            resolve('我从resolve传过去的值');
        }else{
            log(`调用reject函数`);
            reject('我从reject的传过去的值')
        }
    },timeOut*1000);
}

// 一个承诺,一定会执行test函数的
var p1 = new Promise(test);
log('设定承诺之后');
p1.then((result)=>{
    log(result);// 将传过来的数值打印
});
p1.catch((result)=>{
    log(result);
})
// 声明一个promise
var p2 = new Promise(function(resolve1,reject) {
  // 随机数
  let num = Math.random()*10;
  console.log(`num = ${num}`)
  if(num > 5){
    resolve1(num);
  }else{
    reject(num);
  }
});

function printResolve(val) {
  console.log(`resolve 回调 num =${val}`);
}
function printReject(val){
  console.log(`reject 回调 num = ${val}`);
}

// 使用promise
p2.then(printNum).catch(printReject);
