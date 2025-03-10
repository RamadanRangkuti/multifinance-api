package helper

import (
	"fmt"

	"github.com/go-playground/validator"
)

// Format error validasi menjadi map yang lebih deskriptif
func FormatValidationError(err error) map[string]string {
	errors := make(map[string]string)

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			field := e.Field() // Nama field yang error
			tag := e.Tag()     // Jenis validasi yang gagal
			param := e.Param() // Parameter validasi (misalnya min/max value)

			// Custom pesan error berdasarkan jenis validasi
			switch tag {
			case "required":
				errors[field] = fmt.Sprintf("%s harus diisi", field)
			case "len":
				errors[field] = fmt.Sprintf("%s harus memiliki panjang %s karakter", field, param)
			case "min":
				errors[field] = fmt.Sprintf("%s minimal harus %s karakter", field, param)
			case "max":
				errors[field] = fmt.Sprintf("%s maksimal %s karakter", field, param)
			case "numeric":
				errors[field] = fmt.Sprintf("%s harus berupa angka", field)
			case "oneof":
				errors[field] = fmt.Sprintf("%s harus salah satu dari %s", field, param)
			case "gt":
				errors[field] = fmt.Sprintf("%s harus lebih besar dari %s", field, param)
			default:
				errors[field] = fmt.Sprintf("%s tidak valid", field)
			}
		}
	}

	return errors
}
