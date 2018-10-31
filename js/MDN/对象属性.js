var myObj = {},
    str = "myString",
    rand = Math.random(),
    obj = {};

myObj.type              = "属性 => 点语法";
myObj["date created"]   = "属性 => 包含空格的字符串";
myObj[str]              = "属性 => 字符串的值";
myObj[rand]             = "属性 => 随机数字";
myObj[obj]              = "属性 => 对象";
myObj[""]               = "属性 => 空字符串";
console.log(myObj);

var myCar = {};
var propertyName = "make";
myCar[propertyName] = "Ford";
propertyName = "model";
myCar[propertyName] = "Mustang";

console.log(myCar);