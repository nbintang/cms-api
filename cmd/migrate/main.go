package main

import (
	"context"
	"rest-fiber/pkg"

	"github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()
	pkg.LoadEnv()
	if err := InitMigrate(ctx); err != nil {
		logrus.Fatalf("Migration failed: %v", err)
	}
	logrus.Println("Migration Succeed")
}
