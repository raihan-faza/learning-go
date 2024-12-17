package migrate

import (
	"github.com/raihan-faza/learning-go/initializers"
	"github.com/raihan-faza/learning-go/models"
)

func init() {
	initializers.ConnectDB()

}

func main() {
	initializers.DB.AutoMigrate(&models.User{}, &models.Post{})
}
