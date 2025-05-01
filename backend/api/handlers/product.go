package handlers

import (
	"net/http"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/AMETORY/ametory-erp-modules/inventory"
	"github.com/AMETORY/ametory-erp-modules/shared/models"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	ctx          *context.ERPContext
	inventorySrv *inventory.InventoryService
}

func NewProductHandler(ctx *context.ERPContext) *ProductHandler {
	inventorySrv, ok := ctx.InventoryService.(*inventory.InventoryService)
	if !ok {
		panic("inventory service not found")
	}
	return &ProductHandler{
		ctx:          ctx,
		inventorySrv: inventorySrv,
	}
}

func (p *ProductHandler) GetProductHandler(c *gin.Context) {
	id := c.Param("id")
	product, err := p.inventorySrv.ProductService.GetProductByID(id, c.Request)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": product, "message": "Product retrieved successfully"})
}

func (p *ProductHandler) ListProductsHandler(c *gin.Context) {
	products, err := p.inventorySrv.ProductService.GetProducts(*c.Request, c.Query("search"))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": products, "message": "Products retrieved successfully"})
}

func (p *ProductHandler) CreateProductHandler(c *gin.Context) {
	var input models.ProductModel
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	companyID := c.MustGet("companyID").(string)
	input.CompanyID = &companyID
	err = p.inventorySrv.ProductService.CreateProduct(&input)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, gin.H{"message": "Product created successfully", "data": input})
}

func (p *ProductHandler) UpdateProductHandler(c *gin.Context) {
	id := c.Param("id")
	var input models.ProductModel
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if id != input.ID {
		c.JSON(400, gin.H{"error": "ID mismatch"})
	}

	// product, err := p.inventorySrv.ProductService.GetProductByID(id, c.Request)
	// if err != nil {
	// 	c.JSON(500, gin.H{"error": err.Error()})
	// 	return
	// }

	// for _, v := range product.ProductImages {
	// 	p.inventorySrv.ProductService.DeleteImageOfProduct(id, v.ID)
	// }

	err = p.inventorySrv.ProductService.UpdateProduct(&input)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Product updated successfully"})
}

func (p *ProductHandler) DeleteProductHandler(c *gin.Context) {
	id := c.Param("id")
	err := p.inventorySrv.ProductService.DeleteProduct(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Product deleted successfully"})
}

func (p *ProductHandler) DeleteImageProductHandler(c *gin.Context) {
	p.ctx.Request = c.Request
	// Implement logic to delete an image from a product

	id := c.Param("id")
	imageID := c.Param("imageId")
	err := p.inventorySrv.ProductService.DeleteImageOfProduct(id, imageID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Image deleted from product successfully"})
}
