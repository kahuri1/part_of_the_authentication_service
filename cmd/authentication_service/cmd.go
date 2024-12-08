package main

import (
	authServer "github.com/kahuri1/part_of_the_authentication_service"
	"github.com/kahuri1/part_of_the_authentication_service/iternal/handler"
	"github.com/kahuri1/part_of_the_authentication_service/iternal/model"
	"github.com/kahuri1/part_of_the_authentication_service/iternal/repository"
	"github.com/kahuri1/part_of_the_authentication_service/iternal/service"
	"github.com/kahuri1/part_of_the_authentication_service/pkg/email"
	"github.com/kahuri1/part_of_the_authentication_service/pkg/hash"
	"github.com/kahuri1/part_of_the_authentication_service/pkg/otp"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var secretKey = []byte("EFFVKC:NRJKVNPWELND")

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initialization configs: %s", err.Error())
	}
	db, err := repository.NewPostgresDB(model.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLmode:  viper.GetString("db.sslmode"),
		Password: viper.GetString("db.password"),
	})
	if err != nil {
		logrus.Errorf("Failed initialization db: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	salt := "#$%&@^#HDJhHDY#@"
	hasher := hash.NewSHA1Hasher(salt)
	otpGenerator := otp.NewGOTPGenerator()
	emailSender := email.NewSMTPEmailSender("aspmx.l.google.com", "25",
		"vova150820@gmail.com", "nubb nosq kjze jboe")
	services := service.NewService(repos, hasher, otpGenerator, emailSender)
	handlers := handler.Newhandler(services)

	srv := new(authServer.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error running http server: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
