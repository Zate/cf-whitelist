package main

import (
	"log"
	"io/ioutil"
	"strings"
  "net/http"
  "bufio"
  "regexp"
)



func main() {
  log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
  url4 := "https://www.cloudflare.com/ips-v4"
  url6 := "https://www.cloudflare.com/ips-v6"
  stuff := "  whitelistSourceRange = [ "
  resp, err := http.Get(url4)
  if err != nil {
    log.Fatal(err)
  }
  defer resp.Body.Close()
  
  if resp.StatusCode == http.StatusOK {
    
    s := bufio.NewScanner(resp.Body)
    for s.Scan() {
      stuff = stuff + "\"" + s.Text() + "\" "
      //log.Println(s.Text())
    }
    if err := s.Err(); err != nil {
      log.Fatal(err)
    }
  }
  resp, err = http.Get(url6)
  if err != nil {
    log.Fatal(err)
  }
  defer resp.Body.Close()
  
  if resp.StatusCode == http.StatusOK {
    
    s := bufio.NewScanner(resp.Body)
    for s.Scan() {
      stuff = stuff + "\"" + s.Text() + "\" "
      //log.Println(s.Text())
    }
    if err := s.Err(); err != nil {
      log.Fatal(err)
    }
  }
  stuff = stuff + "]"
  stuff = strings.Replace(stuff, "\" \"", "\", \"", -1)
  
    
//  log.Println(stuff)
  

  // Write the body to file
  
  path := "traefik.toml"
  read, err := ioutil.ReadFile(path)
  if err != nil {
    panic(err)
  }
  //fmt.Println(string(read))
 // log.Println(path)
  
 // match, _ := regexp.MatchString("  whitelistSourceRange = \\[.*\\]", string(read))
  //log.Println(match)
  
  re := regexp.MustCompile("  whitelistSourceRange = \\[.*\\]")
	newContents := re.ReplaceAllString(string(read), stuff)

  //newContents := strings.Replace(string(read), "", "new", -1)

  //log.Println(newContents)

  err = ioutil.WriteFile("traefik.toml.new", []byte(newContents), 0600)
		if err != nil {
			log.Panic(err)
	}
}
