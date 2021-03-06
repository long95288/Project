var fs = require('fs');
var data = '';

// 创建可读流
var readerStream = fs.createReadStream('input.txt');

// 设置编码为 utf-8
readerStream.setEncoding('UTF8');

// 处理流事件

readerStream.on('data',function(chunk){
	data += chunk;
})

readerStream.on('end',function(){
	console.log(data);
})

readerStream.on('error',function(){
	console.log(err.stack);
})


// 创建写入流
var writerStream = fs.createWriteStream('ouput.txt');

// 使用utf-8写入数据
data = "写入的内容";
writerStream.write(data,'UTF-8');

writerStream.on('finish',function(){
	console.log("写入完成");
})

writerStream.on('error',function(){
	console.log(err.stack);
})

console.log("程序执行完毕");
