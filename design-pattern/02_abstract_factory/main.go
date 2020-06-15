package main

import "fmt"

type OrderMainDAO interface {
	SaveOrderMain()
}

type OrderDetailDAO interface {
	SaveOrderDetail()
}
//DAOFactory DAO 抽象模式工厂接口
type DAOFactory interface {
	CreateOrderMainDAO() OrderMainDAO
	CreateOrderDetailDAO() OrderDetailDAO
}

type RDBMainDAO struct {

}

func (*RDBMainDAO) SaveOrderMain()  {
	fmt.Println("rdb main save")
}

type RDBDetailDAO struct {

}

func (*RDBDetailDAO) SaveOrderDetail()  {
	fmt.Println("rdb detail save")
}
//RDBDAOFactory 是RDB 抽象工厂实现
type RDBDAOFactory struct {

}

func (*RDBDAOFactory) CreateOrderMainDAO() OrderMainDAO {
	return &RDBMainDAO{}
}

func (*RDBDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &RDBDetailDAO{}
}

type XMLMainDAO struct {

}

func (*XMLMainDAO) SaveOrderMain()  {
	fmt.Println("xml main save")
}

type XMLDetailDAO struct {

}

func (*XMLDetailDAO) SaveOrderDetail()  {
	fmt.Println("xml detail save")
}
//XMLDAOFactory 是XML 抽象工厂实现
type XMLDAOFactory struct {

}
func (*XMLDAOFactory) CreateOrderMainDAO() OrderMainDAO {
	return &XMLMainDAO{}
}

func (*XMLDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &XMLDetailDAO{}
}

func main() {
	
}
