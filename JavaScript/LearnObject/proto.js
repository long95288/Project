/**
 * 面向对象的思想
 */

 /**
  * Student 对象原型
  */
 var Student = {
    name: 'Robot',
    height: 1.2,
    run:function(){
        console.log(this);
        console.log(this.name +'is running ..');
    }
 };

 function createStudent(name) {
    // 基于Student对象原型创建一个新对象
    var s = Object.create(Student);
    console.log(name);
    s.name = name;
    return s;
 }

 var xiaoming = new createStudent('小明');
 xiaoming.run();
 xiaoming.__proto__ === Student;
 var xiaoli = {
    name: '小李'
 };
 // 原型链继承对象
 xiaoli.__proto__ = Student;
 console.log(xiaoli.name);
 xiaoli.run();
 /**
  * 构造函数构造对象
  */
 function Student2(name) {
     this.name = name;
     this.hello = function(){
        console.log('hello,'+this.name+'!');
     }
 }

// 使用new来构造对象
var zhangsan = new Student2('张三');
console.log(zhangsan.name);
zhangsan.hello();