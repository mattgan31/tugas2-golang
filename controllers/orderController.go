package controllers

import (
	"golang-tugas2/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func (idb *InDB) CreateOrder(ctx *gin.Context) {
	order := models.Order{}
	order.Ordered_At = time.Now()

	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	newOrder := models.Order{
		Ordered_At:    order.Ordered_At,
		Customer_Name: order.Customer_Name,
		Items:         order.Items,
	}

	if err := idb.DB.Create(&newOrder).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, newOrder)

}

func (idb *InDB) ShowOrders(ctx *gin.Context) {

	var (
		orders []models.Order
		result gin.H
	)

	idb.DB.Find(&orders)
	if len(orders) <= 0 {
		result = gin.H{
			"resul": nil,
			"count": 0,
		}
	} else {
		result = gin.H{
			"resul": orders,
			"count": len(orders),
		}
	}
	ctx.JSON(http.StatusOK, result)
}

func (idb *InDB) UpdateOrder(ctx *gin.Context) {
	id := ctx.Param("orderID")
	order := models.Order{}
	
	var(
		
		result gin.H
	)

	newOrder := models.Order{
		Ordered_At:    order.Ordered_At,
		Customer_Name: order.Customer_Name,
		Items:         order.Items,
	}

	if err := ctx.ShouldBindJSON(&newOrder); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err:=idb.DB.First(&order, id).Error
	if err!=nil{
		result = gin.H{
			"result":"data not found",
		}
	}
	
	
	err = idb.DB.Model(&order).Updates(newOrder).Error
	if err!=nil{
		result = gin.H{
			"result":"update fail",
		}
	} else {
		result = gin.H{
			"result":"update data success",
		}
	}

	ctx.JSON(http.StatusOK, result)

	// if err := idb.DB.Where("order_id = ?", ctx.Param("orderID")).First(&order).Error; err != nil {
	// 	ctx.JSON(http.StatusNotFound, gin.H{
	// 		"message": err.Error(),
	// 	})

	// 	return
	// }

	// idb.DB.Unscoped().Where("order_id = ?", order.Order_ID).Delete(item)

	// if err := idb.DB.Unscoped().Where("order_id = ?", order.Order_ID).Delete(item).Error; err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{
	// 		"message": err.Error(),
	// 	})

	// 	return
	// }

	// if err := ctx.ShouldBind(&order); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{
	// 		"message": err.Error(),
	// 	})

	// 	return
	// }

	// if err := idb.DB.Save(order).Error; err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{
	// 		"message": err.Error(),
	// 	})

	// 	return
	// }
}

func (idb *InDB) DeleteOrder(ctx *gin.Context) {
	order := models.Order{}

	if err := idb.DB.Where("order_id = ?", ctx.Param("orderID")).First(&order).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := idb.DB.Select(clause.Associations).Delete(&order).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Order Successfully Deleted",
	})
}
