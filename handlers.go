package main

import (
	"encoding/json"
	"os"

	"github.com/labstack/echo/v4"
)

type Data struct {
	Key string `json:"key"`
	Value string `json:"value"`
}

type Handler struct {
	dto *DTO
}

func NewHandler(dto *DTO) *Handler {
	return &Handler{
		dto: dto,
	}
}

func (h *Handler) get(c echo.Context) error {
	key := c.Param("key") 
	data := h.dto.getOne(key)
	if data != nil {
		return c.JSON(200, data)
	}
	return c.JSON(404, data)
}

func (h *Handler) list(c echo.Context) error {
	return c.JSON(200, h.dto.getAll())
}

func (h *Handler) post(c echo.Context) error {
	var data Data
	json.NewDecoder(c.Request().Body).Decode(&data) 

	if data.Key == "" {
		return c.JSON(400, "unable to accept empty value/ key")
	}

	return c.JSON(201, h.dto.insert(data.Key, data.Value))
}

func (h *Handler) delete(c echo.Context) error {
	key := c.Param("key") 
	return c.JSON(200, h.dto.delete(key))
}

func (h *Handler) version(c echo.Context) error {
	return c.JSON(200, os.Getenv("version"))
}