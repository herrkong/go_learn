
#### defer recover panic

panic
函数中遇到panic语句，会立即终止当前函数的执行，在panic所在函数内如果存在要执行的defer函数列表，按照defer的逆序执行

recover
recover函数的返回值报告协程是否正在遭遇panic
有异常时，recover()只能调用一次，后面再次调用则捕获不到任何异常


go中可以抛出一个panic的异常，然后在defer中通过recover捕获这个异常，然后正常处理，从而恢复正常代码的执行




/这里使用defer + recover来捕获处理异常
defer func() {  //defer就是把匿名函数压入到defer栈中，等到执行完毕后或者发生异常后调用匿名函数(必须要先声明defer，否则不能捕获到panic异常)


panic 和 recover 的组合有如下特性：
有 panic 没 recover，程序宕机。
有 panic 也有 recover，程序不会宕机，执行完对应的 defer 后，从宕机点退出当前函数后继续执行。