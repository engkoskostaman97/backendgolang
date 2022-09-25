package handlers

import (
	filmdto "dumbflix/dto/film"
	dto "dumbflix/dto/result"
	"dumbflix/models"
	"dumbflix/repositories"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"context"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"

	"github.com/go-playground/validator/v10"

	"github.com/gorilla/mux"
)

type handlerFilm struct {
	FilmRepository repositories.FilmRepository
}

// Create `path_file` Global variable here ...
// var PathFile = os.Getenv("PATH_FILE")

func HandlerFilm(FilmRepository repositories.FilmRepository) *handlerFilm {
	return &handlerFilm{FilmRepository}
}

func (h *handlerFilm) FindFilms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	films, err := h.FilmRepository.FindFilms()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	// for i, p := range films {
	// 	films[i].ThumbnailFilm = os.Getenv("PATH_FILE") + p.ThumbnailFilm
	// }
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: films}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerFilm) GetFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var film models.Film
	film, err := h.FilmRepository.GetFilm(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// film.ThumbnailFilm = os.Getenv("PATH_FILE") + film.ThumbnailFilm

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseFilm(film)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerFilm) CreateFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// get data user token

	// userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	// userId := int(userInfo["id"].(float64))

	// fmt.Println(userId)

	// dataContex := r.Context().Value("dataFile") // add this code
	// filename := dataContex.(string)             // add this code
	// get image filename
	// get image filepath
	dataContex := r.Context().Value("dataFile")
	filepath := dataContex.(string)

	category_id, _ := strconv.Atoi(r.FormValue("category_id"))
	request := filmdto.FilmRequest{
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
		Year:        r.FormValue("year"),
		CategoryID:  category_id,
		LinkFilm:    r.FormValue("linkfilm"),
	}

	validation := validator.New()

	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	// Add your Cloudinary credentials ...
	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	// Upload file to Cloudinary ...
	resp, err := cld.Upload.Upload(ctx, filepath, uploader.UploadParams{Folder: "dumbflix"})
	if err != nil {
		fmt.Println(err.Error())
	}
	film := models.Film{
		Title:         request.Title,
		Description:   request.Description,
		Year:          request.Year,
		ThumbnailFilm: resp.SecureURL, // Modify store file URL to database from resp.SecureURL ...,
		CategoryID:    request.CategoryID,
		LinkFilm:      request.LinkFilm,
	}

	film, err = h.FilmRepository.CreateFilm(film)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	film, _ = h.FilmRepository.GetFilm(film.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: film}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerFilm) DeleteFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	film, err := h.FilmRepository.GetFilm(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.FilmRepository.DeleteFilm(film)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseFilm(data)}
	json.NewEncoder(w).Encode(response)
}

func convertResponseFilm(u models.Film) models.FilmResponse {
	return models.FilmResponse{
		ID:            u.ID,
		CategoryID:    u.CategoryID,
		Title:         u.Title,
		Description:   u.Description,
		Year:          u.Year,
		ThumbnailFilm: u.ThumbnailFilm,
		Category:      u.Category,
		LinkFilm:      u.LinkFilm,
	}
}
