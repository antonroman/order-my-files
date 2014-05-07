package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	//"reflect"
	"strconv"
	"os"
	"os/exec"
)

func main() {

	//Set path
	path := "/home/anton/Desktop/100MSDCF"

	//Get the list of files of the directory	
	fileList, err := ioutil.ReadDir(path)
	if err != nil {
    		log.Fatal(err)
	}

	var folder_name string

	photo_directories := make(map[string] int)

	for _, f := range fileList {

		fmt.Printf("File name: %v\n",f.Name())
		if (f.IsDir()){
			fmt.Printf("The file is a folder. Let's try with next file.\n")
			continue
		}

		//Get modification time from the file
		year, month, _ := f.ModTime().Date()
		//fmt.Println(year,reflect.TypeOf(year))

		//Create folder name 
		name := []string{strconv.Itoa(int(year)),strconv.Itoa(int(month))}
		if int(month) < 10{
			folder_name= strings.Join(name,"_0")
		}else{			
			folder_name= strings.Join(name,"_")
		}
		
		//Add folder name to map
		photo_directories[folder_name] += 1
		if (photo_directories[folder_name] == 1){
			fmt.Printf("Creating folder name: %v\n",folder_name)
			err := os.Mkdir(path+"/"+folder_name,0777) // Create folder.
			if err != nil {
				log.Fatal(err)
				fmt.Printf("Can not create folder: %v\n",folder_name)
			}
		}

		//Copying file to its directory, if the copy goes fine, delete the file
		fmt.Printf("Copying file "+ path+"/"+f.Name()+" to "+path+"/"+folder_name+"/"+f.Name()+"\n")
		cpCmd := exec.Command("cp", "-rf", path+"/"+f.Name(), path+"/"+folder_name+"/"+f.Name())
		err := cpCmd.Run()
		if err != nil {
			log.Fatal(err)
			fmt.Printf("Can not copy file: %v\n",f.Name())
		}else{
			fmt.Printf("Deleting file: "+path+"/"+f.Name()+"\n")
			err := os.Remove(path+"/"+f.Name()) // Remove file.
                        if err != nil {
                                log.Fatal(err)
                                fmt.Printf("Can not delete file: %v\n",f.Name())
                        }

		}
		  
	}
	//for key, _ := range photo_directories {
	//	fmt.Println("Key:", key)
	//}
}
