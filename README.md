# zero-helloworld

This project is a demo which use go-zero microservice framework to implement a user register platform.

## Component

The demo is made up by two api services (search and user) and one rpc service (user). 

User api service implemenets register and login interface.

Search api service implements get user password interface, which would invoke user rpc service to get password.

User rpc service implements query password function.

## Last

The project is the basic use of go-zero framework. More advanced usage such as rate limit will be added later.

## Reference

[go-zero official document](https://go-zero.dev/cn/api-coding.html)

[go-zero实战之blog系统](https://mp.weixin.qq.com/s/XCm5xTx6dLR3uJFKEFx6Mw)
