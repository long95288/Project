// 全局变量
var num = 1;
function print(){
	console.log("hello world"+num);
	num+=1;	
	if(num == 10) {
		console.log("end");
	 clearInterval(s);
	}
}

var s = setInterval(print,1000);
