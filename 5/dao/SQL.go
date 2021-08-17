package dao

//数据库模型-主要是用于配置存储
type Sql struct {
	SqlName string		//数据库名
	SqlUserName string	//数据库登录用账户名
	SqlUserPwd string	//数据库登录用账户密码
	SqlAddr string		//数据库地址
}