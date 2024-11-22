package handler

import (
	"github.com/Xapsiel/EffectiveMobile"
	"github.com/gin-gonic/gin"
	"strconv"
)

const (
	TimeTemplate = "2006-02-02"
)

/*
		GetSongs(EffectiveMobile.Filter) ([]EffectiveMobile.Song, error)
		GetSongVerse(string, string) (string, error)
		DeleteSong(string, string) (bool, error)
		UpdateSong(EffectiveMobile.Song) (bool, error)
		Add(song EffectiveMobile.Song) (bool, error)
	}
*/

// @Summary Получение списка песен
// @Description Получение списка песен из базы данных
// @Tags songs
// @Accept json
// @Produce json
// @Param song query string true "Название песни"
// @Param group query string true "Группа"
// @Param id query int false "ID песни"
// @Param page query int false "Номер страницы"
// @Param since query string false "Дата с"
// @Param to query string false "Дата по"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /info [get]
func (h *Handler) GetSongs(c *gin.Context) {
	group := c.DefaultQuery("group", "")
	song := c.DefaultQuery("song", "")
	id, err := strconv.Atoi(c.DefaultQuery("id", "0"))
	if err != nil {
		c.AbortWithStatusJSON(404, gin.H{
			"status":   "fail",
			"errorMsg": "id должен быть числовым значением",
		})
		return
	}
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.AbortWithStatusJSON(404, gin.H{
			"status":   "fail",
			"errorMsg": "page должен быть числовым значением",
		})
		return
	}
	filter := EffectiveMobile.Filter{
		ID:    id,
		Song:  song,
		Group: group,
		Since: c.DefaultQuery("since", ""),
		To:    c.DefaultQuery("to", ""),
		Page:  page,
	}
	res, err := h.service.GetSongs(filter)
	if err != nil {
		c.AbortWithStatusJSON(404, gin.H{
			"status":   "fail",
			"errorMsg": err.Error(),
		})
		return
	}
	c.AbortWithStatusJSON(200, res)
}

// @Summary Добавление новой песни
// @Description Добавление новой песни в базу данных
// @Tags songs
// @Accept json
// @Produce json
// @Param song body EffectiveMobile.Song true "Данные песни"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /info [post]
func (h *Handler) AddSong(c *gin.Context) {
	var song EffectiveMobile.Song // Один объект Song
	if err := c.ShouldBindJSON(&song); err != nil {
		c.AbortWithStatusJSON(404, gin.H{
			"status":   "fail",
			"errorMsg": "Ошибка парсинга структуры",
		})
		return
	}
	ok, id, err := h.service.Add(song)
	if err != nil || !ok {
		c.AbortWithStatusJSON(404, gin.H{
			"status":   "fail",
			"errorMsg": err.Error(),
		})
		return
	}
	c.AbortWithStatusJSON(200, gin.H{
		"status": "success",
		"id":     id,
		"text":   "Песня добавлена",
	})
}

// @Summary Получение текста песни
// @Description Получение текста конкретного куплета песни.
// @Tags songs
// @Accept json
// @Produce json
// @Param song query string false "Название песни"
// @Param group query string false "Группа"
// @Param id query int false "ID песни"
// @Param verse query int true "Номер куплета. Передайте с ним либо ID, либо название и группу песни"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /info/verse [get]
func (h *Handler) GetSongVerse(c *gin.Context) {
	group := c.DefaultQuery("group", "")
	song_name := c.DefaultQuery("song", "")
	id, err := strconv.Atoi(c.DefaultQuery("id", "0"))
	if err != nil {
		c.AbortWithStatusJSON(404, gin.H{
			"status":   "fail",
			"errorMsg": "id должен быть числовым значением",
		})
		return
	}
	verseNumber, err := strconv.Atoi(c.DefaultQuery("verse", "1"))
	if err != nil {
		verseNumber = 1
		err = nil
	}
	song := EffectiveMobile.Song{SongName: song_name, Group: group, ID: id}
	verse, id, err := h.service.GetSongVerse(song, verseNumber)
	if err != nil {
		c.AbortWithStatusJSON(404, gin.H{
			"status":   "fail",
			"errorMsg": err.Error(),
		})
		return
	}
	c.AbortWithStatusJSON(200, gin.H{
		"status": "success",
		"id":     id,
		"text":   verse,
	})
}

// @Summary Удаление песни
// @Description Удаление песни по предоставленным данным
// @Tags songs
// @Accept json
// @Produce json
// @Param song body EffectiveMobile.Song true "Данные песни для удаления. Передайте либо ID, либо название и группу песни"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /info [delete]
func (h *Handler) DeleteSong(c *gin.Context) {
	var song EffectiveMobile.Song
	if err := c.ShouldBindJSON(&song); err != nil {
		c.AbortWithStatusJSON(404, gin.H{
			"status":   "fail",
			"errorMsg": "Ошибка парсинга структуры",
		})
		return
	}
	ok, err := h.service.DeleteSong(song)
	if err != nil || !ok {
		c.AbortWithStatusJSON(404, gin.H{
			"status":   "fail",
			"errorMsg": err.Error(),
		})
		return
	}
	c.AbortWithStatusJSON(200, gin.H{
		"status": "success",
		"text":   "Удаление прошло успешно",
	})
}

// @Summary Обновление информации о песне
// @Description Обновление информации о песне по предоставленным данным
// @Tags songs
// @Accept json
// @Produce json
// @Param song body EffectiveMobile.Song true "Данные песни для обновления"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /info [put]
func (h *Handler) UpdateSong(c *gin.Context) {
	var song EffectiveMobile.Song // Один объект Song
	if err := c.ShouldBindJSON(&song); err != nil {
		c.AbortWithStatusJSON(404, gin.H{
			"status":   "fail",
			"errorMsg": "Ошибка парсинга структуры",
		})
		return
	}
	ok, song, err := h.service.UpdateSong(song)
	if err != nil || !ok {
		c.AbortWithStatusJSON(404, gin.H{
			"status":   "fail",
			"errorMsg": err.Error(),
		})
		return
	}
	c.AbortWithStatusJSON(200, gin.H{
		"status": "success",
		"text":   "Обновление прошло успешно",
		"song":   song,
	})
}
