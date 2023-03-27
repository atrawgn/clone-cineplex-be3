package controller

import (
	"github.com/atrawiguna/golang-restapi-gorm/database"
	"github.com/atrawiguna/golang-restapi-gorm/model/entity"
	"github.com/atrawiguna/golang-restapi-gorm/model/request"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"log"
)

func FilmControllerGet(ctx *fiber.Ctx) error {
	var films []entity.Film
	err := database.DB.Find(&films)
	if err != nil {
		log.Println(err)
	}
	return ctx.JSON(films)

}

func FilmControllerCreate(ctx *fiber.Ctx) error {
	film := new(request.FilmCreateRequest)
	if err := ctx.BodyParser(film); err != nil {
		return err
	}

	// VALIDASI REQUEST
	validate := validator.New()
	errValidate := validate.Struct(film)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Gagal",
			"error":   errValidate.Error(),
		})
	}

	newFilm := entity.Film{
		Judul:     film.Judul,
		JenisFilm: film.JenisFilm,
		Produser:  film.Produser,
		Sutradara: film.Sutradara,
		Penulis:   film.Penulis,
		Produksi:  film.Produksi,
		Casts:     film.Casts,
		Sinopsis:  film.Sinopsis,
	}

	errCreateFilm := database.DB.Create(&newFilm).Error
	if errCreateFilm != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Tidak berhasil menyimpan data",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Berhasil",
		"data":    newFilm,
	})
}

func FilmControllerGetById(ctx *fiber.Ctx) error {
	filmId := ctx.Params("id")

	var film entity.Film
	err := database.DB.First(&film, "id = ?", filmId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data tidak ditemukan",
		})
	}

	/*userResponse := response.UserResponse{
		ID:        user.ID,
		Nama:      user.Nama,
		JenisFilm: user.JenisFilm,
		Produser:  user.Produser,
		Sutradara: user.Sutradara,
		Penulis:   user.Penulis,
		Produksi:  user.Produksi,
		Casts:     user.Casts,
		Sinopsis:  user.Sinopsis,
	}*/
	return ctx.JSON(fiber.Map{
		"message": "Sukses",
		"data":    film,
	})
}

func FilmControllerUpdate(ctx *fiber.Ctx) error {
	filmRequest := new(request.FilmUpdateRequest)
	if err := ctx.BodyParser(filmRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var film entity.Film

	userId := ctx.Params("id")
	// CHECK AVAILABLE USER
	err := database.DB.First(&film, "id = ?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "data tidak valid",
		})
	}

	// UPDATE USER DATA
	if filmRequest.Judul != "" {
		film.Judul = filmRequest.Judul
	}
	errUpdate := database.DB.Save(&film).Error
	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	film.Judul = filmRequest.Judul

	return ctx.JSON(fiber.Map{
		"message": "Sukses",
		"data":    film,
	})
}

func FilmControllerDelete(ctx *fiber.Ctx) error {
	filmId := ctx.Params("id")
	var film entity.Film

	// CHECK AVAILABLE USER
	err := database.DB.Debug().First(&film, "id=?", filmId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "film tidak ditemukan",
		})
	}

	errDelete := database.DB.Debug().Delete(&film).Error
	if errDelete != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "film telah dihapus",
	})
}
