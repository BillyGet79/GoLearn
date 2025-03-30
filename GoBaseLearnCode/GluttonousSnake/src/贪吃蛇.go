package main

import (
	"math/rand"
	"time"
)

// 全局常量 界面大小
const WIDE int = 20
const HIGH int = 20

// 定义全局变量存储食物
var food Food

// 初始化父类 坐标
type Position struct {
	X int
	Y int
}

// 初始化蛇子类
type Snake struct {
	size int                   //长度
	dir  byte                  //方向
	pos  [WIDE * HIGH]Position //定义数组存储每一节蛇的坐标
}

// 初始化食物的子类
type Food struct {
	Position
}

// 初始化蛇信息
func (s *Snake) SnakeInit() {
	//蛇的长度
	s.size = 2
	//蛇头位置
	s.pos[0].X = WIDE / 2
	s.pos[0].Y = HIGH / 2

	s.pos[1].X = WIDE/2 - 1
	s.pos[1].Y = HIGH / 2
	//蛇的方向
	//用 U上 D下 L左 R右
	s.dir = 'R'
}

func RandomFood() {
	//随机食物
	//重合之后重新随机
	food.X = rand.Intn(WIDE) //0-19
	food.Y = rand.Intn(HIGH)
}

func main() {
	//设置随机数种子	用作于混淆
	rand.New(rand.NewSource(time.Now().UnixNano()))
	//随机食物
	RandomFood()
	//创建蛇对象
	var s Snake
	//初始化蛇信息
	s.SnakeInit()
}
