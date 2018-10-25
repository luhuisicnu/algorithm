# 通用
## 细节展示
* -s：展示执行步骤细节

## 排序参数
* -d：倒序排列


# 插入排序
## 运行办法
```
$ go run insertion_sort.go -s 3 5 1 4 2
start sort: [3 5 1 4 2]

now insert index 1 value 5
keep status quo
status: [3 5 1 4 2]

now insert index 2 value 1
index: 1 value: 5 back 1 position
index: 0 value: 3 back 1 position
status: [1 3 5 4 2]

now insert index 3 value 4
index: 2 value: 5 back 1 position
status: [1 3 4 5 2]

now insert index 4 value 2
index: 3 value: 5 back 1 position
index: 2 value: 4 back 1 position
index: 1 value: 3 back 1 position
status: [1 2 3 4 5]

[1 2 3 4 5]
```

# 归并排序
## 运行办法
```
$ go run merge_sort.go -d -s 8 6 0 9 5
start sort: [8 6 0 9 5]

split [8 6 0 9 5] after index 1 value 6 left [8 6] right [0 9 5]
split [8 6] after index 0 value 8 left [8] right [6]
merge left [8] and right [6]
select left 8
select right 6
status [8 6]
split [0 9 5] after index 0 value 0 left [0] right [9 5]
split [9 5] after index 0 value 9 left [9] right [5]
merge left [9] and right [5]
select left 9
select right 5
status [9 5]
merge left [0] and right [9 5]
select right 9
select right 5
select left 0
status [9 5 0]
merge left [8 6] and right [9 5 0]
select right 9
select left 8
select left 6
select right 5
select right 0
status [9 8 6 5 0]

[9 8 6 5 0]
```

# 堆排序
## 运行办法
```
$ go run heap_sort.go -s 3 5 1
build: start build a max heap [3 5 1]
max heapifing: index 0 left index 1 right index 2
max heapifing: left value 5 is more than this value 3
max heapifing: source [3 5 1]
max heapifing: exchange value 5 and 3
max heapifing: index 1 left index -1 right index -1
max heapifing: nothing to do
max heapifing: exchange result [5 3 1]
build: build a new max heap [5 3 1]
sort: exchange first value 5 and this value 1
sort: source [1 3 5]
max heapifing: index 0 left index 1 right index -1
max heapifing: left value 3 is more than this value 1
max heapifing: source [1 3 5]
max heapifing: exchange value 3 and 1
max heapifing: index 1 left index -1 right index -1
max heapifing: nothing to do
max heapifing: exchange result [3 1 5]
sort: exchange first value 3 and this value 1
sort: source [1 3 5]
max heapifing: index 0 left index -1 right index -1
max heapifing: nothing to do
[1 3 5]
```

# 快速排序
## 运行办法
```
$ go run quick_sort.go -s -d 3 5 1
sort: source [3 5 1] start index 0 end index 3
partition: source [3 5 1]
partition: find last value as key 1
partition: start part
partition: index 0 value 3 is more than key value
partition: exchange  3 3
partition: middle result [3 5 1]
partition: index 1 value 5 is more than key value
partition: exchange  5 5
partition: middle result [3 5 1]
partition: exchange key value 1 and first bigger value 1
partition: result [3 5 1]
sort: find a middle index 2
sort: source [3 5] start index 0 end index 2
partition: source [3 5]
partition: find last value as key 5
partition: start part
partition: nothing to do with index 0 value 3
partition: exchange key value 5 and first bigger value 3
partition: result [5 3]
sort: find a middle index 0
sort: source [] start index 0 end index 0
sort: nothing to do
sort: source [3] start index 1 end index 2
sort: nothing to do
sort: source [] start index 3 end index 3
sort: nothing to do
[5 3 1]
```

# 计数排序
## 运行办法
```
$ go run counting_sort.go -s 3 2 1
sort: temp counting array init success [0 1 2 3]
sort: result index 1 get value 1
sort: temp counting array state [0 0 2 3]
sort: result state [1 0 0]
sort: result index 2 get value 2
sort: temp counting array state [0 0 1 3]
sort: result state [1 2 0]
sort: result index 3 get value 3
sort: temp counting array state [0 0 1 2]
sort: result state [1 2 3]
[1 2 3]
```

# 查找最大子数组
## 暴力求解
### 运行办法
```
$ go run find_max_subarray_n2.go -s 3 -8 2
start find: [3 -8 2]

get: start index 0 end index 0 current max sum 3
pass: start index 0 end index 1 sum -5 is less than max sum 3
pass: start index 0 end index 2 sum -3 is less than max sum 3
pass: this index 1 element -8  is less than max sum 3
pass: start index 1 end index 2 sum -6 is less than max sum 3
pass: this index 2 element 2  is less than max sum 3

start index: 0 end index: 0 sum: 3
```

## 左右划分递归求解
### 运行办法
```
$ go run find_max_subarray_nlgn.go -s 3 -8 2
start find: [3 -8 2]

split array to left [3] and right [-8 2]
side get: index 0 sum 3
split array to left [-8] and right [2]
side get: index 1 sum -8
side get: index 2 sum 2
cross start: left index 1 right index 3 source [-8 2]
cross get: left extend, index 1 temp sum -8 left sum -9223372036854775808
cross get: find bigger sum, left index 1 mid index 2 sum -8
cross get: right extend, index 2 temp sum 2 right sum -9223372036854775808
cross get: find bigger sum, mid index 2 right index 2 sum 2
cross end: left index 1 right index 2 sum -6
left array [-8] left max sum -8 right array [2] right max sum 2 cross array [-8 2] cross max sum -6
get: low index 2 high index 2 sum 2
cross start: left index 0 right index 3 source [3 -8 2]
cross get: left extend, index 0 temp sum 3 left sum -9223372036854775808
cross get: find bigger sum, left index 0 mid index 1 sum 3
cross get: right extend, index 1 temp sum -8 right sum -9223372036854775808
cross get: find bigger sum, mid index 1 right index 1 sum -8
cross get: right extend, index 2 temp sum -6 right sum -8
cross get: find bigger sum, mid index 1 right index 2 sum -6
cross end: left index 0 right index 2 sum -3
left array [3] left max sum 3 right array [-8 2] right max sum 2 cross array [3 -8 2] cross max sum -3
get: low index 0 high index 0 sum 3

start index: 0 end index: 0 sum: 3
```