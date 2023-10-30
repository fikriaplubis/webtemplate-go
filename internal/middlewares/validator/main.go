package validator

import (
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"

	"webtemplate/internal/middlewares/validator/direct_transfer"
	"webtemplate/internal/models/dto"
)

type registeredValidators struct {
	general         *Validator
	direct_transfer *direct_transfer.Validator
}

func setupValidators() *registeredValidators {
	return &registeredValidators{
		general:         NewValidator(),
		direct_transfer: direct_transfer.NewValidator(&dto.DirectTransferRequest{}),
	}
}

func initialize() map[string]interface{} {
	validators := setupValidators()
	type mapping map[string]interface{}

	var storage = mapping{
		// General
		"extension":    validators.general.FileExtValidation,
		"content-type": validators.general.ContentTypeValidation,
		"binary-file":  validators.general.BinaryFileValidation,
		"filename":     validators.general.SpecialCharacterFileNameValidation,

		// Direct Transfer
		"filename-DT":      validators.direct_transfer.RuleFilenameValidation,
		"amount-DT":        validators.direct_transfer.SpecialCharacterAmountValidation,
		"accountnumber-DT": validators.direct_transfer.AccountNumberValidation,
		"description-DT":   validators.direct_transfer.FormatDescriptionValidation,
	}

	return storage
}

func File(names []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var errorMessage []string
		storage := initialize()
		_, header, err := c.Request.FormFile("file")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Internal Server Error",
				"message": err.Error(),
			})
			c.Abort()
			return
		}

		for _, name := range names {
			f := reflect.ValueOf(storage[name])
			in := make([]reflect.Value, 1)
			in[0] = reflect.ValueOf(header)

			res := f.Call(in)
			result := res[0].Interface()
			resultValue := result.(string)

			if len(resultValue) > 0 {
				errorMessage = append(errorMessage, resultValue)
			}
		}

		if len(errorMessage) > 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad request",
				"message": strings.Join(errorMessage[:], ", "),
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func Data(names []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var errorMessage []string
		storage := initialize()

		for _, name := range names {
			f := reflect.ValueOf(storage[name])
			in := make([]reflect.Value, 1)
			in[0] = reflect.ValueOf(c)

			res := f.Call(in)
			result := res[0].Interface()
			resultValue := result.(string)

			if len(resultValue) > 0 {
				errorMessage = append(errorMessage, resultValue)
			}
		}

		if len(errorMessage) > 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad request",
				"message": strings.Join(errorMessage[:], ", "),
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
