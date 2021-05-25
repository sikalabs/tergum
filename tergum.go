package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyz"

func randLetters(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func mysqlDump(mysqlHost string, mysqlPort string, mysqlUser string, mysqlPassword string, mysqlDatabase string) ([]byte, error) {
	cmd := exec.Command("mysqldump", "-h", mysqlHost, "-P", mysqlPort, "-u", mysqlUser, "-p"+mysqlPassword, mysqlDatabase)
	out, err := cmd.Output()
	return out, err
}

func saveFile(fileName string, fileContent []byte) error {
	dir := filepath.Dir(fileName)
	os.MkdirAll(dir, 0755)
	err := ioutil.WriteFile(fileName, fileContent, 0644)
	return err
}

func getOutputPath(dstFilePath string, dstFileDir string, dstFilePrefix string, dstFileSuffix string) string {
	if dstFilePath != "" {
		return dstFilePath
	}
	nowFormatted := time.Now().UTC().Format("2006-01-02_15-04-05")
	return path.Join(dstFileDir, dstFilePrefix+"."+nowFormatted+"_"+randLetters(3)+"."+dstFileSuffix)
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	// Backup Source Parameters
	src := flag.String("src", "", "choose backup source form: {mysql}")

	srcMysqlHost := flag.String("src-mysql-host", "", "mysql host, eg.: 127.0.0.1")
	srcMysqlPort := flag.String("src-mysql-port", "", "mysql port, eg.: 3306")
	srcMysqlUser := flag.String("src-mysql-user", "", "mysql user, eg.: root")
	srcMysqlPassword := flag.String("src-mysql-password", "", "mysql host, eg.: root_password")
	srcMysqlDatabase := flag.String("src-mysql-database", "", "mysql host, eg.: default")

	// Backup Destination Parameters
	dst := flag.String("dst", "", "choose backup destination form: {stdout file}")
	dstFilePath := flag.String("dst-file-path", "", "output full file path, eg.: ~/backup/backup.sql")
	dstFileDir := flag.String("dst-file-dir", "", "output directory, eg.: ~/backup")
	dstFilePrefix := flag.String("dst-file-prefix", "", "output file path, eg.: default")
	dstFileSuffix := flag.String("dst-file-suffix", "", "output file path, eg.: sql")

	flag.Parse()

	var out []byte
	var err error

	switch *src {
	case "mysql":
		if *srcMysqlHost == "" {
			log.Fatal("arg -src-mysql-host must be set")
		}
		if *srcMysqlPort == "" {
			log.Fatal("arg -src-mysql-port must be set")
		}
		if *srcMysqlUser == "" {
			log.Fatal("arg -src-mysql-user must be set")
		}
		if *srcMysqlPassword == "" {
			log.Fatal("arg -src-mysql-password must be set")
		}
		if *srcMysqlDatabase == "" {
			log.Fatal("arg -src-mysql-database must be set")
		}
		out, err = mysqlDump(*srcMysqlHost, *srcMysqlPort, *srcMysqlUser, *srcMysqlPassword, *srcMysqlDatabase)
		if err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal("no src selected")
	}

	switch *dst {
	case "stdout":
		fmt.Printf("%s\n", out)
	case "file":
		if (*dstFilePath == "") && (*dstFileDir == "" || *dstFilePrefix == "" || *dstFileSuffix == "") {
			log.Fatal("args (-dst-file-path) OR (-dst-file-dir AND -dst-file-prefix AND -dst-file-suffix) must be set")
		}
		finaldstFilePath := getOutputPath(*dstFilePath, *dstFileDir, *dstFilePrefix, *dstFileSuffix)
		err = saveFile(finaldstFilePath, out)
		if err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal("no dst selected")
	}
}
