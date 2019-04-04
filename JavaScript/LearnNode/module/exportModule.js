function Hello(){
	var name;
	this.setName = function(thyname){
		name = thyname; 
	};
	this.sayHello = function(){
		console.log('Hello'+name);
	}
}

function Print2(){
	let out = "第二个对象";
	this.print =()=>{
		console.log(out);
	}
}

// 暴露出去两个对象
module.exports = {Hello,Print2};
