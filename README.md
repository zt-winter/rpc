# 说明

2022.6.5
目前的实现的rpc还是功能非常有限的轮子，后续不断更新，目标
* 支持并发，添加并发控制，将map[string]interface{}改为sync.map
* 支持同步、异步操作，将执行代码分割打包

    客户端传输方法名，找到对应的方法，这里是依靠map[string]interface{}实现的，最开始想法是直接使用map[string]的结果，但interface{}使用还需要下断言，但这样就需要提前知道方法类型，无法有效解决。
		最后这里使用reflect.method.func的方法call，直接调用自身函数，详情见[type method](https://pkg.go.dev/reflect@go1.18.3#Method)和[func (Value) Call](https://pkg.go.dev/reflect@go1.18.3#Value)
		这里还可以考虑使用golang相关注入依赖的库，通过控制反转实现。[相关开源库](https://github.com/facebookarchive/inject)
