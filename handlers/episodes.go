package handlers

import (
	episodesdto "dumbflix/dto/episodes"
	dto "dumbflix/dto/result"
	"dumbflix/models"
	"dumbflix/repositories"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerEpisode struct {
	EpisodeRepository repositories.EpisodeRepository
}

func HandlerEpisode(EpisodeRepository repositories.EpisodeRepository) *handlerEpisode {
	return &handlerEpisode{EpisodeRepository}
}

func (h *handlerEpisode) FindEpisodes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	episodes, err := h.EpisodeRepository.FindEpisodes()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: episodes}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerEpisode) GetEpisode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var episode models.Episode
	episode, err := h.EpisodeRepository.GetEpisode(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseEpisode(episode)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerEpisode) CreateEpisode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// get data user token
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	fmt.Println(userId)

	dataContex := r.Context().Value("dataFile") // add this code
	filename := dataContex.(string)             // add this code
	film_id, _ := strconv.Atoi(r.FormValue("film_id"))
	request := episodesdto.EpisodeRequest{
		Title:            r.FormValue("title"),
		ThumbnailEpisode: filename,
		LinkFilm:         r.FormValue("linkfilm"),
		FilmID:           film_id,
		Description: r.FormValue("description"),
	}
	// request := new(episodesdto.CreateEpisodeRequest)
	// if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	// 	json.NewEncoder(w).Encode(response)
	// 	return
	// }

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	episode := models.Episode{
		Title:            request.Title,
		ThumbnailEpisode: filename,
		LinkFilm:         request.LinkFilm,
		FilmID:           request.FilmID,
		Description:      request.Description,
	}

	episode, err = h.EpisodeRepository.CreateEpisode(episode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	episode, _ = h.EpisodeRepository.GetEpisode(episode.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: episode}
	json.NewEncoder(w).Encode(response)

}

// func (h *handlerEpisode) UpdateEpisode(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	request := new(episodesdto.UpdateEpisodeRequest)
// 	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	id, _ := strconv.Atoi(mux.Vars(r)["id"])
// 	episode, err := h.EpisodeRepository.GetEpisode(int(id))
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	if request.Title != "" {
// 		episode.Title = request.Title
// 	}
// 	if request.ThumbnailEpisode != "" {
// 		episode.ThumbnailEpisode = request.ThumbnailEpisode
// 	}
// 	if request.LinkFilm != "" {
// 		episode.LinkFilm = request.LinkFilm
// 	}
// 	if request.FilmID != {
// 		episode.FilmID = request.FilmID
// 	}
// 	if request.Description != "" {
// 		episode.Description = request.Description
// 	}

// 	data, err := h.EpisodeRepository.UpdateEpisode(episode)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	response := dto.SuccessResult{Code: http.StatusOK, Data: convertEpisodeResponse(episode)}
// 	json.NewEncoder(w).Encode(response)
// }

func (h *handlerEpisode) DeleteEpisode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	episode, err := h.EpisodeRepository.GetEpisode(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.EpisodeRepository.DeleteEpisode(episode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseEpisode(data)}
	json.NewEncoder(w).Encode(response)
}

func convertResponseEpisode(u models.Episode) models.EpisodeResponse {
	return models.EpisodeResponse{
		ID:               u.ID,
		Title:            u.Title,
		ThumbnailEpisode: u.ThumbnailEpisode,
		LinkFilm:         u.LinkFilm,
		Description:      u.Description,
		FilmID:           u.FilmID,
		Film:             u.Film,
	}
}
