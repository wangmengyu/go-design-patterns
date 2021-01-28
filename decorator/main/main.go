package main

import (
	"design_patterns/decorator"
	"fmt"
)

func main() {

	//首先获得一个蔬菜大饼
	pizza := &decorator.VeggeMania{} // 15

	//加入cheese浇头
	cheesePizza := &decorator.CheeseTopping{Pizza: pizza} // 15+10
	//加入番茄浇头

	tomatoCheesePizza := &decorator.TomatoTopping{ // 15+7
		Pizza: cheesePizza,
	}

	//获得整体的披萨的价格
	fmt.Println(tomatoCheesePizza.GetPrice())
}
