package migration

import (
	"github.com/amirazad1/ELearning/config"
	"github.com/amirazad1/ELearning/constant"
	database "github.com/amirazad1/ELearning/data/db"
	models "github.com/amirazad1/ELearning/data/model"
	"github.com/amirazad1/ELearning/pkg/logging"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const countStarExp = "count(*)"

var logger = logging.NewLogger(config.GetConfig())

func Up1() {
	db := database.GetDb()

	createTables(db)
	createDefaultUserInformation(db)
}

func createTables(database *gorm.DB) {
	tables := []interface{}{}

	// User
	tables = addNewTable(database, models.User{}, tables)
	tables = addNewTable(database, models.Role{}, tables)
	tables = addNewTable(database, models.UserRole{}, tables)

	err := database.Migrator().CreateTable(tables...)
	if err != nil {
		logger.Error(logging.Postgres, logging.Migration, err.Error(), nil)
	}
	logger.Info(logging.Postgres, logging.Migration, "tables created", nil)
}

func addNewTable(database *gorm.DB, model interface{}, tables []interface{}) []interface{} {
	if !database.Migrator().HasTable(model) {
		tables = append(tables, model)
	}
	return tables
}

func createDefaultUserInformation(database *gorm.DB) {

	adminRole := models.Role{Name: constant.AdminRoleName}
	createRoleIfNotExists(database, &adminRole)

	defaultRole := models.Role{Name: constant.DefaultRoleName}
	createRoleIfNotExists(database, &defaultRole)

	u := models.User{Username: constant.DefaultUserName, FirstName: "Amir", LastName: "Azad",
		MobileNumber: "09131112222", Email: "amir@azad.com"}
	pass := "12345678"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	u.Password = string(hashedPassword)

	createAdminUserIfNotExists(database, &u, adminRole.Id)

}

func createRoleIfNotExists(database *gorm.DB, r *models.Role) {
	exists := 0
	database.
		Model(&models.Role{}).
		Select("1").
		Where("name = ?", r.Name).
		First(&exists)
	if exists == 0 {
		database.Create(r)
	}
}

func createAdminUserIfNotExists(database *gorm.DB, u *models.User, roleId int) {
	exists := 0
	database.
		Model(&models.User{}).
		Select("1").
		Where("username = ?", u.Username).
		First(&exists)
	if exists == 0 {
		database.Create(u)
		ur := models.UserRole{UserId: u.Id, RoleId: roleId}
		database.Create(&ur)
	}
}

func Down1() {
	// nothing
}
