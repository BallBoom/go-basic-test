# GO 结构

## 一、基础定义

### 结构struct

1. 使用
***type Name struct{}*** 来定义结构，名称遵循可见性原则；

    ``` GO
    type Person struct{
        Name string
        Age int
    }
    ```

2. 支持指向自身的指针类型成员；

3. 支持匿名结构，可用作成员或定义成员变量；

    ``` GO
     Humen := struct{
        Sex int
    }{
        Sex : 1, //给结构赋值的时候，最后一个字段后面也需要加上“,”否则报错
    }
    ```

4. 匿名结构也可用于map的值

5. 可以使用字面值进行初始化；

    ``` GO
    a := Person{}
    a.Name = "jack"
    a.Age = 20
    ```

6. 允许通过指针对结构成员进行读写；

    ``` GO
    b := &Person{} //取地址，b 即为指向Person的指针，无需加*号
    b.Name = "jack"
    b.Age = 20
    ```

7. 相同类型的成员可以直接进行拷贝赋值；

    ``` GO
    c := Person{}
    c == a
    ```

8. 支持==或!= 运算符，不支持< 或 >；

    ``` GO
    fmt.Println(c==a) //true
    ```

9. 支持匿名字段，本质上是定义了以一个类型名为字段的字段；

    ``` GO
    type Person struct{
         string
         int
    }
    给匿名字段初始化、赋值时，必须按照类型顺序赋值，否则会报错
    p := Person{"jack",18}

    ```
