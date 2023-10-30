package dto

func RequestMapping() map[string]interface{} {
	type mapping map[string]interface{}

	dto := mapping{
		"direct_transfer": DirectTransferRequest.Get,
	}

	return dto
}
