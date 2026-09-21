package dto

type LoginRequestDto struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=4"`
}

type LoginResponseDto struct {
	Token      string `json:"token"`
	ExpireDate string `json:"expire_date"`
}

type LoginTempResponseDto struct {
	Token                     string         `json:"token"`
	ExpireDate                string         `json:"expire_date"`
	RequireWarehouseSelection bool           `json:"require_warehouse_selection"`
	Warehouses                []*UserRoleDto `json:"warehouses"`
}

type RegisterRequestDto struct {
	FullName string `json:"full_name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=4"`
	Email    string `json:"email" validate:"required,email"`
}

type SelectWarehouseRequestDto struct {
	WarehouseId string `json:"warehouse_id" validate:"required"`
}

type SelectWarehouseResponseDto struct {
	Token      string `json:"token"`
	ExpireDate string `json:"expire_date"`
}
