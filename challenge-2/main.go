package main

import (
	"be/api/aws"
	"be/api/aws/s3"
	"be/configs"
	"be/delivery/controllers/auth"
	"be/delivery/controllers/product"
	"be/delivery/controllers/user"
	productLogic "be/delivery/logic/product"
	userLogic "be/delivery/logic/user"
	"be/delivery/routes"
	authRepo "be/repository/auth"
	productRepo "be/repository/product"
	userRepo "be/repository/user"
	"be/utils"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	var config = configs.GetConfig()
	var db = utils.InitDB(config)

	var awsS3Conf = aws.InitS3(config.S3_REGION, config.S3_ID, config.S3_SECRET)

	var awsS3 = s3.New(awsS3Conf)

	var authRepo = authRepo.New(db)
	var authCont = auth.New(authRepo)

	var userRepo = userRepo.New(db)
	var userLogic = userLogic.New()
	var userCont = user.New(userRepo, awsS3, userLogic)

	var productRepo = productRepo.New(db)
	var productLogic = productLogic.New()
	var productCont = product.New(productRepo, awsS3, productLogic)
	var e = echo.New()

	routes.Routes(e, authCont, userCont, productCont)
	log.Fatal(e.Start(fmt.Sprintf(":%d", config.PORT)))
}
