var server = require('./LearnServe');
var router = require('./router');

server.start(router.route);
