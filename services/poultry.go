package services

import (
	"backend/interfaces"
	"backend/models"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TransactionService struct {
	PoultryCollection  *mongo.Collection
	HatchingCollection *mongo.Collection
	order              *mongo.Collection
	CustomerCollection *mongo.Collection
	CheckCollection    *mongo.Collection
	UserCollection     *mongo.Collection
	AdminCollection    *mongo.Collection
	GraphCollection *mongo.Collection
	ctx                context.Context
}

func PoultryServiceInit(poultry *mongo.Collection, Hatching *mongo.Collection, order *mongo.Collection, customer *mongo.Collection, checkcollection *mongo.Collection, UserCollection *mongo.Collection, AdminCollection *mongo.Collection, GraphCollection *mongo.Collection,ctx context.Context) interfaces.IPoultry {
	return &TransactionService{poultry, Hatching, order, customer, checkcollection, UserCollection, AdminCollection, GraphCollection,ctx}
}

func (p *TransactionService) CreateService(customer *models.Poultry) (string, error) {
   var checker models.User
	filter := bson.M{"mailid": customer.MailFind}
	val := p.UserCollection.FindOne(context.Background(), filter).Decode(&checker)
	fmt.Println(val)
	if checker.Poultry_Count>10{
		    return "greater",nil
	}
	update := bson.M{
		"$push": bson.M{
			"poultry": customer,
		},
	}

	// Execute the update operation
	_, err := p.UserCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Println("Transaction not updated")
		return "Transaction not updated", err
	}

	update1 := bson.M{"$inc": bson.M{"poultry_count": 1}} // Replace "your_field_to_increment" with the actual field name
	_, err2 := p.UserCollection.UpdateOne(p.ctx, filter, update1)
	if err2 != nil {
		log.Fatal(err2)
	}

	return "hello", nil
}

func (p *TransactionService) GetEmployeeService(filter bson.M) ([]models.Poultry, error) {
	cursor, err := p.UserCollection.Find(p.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(p.ctx) // Don't forget to close the cursor when done

	var results []models.Poultry
	for cursor.Next(p.ctx) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			log.Fatal(err)
		}
		results = append(results, user.Poultry...)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (p *TransactionService) CreateServiceHatching(customer *models.Hatching,val string) (string, error) {
	var checker models.User

	
	update1 := bson.M{
        "$inc": bson.M{
            "egg." + val:   customer.Egg,
            "meet." + val: customer.Meet,
            "chick." + val: customer.Chick,
        },
    }

    filter1 := bson.M{"mailid": customer.MailFind}
    _, err1 := p.GraphCollection.UpdateOne(context.Background(), filter1, update1)
    if err1 != nil {
        fmt.Println("Transaction not inserted")
        return "Transaction not inserted", err1
    }
	filter := bson.M{"mailid": customer.MailFind}
	val1 := p.UserCollection.FindOne(context.Background(), filter).Decode(&checker)
	fmt.Println(val1)
	if checker.Hatching_Count>5{
		    return "greater",nil
	}
	update := bson.M{
		"$push": bson.M{
			"hatching": customer,
		},
	}
	_, err := p.UserCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Println("Transaction not inserted")
		return "Transaction not inserted", err
	}
	update2 := bson.M{"$inc": bson.M{"hatching_count": 1}} // Replace "your_field_to_increment" with the actual field name
	_, err2 := p.UserCollection.UpdateOne(p.ctx, filter, update2)
	if err2 != nil {
		log.Fatal(err2)
	}

	return "hello", nil
}

func (p *TransactionService) GetHatchingService(filter bson.M) ([]models.Hatching, error) {
	cursor, err := p.UserCollection.Find(p.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(p.ctx) // Don't forget to close the cursor when done

	var results []models.Hatching
	for cursor.Next(p.ctx) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			log.Fatal(err)
		}
		results = append(results, user.Hatching...)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (p *TransactionService) GetEditEmployeeService(filter bson.M, update bson.M) (string, error) {
	_, err := p.UserCollection.UpdateOne(p.ctx, filter, update)
	if err != nil {
		return "failed to update", err
	}
	return "successfully updated", nil
}
func (p *TransactionService) GetEditHatchService(filter, update bson.M) (*mongo.UpdateResult, error) {
	res, err := p.UserCollection.UpdateOne(p.ctx, filter, update)
	if err != nil {
		return res, err
	}
	fmt.Println(res)
	return res, nil
}
func (p *TransactionService) GetDeleteEmployeeService(filter bson.M, update bson.M,filter1 bson.M) (string, error) {
	_, err := p.UserCollection.UpdateOne(p.ctx, filter, update)
	if err != nil {
		return "failed to delete", err
	}
	update1 := bson.M{"$inc": bson.M{"poultry_count": -1}} // Replace "your_field_to_increment" with the actual field name
	_, err2 := p.UserCollection.UpdateOne(p.ctx, filter1, update1)
	if err2 != nil {
		log.Fatal(err2)
	}
	return "successfully delete", nil
}
func (p *TransactionService) GetDeleteHatchService(customer *models.Hatching,filter bson.M, update bson.M,filter2 bson.M,val string) (string, error) {
	update1 := bson.M{
        "$inc": bson.M{
            "egg." + val:   -(customer.Egg),
            "meet." + val: -(customer.Meet),
            "chick." + val: -(customer.Chick),
        },
    }

    filter1 := bson.M{"mailid": customer.MailFind}
    _, err1 := p.GraphCollection.UpdateOne(context.Background(), filter1, update1)
    if err1 != nil {
        fmt.Println("Transaction not inserted")
        return "Transaction not inserted", err1
    }
	
	_, err := p.UserCollection.UpdateOne(p.ctx, filter, update)
	if err != nil {
		return "failed to delete", err
	}

	update2 := bson.M{"$inc": bson.M{"hatching_count": -1}} // Replace "your_field_to_increment" with the actual field name
	_, err2 := p.UserCollection.UpdateOne(p.ctx, filter2, update2)
	if err2 != nil {
		log.Fatal(err2)
	}
	return "successfully delete", nil
}
func (p *TransactionService) GetDeleteOrderService(filter bson.M, update bson.M,filter2 bson.M) (string, error) {
	_, err := p.UserCollection.UpdateOne(p.ctx, filter, update)
	if err != nil {
		return "failed to delete", err
	}
	update1 := bson.M{"$inc": bson.M{"order_count": -1}} // Replace "your_field_to_increment" with the actual field name
	_, err2 := p.UserCollection.UpdateOne(p.ctx, filter2, update1)
	if err2 != nil {
		log.Fatal(err2)
	}
	return "successfully delete", nil
}

func (p *TransactionService) CreateServiceOrder(customer *models.Order,val string) (string, error) {
	var checker models.User
	update11 := bson.M{
        "$inc": bson.M{
            "order." + val:   customer.Egg,
          
        },
    }

    filter11 := bson.M{"mailid": customer.MailFind}
    _, err11 := p.GraphCollection.UpdateOne(context.Background(), filter11, update11)
    if err11 != nil {
        fmt.Println("Transaction not inserted")
        return "Transaction not inserted", err11
    }
	filter := bson.M{"mailid": customer.MailFind}
	val1 := p.UserCollection.FindOne(context.Background(), filter).Decode(&checker)
	fmt.Println(val1)
	
	update := bson.M{
		"$push": bson.M{
			"order": customer,
		},
	}

	// Execute the update operation
	_, err := p.UserCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Println("Transaction not inserted")
		return "Transaction not inserted", err
	}
	update1 := bson.M{"$inc": bson.M{"order_count": 1}} // Replace "your_field_to_increment" with the actual field name
	_, err2 := p.UserCollection.UpdateOne(p.ctx, filter, update1)
	if err2 != nil {
		log.Fatal(err2)
	}

	return "hello", nil
}

func (p *TransactionService) GetOrderService(filter bson.M) ([]models.Order, error) {
	cursor, err := p.UserCollection.Find(p.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(p.ctx) // Don't forget to close the cursor when done

	var results []models.Order
	for cursor.Next(p.ctx) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			log.Fatal(err)
		}
		results = append(results, user.Order...)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (p *TransactionService) GetEditOrderService(filter bson.M, update bson.M) (string, error) {
	_, err := p.UserCollection.UpdateOne(p.ctx, filter, update)
	if err != nil {
		return "failed to update", err
	}
	return "successfully updated", nil
}

//here we start

func (p *TransactionService) CreateCustomer(user *models.User) (int, error) {
	filter := bson.M{"mailid": user.MailID}
	var val int
	val = 1

	// Perform the find query.
	err1 := p.UserCollection.FindOne(context.Background(), filter)
	if err1 != nil {
		if err1.Err() == mongo.ErrNoDocuments {
			val = 0
			fmt.Println("No matching document found.")
			fmt.Println(val)
		} else {
			// Handle other errors gracefully
			return val, err1.Err()
		}
	}

	user1 := models.User{
		MailID:   user.MailID,
		Password: user.Password,
		Poultry:  []models.Poultry{},
		Hatching: []models.Hatching{},
		Order:    []models.Order{},
	}

	_, err := p.UserCollection.InsertOne(p.ctx, &user1)

	if err != nil {
		return val, err
	}
	emptyGraph := models.Graph{
		Mailid: user.MailID, // Set the mailid field to an appropriate value
		Egg: models.Graph_Egg{
			Jan: 0,
			Feb: 0,
			Mar: 0,
			Apr: 0,
			May: 0,
			Jun: 0,
			Jul: 0,
			Aug: 0,
			Sep: 0,
			Oct: 0,
			Nov: 0,
			Dec: 0,
		},
		Meet: models.Graph_Meet{
			Jan: 0,
			Feb: 0,
			Mar: 0,
			Apr: 0,
			May: 0,
			Jun: 0,
			Jul: 0,
			Aug: 0,
			Sep: 0,
			Oct: 0,
			Nov: 0,
			Dec: 0,
		},
		Chick: models.Graph_Chick{
			Jan: 0,
			Feb: 0,
			Mar: 0,
			Apr: 0,
			May: 0,
			Jun: 0,
			Jul: 0,
			Aug: 0,
			Sep: 0,
			Oct: 0,
			Nov: 0,
			Dec: 0,
		},
		Order: models.Graph_Order{
			Jan: 0,
			Feb: 0,
			Mar: 0,
			Apr: 0,
			May: 0,
			Jun: 0,
			Jul: 0,
			Aug: 0,
			Sep: 0,
			Oct: 0,
			Nov: 0,
			Dec: 0,
		},
	}
	
	// Insert the emptyGraph into the MongoDB collection
	_, err3 := p.GraphCollection.InsertOne(p.ctx, &emptyGraph)
	if err != nil {
		return val, err3
	}

	return val, nil
}

func (p *TransactionService) FindService(filter bson.M) (int, error) {
	cursor, err := p.UserCollection.Find(p.ctx, filter)
	fmt.Println("Good",cursor)
	var results []models.User
	var p1 = 1
	for cursor.Next(context.TODO()) {

		var result models.User // Replace YourStruct with the type of documents in your collection
		if err := cursor.Decode(&result); err != nil {
			log.Fatal(err)
		}
		if p1 == 1 {
			var temp models.Login
			temp.Signup = result.MailID
			temp.Password = result.Password
			_, err1 := p.CheckCollection.InsertOne(context.Background(), temp)
			if err1 != nil {
				return 0, err
			}
			p1++
		}
		results = append(results, result)
	}
	if err != nil {
		return 0, err
	}

	if len(results) != 0 {
		return 1, nil
	}
	return 0, nil
}
func (p *TransactionService) GetAdminService(filter bson.M) (int, error) {
	cursor, _ := p.AdminCollection.Find(p.ctx, filter)
	fmt.Println(cursor)
	for cursor.Next(context.TODO()) {
		return 1, nil
	}
	return 0, nil
}
func (p *TransactionService) FindLastService() (*models.Login, error) {
	filter := bson.M{}
	cursor, err := p.CheckCollection.Find(p.ctx, filter, options.Find().SetSort(bson.M{"timer": -1}))
	if err != nil {
		// Handle the error appropriately, e.g., log.Fatal(err) or return an error response.
		log.Fatal(err)
		return nil, err
	}
	for cursor.Next(p.ctx) {
		var result models.Login
		if err := cursor.Decode(&result); err != nil {
			// Handle the decoding error, e.g., log.Fatal(err) or return an error response.
			log.Fatal(err)
			return nil, err
		}

		return &result, nil
	}
	return nil, nil

}

func (p *TransactionService) UpdateService(filter bson.M, update bson.M) (int, error) {
	_, err := p.CustomerCollection.UpdateOne(p.ctx, filter, update)
	if err != nil {
		return 0, err // Return a more descriptive error message
	}

	return 1, nil
}

func (p *TransactionService) GetAdminInfoService(filter bson.M) ([]models.Admin, error)  {

	
	cursor, err := p.UserCollection.Find(p.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(p.ctx) // Don't forget to close the cursor when done

	var results []models.Admin
	for cursor.Next(p.ctx) {
		var user models.Admin
		if err := cursor.Decode(&user); err != nil {
			log.Fatal(err)
		}
		results = append(results, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (p *TransactionService)AdminUpdateService(filter bson.M,update bson.M) (error)  {
	_, err := p.AdminCollection.UpdateOne(p.ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}


func (p *TransactionService) AdminDataService(filter bson.M) (models.Admin1, error) {
	
	var Admin models.Admin1
	err := p.AdminCollection.FindOne(p.ctx, filter).Decode(&Admin)
	if err != nil {
		return models.Admin1{}, err
}
return Admin, nil
}

func (p *TransactionService) DeleteUserService(filter bson.M) ( error) {
	
	
	_,err := p.UserCollection.DeleteOne(p.ctx, filter)
	if err != nil {
		return  err
}
return nil
}

func (p *TransactionService) MonthService(filter bson.M) (models.Graph, error) {
	var modelgraph models.Graph
	err := p.GraphCollection.FindOne(p.ctx, filter).Decode(&modelgraph)

	if err != nil {
		return models.Graph{}, err
	}

	return modelgraph,nil
}

func (p *TransactionService) GetCustomerOneService(filter bson.M) (models.User, error) {
	var findonemodel models.User
	err := p.UserCollection.FindOne(p.ctx, filter).Decode(&findonemodel)
	if err != nil {
		return models.User{}, err
	}

	

	return findonemodel, nil
}