# Тестовое задание для [Addreality](https://addreality.com/ru)

Текст самого задания: [скачать](/task.docx?raw=true)

# Установка
```go get github.com/Madredix/addreality```

# Пример использования
```go
package main

import (
	"fmt"
	"github.com/Madredix/addreality"
)

const countLines = 2
const countArgs = 0

func main() {
	rows := []struct {
		Name       string
		GroupID    uint
		PlatformID uint
	}{
		{Name: "device 1", GroupID: 1, PlatformID: 5281},
		{Name: "device 2", GroupID: 2, PlatformID: 5281},
		{Name: "device 3", GroupID: 3, PlatformID: 5281},
	}

	addreality, err := addreality.NewInsertBuilderFactory(countLines, countArgs)
	if err != nil {
		fmt.Println(err)
		return
	}
	bulk, err := addreality.CreateInsertBuilder(`devices`, `name`, `group_id`, `platform_id`)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i, r := range rows {
		if err = bulk.Append(r.Name, r.GroupID, r.PlatformID); err != nil {
			fmt.Printf(`error adding row №%d: %s`, i, err.Error())
		}
	}

	batches := bulk.ToSQL()
	for i, b := range batches {
		fmt.Printf("Batch: %d\nSQL: %s\nParams: %+v\n\n", i+1, b.Query, b.Args)
	}
}
```

## Отклонение от задачи
Поскольку нужно контролировать, что каждая добавляемая строка содержит одинаковое кол-во аргументов, то возврат ошибки добавлен в метод _Append_ и убран из метода _ToSQL_
