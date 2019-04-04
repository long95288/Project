process.on('exit',(code)=>{
    
    // timeout不会被执行 
    setTimeout(()=>{console.log('该代码不会被执行')},0);

    console.log(`退出码${code}`);
});

process.on('beforeExit',()=>{
    console.log('退出之前调用');
});

process.on('uncaughtException',()=>{
    
});