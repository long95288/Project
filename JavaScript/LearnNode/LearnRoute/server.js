// 服务器模块
var http = require('http');
var url = require('url')

function start(route){
    function OnRequest(request,response){
        var pathname = url.parse(request.url).pathname;
        console.log(`从${pathname}处发送请求`);
        // 路由函数获得路由信息
        route(pathname);

        response.writeHead(200,{'Content-Type':'text/plain'});
        response.write('Hello word');
        response.end();
    }
    
    http.createServer(OnRequest).listen(8888);
    console.log("http://localhost:8888");
}

// 暴露接口
exports.start=start;
