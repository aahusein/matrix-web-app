package main
import (
   "fmt"
   "gonum.org/v1/gonum/mat"
   "net/http"
   "strconv"
   "strings"
)

// root url html
const inputForm = `
<html>
   <body>
      <form method="post">
         <textarea rows="10" cols="30" input type="text" name="input"></textarea><br>
         <input type="submit" value="Display" formaction="/DisplayMatrix"><br>
         <input type="submit" value="Inverse" formaction="/InverseMatrix"><br>
         <input type="text" name="scalar">
         <input type="submit" value="Scale" formaction="/ScaleMatrix"><br>
         <textarea rows="10" cols="30" input type="text" name="matrix"></textarea><br>
         <input type="submit" value="Multiply" formaction="/MultiplyMatrix"><br>
      </form>
   </body>
</html>
`

// handle the root url
func home(w http.ResponseWriter, r *http.Request) {
   //fmt.Fprint(w, inputForm)
   w.Write([]byte(inputForm))
}

func displayMat(w http.ResponseWriter, r *http.Request) {
   //fmt.Fprint(w, inputForm)
   input := r.FormValue("input")
   
   A := makeMat(input)
   
   printMat(w, A)
}

func inverseMat(w http.ResponseWriter, r *http.Request) {
   input := r.FormValue("input")
   
   A := makeMat(input)
   A.Inverse(A)
    
   printMat(w, A)
}

func scaleMat(w http.ResponseWriter, r *http.Request) {
   input := r.FormValue("input")
   scalar := r.FormValue("scalar")
   
   A := makeMat(input)
   c, _ := strconv.ParseFloat(scalar, 64)
   
   A.Scale(c, A)
   printMat(w, A)
}

func multiplyMat(w http.ResponseWriter, r *http.Request) {
   input := r.FormValue("input")
   matrix := r.FormValue("matrix")
   
   A := makeMat(input)
   B := makeMat(matrix)

   m := A.ColView(0).Len()
   n := B.RowView(0).Len()
   C := mat.NewDense(m, n, nil)
   
   C.Product(A, B)
   printMat(w, C)
}

func makeMat(input string) (*mat.Dense) {
   rows := strings.Split(input, "\r\n")
   m := len(rows)
   n := len(strings.Split(rows[0], " "))
   
   var v []float64
   for _ , i := range rows {
      i := strings.Split(i, " ")
      for _ , j := range i {
         a, _ := strconv.ParseFloat(j, 64)
         v = append(v, a)
      }
   }
   
   A := mat.NewDense(m, n, v)
   return A
}

func printMat(w http.ResponseWriter, A *mat.Dense) {
   fa := mat.Formatted(A, mat.Prefix(""), mat.Squeeze())
   matString := fmt.Sprintf("%v\n", fa)
   w.Write([]byte(matString))
   //w.Write([]byte("<span style=\"white-space: pre-line\">" + matString + "</span>"))
}

func main() {
   http.HandleFunc("/", home)
   http.HandleFunc("/DisplayMatrix", displayMat)
   http.HandleFunc("/InverseMatrix", inverseMat)
   http.HandleFunc("/ScaleMatrix", scaleMat)
   http.HandleFunc("/MultiplyMatrix", multiplyMat)
   if err := http.ListenAndServe(":8080", nil); err != nil {
      panic(err)
   }
}