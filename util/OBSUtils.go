package util

import "fmt"

func uploadHuawei()  {
	input := &obs.PutFileInput{}
	input.Bucket = "bucketname"
	input.Key = "objectname"
	input.SourceFile = "localfile"  // localfile为待上传的本地文件路径，需要指定到具体的文件名
	output, err := obsClient.PutFile(input)

	if err == nil {
		fmt.Printf("RequestId:%s\n", output.RequestId)
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println(obsError.Code)
			fmt.Println(obsError.Message)
		} else {
			fmt.Println(err)
		}
	}
}
