// 引入server和router并调用函数
var server = require('./server');
var router = require('./router');

server.start(router.route);
