package uploads

import(
	"io"
	"net/http"
	"os"
	"errors"
	CO "../config"
	"mime/multipart"
)

/*
	SingleFileUpload
	This function will read an image file and 
	upload it to the server
	the function will return an error type and the file directory string
	
	This function can be modified for other file types i.e images,audio,etc
*/
func SingleFileUpload(file multipart.File,handler *multipart.FileHeader)(string,error){
	var imageType string	
	if handler.Filename == ""{
		err := errors.New("File does not exist")
		return "",err
	}
	
	//Create a buffer that will store the header of the file
	fileHeader := make([]byte,512)

	//Copy the headers into the FileHeader buffer
	_,err := file.Read(fileHeader)
	if err != nil {
		return "",err
	}

	//Set position back to start.
	if _,err := file.Seek(0,0); err != nil{
		return "",err
	}
		
	if http.DetectContentType(fileHeader) == "image/png"{
		//Handle upload
		imageType = ".png"				
	}else if http.DetectContentType(fileHeader) == "image/jpg"{
		//Handle upload
		imageType = ".jpg"
	}else if http.DetectContentType(fileHeader) == "image/gif"{
		//Handle upload
		imageType = ".gif"
	}else if http.DetectContentType(fileHeader) == "image/jpeg"{
		//Handle upload
		imageType = ".jpg"
	}else{
		err = errors.New("Wrong file type")
		return "",err
	}
	
	//Use GenerateID to create a unique string
	imageName := CO.GenerateID("pro_pic",16)
	if err != nil{
		return "",err
	}
	//This is the directory name of the folder that we want to save our image to
	serverFileName := "./data/pictures/"+imageName+imageType
	//This will store the picture directory to the database
	dbFileName := "/data/pictures/"+imageName+imageType
	
	out,err := os.Create(serverFileName)		
	if err != nil{
		return "",err
	}
		
	defer out.Close()
	//Copy the image to the server
	io.Copy(out,file)
	/*
		If everything goes well with the upload we return
		the dbFileName string and also a nil error
	*/		
	return dbFileName,nil
}

