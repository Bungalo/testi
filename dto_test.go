package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllEmpty(t *testing.T) {
	dto := NewDTO(data{})
	expected := data{}
	data := dto.getAll()

	assert.Equal(t, expected, data)
}

func TestGetAllWithData(t *testing.T) {
	dto := NewDTO(data{
		"some": "value",
	})

	expected := data{
		"some": "value",
	}

	data := dto.getAll()

	assert.Equal(t, expected, data)
}

func TestGetOne(t *testing.T) {
	dto := NewDTO(data{
		"some": "value",
	})

	expected := data{
		"some": "value",
	}

	data := dto.getOne("some")

	assert.Equal(t, expected, data)
}

func TestInsert(t *testing.T) {
	dto := NewDTO(data{})

	expected := data{
		"new": "value",
	}

	data := dto.insert("new", "value")

	assert.Equal(t, expected, data)
}

func TestDelete(t *testing.T) {
	dto := NewDTO(data{
		"some": "value",
	})

	expected := data{}

	data := dto.delete("some")

	assert.Equal(t, expected, data)
}

func TestDeleteFail(t *testing.T) {
	dto := NewDTO(data{
		"some": "value",
	})

	expected := data{
		"some": "value",
	}

	data := dto.delete("wrongkey")

	assert.Equal(t, expected, data)
}