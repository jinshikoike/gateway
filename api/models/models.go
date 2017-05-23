package models

// Gateway table
type Gateway struct {
	//gorm.Model
	ID        int64
	Path      string
	HostID    int64   `gorm:"column:host_id"`
	Host      Host    `gorm:"ForeignKey:HostID;AssociationForeignKey:Refer"`
	ClaptonID int64   `gorm:"column:clapton_id"`
	Clapton   Clapton `gorm:"ForeignKey:ClaptonID;AssociationForeignKey:Refer"`
}

// TableName return table name for Gateway
func (table Gateway) TableName() string {
	return "gateways"
}

// Host table
type Host struct {
	//gorm.Model
	ID   int
	Name string
}

// TableName return table name for Host
func (table Host) TableName() string {
	return "hosts"
}

// Clapton table
type Clapton struct {
	ID   int
	UUID string `gorm:"column:uuid"`
}

// TableName return table name for Clapton
func (table Clapton) TableName() string {
	return "claptons"
}
