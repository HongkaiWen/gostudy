package main

import (
	"fmt"
	"sync"
)

type Student struct {
	sync.RWMutex
	age uint8
}

//构造方法
func NewStudent(studentAge uint8) *Student {
	t := &Student{
		age: studentAge,
	}
	return t
}

//成员方法
func (s *Student) AgeAdd() {
	s.Lock()
	s.age += 1
	s.Unlock()
}

//成员方法
func (s *Student) study() {
	fmt.Println("I'm studying...")
}

//普通方法
func AgeAdd(student *Student) {
	student.Lock()
	student.age += 1
	student.Unlock()
}

func teacher(student *Student) {
	student.study()
}

//Person has method study
type Person struct {
	Student
}

//type Dog struct {
//}

//func (s *Dog) study() {
//	fmt.Println("dog studying...")
//}

func main() {
	student := NewStudent(2)
	AgeAdd(student)
	student.AgeAdd()
	fmt.Println(student.age)

	person := &Person{}
	person.study()

	//	dog := &Dog{}

	//	teacher(dog) //error
}
