package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 学生结构体
type Student struct {
	ID     int       `json:"id"`
	Name   string    `json:"name"`
	Age    int       `json:"age"`
	Scores []float64 `json:"scores"`
}

// 计算平均分
func (s Student) Average() float64 {
	if len(s.Scores) == 0 {
		return 0
	}

	total := 0.0
	for _, score := range s.Scores {
		total += score
	}
	return total / float64(len(s.Scores))
}

// 添加成绩
func (s *Student) AddScore(score float64) {
	s.Scores = append(s.Scores, score)
}

// 学生管理器
type StudentManager struct {
	students map[int]*Student
	nextID   int
	filename string
}

// 创建新的学生管理器
func NewStudentManager(filename string) *StudentManager {
	sm := &StudentManager{
		students: make(map[int]*Student),
		nextID:   1,
		filename: filename,
	}
	sm.loadFromFile()
	return sm
}

// 添加学生
func (sm *StudentManager) AddStudent(name string, age int) *Student {
	student := &Student{
		ID:     sm.nextID,
		Name:   name,
		Age:    age,
		Scores: make([]float64, 0),
	}

	sm.students[sm.nextID] = student
	sm.nextID++
	return student
}

// 根据ID查找学生
func (sm *StudentManager) FindByID(id int) (*Student, bool) {
	student, exists := sm.students[id]
	return student, exists
}

// 根据姓名查找学生
func (sm *StudentManager) FindByName(name string) []*Student {
	var results []*Student
	for _, student := range sm.students {
		if strings.Contains(strings.ToLower(student.Name), strings.ToLower(name)) {
			results = append(results, student)
		}
	}
	return results
}

// 删除学生
func (sm *StudentManager) DeleteStudent(id int) error {
	if _, exists := sm.students[id]; !exists {
		return fmt.Errorf("student with ID %d not found", id)
	}
	delete(sm.students, id)
	return nil
}

// 获取所有学生
func (sm *StudentManager) GetAllStudents() []*Student {
	students := make([]*Student, 0, len(sm.students))
	for _, student := range sm.students {
		students = append(students, student)
	}
	return students
}

// 保存到文件
func (sm *StudentManager) SaveToFile() error {
	data, err := json.MarshalIndent(sm.students, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}

	err = os.WriteFile(sm.filename, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

// 从文件加载
func (sm *StudentManager) loadFromFile() error {
	data, err := os.ReadFile(sm.filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // 文件不存在，正常情况
		}
		return fmt.Errorf("failed to read file: %w", err)
	}

	err = json.Unmarshal(data, &sm.students)
	if err != nil {
		return fmt.Errorf("failed to unmarshal data: %w", err)
	}

	// 更新nextID
	for id := range sm.students {
		if id >= sm.nextID {
			sm.nextID = id + 1
		}
	}

	return nil
}

// 应用程序
type App struct {
	manager *StudentManager
	scanner *bufio.Scanner
}

func NewApp() *App {
	return &App{
		manager: NewStudentManager("students.json"),
		scanner: bufio.NewScanner(os.Stdin),
	}
}

func (app *App) showMenu() {
	fmt.Println("\n=== 学生管理系统 ===")
	fmt.Println("1. 添加学生")
	fmt.Println("2. 查看所有学生")
	fmt.Println("3. 查找学生（按ID）")
	fmt.Println("4. 查找学生（按姓名）")
	fmt.Println("5. 添加成绩")
	fmt.Println("6. 删除学生")
	fmt.Println("7. 保存数据")
	fmt.Println("8. 退出")
	fmt.Print("请选择操作（1-8）: ")
}

func (app *App) readLine() string {
	app.scanner.Scan()
	return strings.TrimSpace(app.scanner.Text())
}

func (app *App) readInt() (int, error) {
	input := app.readLine()
	return strconv.Atoi(input)
}

func (app *App) readFloat() (float64, error) {
	input := app.readLine()
	return strconv.ParseFloat(input, 64)
}

func (app *App) addStudent() {
	fmt.Print("请输入学生姓名: ")
	name := app.readLine()

	fmt.Print("请输入学生年龄: ")
	age, err := app.readInt()
	if err != nil {
		fmt.Printf("无效的年龄: %v\n", err)
		return
	}

	student := app.manager.AddStudent(name, age)
	fmt.Printf("成功添加学生: ID=%d, 姓名=%s, 年龄=%d\n", student.ID, student.Name, student.Age)
}

func (app *App) showAllStudents() {
	students := app.manager.GetAllStudents()
	if len(students) == 0 {
		fmt.Println("没有学生记录")
		return
	}

	fmt.Println("\n所有学生信息:")
	fmt.Println("ID\t姓名\t年龄\t成绩\t平均分")
	fmt.Println("-------------------------------------------")
	for _, student := range students {
		scoresStr := fmt.Sprintf("%v", student.Scores)
		if len(student.Scores) == 0 {
			scoresStr = "无"
		}
		fmt.Printf("%d\t%s\t%d\t%s\t%.2f\n",
			student.ID, student.Name, student.Age, scoresStr, student.Average())
	}
}

func (app *App) findStudentByID() {
	fmt.Print("请输入学生ID: ")
	id, err := app.readInt()
	if err != nil {
		fmt.Printf("无效的ID: %v\n", err)
		return
	}

	student, exists := app.manager.FindByID(id)
	if !exists {
		fmt.Printf("未找到ID为 %d 的学生\n", id)
		return
	}

	fmt.Printf("学生信息: ID=%d, 姓名=%s, 年龄=%d, 成绩=%v, 平均分=%.2f\n",
		student.ID, student.Name, student.Age, student.Scores, student.Average())
}

func (app *App) findStudentByName() {
	fmt.Print("请输入学生姓名（支持模糊搜索）: ")
	name := app.readLine()

	students := app.manager.FindByName(name)
	if len(students) == 0 {
		fmt.Printf("未找到姓名包含 '%s' 的学生\n", name)
		return
	}

	fmt.Printf("找到 %d 个匹配的学生:\n", len(students))
	for _, student := range students {
		fmt.Printf("ID=%d, 姓名=%s, 年龄=%d, 平均分=%.2f\n",
			student.ID, student.Name, student.Age, student.Average())
	}
}

func (app *App) addScore() {
	fmt.Print("请输入学生ID: ")
	id, err := app.readInt()
	if err != nil {
		fmt.Printf("无效的ID: %v\n", err)
		return
	}

	student, exists := app.manager.FindByID(id)
	if !exists {
		fmt.Printf("未找到ID为 %d 的学生\n", id)
		return
	}

	fmt.Print("请输入成绩: ")
	score, err := app.readFloat()
	if err != nil {
		fmt.Printf("无效的成绩: %v\n", err)
		return
	}

	if score < 0 || score > 100 {
		fmt.Println("成绩应该在0-100之间")
		return
	}

	student.AddScore(score)
	fmt.Printf("成功为学生 %s 添加成绩 %.2f，当前平均分: %.2f\n",
		student.Name, score, student.Average())
}

func (app *App) deleteStudent() {
	fmt.Print("请输入要删除的学生ID: ")
	id, err := app.readInt()
	if err != nil {
		fmt.Printf("无效的ID: %v\n", err)
		return
	}

	student, exists := app.manager.FindByID(id)
	if !exists {
		fmt.Printf("未找到ID为 %d 的学生\n", id)
		return
	}

	fmt.Printf("确认删除学生 %s（ID: %d）吗？(y/N): ", student.Name, student.ID)
	confirm := app.readLine()

	if strings.ToLower(confirm) == "y" {
		err := app.manager.DeleteStudent(id)
		if err != nil {
			fmt.Printf("删除失败: %v\n", err)
		} else {
			fmt.Printf("成功删除学生 %s\n", student.Name)
		}
	} else {
		fmt.Println("取消删除操作")
	}
}

func (app *App) saveData() {
	err := app.manager.SaveToFile()
	if err != nil {
		fmt.Printf("保存失败: %v\n", err)
	} else {
		fmt.Println("数据保存成功!")
	}
}

func (app *App) Run() {
	fmt.Println("欢迎使用学生管理系统!")

	defer func() {
		// 程序退出时自动保存
		app.manager.SaveToFile()
	}()

	for {
		app.showMenu()
		choice := app.readLine()

		switch choice {
		case "1":
			app.addStudent()
		case "2":
			app.showAllStudents()
		case "3":
			app.findStudentByID()
		case "4":
			app.findStudentByName()
		case "5":
			app.addScore()
		case "6":
			app.deleteStudent()
		case "7":
			app.saveData()
		case "8":
			fmt.Println("感谢使用学生管理系统，再见!")
			return
		default:
			fmt.Println("无效的选择，请输入1-8之间的数字")
		}
	}
}

func main() {
	app := NewApp()
	app.Run()
}
