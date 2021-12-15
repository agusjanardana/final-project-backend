package exceptions

func PanicIfError(err error) {
	//defer EndApp()
	if err != nil {
		panic(err)
	}
	//fmt.Println("Aplikasi Berjalan")
}

//func EndApp(){
//	message := recover()
//	if message != nil {
//		fmt.Println("Message error", message)
//	}
//	fmt.Println("Done")
//}