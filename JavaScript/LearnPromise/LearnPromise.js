console.log('Hello World');
// new Promise(function(){});
function log(text){
  console.log(text);
}
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
