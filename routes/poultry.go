package routes

import (
	"backend/controllers"
    //  "net/http"
	"github.com/gin-gonic/gin"
)
var(
	app *gin.Engine
)

func PoultryRoute(router *gin.Engine, controller controllers.PoultryController) {
	router.POST("api/profile/create", controller.CreateEmployee)
	router.POST("api/profile/create/hatching", controller.CreateEmployeeHatching)
	router.POST("api/profile/get", controller.GetEmployee) //
	router.POST("api/profile/get/hatching", controller.GetHatching)
	router.POST("api/profile/create/edit", controller.EditEmployee)
	router.POST("api/profile/hatch/edit", controller.EditHatch)
	router.POST("api/profile/delete", controller.DeleteEmployee)
	router.POST("api/profile/delete/hatch", controller.DeleteHatch)

	router.POST("api/order/post", controller.CreateOrder)
	router.POST("api/order/get", controller.GetOrder)
	router.POST("api/order/update", controller.UpdateOrder)
	router.POST("api/order/delete", controller.DeleteOrder)

	router.POST("api/customer/create", controller.CreateCustomer)
	router.POST("api/customer/find", controller.FindCustomer)
	router.PUT("api/customer/update", controller.UpdateCustomer)
	router.GET("api/customer/latest", controller.FindLastController)

	router.POST("api/admin", controller.GetAdmin)

	router.GET("admin/userinfo", controller.AdminInfo)

	router.POST("admin/value/update", controller.AdminUpdate)

	router.POST("admin/value/get", controller.AdminData)

	router.DELETE("api/delete/user", controller.DeleteUser)

	router.POST("api/month",controller.GraphMonth)


	router.POST("api/getone/customer",controller.GetCustomerOne)
}
// func PreflightRoute(c *gin.Context) {
// 	c.Header("Access-Control-Allow-Origin", "https://poultry-front.vercel.app")
// 	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
// 	c.Header("Access-Control-Allow-Headers", "Origin, Content-Type")
  
// 	c.Status(200)
//   }

// func init() {
//     app = gin.New()
//     controller := controllers.PoultryController{} // Instantiate your controller here if you haven't already.
//     PoultryRoute(app, controller) // Pass the Gin router and the controller to PoultryRoute.
// }
// func Handler(w http.ResponseWriter,r *http.Request){
// 	app.ServeHTTP(w,r);
// }
