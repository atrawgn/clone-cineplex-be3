package request

type FilmCreateRequest struct {
	Judul     string `json:"judul" validate:"required"`
	JenisFilm string `json:"jenis_film" validate:"required"`
	Produser  string `json:"produser" validate:"required"`
	Sutradara string `json:"sutradara" validate:"required"`
	Penulis   string `json:"penulis" validate:"required"`
	Produksi  string `json:"produksi" validate:"required"`
	Casts     string `json:"casts" validate:"required"`
	Sinopsis  string `json:"sinopsis" validate:"required"`
}

type FilmUpdateRequest struct {
	Judul     string `json:"judul"`
	JenisFilm string `json:"jenis_film"`
	Produser  string `json:"produser"`
	Sutradara string `json:"sutradara"`
	Penulis   string `json:"penulis"`
	Produksi  string `json:"produksi"`
	Casts     string `json:"casts"`
	Sinopsis  string `json:"sinopsis"`
}

// USER SECTION
type UserCreateRequest struct {
	Nama     string `json:"nama"`
	Email    string `json:"email"validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
}

type UserUpdateRequest struct {
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

/*type UserNamaRequest struct {
	Nama     string `json:"nama" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}*/
