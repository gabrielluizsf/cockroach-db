package fatal

import  "log"

func PrintFatalError(err error, message string){
	if err != nil {
		log.Fatal(message, err);
	  }
}