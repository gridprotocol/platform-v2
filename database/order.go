package database

import "time"

type Order struct {
	User         string
	Provider     string
	Id           int
	ActivateTime time.Time `gorm:"column:activate"`
	StartTime    time.Time `gorm:"column:start"`
	EndTime      time.Time `gorm:"column:end"`
	Probation    int64
	Duration     int64
}

func InitOrder() error {
	return GlobalDataBase.AutoMigrate(&Order{})
}

// store order info to db
func (o *Order) CreateOrder() error {
	o.StartTime = o.ActivateTime.Add(time.Duration(o.Probation) * time.Second)
	o.EndTime = o.StartTime.Add(time.Duration(o.Duration) * time.Second)
	return GlobalDataBase.Create(o).Error
}

// get order by order id
func GetOrderById(id uint64) (Order, error) {
	var order Order
	err := GlobalDataBase.Model(&Order{}).Where("id = ?", id).Last(&order).Error
	if err != nil {
		return Order{}, err
	}

	return order, nil
}

// get order list of an user
func GetOrdersByUser(user string) ([]Order, error) {
	var orders []Order
	err := GlobalDataBase.Model(&Order{}).Where("user = ?", user).Find(&orders).Error
	if err != nil {
		return nil, err
	}

	return orders, nil
}

// get all provider address of an user
func GetProsByUser(user string) ([]string, error) {
	var pros []string
	err := GlobalDataBase.Model(&Order{}).Select("provider").Where("user = ?", user).Find(&pros).Error
	if err != nil {
		return nil, err
	}

	return pros, nil
}

func ListAllActivedOrder() ([]Order, error) {
	var now = time.Now()
	var orders []Order
	err := GlobalDataBase.Model(&Order{}).Where("start < ? AND end > ?", now, now).Find(&orders).Error
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func ListAllActivedOrderByUser(address string) ([]Order, error) {
	var now = time.Now()
	var orders []Order
	err := GlobalDataBase.Model(&Order{}).Where("user = ? AND start < ? AND end > ?", address, now, now).Find(&orders).Error
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func ListAllOrderedProvider(address string) ([]Provider, error) {
	var now = time.Now()
	var provider []Provider
	err := GlobalDataBase.Model(&Order{}).Where("user = ? AND start < ? AND end > ?", address, now, now).
		Joins("right join provider on order.provider = provider.addresss").Find(&provider).Error
	if err != nil {
		return nil, err
	}

	return provider, nil
}
