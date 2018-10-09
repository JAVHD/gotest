/*''
* @Author: bear
* @Date:   2018-09-14 09:03:39
* @Last Modified by:   bear
* @Last Modified time: 2018-09-14 09:09:29
 */

package main

import (
	"fmt"
	"net/http"
)

func WebHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "hello", "bear")

}

func main() {

	http.HandleFunc("/api/list", WebHandler)

	http.ListenAndServe(":8082", nil)

}
