# snap7-go
snap7 golang windows binding

# 步骤
1、生成snap7静态库
修改snap7的Mingw64的Makefile：
all-after:
	ar rcs libsnap7.a $(LINKOBJ)

2、修改snap7.go
增加库stdc++、winmm、ws2_32的引用
并将库拷贝到当前目录下


