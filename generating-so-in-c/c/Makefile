#定义变量
CC = gcc
LIB = libtest.so
EXE = test

#编译所有
all:test $(LIB)

#执行文件
$(EXE):main.c
	$(CC) main.c calculator.c  -o $(EXE)

#动态库文件
$(LIB): calculator.c
	$(CC)  calculator.c -fPIC -shared -o $(LIB)

#清理
.PYONY:
clean:
	rm -rf $(EXE) $(LIB)