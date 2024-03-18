package infrastructure

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
	"product-app/common/postgresql"
	"product-app/persistence"
	"testing"
)

var poroductRepository persistence.IProductRepository
var dbPool *pgxpool.Pool

func TestMain(m *testing.M) {
	ctx := context.Background()
	dbPool = postgresql.GetConnectionPool(ctx, postgresql.Config{
		Host:                  "localhost",
		Port:                  "6432",
		DbName:                "productapp",
		UserName:              "postgres",
		Password:              "postgres",
		MaxConnections:        "100",
		MaxConnectionIdleTime: "30s",
	})
	poroductRepository = persistence.NewProductRepository(dbPool)
	exitCode := m.Run()
	os.Exit(exitCode)
}
func TestGetAllProducts(t *testing.T) {
	fmt.Println(dbPool)
	fmt.Println(poroductRepository)
}
