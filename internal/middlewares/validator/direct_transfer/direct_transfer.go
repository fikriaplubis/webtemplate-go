package direct_transfer

import (
	"fmt"
	"mime/multipart"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"webtemplate/internal/models/dto"
)

type Validator struct {
	request *dto.DirectTransferRequest
}

func NewValidator(request *dto.DirectTransferRequest) *Validator {
	return &Validator{request: request}
}

func (v *Validator) RuleFilenameValidation(header *multipart.FileHeader) string {
	var message string

	txType := header.Filename[4:6]
	if txType != "PR" && txType != "DD" && txType != "DC" {
		message = "Invalid filename, missing transactiton type identifier"
	}

	return message
}

func (v *Validator) SpecialCharacterAmountValidation(c *gin.Context) string {
	var message string

	err := c.ShouldBindWith(&v.request, binding.FormPost)
	if err != nil {
		message = err.Error() + "- Amount Validation"
	}

	regex, _ := regexp.Compile(`^[0-9]\d*[\.\,]\d*?$`)
	amountStr := fmt.Sprintf("%f", v.request.Amount)
	clearAmount := strings.ReplaceAll(amountStr, ",", "")

	if !regex.MatchString(clearAmount) {
		message = "Invalid Amount"
	}

	return message
}

func (v *Validator) AccountNumberValidation(c *gin.Context) string {
	var message string

	err := c.ShouldBindWith(&v.request, binding.FormPost)
	if err != nil {
		message = err.Error() + "- Account Number Validation"
	}

	regex, _ := regexp.Compile(`^[A-Za-z0-9._]*[\.\,]*?$`)
	clearAccountNo := strings.ReplaceAll(v.request.AccountNo, ",", "")

	if !regex.MatchString(clearAccountNo) {
		message = "Invalid Account Number"
	}

	return message
}

// func (v *Validator) SpecialCharacterChargeValidation(c *gin.Context) string

// func (v *Validator) MaximumChargeValidation(c *gin.Context) string

func (v *Validator) FormatDescriptionValidation(c *gin.Context) string {
	var message string

	err := c.ShouldBindWith(&v.request, binding.FormPost)
	if err != nil {
		message = err.Error() + "- Description Validation"
	}

	regex, _ := regexp.Compile("^[A-Za-z0-9-_ ]*$")

	if !regex.MatchString(v.request.Description) {
		message = "Invalid Description"
	}

	return message
}
