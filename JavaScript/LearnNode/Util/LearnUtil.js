var util = require('util')
// util.inherits
function Base(){
    this.name='base'
    this.base=1991
    this.sayHello=()=>{
        console.log(`hello:${this.name}`)
    }
}
Base.prototype.showName=()=>{
    console.log(this.name)
}
function Sub(){this.name='sub'}

util.inherits(Sub,Base)

var objBase = new Base();
objBase.showName();
objBase.sayHello();
console.log(objBase);
var objSub = new Sub();
objSub.showName();
console.log(objSub);

// util.inspect(obj)
function Person(){
    this.name='byvoid';
    this.toString =()=>{
        return this.name;
    }
}
var obj = new Person();
console.log(util.inspect(obj,true))

// util.isArray(object)

console.log(util.isArray([]));

console.log(util.isArray(new Array));

console.log(util.isArray({}));

