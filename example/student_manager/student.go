package student_manager

import "fmt"

type Student struct {
	Id int64
	Name string
}

type Manager struct {
	AllStudent map[int64]Student
}

// 查看所有学生
func (m Manager) show() {
	for _, val := range m.AllStudent {
		fmt.Printf("学号：%s  姓名： %s  \n", val.Id, val.Name)
	}
}

// 新增

func (m Manager) add(){
	var (
		stuID int64
		stuName string
	)
	fmt.Print("请输入学号：")
	fmt.Scanln(&stuID)

	fmt.Print("请输入姓名：")
	fmt.Scanln(&stuName)
	newStu := Student{
		Id:stuID,
		Name:stuName,
	}
	m.AllStudent[newStu.Id] = newStu
	fmt.Println("添加成功")
}

func (m Manager) edit(){
	var stuID int64
	fmt.Print("请输入要修改学生的学号：")
	fmt.Scanln(&stuID)
	stuObj,ok := m.AllStudent[stuID]
	if !ok {
		fmt.Println("不存在此学号")
		return
	}
	fmt.Printf("你要修改的学生信息是： 学号:%s  姓名：%s \n", stuObj.Id, stuObj.Name)
	fmt.Print("请输入需要修改的新姓名")
	var newName string
	fmt.Scanln(&newName)
	stuObj.Name = newName
	m.AllStudent[stuID] = stuObj
	fmt.Println("修改成功")
}

func (m Manager) del(){
	var stuID int64
	fmt.Print("请输入需要删除的学号：")
	fmt.Scanln(&stuID)
	_,ok := m.AllStudent[stuID]
	if !ok {
		fmt.Println("不存在此学号")
		return
	}
	delete(m.AllStudent, stuID)
}