
### <strong style="color:#f30">c语言生成so文件，其他语言调用(generating-so-in-c)</strong>

### c语言生成so文件

	make all

清理:

	make clean


### python调用so文件

	[root@localhost python]# python test.py 
	result1:165,result2:78


### golang调用so文件


	[root@localhost golang]# go run test.go 
	result1:165,result2:78



----


### <strong style="color:#f30">golang语言生成so文件，其他语言调用(generating-so-in-golang)</strong>



### golang生成so库


	go build -v -x -gcc test.c -o test ./legendtkl.so -o libtest.so


生成:

	[root@localhost golang]# ls
	libtest.h  libtest.so  main.go

生成libtest.so和libtest.h 2个文件

### c调用so文件

C调用GO的动态链接库


	#gcc test.c -o t1 -I./ -L./ -ltest
	gcc test.c -o test ./libtest.so

输出:

	[root@localhost c]# gcc test.c -o test ./libtest.so
	[root@localhost c]# ./test
	143
	11

### python调用so文件

	[root@localhost python]# python test.py 
	result1:143,result2:11

	