package main

type data map[string]string

type DTO struct {
	Data data
}

func NewDTO(data data) *DTO {
	return &DTO{
		Data: data,
	}
}

func (d *DTO) getOne(key string) data {
	v, ok := d.Data[key]
	if ok {
		return data{
			key: v,
		}
	}
	return nil
}

func (d *DTO) getAll() data {
	return d.Data
}

func (d *DTO) insert(key, value string) data {
	d.Data[key] = value
	return d.Data
}

func (d *DTO) delete(key string) data {
	delete(d.Data, key)
	return d.Data
}