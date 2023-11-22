package controllers

import (
	"backend/interfaces"
	"backend/models"
	"backend/validation"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type PoultryController struct {
	PoultryService interfaces.IPoultry
}

func InitPoultryController(PoultryService interfaces.IPoultry) PoultryController {
	return PoultryController{PoultryService} //DI(dependency injection) pattern
}

func (pc *PoultryController) CreateEmployee(ctx *gin.Context) {

	var customer *models.Poultry
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	val, err := pc.PoultryService.CreateService(customer)
	if val=="greater"{
		ctx.JSON(http.StatusOK, gin.H{"status":"ok","message": "greater"})
		return
	}else if err != nil {

		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	

}

func (pc *PoultryController) GetEmployee(ctx *gin.Context) {
	type value struct {
		Mailid string `json:"mailid" bson:"mailid"`
	}
	var val1 value
	if err := ctx.ShouldBindJSON(&val1); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	filter := bson.M{"mailid": val1.Mailid}
	results, err := pc.PoultryService.GetEmployeeService(filter)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail"})
		return
	}
	fmt.Println(results)
	ctx.JSON(http.StatusOK, gin.H{"array": results})

}

func (pc *PoultryController) CreateEmployeeHatching(ctx *gin.Context) {

	var customer *models.Hatching
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	dateParsed, err := time.Parse("2006-01-02", customer.Date)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid date format"})
		return
	}
	month := dateParsed.Month()

	// Define the field name to increment based on the month
	fieldToIncrement := ""
	switch month {
	case time.January:
		fieldToIncrement = "jan"
	case time.February:
		fieldToIncrement = "feb"
	case time.March:
		fieldToIncrement = "mar"
	case time.April:
		fieldToIncrement = "apr"
	case time.May:
		fieldToIncrement = "may"
	case time.June:
		fieldToIncrement = "jun"
	case time.July:
		fieldToIncrement = "jul"
	case time.August:
		fieldToIncrement = "aug"
	case time.September:
		fieldToIncrement = "sep"
	case time.October:
		fieldToIncrement = "oct"
	case time.November:
		fieldToIncrement = "nov"
	case time.December:
		fieldToIncrement = "dec"
	default:

		fmt.Println("Invalid month")

	}
	fmt.Println(fieldToIncrement)

	// Set the parsed date into the customer struct

	customer.Date = dateParsed.Format("2006-01-02") // Convert it back to string if needed

	val, err := pc.PoultryService.CreateServiceHatching(customer, fieldToIncrement)

	if val=="greater"{
		ctx.JSON(http.StatusOK, gin.H{"status":"ok","message": "greater"})
		return
	}else if err != nil {

		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

}

func (pc *PoultryController) GetHatching(ctx *gin.Context) {
	type value struct {
		Mailid string `json:"mailid" bson:"mailid"`
	}
	var val1 value
	if err := ctx.ShouldBindJSON(&val1); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	filter := bson.M{"mailid": val1.Mailid}
	results, err := pc.PoultryService.GetHatchingService(filter)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail"})
		return
	}
	fmt.Println(results)
	ctx.JSON(http.StatusOK, gin.H{"array": results})
}

func (pc *PoultryController) EditEmployee(ctx *gin.Context) {
	var customer *models.Poultry
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	filter := bson.M{
		"mailid": customer.MailFind,
		"poultry": bson.M{
			"$elemMatch": bson.M{
				"id": customer.ID,
			},
		},
	}

	update := bson.M{
		"$set": bson.M{
			"poultry.$.firstname":  customer.FirstName,
			"poultry.$.lastname":   customer.LastName,
			"poultry.$.gender":     customer.Gender,
			"poultry.$.age":        customer.Age,
			"poultry.$.department": customer.Department,
			"poultry.$.phone":      customer.Phone,
			"poultry.$.salary":     customer.Salary,
		},
	}

	_, err := pc.PoultryService.GetEditEmployeeService(filter, update)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, "Employee updated successfully")
}

func (pc *PoultryController) DeleteEmployee(ctx *gin.Context) {
	type EmpDel struct {
		ID     int    `json:"id" bson:"id"`
		Mailid string `json:"mailid" bson:"mailid"`
	}
	var customer = new(EmpDel)
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// Define the filter to find the user based on mailid and the specific poultry item based on ID
	filter := bson.M{
		"mailid":     customer.Mailid,
		"poultry.id": customer.ID,
	}
	filter1:=bson.M{
		"mailid":customer.Mailid,
	}

	// Define the update operation to pull the specific poultry item from the poultry array
	update := bson.M{
		"$pull": bson.M{
			"poultry": bson.M{"id": customer.ID},
		},
	}

	_, err := pc.PoultryService.GetDeleteEmployeeService(filter, update,filter1)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	

	ctx.JSON(http.StatusOK, "Employee deleted successfully")
}

func (pc *PoultryController) DeleteHatch(ctx *gin.Context) {
  var customer models.Hatching
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println("Deleting Hatching",customer)
	dateParsed, err := time.Parse("2006-01-02", customer.Date)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid date format"})
		return
	}
	month := dateParsed.Month()

	// Define the field name to increment based on the month
	fieldToIncrement := ""
	switch month {
	case time.January:
		fieldToIncrement = "jan"
	case time.February:
		fieldToIncrement = "feb"
	case time.March:
		fieldToIncrement = "mar"
	case time.April:
		fieldToIncrement = "apr"
	case time.May:
		fieldToIncrement = "may"
	case time.June:
		fieldToIncrement = "jun"
	case time.July:
		fieldToIncrement = "jul"
	case time.August:
		fieldToIncrement = "aug"
	case time.September:
		fieldToIncrement = "sep"
	case time.October:
		fieldToIncrement = "oct"
	case time.November:
		fieldToIncrement = "nov"
	case time.December:
		fieldToIncrement = "dec"
	default:

		fmt.Println("Invalid month")

	}
	fmt.Println(fieldToIncrement)

	// Set the parsed date into the customer struct

	customer.Date = dateParsed.Format("2006-01-02")

	// Define the filter to find the user based on mailid and the specific poultry item based on ID
	filter := bson.M{
		"mailid":            customer.MailFind,
		"hatching.batch_id": customer.Batch_ID,

	}
	filter2:=bson.M{
		"mailid": customer.MailFind,

	}

	// Define the update operation to pull the specific poultry item from the poultry array
	update := bson.M{
		"$pull": bson.M{
			"hatching": bson.M{"batch_id": customer.Batch_ID},
		},
	}

	_, err1 := pc.PoultryService.GetDeleteHatchService(&customer,filter, update,filter2,fieldToIncrement)
	if err1 != nil {
		ctx.JSON(http.StatusInternalServerError, err1.Error())
		return
	}

	ctx.JSON(http.StatusOK, "Employee deleted successfully")
}
func (pc *PoultryController) DeleteOrder(ctx *gin.Context) {
	type EmpDel struct {
		Phone  int    `json:"phone" bson:"phone"`
		Mailid string `json:"mailid" bson:"mailid"`
	}
	var customer = new(EmpDel)
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// Define the filter to find the user based on mailid and the specific poultry item based on ID
	filter := bson.M{
		"mailid":      customer.Mailid,
		"order.phone": customer.Phone,
	}
	filter2:=bson.M{
		"mailid":      customer.Mailid,

	}

	// Define the update operation to pull the specific poultry item from the poultry array
	update := bson.M{
		"$pull": bson.M{
			"order": bson.M{"phone": customer.Phone},
		},
	}

	_, err := pc.PoultryService.GetDeleteOrderService(filter, update,filter2)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, "Order deleted successfully")
}
func (pc *PoultryController) EditHatch(ctx *gin.Context) {
	var customer *models.Hatching
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	filter := bson.M{
		"mailid": customer.MailFind,
		"hatching": bson.M{
			"$elemMatch": bson.M{
				"batch_id": customer.Batch_ID,
			},
		},
	}

	update := bson.M{
		"$set": bson.M{

			"hatching.$.date":  customer.Date,
			"hatching.$.egg":   customer.Egg,
			"hatching.$.meet":  customer.Meet,
			"hatching.$.chick": customer.Chick,
		},
	}
	dateParsed, err := time.Parse("2006-01-02", customer.Date)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid date format"})
		return
	}

	// Set the parsed date into the customer struct

	customer.Date = dateParsed.Format("2006-01-02") // Convert it back to string if needed

	_, err1 := pc.PoultryService.GetEditHatchService(filter, update)
	if err1 != nil {
		ctx.JSON(http.StatusInternalServerError, err1.Error())
		return
	}

	ctx.JSON(http.StatusOK, "hATCHING updated successfully")
}

func (pc *PoultryController) CreateOrder(ctx *gin.Context) {

	var customer *models.Order
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	dateParsed, err := time.Parse("2006-01-02", customer.Date)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid date format"})
		return
	}
	month := dateParsed.Month()

	// Define the field name to increment based on the month
	fieldToIncrement := ""
	switch month {
	case time.January:
		fieldToIncrement = "jan"
	case time.February:
		fieldToIncrement = "feb"
	case time.March:
		fieldToIncrement = "mar"
	case time.April:
		fieldToIncrement = "apr"
	case time.May:
		fieldToIncrement = "may"
	case time.June:
		fieldToIncrement = "jun"
	case time.July:
		fieldToIncrement = "jul"
	case time.August:
		fieldToIncrement = "aug"
	case time.September:
		fieldToIncrement = "sep"
	case time.October:
		fieldToIncrement = "oct"
	case time.November:
		fieldToIncrement = "nov"
	case time.December:
		fieldToIncrement = "dec"
	default:

		fmt.Println("Invalid month")

	}
	fmt.Println(fieldToIncrement)

	// Set the parsed date into the customer struct

	customer.Date = dateParsed.Format("2006-01-02") // Convert it back to string if needed

	val, err := pc.PoultryService.CreateServiceOrder(customer,fieldToIncrement)

	if val=="greater"{
		ctx.JSON(http.StatusOK, gin.H{"status":"ok","message": "greater"})
		return
	}else if err != nil {

		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

}

func (pc *PoultryController) GetOrder(ctx *gin.Context) {
	type value struct {
		Mailid string `json:"mailid" bson:"mailid"`
	}
	var val1 value
	if err := ctx.ShouldBindJSON(&val1); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	filter := bson.M{"mailid": val1.Mailid}
	results, err := pc.PoultryService.GetOrderService(filter)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail"})
		return
	}
	fmt.Println(results)
	ctx.JSON(http.StatusOK, gin.H{"array": results})

}

func (pc *PoultryController) UpdateOrder(ctx *gin.Context) {
	var customer *models.Order
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	filter := bson.M{
		"mailid": customer.MailFind,
		"order": bson.M{
			"$elemMatch": bson.M{
				"phone": customer.Phone,
			},
		},
	}

	update := bson.M{
		"$set": bson.M{

			"order.$.name":    customer.Name,
			"order.$.egg":     customer.Egg,
			"order.$.address": customer.Address,
		},
	}

	_, err := pc.PoultryService.GetEditOrderService(filter, update)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, "hATCHING updated successfully")
}

//here

func (pc *PoultryController) CreateCustomer(ctx *gin.Context) {
	e := 0
	p := 0
	var customer *models.User
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if validation.Email(customer.MailID) == 1 {
		e = 1
		fmt.Println("Email is valid", customer.MailID)
		// ctx.JSON(http.StatusOK, gin.H{"message": "Email is Valid"})
		if validation.Password(customer.Password) == 1 {
			fmt.Println("Password is valid", customer.Password)
			// ctx.JSON(http.StatusOK, gin.H{"message": "Password is Valid"})
			p = 1

		}

	} else {
		fmt.Println("Email is not valid", customer.MailID)
		// ctx.JSON(http.StatusOK, gin.H{"message": "Email is Not Valid"})
		if validation.Password(customer.Password) != 1 {
			fmt.Println("Password is not valid", customer.Password)
			// ctx.JSON(http.StatusOK, gin.H{"message": "Password is Not Valid"})

		}
	}
	if p == 1 && e == 1 {
		ctx.JSON(http.StatusOK, gin.H{"message": "Email is Valid", "password": "Password is Valid", "number": ""})
	} else if p != 1 && e == 1 {
		ctx.JSON(http.StatusOK, gin.H{"message": "Email is Valid", "password": "Password is Not Valid", "number": ""})
		return
	} else if p == 1 && e != 1 {
		ctx.JSON(http.StatusOK, gin.H{"message": "Email is Not Valid", "password": "Password is Valid", "number": ""})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "Email is Not Valid", "password": "Password is Not Valid", "number": ""})
		return
	}

	val, _ := pc.PoultryService.CreateCustomer(customer)

	if val == 1 {
		ctx.JSON(http.StatusConflict, gin.H{"exist": "Mail ID already exist"})
	} else {
		ctx.JSON(http.StatusConflict, gin.H{"message": "succesfully inserted"})
	}

}

func (pc *PoultryController) FindCustomer(ctx *gin.Context) {

	var requestData models.User
	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filter := bson.M{
		"$and": []bson.M{
			{"mailid": requestData.MailID},
			{"password": requestData.Password},
		},
	}

	result, err1 := pc.PoultryService.FindService(filter)
	if err1 != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail"})
		return
	}
	if result != 0 {
		fmt.Println("correct")
		ctx.JSON(http.StatusOK, gin.H{"message": requestData.MailID, "password": requestData.Password, "number": "equal"})
	} else {
		fmt.Println("wrong")
		ctx.JSON(http.StatusOK, gin.H{"message": "Does Not Exist", "password": "Invalid", "number": "not"})
	}

}
func (pc *PoultryController) GetAdmin(ctx *gin.Context) {

	var requestData models.Customer
	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filter := bson.M{
		"$and": []bson.M{
			{"mailid": requestData.Signup},
			{"password": requestData.Password},
		},
	}

	result, err1 := pc.PoultryService.GetAdminService(filter)
	if err1 != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail"})
		return
	}
	if result != 0 {
		fmt.Println("correct")
		ctx.JSON(http.StatusOK, gin.H{"message": requestData.Signup, "password": requestData.Password, "number": "equal"})
	} else {
		fmt.Println("wrong")
		ctx.JSON(http.StatusOK, gin.H{"message": "Does Not Exist", "password": "Invalid", "number": "not"})
	}

}

func (pc *PoultryController) FindLastController(ctx *gin.Context) {
	result, err4 := pc.PoultryService.FindLastService()
	fmt.Println("Mail ID " + result.Signup)
	fmt.Println("Password " + result.Password)
	if err4 == nil {
		ctx.JSON(http.StatusOK, gin.H{"mail": result.Signup, "password": result.Password})

		return
	}
}

func (pc *PoultryController) UpdateCustomer(ctx *gin.Context) {

	var requestData models.Customer
	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filter := bson.M{
		"signup": requestData.Signup,
	}
	update := bson.M{"$set": bson.M{
		"password": requestData.Password,
	},
	}

	result, err4 := pc.PoultryService.UpdateService(filter, update)
	if err4 != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"statu": "failed"})
	}
	if result != 1 {
		fmt.Println("password updated")
	} else {
		fmt.Println("password not updated")
	}

}
func (pc *PoultryController) AdminInfo(ctx *gin.Context) {

	filter := bson.M{}
	results, err := pc.PoultryService.GetAdminInfoService(filter)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail"})
		return
	}
	fmt.Println(results)
	ctx.JSON(http.StatusOK, gin.H{"array": results})
}

func (pc *PoultryController) AdminUpdate(ctx *gin.Context) {
	type Admin struct {
		MailID           string `json:"mailid" bson:"mailid"`
		Password         string `json:"password" bson:"password"`
		Employee_Count   int    `json:"employee_count" bson:"employee_count"`
		Production_Count int    `json:"production_count" bson:"production_count"`
		Order_Count      int    `json:"order_count" bson:"order_count"`
		Doctor_Count     int    `json:"doctor_count" bson:"doctor_count"`
	}
	var adminmod Admin
	if err := ctx.ShouldBindJSON(&adminmod); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	filter := bson.M{
		"$and": []bson.M{
			{"mailid": adminmod.MailID},
			{"password": adminmod.Password},
		},
	}
	update := bson.M{
		"$set": bson.M{
			"employee_count":   adminmod.Employee_Count,
			"production_count": adminmod.Production_Count,
			"order_count":      adminmod.Order_Count,
			"doctor_count":     adminmod.Doctor_Count,
		},
	}
	err := pc.PoultryService.AdminUpdateService(filter, update)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, "Admin updated successfully")
}

func (pc *PoultryController) AdminData(ctx *gin.Context) {
	type model struct {
		MailID   string `json:"mailid" bson:"mailid"`
		Password string `json:"password" bson:"password"`
	}
	var requestData model
	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filter := bson.M{
		"$and": []bson.M{
			{"mailid": requestData.MailID},
			{"password": requestData.Password},
		},
	}

	result, err1 := pc.PoultryService.AdminDataService(filter)
	if err1 != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"array": result})
}

func (pc *PoultryController) DeleteUser(ctx *gin.Context) {
	type model struct {
		MailID   string `json:"mailid" bson:"mailid"`
		Password string `json:"password" bson:"password"`
	}
	var requestData model
	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filter := bson.M{
		"$and": []bson.M{
			{"mailid": requestData.MailID},
			{"password": requestData.Password},
		},
	}

	err1 := pc.PoultryService.DeleteUserService(filter)
	if err1 != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail"})
		return
	}
}

func (pc *PoultryController) GraphMonth(ctx *gin.Context) {

	var requestData models.User
	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filter := bson.M{
		"mailid": requestData.MailID,
			
		
	}

	result, err1 := pc.PoultryService.MonthService(filter)
	
	if err1 != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"array": result})

}


func (pc *PoultryController) GetCustomerOne(ctx *gin.Context) {

	var requestData models.User
	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filter := bson.M{"mailid": requestData.MailID}
	result, err1 := pc.PoultryService.GetCustomerOneService(filter)
	if err1 != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"array": result})

}

