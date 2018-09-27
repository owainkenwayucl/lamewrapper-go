// This program wraps the LAME encoder with my standard options.
// Dr Owain Kenway

// Where it is distributed, it is done so under a 4 clause,
// BSD-style license (see LICENSE)

package main

import (
        "fmt"
        "os"
        "os/exec"
        "bufio"
        "strings"
       )

// Interactively ask the user for the details of the file we want to encode.
func interactive() {

    title("interactive")

    fmt.Print(" File name: ")
    file := getinput()

    fmt.Print(" Song name: ")
    song := getinput()

    fmt.Print("    Artist: ")
    artist := getinput()

    fmt.Print("     Album: ")
    album := getinput()

    fmt.Print("     Genre: ")
    genre := getinput()

    fmt.Println("")

    encode(file, song, artist, album, genre)
}

// This is cumbersome and we do it multiple times so seperate out.
func getinput() string {
    stdin_r := bufio.NewReader(os.Stdin)
    ret_val, _ := stdin_r.ReadString('\n')
    ret_val = strings.Replace(ret_val, "\n", "", -1)
    return ret_val
}

// Encode stuff with lame.
func encode(file string, song string, artist string, album string, genre string) {

    cmd  := exec.Command("lame", "-b", "320", "-B", "320", "-q", "0", "--tt", song, "--ta", artist, "--tl", album, "--tg", genre, file )

    fmt.Print("Encoding... ")
    err := cmd.Run()
    if err != nil {
        fmt.Println("[ERROR]")
        fmt.Println(err)
    } else {
        fmt.Println("[Done]")
    }
}

// Print header
func title(mode string) {
    fmt.Println("")
    fmt.Println("Owain Kenway's LAME wrapper version (ii) [" + mode + " mode]")
    fmt.Println("========================================")
    fmt.Println("")
}

// Print out what we are doing.
func debug(file string, song string, artist string, album string, genre string) {
    fmt.Println(" File name: " + file)
    fmt.Println(" Song name: " + song)
    fmt.Println("    Artist: " + artist)
    fmt.Println("     Album: " + album)
    fmt.Println("     Genre: " + genre)
    fmt.Println("")
}

// Decide whether the user has invoked interactive mode
// or whether to parse arguemnts.
func main () {
    if (len(os.Args) == 1) {
        // if there are no arguments enter interactive mode.
	interactive()
    } else if (len(os.Args) == 6) {
        // if there are five arguments use those.
        file := os.Args[1]
        song := os.Args[2]
        artist := os.Args[3]
        album := os.Args[4]
        genre := os.Args[5]

        title("batch")
        debug(file, song, artist, album, genre)
        encode(file, song, artist, album, genre)
    } else {
        // print usage.
        cmd := os.Args[0]
        usage := "Usage:\n  " + cmd + " <file> <song> <artist> <album> <genre>"
        fmt.Println(usage)
    }
}
