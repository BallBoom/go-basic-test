# GO MAP

## 一、map

### 基础定义

    1、类似于其他表的哈希表或字典，以key-value形式存储；
    2、Key必须是支持== 或!= 比较运算的类型，不可以是函数、map或者slice
    3、Map可以使用make()创建，也支持 := 这种简写；
        初始化后再赋值：  m := make(map[int]string,5)   m[1] = "a" m[2] = "b"
                         var m map[int]string = make(map[int]string, 5) m[1] = "a" m[2] = "b"
        初始化直接赋值：  m := map[int]string{1:"a",2:"b",3:"c"}
    4、make([keyType]valueType,cap),cap表示容量，可忽略；5、超出容量是自动扩容，最好是提供合适的初始值；
    6、使用len()获取元素个数；
    7、map是无序的，每次迭代的序号都是不一样的，可以使用slice进行间接排序
