package models

type Tek struct {
	ID uint
}

type Cok struct {
	ID  uint
	Tid uint
	Tek Tek `gorm:"foreignKey:Tid"`
}

type Cargos struct {
	ID        int32 `gorm:"primaryKey"`
	Wcid      int32
	WorkCargo WorkCargo `gorm:"foreignKey:Wcid"`
}

type CV struct {
	ID int32
}

type Vehicles struct {
	ID     int32 `gorm:"primaryKey"`
	Cid    int32
	Cargos Cargos `gorm:"foreignKey:Cid"`
}

type Packages struct {
	ID     int32 `gorm:"primaryKey"`
	Cid    int32
	Cargos Cargos `gorm:"foreignKey:Cid"`
}
type WorkCargo struct {
	ID int32 `gorm:"primaryKey"`
}
type Workmans struct {
	ID        uint `gorm:"primaryKey"`
	Wcid      int32
	WorkCargo WorkCargo `gorm:"foreignKey:Wcid"`
}

type Admins struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required" gorm:"unique"`
	Password string `json:"password" binding:"required"`
}
