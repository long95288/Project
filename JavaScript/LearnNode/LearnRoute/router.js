// 路由模块，打印路由信息

function route(pathname){
	console.log("About to route a request for"+pathname);
}
// 暴露函数
exports.route = route;
