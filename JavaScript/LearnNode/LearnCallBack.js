var fs = require("fs");
var events = require("events");

console.log("hello word");

var eventEmitter = new events.EventEmitter();

// 绑定事件及处理程序
eventEmitter.on("print",function(){
fs.readFile('input.txt',function(err,data){
	if(err) return console.error(err);
	console.log("文件数据:"+data.toString());
});	
});

var print = function print(filename){
fs.readFile(filename,function(err,data){
	if(err) return console.error(err);
	console.log("文件数据:"+data.toString());

})
}

// 触发事件
eventEmitter.emit('print');

eventEmitter.on('printfile',print);

eventEmitter.emit('printfile');
