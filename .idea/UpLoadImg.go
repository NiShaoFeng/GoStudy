func upLoads(ctx *gin.Context){
	file,err := ctx.FormFile("file")
	if err == nil{
		extName := path.Ext(file.Filename) 
		allowExtMap := Map[string]bool{
			".jpg":true,
			".png":true,
			".jpeg":true,
		}
	}
	if _, ok :=allowExtMap[extName];!ok{
		
	}
}