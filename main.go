package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func main() {
	if len(os.Args) > 2 {
		args := os.Args[1:]
		paths, err := getDirContents(args)
		checkError(err)
		paths, err = filterDirContents(args, paths)
		commentToken := getCommentToken(args[1])
		checkError(err)
		var contents []string
		contents, err = getContentsFromFiles(paths, args[0])
		checkError(err)
		paths = isolateFileNames(paths)
		resultStr := createResultString(paths, contents, commentToken)
		err = os.WriteFile("summary.txt", []byte(resultStr), 0644)
		checkError(err)
		switch len(paths) {
		case 0:
			fmt.Println("No matching files found.")
		case 1:
			fmt.Println("File written to 'summary.txt'!")
		default:
			fmt.Println("Files written to 'summary.txt'!")
		}
	} else {
		printUsage()
	}
}

func printUsage() {
	fmt.Println(
		"Usage: docsum <dir path> <file ext>",
	)
}

func getDirContents(args []string) ([]string, error) {
	files := make([]string, 0)
	var err error
	if len(args) > 1 {
		root := args[0]
		fileSystem := os.DirFS(root)
		err = fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			files = append(files, path)
			return nil
		})
	} else {
		err = fmt.Errorf("Not enough arguments.")
	}
	return files, err
}

func filterDirContents(args []string, contents []string) ([]string, error) {
	result := make([]string, 0)
	var err error
	if len(args) > 1 {
		for _, entry := range contents {
			ext := fmt.Sprintf(".%s", args[1])
			if strings.HasSuffix(entry, ext) {
				result = append(result, entry)
			}
		}
	} else {
		err = fmt.Errorf("Not enough arguments. No entries removed.")
		result = contents
	}
	return result, err
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func getContentsFromFiles(filepaths []string, root string) ([]string, error) {
	result := make([]string, 0)
	var err error

	for _, path := range filepaths {
		dirFS := os.DirFS(root)
		contents, err := fs.ReadFile(dirFS, path)
		if err != nil {
			fmt.Println(err)
			continue
		}
		result = append(result, string(contents))
	}
	return result, err
}

func isolateFileNames(filepaths []string) []string {
	result := make([]string, len(filepaths))
	for i, path := range filepaths {
		if filepath.Base(path) != path {
			result[i] = filepath.Base(path)
		} else {
			result[i] = path
		}
	}
	return result
}

func getCommentToken(ext string) string {
	doubleSlashLangs := []string{"c", "cpp", "cc", "java", "js", "ts", "go", "odin", "rust", "zig"}
	hashLangs := []string{"py", "rb"}
	var result string
	switch {
	case slices.Contains(doubleSlashLangs, ext):
		result = "//"
	case slices.Contains(hashLangs, ext):
		result = "#"
	default:
		result = "-"
	}
	return result
}

func createResultString(paths []string, contents []string, commentToken string) string {
	result := ""
	for i := range len(paths) {
		temp := fmt.Sprintf("%s %s\n", commentToken, paths[i])
		result += temp
		result += strings.Trim(contents[i], " \n")
		result += "\n\n"
	}
	result = strings.Trim(result, " \n")
	result += "\n"
	return result
}
