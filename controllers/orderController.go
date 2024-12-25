package controllers

import (
	"context"
	"golang-restaurant-management/database"
	"golang-restaurant-management/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var orderCollection *mongo.Collection = database.OpenCollection(database.Client, "order")

func GetOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		result, err := orderCollection.Find(ctx, nil)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occured while fetching order items"})
			defer cancel()
			return
		}

		var allOrders []bson.M
		if err = result.All(ctx, &allOrders); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occured while decoding order items"})
			defer cancel()
			return
		}

		defer cancel()
		c.JSON(http.StatusOK, gin.H{
			"data": allOrders,
		})
	}
}

func GetOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		orderId := c.Param("order_id")

		var order bson.M

		err := orderCollection.FindOne(ctx, bson.M{"order_id": orderId}).Decode(&order)
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occured while decoding order itmes"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": order,
		})
	}
}

func CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var order models.Order

		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in Binding the JSON value"})
			defer cancel()
			return
		}

		validateErr := validate.Struct(order)

		if validateErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error in validation of provide data",
			})
			defer cancel()
			return
		}

		order.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		order.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		order.ID = primitive.NewObjectID()
		order.Order_id = order.ID.Hex()

		res, err := orderCollection.InsertOne(ctx, order)
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error while inserting order object",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": res,
		})
	}
}

func UpdateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var table models.Table
		var order models.Order

		var updateObj primitive.D

		orderId := c.Param("order_id")

		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			defer cancel()
			return
		}

		if order.Table_id != nil {
			err := orderCollection.FindOne(ctx, bson.M{"table_id": order.Table_id}).Decode(&table)

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "message:Table was not found"})
				defer cancel()
				return
			}

			updateObj = append(updateObj, bson.E{Key: "table_id", Value: order.Table_id})

		}

		order.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{Key: "updated_at", Value: order.Updated_at})

		upsert := true

		filter := bson.M{"order_id": orderId}
		opt := options.UpdateOptions{
			Upsert: &upsert,
		}

		result, err := orderCollection.UpdateOne(
			ctx,
			filter,
			bson.D{
				{Key: "$set", Value: updateObj},
			},
			&opt,
		)
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "order item update failed"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": result,
		})
	}
}
