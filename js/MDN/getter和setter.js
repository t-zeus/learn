/**
 * getter：获取某个属性的值
 *      语法：get 属性() {}
 * setter：设置某个属性的值
 *      语法：set 属性(x) {}
 */

var o = {
    a: 7,
    get b() {
        return this.a + 1;
    },
    set c(x) {
        this.a = x / 2;
    }
};

console.log(o.a);
console.log(o.b);
o.c = 100;
console.log(o.a);

