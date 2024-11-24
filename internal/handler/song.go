package handler

import (
	"fmt"
	"github.com/Xapsiel/EffectiveMobile/internal/models"
	. "github.com/Xapsiel/EffectiveMobile/pkg/log"
	"github.com/gin-gonic/gin"

	"net/http"
	"strconv"
)

// @Summary		Получение списка песен
// @Tags			songs
// @Description	Получение списка песен из базы данных с фильтрацией по параметрам
// @Accept			json
// @Produce		json
// @Param			song	query		string	false	"Название песни"	default(Supermassive Black Hole)
// @Param			group	query		string	false	"Группа"			default(Muse)
// @Param			id		query		int		false	"ID песни"
// @Param			page	query		int		false	"Номер страницы"
// @Param			since	query		string	false	"Дата с" example(2006-07-19)
// @Param			to		query		string	false	"Дата по" example(2006-07-19)
// @Success		200		{object}	models.Song
// @Failure		400		{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Router			/info [get]
func (h *Handler) GetSongs(c *gin.Context) {
	group := c.DefaultQuery("group", "")
	song := c.DefaultQuery("song", "")
	id, err := strconv.Atoi(c.DefaultQuery("id", "0"))
	if err != nil {
		newErrorResponce(c, http.StatusUnprocessableEntity, "id должен быть числовым значением")
		return
	}
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		newErrorResponce(c, http.StatusUnprocessableEntity, "page должен быть числовым значением")
		return
	}
	filter := models.Filter{
		ID:    id,
		Song:  song,
		Group: group,
		Since: c.DefaultQuery("since", ""),
		To:    c.DefaultQuery("to", ""),
		Page:  page,
	}
	res, err := h.service.GetSongs(filter)
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}
	if len(res) == 0 {
		res = make([]models.Song, 0)
		c.AbortWithStatusJSON(200, res)
		return
	}
	Logger.Debug(fmt.Sprintf("Получен список песен по фильтру:\n\t%+v", filter))

	c.AbortWithStatusJSON(200, res)
}

// @Summary		Добавление новой песни
// @Description	Добавление новой песни в базу данных(Обязательные параметры - song,group)
// @Tags			songs
// @Accept			json
// @Produce		json
// @Param			song	body		models.Song	true	"Данные песни"	default({ "group": "Muse", "song_name": "Supermassive Black Hole", "release_date": "16-07-2006", "link": "https://www.youtube.com/watch?v=Xsp3_a-PMTw", "text": "Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?" })
// @Success		200		{object}	resultResponse
// @Failure		400		{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Router			/songs [post]
func (h *Handler) AddSong(c *gin.Context) {
	var song models.Song
	if err := c.ShouldBindJSON(&song); err != nil {
		newErrorResponce(c, http.StatusUnprocessableEntity, "Ошибка парсинга структуры")
		return
	}
	ok, id, err := h.service.Add(song)
	if err != nil || !ok {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}
	Logger.Debug(fmt.Sprintf("Добавлена песня песня:\n\t%+v", song))

	c.AbortWithStatusJSON(200, resultResponse{
		Status: "success",
		Id:     id,
		Text:   "Песня добавлена",
	})
}

// @Summary		Получение текста куплета песни
// @Description	Получение текста конкретного куплета песни(Обязательные параметры - song,group или id)
// @Tags			songs
// @Accept			json
// @Produce		json
// @Param			song	query		string	false	"Название песни"	default(Supermassive Black Hole)
// @Param			group	query		string	false	"Группа"			default(Muse)
// @Param			id		query		int		false	"ID песни"			default(1)
// @Param			verse	query		int		false	"Номер куплета"		example(1)
//
// @Success		200		{object}	resultResponse
//
// @Failure		400		{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Router			/info/verse [get]
func (h *Handler) GetSongVerse(c *gin.Context) {
	group := c.DefaultQuery("group", "")
	song_name := c.DefaultQuery("song", "")
	id, err := strconv.Atoi(c.DefaultQuery("id", "0"))
	if err != nil {
		newErrorResponce(c, http.StatusUnprocessableEntity, "id должен быть числовым значением")

		return
	}
	verseNumber, err := strconv.Atoi(c.DefaultQuery("verse", "1"))
	if err != nil {
		verseNumber = 1
		err = nil
	}
	song := models.Song{SongName: song_name, Group: group, ID: id}
	verse, id, err := h.service.GetSongVerse(song, verseNumber)
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}
	c.AbortWithStatusJSON(200, resultResponse{
		Status: "success",
		Id:     id,
		Text:   verse,
	})
	Logger.Debug(fmt.Sprintf("Получен куплет песни %d:\n\t", id))

}

// @Summary		Удаление песни
// @Description	Удаление песни по предоставленным данным(Обязательные параметры- song,group или id)
// @Tags			songs
// @Accept			json
// @Produce		json
// @Param			song	body		models.Song	true	"Данные песни для удаления"
// @Success		200		{object}	resultResponse
// @Failure		400		{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Router			/songs [delete]
func (h *Handler) DeleteSong(c *gin.Context) {
	var song models.Song
	if err := c.ShouldBindJSON(&song); err != nil {
		newErrorResponce(c, http.StatusUnprocessableEntity, "Ошибка парсинга структуры")

		return
	}
	ok, err := h.service.DeleteSong(song)
	if err != nil || !ok {
		newErrorResponce(c, http.StatusBadRequest, err.Error())

		return
	}
	Logger.Debug(fmt.Sprintf("Удалена песня \"%s\"-%s с id=%d", song.SongName, song.Text, song.ID))
	c.AbortWithStatusJSON(200, resultResponse{
		Status: "success",
		Text:   "Удаление прошло успешно",
	})
}

// @Summary		Обновление информации о песне
// @Description	Обновление информации о песне по предоставленным данным(Обязательные параметры- song,group или id)
// @Tags			songs
// @Accept			json
// @Produce		json
// @Param			song	body		models.Song	true	"Данные песни для обновления"
// @Success		200		{object}	resultResponse
// @Failure		400		{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Router			/songs [put]
func (h *Handler) UpdateSong(c *gin.Context) {
	var song models.Song
	if err := c.ShouldBindJSON(&song); err != nil {
		newErrorResponce(c, http.StatusUnprocessableEntity, "Ошибка парсинга структуры")

		return
	}
	ok, song, err := h.service.UpdateSong(song)
	if err != nil || !ok {
		newErrorResponce(c, http.StatusBadRequest, err.Error())

		return
	}
	Logger.Debug(fmt.Sprintf("Обновлена информация о песня \"%s\"-%s с id=%d", song.SongName, song.Group, song.ID))
	c.AbortWithStatusJSON(200, resultResponse{
		Status: "success",
		Text:   "Обновление прошло успешно",
		Id:     song.ID,
	})
}
