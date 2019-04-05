process.on('exit',(code)=>{
    
    // timeout不会被执行 
    setTimeout(()=>{console.log('该代码不会被执行')},0);

    console.log(`退出码${code}`);
});

process.on('beforeExit',()=>{
    console.log('退出之前调用');
});

process.on('uncaughtException',()=>{
   console.log('异常处理');
});

// stdout
process.stdout.write("Hello stdout");

// argv
process.argv.forEach((val,index,array) => {
    console.log(`参数${index}:${val}`);
});

// 获得执行路径
console.log(`路径:${process.execPath}`);

// 平台信息
console.log(`平台:${process.platform}`);

// cwd
console.log(`当前目录:${process.cwd()}`);

// 当前版本
console.log(`当前版本:${process.version}`);

// 内存使用情况
console.log(`内存使用情况${process.memoryUsage()}`);
console.log(process.memoryUsage());

// 