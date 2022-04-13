package main

import (
	"be/api/aws"
	"be/api/aws/s3"
	"be/configs"
	"be/delivery/controllers/movie"
	movieLogic "be/delivery/logic/movie"
	"be/delivery/routes"
	movieRepo "be/repository/movie"
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

	var movieRepo = movieRepo.New(db)
	var movieLogic = movieLogic.New()
	var movieCont = movie.New(movieRepo, awsS3, movieLogic)
	var e = echo.New()

	routes.Routes(e, movieCont)
	log.Fatal(e.Start(fmt.Sprintf(":%d", config.PORT)))
}
