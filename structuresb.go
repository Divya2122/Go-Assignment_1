package main

import "fmt"

type Books struct {
   title string
   author string
   subject string
   book_id int
}
func main() {
   var Book1 Books    
   var Book2 Books   
    var Book3 Books   
 
   
   Book1.title = "java"
   Book1.author = "james gosling"
   Book1.subject = "java Tutorials"
   Book1.book_id = 1998
   
   Book2.title = "go"
   Book2.author = "divya"
   Book2.subject = "go Tutorials"
   Book2.book_id = 1999

   
   Book3.title = "c"
   Book3.author = "dennis ritchie"
   Book3.subject = "c Learning"
   Book3.book_id = 1999
 
   
   fmt.Printf( "Book 1 title : %s\n", Book1.title)
   fmt.Printf( "Book 1 author : %s\n", Book1.author)
   fmt.Printf( "Book 1 subject : %s\n", Book1.subject)
   fmt.Printf( "Book 1 book_id : %d\n", Book1.book_id)

   
   fmt.Printf( "Book 2 title : %s\n", Book2.title)
   fmt.Printf( "Book 2 author : %s\n", Book2.author)
   fmt.Printf( "Book 2 subject : %s\n", Book2.subject)
   fmt.Printf( "Book 2 book_id : %d\n", Book2.book_id)

   fmt.Printf( "Book 3 title : %s\n", Book2.title)
   fmt.Printf( "Book 3 author : %s\n", Book2.author)
   fmt.Printf( "Book 3 subject : %s\n", Book2.subject)
   fmt.Printf( "Book 3 book_id : %d\n", Book2.book_id)
}