# GO切片

## 一、基础定义

### 切片slice

1. 本身不是数组，是指向底层的数组；切片是引用类型

2. 作为变长数组的替代方案，可以关联底层数组的局部或全部；

3. 可以直接创建或从底层数组获取生成；使用len()获取长度，使用cap()获取容量；
4. 一般使用make创建：
   1. make([]T,len,cap)
   2. cap可以省略，则和len值一样
   3. len表示元素个数，cap表示容量
   4. s3 := make([]int, 3, 5)
   5. 当所存放的值超过了创建的元素个数，cap容量则会自动翻倍增加，已容纳更多的值，
        但是是创建新的slice指向之前的那个slice
   6. 如果多个slice指向同一个底层数组，其中一个的值改变会影响全部；但是如果其中一个先append()，则不会影响到另一个slice

### 函数

1. Reslice

2. append
        slice 使用append()时，如果超过了一开始创建的容量，则会重新分配一个内存地址，然后将之前的数据拷贝过去，随后再新增添加的数据
        s1 := []int{1,2,3}
        s1 = append(s1,4,5,6,7)

        s3 := make([]int, 3, 5)  &s3 = 0xc000084030
        s3 = append(s3, 5, 6, 7, 8)  &s3 = 0xc00008c0f0
        扩展后的s3的地址已经不一样；

3. copy
        s1 := []int{1,2,3,4,5,6}
        s2 := []int{7,8,9}
        copy(s1,s2)   将s2拷贝到s1
        s1的值变为{7,8,9,4,5,6}，因为s2只有三个值，那么拷贝过去的值也只是前三个
        copy(s2,s1)   将s1拷贝到s2
        s2的值变为{1,2,3}        因为s2只有三个值，s1只会把前三个值拷贝到s2
        所以，拷贝的时候 是看数值较少的那个slice

4. 使s2与s1的数据完全一致的写法

    ``` GO
        s1 := [1,2,3,4,5,6]
        s2 := s1[0:6] //s1下标从0到5的数据拷贝到s2  or s2 := s1[:]  最简洁的写法
    ```
