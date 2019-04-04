// __filename 文件名
console.log("filename:"+__filename);

// __dirname所在的目录
console.log(`dirname:${__dirname}`);

// setTimeout(cb,ms)
var t=setTimeout(()=>{console.log(`2秒后了`)},2000);

// clearTimeout(t)
var i =1;
var interval = setInterval(()=>{
    if(i==10){
        clearInterval();
    }else{
        i+=1;
        console.log(`i=${i}`);
    }
},2000);
console.log('结束');

// console使用
console.info("消息");

console.error("错误");

console.time("正在执行");
console.timeEnd("正在执行");