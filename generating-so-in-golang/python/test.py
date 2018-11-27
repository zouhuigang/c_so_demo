#!/usr/bin/python
from ctypes import cdll
libtest = cdll.LoadLibrary("./libtest.so")
result1=libtest.Add(77,66)
result2=libtest.Minus(77,66)

print("result1:%d,result2:%d" % (result1,result2))