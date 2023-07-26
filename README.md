# go_translation

## API

- `http://localhost:8080/translate`
```bash
POST：

package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

func (p *Person) SayHello() {
	fmt.Println("Hello, my name is", p.Name, "and I'm", p.Age, "years old")
}

func main() {
	person := NewPerson("John", 30)
	if person != nil {
		person.SayHello()
	} else {
		fmt.Println("Error creating person")
	}
}
```
```bash
ANSWER：

package main

パッケージをインポート
"fmt"
)

Personという名前の構造体を定義
Name string
Age  int
}

NewPersonという名前の関数を定義し、パラメータはname string, age intです
return &Person{
Name: name,
Age:  age,
}
}

(p *Person) SayHelloという名前の関数を定義し、パラメータはです
「Hello, my name is", p.Name, "and I'm", p.Age, "years old」を出力
}

mainという名前の関数を定義し、パラメータはです
personという名前の変数を定義し、値はNewPerson("John", 30)です
もしperson != nil {の場合
person.SayHello()
} else {
「Error creating person」を出力
}
}
```