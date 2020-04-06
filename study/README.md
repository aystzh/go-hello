#### golang语法学习笔记
##### go func 函数详解
```gotemplate
func (p myType) funcName(a,b,int ,c string)(r,s int){
return
}
```
##### 相关解释：
###### func 关键字，定义函数的关键字
###### (p myType) 函数拥有者，go可以为特定类型定义函数，即为类型对象定义方法
###### funcName 函数名
###### a,b,int ,c string) 入参
###### (r,s int)返回值
###### {}函数体