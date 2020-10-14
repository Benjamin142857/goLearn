package main

import "fmt"

/*
	学生管理系统 1.0
	支持功能：
		1. 查看学生列表
		2. 增加学生
		3. 删除学生
*/

// School 学校类型
type School struct {
	name     string
	capacity int
	students []*Student
}

// School 构造函数
func newSchool(name string, capacity int) School {
	return School{
		name:     name,
		capacity: capacity,
		students: make([]*Student, 0, capacity),
	}
}

// School 增加学生
func (s *School) addStudent(pStudent *Student) {
	if len(s.students) < s.capacity {
		s.students = append(s.students, pStudent)
		fmt.Printf("添加成功, 已成功添加学生 %v, 学号为 %v.\n", pStudent.name, pStudent.id)
	} else {
		fmt.Printf("添加失败, 该学校学生已满员 %v 人.\n", s.capacity)
	}
}

// School 删除学生
func (s *School) delStudent(id int) {
	for idx, pStd := range s.students {
		if pStd.id == id {
			s.students = append(s.students[:idx], s.students[idx+1:]...)
			fmt.Printf("删除成功, 已将id为 %v 的学生 '%v' 从学校中剔除学籍.", pStd.id, pStd.name)
			goto END
		}
	}
	fmt.Printf("删除失败, 未找到id为 %v 的学生.\n", id)
END:
}

// Student 学生类型
type Student struct {
	id   int
	name string
	age  int
}

// Student 构造函数
func newStudent(idHWM *int, name string, age int) (std *Student) {
	std = &Student{
		id:   *idHWM,
		name: name,
		age:  age,
	}
	*idHWM++
	return
}

// 标准输入捕获 int
func inputInt(input *int, text string) {
	for {
		fmt.Printf("\n%v", text)
		_, err := fmt.Scanf("%d", input)
		if err == nil {
			break
		} else {
			fmt.Println("无效输入, 请重新按要求输入...")
		}
	}
}

// 标准输入捕获 str
func inputStr(input *string, text string) {
	for {
		fmt.Printf("\n%v", text)
		_, err := fmt.Scanf("%s", input)
		if err == nil {
			break
		} else {
			fmt.Println("无效输入, 请重新按要求输入...")
		}
	}
}

// cmdSet - help
func cmdHelp() {
	fmt.Println("\n  系统指令集:")
	for k, v := range cmdSet {
		fmt.Printf("    %v: \t %v \n", k, v.description)
	}
	fmt.Println("")
}

// cmdSet - exit
func cmdExit() {
	sysExit = true
}

// cmdSet - show
func cmdShow() {
	fmt.Printf("  当前学校人数: %v, 学校总容量: %v:\n\n", len(mySchool.students), mySchool.capacity)
	fmt.Printf("  学号\t\t姓名\t\t年龄\n")
	for _, pStd := range mySchool.students {
		fmt.Printf("  %v\t\t%v\t\t%v\n", pStd.id, pStd.name, pStd.age)
	}
	fmt.Println("")
}

// cmdSet - add
func cmdAdd() {
	name := ""
	age := 0
	fmt.Printf("添加学生(输入 '-1' 中断)>>>>>>>>>>>>>>>>>>>>>\n")
	inputStr(&name, "[str] 请输入学生姓名: ")
	if name == "-1" {
		return
	}
	inputInt(&age, "[int] 请输入学生年龄: ")
	if age == -1 {
		return
	}
	mySchool.addStudent(newStudent(&idHWM, name, age))
}

// cmdSet - del
func cmdDel() {
	id := 0
	fmt.Printf("删除学生(输入 '-1' 中断)>>>>>>>>>>>>>>>>>>>>>\n")
	inputInt(&id, "[int] 请输入学生ID: ")
	if id == -1 {
		return
	}
	mySchool.delStudent(id)
}

// cmdSet 命令集注册
var cmdSet = map[string]struct {
	description string
	exec        func()
}{
	"exit": {"退出系统", cmdExit},
	"show": {"查看学生列表", cmdShow},
	"add":  {"添加学生", cmdAdd},
	"del":  {"删除学生", cmdDel},
}

// 系统配置常量
const (
	schoolName        = "GUDT"
	schoolMaxStudents = 10
)

// 系统退出 Flag
var sysExit = false

// 学校
var mySchool = newSchool(schoolName, schoolMaxStudents)

// ID高水位线
var idHWM = 1001

func main() {
	cmd := ""

	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	fmt.Printf("  欢迎来到 %v 学校的学生管理系统\n", mySchool.name)
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	fmt.Printf("\n  您可以通过输入 help 查看当前可用指令...\n")

	for !sysExit {
		inputStr(&cmd, ">>> ")

		if cmd == "help" {
			cmdHelp()
		} else {
			cmdBody, ok := cmdSet[cmd]
			if ok {
				cmdBody.exec()
			} else {
				fmt.Printf("指令 %v 不存在, 请通过 help 查看有效指令...\n", cmd)
			}
		}
	}
}
