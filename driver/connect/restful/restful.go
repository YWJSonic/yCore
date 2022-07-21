package restful

import "github.com/gin-gonic/gin"

func New() *RestfulDriver {

	driverObj := &RestfulDriver{}
	driverObj.engine = gin.Default()
	return driverObj
}

func (self *RestfulDriver) Run(Addr string) error {
	return self.engine.Run(Addr)
}

func (self *RestfulDriver) Handle(httpMethod, relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return self.engine.Handle(httpMethod, relativePath, handlers...)
}
