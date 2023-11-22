package interfaces

import ("backend/models"
"go.mongodb.org/mongo-driver/bson"
"go.mongodb.org/mongo-driver/mongo"
)

type IPoultry interface{
	CreateService(customer*models.Poultry) (string, error)
	CreateServiceHatching(customer *models.Hatching,val string) (string, error) 
	GetEmployeeService(filter bson.M)([] models.Poultry,error)
	GetHatchingService(filter bson.M)([] models.Hatching,error)
	GetEditEmployeeService(filter bson.M,update bson.M)(string,error)
	GetDeleteEmployeeService(filter bson.M, update bson.M,filter1 bson.M) (string, error)
	GetEditHatchService(filter,update bson.M)(*mongo.UpdateResult,error)
	GetDeleteHatchService(customer *models.Hatching,filter bson.M, update bson.M,filter2 bson.M,val string) (string, error)
	CreateServiceOrder(customer *models.Order,val string) (string, error)
	GetOrderService(filter bson.M)([] models.Order,error)
	GetEditOrderService(filter bson.M,update bson.M)(string,error)
	GetDeleteOrderService(filter bson.M, update bson.M,filter2 bson.M) (string, error)
	CreateCustomer(user *models.User) (int, error)
	FindService(filter bson.M)(int,error)
	UpdateService(filter bson.M,update bson.M)(int,error)
	FindLastService()(*models.Login,error)

	GetAdminService(filter bson.M)(int,error)


	GetAdminInfoService(filter bson.M) ([]models.Admin, error) 

	AdminUpdateService(filter bson.M,update bson.M) (error) 

	AdminDataService(filter bson.M) (models.Admin1, error)

	DeleteUserService(filter bson.M) ( error) 

	MonthService(filter bson.M) (models.Graph, error) 
	GetCustomerOneService(filter bson.M) (models.User, error)

	
}