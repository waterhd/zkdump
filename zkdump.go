package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"syscall"
	"time"

	"github.com/samuel/go-zookeeper/zk"
	"golang.org/x/crypto/ssh/terminal"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	err       error
	c         = &zk.Conn{}
	app       = kingpin.New("zkdump", "A command-line utility to dump Zookeeper data.").Author("Dennis Waterham <dennis.waterham@oracle.com>").Version("1.0")
	servers   = app.Flag("server", "Host name and port to connect to (host:port)").Required().Short('s').Strings()
	verbose   = app.Flag("verbose", "Print verbose.").Short('v').Bool()
	user      = app.Flag("user", "Username to use for digest authentication.").Short('u').String()
	password  = app.Flag("password", "Password to use for digest authentication (will read from TTY if not given).").Short('p').String()
	recursive = app.Flag("recursive", "Get nodes recursively.").Short('r').Bool()
	rootpath  = app.Arg("path", "Root path (default: \"/\").").Default("/").String()
)

type zkNode struct {
	Name     string
	Path     string
	Data     string   `json:",omitempty"`
	Children []zkNode `json:",omitempty"`
}

func (z *zkNode) getChildren() {
	items, st, err := c.Children(z.Path)

	if err != nil {
		log.Fatal(err)
	}

	if st.NumChildren == 0 {
		return
	}

	for _, child := range items {
		z.Children = append(z.Children, *getZkNode(path.Join(z.Path, child), child))
	}
}

func main() {

	kingpin.MustParse(app.Parse(os.Args[1:]))

	if *user != "" {
		if *password == "" {
			*password = readPassword()
		}
	}

	c, _, err = zk.Connect(*servers, time.Second, zk.WithLogInfo(*verbose))
	defer c.Close()
	check(err)

	if *user != "" {
		verboseLog("Adding digest authentication for user %s", *user)
		c.AddAuth("digest", []byte(*user+":"+*password))
	}

	verboseLog("Checking if root path %s exists", *rootpath)
	exists, _, err := c.Exists(*rootpath)
	check(err)

	if !exists {
		log.Fatalf("ERROR: Path %s doesn't exist", *rootpath)
	}

	// Get Root node
	rootNode := getZkNode(*rootpath, path.Base(*rootpath))
	//	saveJSON(rootNode, "test.json")

	bin, err := json.MarshalIndent(&rootNode, "", "  ")
	fmt.Println(string(bin))
}

func verboseLog(s string, p string) {
	if *verbose {
		log.Printf(s+"\n", p)
	}
}

func getZkNode(path, name string) *zkNode {
	bin, st, err := c.Get(path)
	check(err)

	zkNode := &zkNode{Path: path, Name: name, Data: string(bin)}

	if st.NumChildren > 0 {
		zkNode.getChildren()
	}

	return zkNode
}

func readPassword() string {
	fmt.Print("Enter Password: ")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	check(err)

	fmt.Printf("\n")
	return string(bytePassword)
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func saveJSON(data interface{}, file string) {
	outFile, err := os.Create(file)
	defer outFile.Close()
	if err != nil {
		log.Fatalln("Error occurred")
	}

	jsonWriter := json.NewEncoder(outFile)
	jsonWriter.SetIndent("", "   ")
	jsonWriter.Encode(&data)
}
