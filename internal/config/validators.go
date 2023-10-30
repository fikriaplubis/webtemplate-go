package config

type fileValidation struct {
	direct_transfer []string
}

type dataValidation struct {
	direct_transfer []string
}

func (s *server) ImplementValidators() (*fileValidation, *dataValidation) {
	implementFile := &fileValidation{
		direct_transfer: []string{"extension", "content-type", "binary-file", "filename", "filename-DT"},
	}

	implementData := &dataValidation{
		direct_transfer: []string{"amount-DT", "accountnumber-DT", "description-DT"},
	}

	return implementFile, implementData
}
