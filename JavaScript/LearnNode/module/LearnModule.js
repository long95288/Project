// 引用hello模块
var hello1 = require('./hello');

// Hello2是一个对象模版，需要实例化才能使用
var Multi = require('./exportModule');

// 未实例化便调用，发生错误
// Hello2.setName("ni");
// Hello2.sayHello();

// 调用函数
hello1.world();
// 实例化1
var Hello3 = new Multi.Hello();
Hello3.setName("ff");
Hello3.sayHello();

// 实例化2
var Hello4 = new Multi.Hello();
Hello4.setName("Hello4");
Hello4.sayHello();

// 实例化Print2
print = new Multi.Print2();
print.print();