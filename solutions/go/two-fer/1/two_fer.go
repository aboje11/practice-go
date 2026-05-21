package twofer

        

          
          

        

          
          
import "fmt"

        

          
          

        

          
          
// ShareWith needs a comment documenting it.

        

          
          
func ShareWith(name string) string {

        

          
          
	if (name == "") {

        

          
          
		return "One for you, one for me."

        

          
          
	} else {

        

          
          
		return fmt.Sprintf("One for %s, one for me.", name)

        

          
          
	}

        

          
          
}

        