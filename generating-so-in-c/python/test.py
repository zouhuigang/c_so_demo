#!/usr/bin/python
from ctypes import cdll
libtest = cdll.LoadLibrary("./libtest.so")
result1=libtest.add(77,88)
result2=libtest.minus(100,22)

print("result1:%d,result2:%d" % (result1,result2))