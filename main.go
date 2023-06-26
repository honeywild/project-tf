package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Определение флага для имени проекта
	defaultProjectName := "my_terraform_project"
	projectName := flag.String("d", defaultProjectName, "Имя проекта")
	flag.Parse()

	// Создание директории проекта
	if err := os.Mkdir(*projectName, 0755); err != nil {
		fmt.Printf("Не удалось создать директорию проекта: %v\n", err)
		return
	}

	// Создание файлов в корне проекта
	rootFiles := []string{"main.tf", "variables.tf", "outputs.tf"}
	for _, file := range rootFiles {
		if _, err := os.Create(fmt.Sprintf("%s/%s", *projectName, file)); err != nil {
			fmt.Printf("Не удалось создать файл %s: %v\n", file, err)
		}
	}

	// Создание директории modules
	modulesDir := fmt.Sprintf("%s/modules", *projectName)
	if err := os.Mkdir(modulesDir, 0755); err != nil {
		fmt.Printf("Не удалось создать директорию modules: %v\n", err)
		return
	}

	// Создание директории digitalocean внутри modules
	doModuleDir := fmt.Sprintf("%s/digitalocean", modulesDir)
	if err := os.Mkdir(doModuleDir, 0755); err != nil {
		fmt.Printf("Не удалось создать директорию digitalocean: %v\n", err)
		return
	}

	// Создание файлов внутри digitalocean
	doModuleFiles := []string{"main.tf", "variables.tf", "outputs.tf"}
	for _, file := range doModuleFiles {
		if _, err := os.Create(fmt.Sprintf("%s/%s", doModuleDir, file)); err != nil {
			fmt.Printf("Не удалось создать файл %s: %v\n", file, err)
		}
	}

	// Создание директории aws внутри modules
	awsModuleDir := fmt.Sprintf("%s/aws", modulesDir)
	if err := os.Mkdir(awsModuleDir, 0755); err != nil {
		fmt.Printf("Не удалось создать директорию aws: %v\n", err)
		return
	}

	// Создание файлов внутри aws
	awsModuleFiles := []string{"main.tf", "variables.tf", "outputs.tf"}
	for _, file := range awsModuleFiles {
		if _, err := os.Create(fmt.Sprintf("%s/%s", awsModuleDir, file)); err != nil {
			fmt.Printf("Не удалось создать файл %s: %v\n", file, err)
		}
	}

	fmt.Println("Структура проекта успешно создана.")
}

