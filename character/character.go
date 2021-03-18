package character

import "database/sql"

// Character is a Model of a character
type Character struct {
	// Id
	// sql.NullInt64 だと、ProductID == nil という比較でエラーになるので、*sql.NullInt64 にしている
	CharacterID *int   `json:"characterId"`
	UserID      string `json:"userID"`

	// Appearance
	CharacterType   string `json:"characterType"`
	CircleColorName string `json:"circleColorName"`
	SavedFileName   string `json:"savedFileName"`
	BlobImage       []byte `json:"blobImage"`
	IsPicUsed       bool   `json:"isPicUsed"`

	// Name
	NickName     string `json:"nickname"`
	FirstName    sql.NullString         `json:"firstName"`
	MiddleName   sql.NullString         `json:"middleName"`
	LastName     sql.NullString         `json:"lastName"`
	OriginalName sql.NullString         `json:"originalName"`

	// Details
	IsImaginary bool           `json:"isImaginary"`
	BirthDate   int            `json:"birthDate"`
	Description sql.NullString `json:"description"`
	Level       int            `json:"level"`
	IsLiked     bool           `json:"isLiked"`

	//Date
	RegisteredDate int `json:"registeredDate"`
	UpdatedDate    int `json:"updatedDate"`

	// Delete status
	OnDeleteLock bool `json:"onDeleteLock"`
	InLimbo      bool `json:"inLimbo"`
	// Reserved
	ReserveNumber01 int `json:"reserveNumber01"`
	ReserveNumber02 int `json:"reserveNumber02"`
	ReserveNumber03 int `json:"reserveNumber03"`
	ReserveNumber04 int `json:"reserveNumber04"`
	ReserveNumber05 int `json:"reserveNumber05"`
	ReserveNumber06 int `json:"reserveNumber06"`
	ReserveNumber07 int `json:"reserveNumber07"`
	ReserveNumber08 int `json:"reserveNumber08"`
	ReserveNumber09 int `json:"reserveNumber09"`
	ReserveNumber10 int `json:"reserveNumber10"`
	ReserveNumber11 int `json:"reserveNumber11"`
}
