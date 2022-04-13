package auth

import (
	"be/configs"
	"be/entities"
	"be/repository/user"
	"be/utils"
	"testing"

	"github.com/lithammer/shortuuid"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	var config = configs.GetConfig()
	var db = utils.InitDB(config)
	var r = New(db)
	db.Migrator().DropTable(&entities.User{})
	db.Migrator().DropTable(&entities.Product{})
	db.AutoMigrate(&entities.User{})

	t.Run("success run login", func(t *testing.T) {
		var mock1 = entities.User{UserName: shortuuid.New(), Email: shortuuid.New(), Password: shortuuid.New(), Name: shortuuid.New()}

		if _, err := user.New(db).Create(mock1); err != nil {
			t.Log(err)
			t.Fatal()
		}

		var res, err = r.Login(mock1.UserName, mock1.Password)

		assert.Nil(t, err)
		assert.NotNil(t, res)
	})

	t.Run("error userName", func(t *testing.T) {
		var mock1 = entities.User{UserName: shortuuid.New(), Email: shortuuid.New(), Password: shortuuid.New(), Name: shortuuid.New()}

		if _, err := user.New(db).Create(mock1); err != nil {
			t.Log(err)
			t.Fatal()
		}

		var _, err = r.Login("anonim1", mock1.Password)
		t.Log(err)
		assert.NotNil(t, err)
	})

	t.Run("error password", func(t *testing.T) {
		var mock1 = entities.User{UserName: shortuuid.New(), Email: shortuuid.New(), Password: shortuuid.New(), Name: shortuuid.New()}

		if _, err := user.New(db).Create(mock1); err != nil {
			t.Log(err)
			t.Fatal()
		}

		var _, err = r.Login(mock1.UserName, "anonim1")

		assert.NotNil(t, err)
	})
}
