# jsonrpc-demo

这个程序演示了如何输出Go的函数， 并且在C语言中调用这个Go函数。

主要用来调研CGO的其中的一个应用场景：
- GO实现函数， C语言调用

演示程序实现了一个`call`函数， 它接收参数，然后调用steem的API(websocket/jsonrpc 2),然后将结果返回。

C程序调用这个Go实现的`call`函数。


