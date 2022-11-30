package v1

type Contract struct {
	Id           int    `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	ChainId      uint32 `json:"chain_id" gorm:"type:int(20)"` //外键
	ContractAddr string `json:"contract_addr" gorm:"type:int(20)"`
	Symbol       string `json:"symbol"`
	Name         string `json:"name"`
	Decimals     string `json:"decimals"`
	Type         string `json:"type"`
}

func (Contract) TableName() string {
	return "contract"
}
