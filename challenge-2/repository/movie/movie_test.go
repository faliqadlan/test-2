package movie

import (
	"be/configs"
	"be/entities"
	"be/utils"
	"fmt"
	"testing"

	"github.com/lithammer/shortuuid"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	var config = configs.GetConfig()
	var db = utils.InitDB(config)
	var r = New(db)
	db.Migrator().DropTable(&entities.Movie{})
	db.AutoMigrate(&entities.Movie{})

	t.Run("success create", func(t *testing.T) {
		var mock1 = entities.Movie{Title: shortuuid.New(), Description: shortuuid.New(), Artist: shortuuid.New(), Genres: shortuuid.New()}

		res, err := r.Create(mock1)

		assert.Nil(t, err)
		assert.NotNil(t, res)
	})
}

func TestDelete(t *testing.T) {
	var config = configs.GetConfig()
	var db = utils.InitDB(config)
	var r = New(db)
	db.Migrator().DropTable(&entities.Movie{})
	db.AutoMigrate(&entities.Movie{})

	t.Run("success delete", func(t *testing.T) {
		var mock1 = entities.Movie{Title: shortuuid.New(), Description: shortuuid.New(), Artist: shortuuid.New(), Genres: shortuuid.New()}

		res, err := r.Create(mock1)

		if err != nil {
			t.Log(err)
			t.Fatal()
		}

		res2, err := r.Delete(res.Movie_uid)

		assert.Nil(t, err)
		assert.Equal(t, true, res2.DeletedAt.Valid)
	})
}

func TestUpate(t *testing.T) {
	var config = configs.GetConfig()
	var db = utils.InitDB(config)
	var r = New(db)
	db.Migrator().DropTable(&entities.Movie{})
	db.AutoMigrate(&entities.Movie{})

	t.Run("success Update", func(t *testing.T) {

		var mock1 = entities.Movie{Title: shortuuid.New(), Description: shortuuid.New(), Artist: shortuuid.New(), Genres: shortuuid.New()}

		res, err := r.Create(mock1)

		if err != nil {
			t.Log(err)
			t.Fatal()
		}

		res2, err := r.Update(res.Movie_uid, entities.Movie{Title: shortuuid.New(), Description: shortuuid.New(), Artist: shortuuid.New(), Genres: shortuuid.New()})

		assert.Nil(t, err)
		assert.NotNil(t, res2)
	})
}

func TestGet(t *testing.T) {
	var config = configs.GetConfig()
	var db = utils.InitDB(config)
	var r = New(db)
	db.Migrator().DropTable(&entities.Movie{})
	db.AutoMigrate(&entities.Movie{})

	t.Run("success get", func(t *testing.T) {

		for i := 1; i < 11; i++ {

			if i%2 == 0 {
				if _, err := r.Create(entities.Movie{Title: shortuuid.New() + "judul" + "genap" + fmt.Sprintf("%d", i), Description: shortuuid.New() + "description" + "genap" + fmt.Sprintf("%d", i), Artist: shortuuid.New() + "artist" + "genap" + fmt.Sprintf("%d", i), Genres: shortuuid.New() + "genres" + "genap" + fmt.Sprintf("%d", i)}); err != nil {
					t.Log(err)
					t.Fatal()
				}
			} else {
				if _, err := r.Create(entities.Movie{Title: shortuuid.New() + "judul" + "ganjil" + fmt.Sprintf("%d", i), Description: shortuuid.New() + "description" + "ganjil" + fmt.Sprintf("%d", i), Artist: shortuuid.New() + "artist" + "ganjil" + fmt.Sprintf("%d", i), Genres: shortuuid.New() + "genres" + "ganjil" + fmt.Sprintf("%d", i)}); err != nil {
					t.Log(err)
					t.Fatal()
				}
			}

		}

		res1, err := r.Create(entities.Movie{Title: shortuuid.New() + "anonim" + "genap" + "11", Description: shortuuid.New() + "description" + "genap" + "11", Artist: shortuuid.New() + "artist" + "genap" + "11", Genres: shortuuid.New() + "genres" + "genap" + "11"})

		if err != nil {
			t.Log(err)
			t.Fatal()
		}

		res, err := r.Get("", "", "", "", "", 2, 4)

		assert.Nil(t, err)
		assert.NotNil(t, res)

		res, err = r.Get("genap2", "", "", "", "", 0, 2)

		assert.Nil(t, err)
		assert.Contains(t, res.Responses[0].Title, "genap2")

		res, err = r.Get("", "genap2", "", "", "", 0, 2)

		assert.Nil(t, err)
		assert.Contains(t, res.Responses[0].Description, "genap2")

		res, err = r.Get("", "", "genap2", "", "", 0, 2)

		assert.Nil(t, err)
		assert.Contains(t, res.Responses[0].Artist, "genap2")

		res, err = r.Get("", "", "", "genap", "", 0, 2)

		assert.Nil(t, err)
		assert.Contains(t, res.Responses[0].Genres, "genap")

		res, err = r.Get("", "", "", "", res1.Movie_uid, 0, 2)

		assert.Nil(t, err)
		assert.Equal(t, 1, len(res.Responses))

		res, err = r.Get("anonim", "", "", "", res1.Movie_uid, 0, 1)

		assert.Nil(t, err)
		assert.Equal(t, 1, len(res.Responses))
	})
}
