package models



type Poultry struct{

	ID int  `json:"id" bson:"id"`
	FirstName string `json:"firstname" bson:"firstname"`
	LastName string `json:"lastname" bson:"lastname"`
	Gender string `json:"gender" bson:"gender"`
	Age int `json:"age" bson:"age"`
	Department string `json:"department" bson:"department"`
	Phone int `json:"phone" bson:"phone"`
	Salary int `json:"salary" bson:"salary"`
	MailFind string `json:"mailfind" bson:"mailfind"`
}
type Hatching struct{
	Batch_ID int `json:"batch_id" bson:"batch_id"`
	Date string `json:"date" bson:"date"`
	Egg int `json:"egg" bson:"egg"`
	Meet int `json:"meet" bson:"meet"`
	Chick int `json:"chick" bson:"chick"`
	MailFind string `json:"mailfind" bson:"mailfind"`
}

type Order struct{
	Name string `json:"name" bson:"name"`
	Date string `json:"date" bson:"date"`
	Phone int `json:"phone" bson:"phone"`
	Egg int `json:"egg" bson:"egg"`
	Address string `json:"address" bson:"address"`
	MailFind string `json:"mailfind" bson:"mailfind"`
}
type Customer struct{

    Signup string `json:"signup" bson:"signup"`
	Password string `json:"password" bson:"password"`
}

type Login struct{
    Signup string `json:"signup" bson:"signup"`
	Password string `json:"password" bson:"password"`
}



type User struct{
	MailID string `json:"mailid" bson:"mailid"`
	Password string `json:"password" bson:"password"`
    Image []byte `json:"image" bson:"image`
	Poultry []Poultry 
	Poultry_Count int `json:"poultry_count" bson:"poultry_count"`
	Hatching []Hatching 
	Hatching_Count int `json:"hatching_count" bson:"hatching_count"`
	Order []Order 
	Order_Count int `json:"order_count" bson:"order_count"`
}
type Admin struct {
	MailID string `json:"mailid" bson:"mailid"`
	Password string `json:"password" bson:"password"`
	Poultry_Count int `json:"poultry_count" bson:"poultry_count"`
	Hatching_Count int `json:"hatching_count" bson:"hatching_count"`
	Order_Count int `json:"order_count" bson:"order_count"`
	Doctor_Count int `json:"doctor_count" bson:"doctor_count"`
}
type Admin1 struct {
	MailID string `json:"mailid" bson:"mailid"`
	Password string `json:"password" bson:"password"`
	Employee_Count int `json:"employee_count" bson:"employee_count"`
	Production_Count int `json:"production_count" bson:"production_count"`
	Order_Count int `json:"order_count" bson:"order_count"`
	Doctor_Count int `json:"doctor_count" bson:"doctor_count"`
}
type Graph_Egg struct {
    Jan int `json:"jan" bson:"jan"`
    Feb int `json:"feb" bson:"feb"`
    Mar int `json:"mar" bson:"mar"`
    Apr int `json:"apr" bson:"apr"`
    May int `json:"may" bson:"may"`
    Jun int `json:"jun" bson:"jun"`
    Jul int `json:"jul" bson:"jul"`
    Aug int `json:"aug" bson:"aug"`
    Sep int `json:"sep" bson:"sep"`
    Oct int `json:"oct" bson:"oct"`
    Nov int `json:"nov" bson:"nov"`
    Dec int `json:"dec" bson:"dec"`
}

type Graph_Meet struct {
    Jan int `json:"jan" bson:"jan"`
    Feb int `json:"feb" bson:"feb"`
    Mar int `json:"mar" bson:"mar"`
    Apr int `json:"apr" bson:"apr"`
    May int `json:"may" bson:"may"`
    Jun int `json:"jun" bson:"jun"`
    Jul int `json:"jul" bson:"jul"`
    Aug int `json:"aug" bson:"aug"`
    Sep int `json:"sep" bson:"sep"`
    Oct int `json:"oct" bson:"oct"`
    Nov int `json:"nov" bson:"nov"`
    Dec int `json:"dec" bson:"dec"`
}

type Graph_Chick struct {
    Jan int `json:"jan" bson:"jan"`
    Feb int `json:"feb" bson:"feb"`
    Mar int `json:"mar" bson:"mar"`
    Apr int `json:"apr" bson:"apr"`
    May int `json:"may" bson:"may"`
    Jun int `json:"jun" bson:"jun"`
    Jul int `json:"jul" bson:"jul"`
    Aug int `json:"aug" bson:"aug"`
    Sep int `json:"sep" bson:"sep"`
    Oct int `json:"oct" bson:"oct"`
    Nov int `json:"nov" bson:"nov"`
    Dec int `json:"dec" bson:"dec"`
}
type Graph_Order struct{
	Jan int `json:"jan" bson:"jan"`
    Feb int `json:"feb" bson:"feb"`
    Mar int `json:"mar" bson:"mar"`
    Apr int `json:"apr" bson:"apr"`
    May int `json:"may" bson:"may"`
    Jun int `json:"jun" bson:"jun"`
    Jul int `json:"jul" bson:"jul"`
    Aug int `json:"aug" bson:"aug"`
    Sep int `json:"sep" bson:"sep"`
    Oct int `json:"oct" bson:"oct"`
    Nov int `json:"nov" bson:"nov"`
    Dec int `json:"dec" bson:"dec"`
}

type Graph struct {
	Mailid string `json:"mailid" bson:"mailid"`
    Egg Graph_Egg `json:"egg" bson:"egg"`
	Meet Graph_Meet `json:"meet" bson:"meet"`
	Chick Graph_Chick `json:"chick" bson:"chick"`
	Order Graph_Order `json:"order" bson:"order"`
}