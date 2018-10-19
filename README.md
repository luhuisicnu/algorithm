# 通用
## 参数
* -s：展示排序步骤
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