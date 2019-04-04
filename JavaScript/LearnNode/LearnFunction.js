
// 定义一个函数传递体
function execute(someFunction,value){someFunction(value)}

// 非匿名函数
function say(word){console.log('hello:'+word)}

// 使用
execute(say,"world");

// 使用非匿名函数
execute((word)=>{console.log(word)},'hello world');

// 服务器
var http = require('http');
function OnRequest(request,response){
    response.writeHead(200,{'Content-Type':'text/plain'});
    response.write('hello serve');
    response.end();
}

http.createServer(OnRequest).listen(8888);
console.log("serve:http://127.0.0.1:8888/");