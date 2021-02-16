package sysinit

/**
初始化全局数据链接
 */
func init() {
	sysinit()
	dbinit() // 初始化主库
	//dbinit("w","r") // 初始化主从库
}