package product

import (
	"be/configs"
	"be/entities"
	"be/repository/user"
	"be/utils"
	"testing"

	"github.com/lithammer/shortuuid"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	var config = configs.GetConfig()
	var db = utils.InitDB(config)
	var r = New(db)
	db.Migrator().DropTable(&entities.User{})
	db.Migrator().CreateTable(&entities.Product{})
	db.AutoMigrate(&entities.Product{})

	t.Run("success create", func(t *testing.T) {
		var mock1 = entities.User{UserName: shortuuid.New(), Email: shortuuid.New(), Password: shortuuid.New(), Name: shortuuid.New()}

		res, err := user.New(db).Create(mock1)

		if err != nil {
			t.Log(err)
			t.Fatal()
		}

		var mock2 = entities.Product{Name: shortuuid.New(), Price: shortuuid.New(), Description: shortuuid.New(), Stock: 1}

		res2, err := r.Create(res.User_uid, mock2)

		assert.Nil(t, err)
		assert.NotNil(t, res2)
	})
}

func TestDelete(t *testing.T) {
	var config = configs.GetConfig()
	var db = utils.InitDB(config)
	var r = New(db)
	db.Migrator().DropTable(&entities.User{})
	db.Migrator().CreateTable(&entities.Product{})
	db.AutoMigrate(&entities.Product{})

	t.Run("success create", func(t *testing.T) {
		var mock1 = entities.User{UserName: shortuuid.New(), Email: shortuuid.New(), Password: shortuuid.New(), Name: shortuuid.New()}

		res, err := user.New(db).Create(mock1)

		if err != nil {
			t.Log(err)
			t.Fatal()
		}

		var mock2 = entities.Product{Name: shortuuid.New(), Price: shortuuid.New(), Description: shortuuid.New(), Stock: 1}

		res2, err := r.Create(res.User_uid, mock2)

		if err != nil {
			t.Log(err)
			t.Fatal()
		}

		res3, err := r.Delete(res2.Product_uid)

		assert.Nil(t, err)
		assert.Equal(t, true, res3.DeletedAt.Valid)
	})
}

func TestUpate(t *testing.T) {
	var config = configs.GetConfig()
	var db = utils.InitDB(config)
	var r = New(db)
	db.Migrator().DropTable(&entities.User{})
	db.Migrator().CreateTable(&entities.Product{})
	db.AutoMigrate(&entities.Product{})

	t.Run("success create", func(t *testing.T) {
		var mock1 = entities.User{UserName: shortuuid.New(), Email: shortuuid.New(), Password: shortuuid.New(), Name: shortuuid.New()}

		res, err := user.New(db).Create(mock1)

		if err != nil {
			t.Log(err)
			t.Fatal()
		}

		var mock2 = entities.Product{Name: shortuuid.New(), Price: shortuuid.New(), Description: shortuuid.New(), Stock: 1}

		res2, err := r.Create(res.User_uid, mock2)

		if err != nil {
			t.Log(err)
			t.Fatal()
		}

		res3, err := r.Update(res2.Product_uid, entities.Product{Name: shortuuid.New(), Price: shortuuid.New(), Description: shortuuid.New(), Stock: 10})

		assert.Nil(t, err)
		assert.NotNil(t, res3)
	})
}

func TestGet(t *testing.T) {
	var config = configs.GetConfig()
	var db = utils.InitDB(config)
	var r = New(db)
	db.Migrator().DropTable(&entities.User{})
	db.Migrator().CreateTable(&entities.Product{})
	db.AutoMigrate(&entities.Product{})

	t.Run("success create", func(t *testing.T) {
		var mock1 = entities.User{UserName: shortuuid.New(), Email: shortuuid.New(), Password: shortuuid.New(), Name: shortuuid.New()}

		res, err := user.New(db).Create(mock1)

		if err != nil {
			t.Log(err)
			t.Fatal()
		}

		var mock2 = entities.Product{Name: shortuuid.New(), Price: shortuuid.New(), Description: shortuuid.New(), Stock: 1}

		res2, err := r.Create(res.User_uid, mock2)

		if err != nil {
			t.Log(err)
			t.Fatal()
		}

		if _, err := r.Create(res.User_uid, entities.Product{Name: shortuuid.New(), Price: shortuuid.New(), Description: shortuuid.New(), Stock: 1}); err != nil {
			t.Log(err)
			t.Fatal()
		}

		if _, err := r.Create(res.User_uid, entities.Product{Name: shortuuid.New(), Price: shortuuid.New(), Description: shortuuid.New(), Stock: 1}); err != nil {
			t.Log(err)
			t.Fatal()
		}

		if _, err := r.Create(res.User_uid, entities.Product{Name: shortuuid.New(), Price: shortuuid.New(), Description: shortuuid.New(), Stock: 1}); err != nil {
			t.Log(err)
			t.Fatal()
		}

		res3, err := r.Get("", "")
		assert.Nil(t, err)
		assert.Equal(t, 4, len(res3.Responses))
		res3, err = r.Get(res.User_uid, "")
		assert.Nil(t, err)
		assert.Equal(t, 4, len(res3.Responses))
		res3, err = r.Get(res.User_uid, res2.Product_uid)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(res3.Responses))
		res3, err = r.Get("", res2.Product_uid)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(res3.Responses))
	})
}
