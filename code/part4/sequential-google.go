// +build ignore

package main

import 

func Search(query string) ([]Result, error) {
    results := []Result{
        Web(query),
        Image(query),
        Video(query),
    }
    return results, nil
}

func main() {

}