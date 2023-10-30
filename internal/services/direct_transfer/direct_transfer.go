package direct_transfer

import (
	"net/http"

	"gorm.io/gorm"

	"github.com/devfeel/mapper"

	// "webtemplate/internal/constant"
	"webtemplate/internal/models/dao"
	"webtemplate/internal/models/dto"

	// "webtemplate/internal/services/file"
	"webtemplate/pkg/generate/direct_transfer"
)

type Service interface {
	Create(request dto.DirectTransferRequest) dto.DirectTransferResponseWrite
}

type service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *service {
	return &service{
		db: db.Session(&gorm.Session{PrepareStmt: true}),
	}
}

func (s *service) Create(request dto.DirectTransferRequest) dto.DirectTransferResponseWrite {
	// fileServices := file.NewService(s.db)
	var newDirectTransfer dao.DirectTransfer

	if err := mapper.AutoMapper(&request, &newDirectTransfer); err != nil {
		return request.ErrorWrite(http.StatusInternalServerError, "Error Mapper", err.Error())
	}

	// if err, message := fileServices.Save(*request.File, constant.GetPath("DirectTransferOri"), request.File.Filename); err {
	// 	return request.ErrorWrite(http.StatusInternalServerError, "Error Saving File", message)
	// }

	// var decrypted decrypted
	// decrypted = fileServices.Decrypt(*request.File)
	// trx_path, trx_type := decrypted.t24OrSunline()

	// Set Others Value
	no_reff := direct_transfer.NewNoReff()
	newDirectTransfer.SetNoReff(no_reff)

	if err := s.db.Create(newDirectTransfer).Error; err != nil {
		return request.ErrorWrite(http.StatusInternalServerError, "Error when trying to upload data in the direct_transfer table", err.Error())
	}

	return request.SuccessWrite("Success upload new direct transfer", &newDirectTransfer)
}
