/*
	版权所有，侵权必究
	署名-非商业性使用-禁止演绎 4.0 国际
	警告： 以下的代码版权归属hunterhug，请不要传播或修改代码
	你可以在教育用途下使用该代码，但是禁止公司或个人用于商业用途(在未授权情况下不得用于盈利)
	商业授权请联系邮箱：gdccmcm14@live.com QQ:569929309

	All right reserved
	Attribution-NonCommercial-NoDerivatives 4.0 International
	Notice: The following code's copyright by hunterhug, Please do not spread and modify.
	You can use it for education only but can't make profits for any companies and individuals!
	For more information on commercial licensing please contact hunterhug.
	Ask for commercial licensing please contact Mail:gdccmcm14@live.com Or QQ:569929309

	2017.7 by hunterhug
*/
package core

// not use!!!!!down!
import (
	"github.com/hunterhug/GoSpider/util"
	"math/rand"
)

var (
	num  int
	cnum chan string
)

func collectasintask(taskname string, files []string) {
	second := rand.Intn(5)
	AmazonListLog.Debugf("%s:%d second", taskname, second)
	util.Sleep(second)
	CollectAsin(files)
	cnum <- taskname
}

func LocalListParseTask() {
	OpenMysql()
	err := CreateAsinTables()
	if err != nil {
		AmazonListLog.Errorf("createtables:%s,error:%s", Today, err.Error())
	}
	err = CreateAsinRankTables()
	if err != nil {
		AmazonListLog.Errorf("createtables:%s,error:%s", "Asin"+Today, err.Error())
	}
	num = MyConfig.Localtasknum
	cnum = make(chan string, num)
	stoptimes := 10
	// loop loop loop loop
	for {

		if stoptimes == 0 {
			break
		}
		files, err := util.ListDir(DataDir+"/list/"+*day, "html")
		if err != nil {
			AmazonListLog.Errorf("open dir error:%s", err.Error())
			break
		}
		filess, err := util.DevideStringList(files, num)
		if err != nil {
			if len(files) > 0 {
				collectasintask("dudu", files)
				<-cnum
			}
			AmazonListLog.Error(err.Error())
			stoptimes--
			// sleep 10 minute 60/10*24=144
			AmazonListLog.Error("wait for 10 minute。。")
			util.Sleep(600)
			continue
		}
		for i := 0; i < num; i++ {
			go collectasintask("Collect Asin task"+util.IS(i), filess[i])
		}
		for i := 0; i < num; i++ {
			AmazonListLog.Log("Collect Asin %s done!\n", <-cnum)
		}

	}
	AmazonListLog.Log("Collect Asin All done")

}
/*
	版权所有，侵权必究
	署名-非商业性使用-禁止演绎 4.0 国际
	警告： 以下的代码版权归属hunterhug，请不要传播或修改代码
	你可以在教育用途下使用该代码，但是禁止公司或个人用于商业用途(在未授权情况下不得用于盈利)
	商业授权请联系邮箱：gdccmcm14@live.com QQ:569929309

	All right reserved
	Attribution-NonCommercial-NoDerivatives 4.0 International
	Notice: The following code's copyright by hunterhug, Please do not spread and modify.
	You can use it for education only but can't make profits for any companies and individuals!
	For more information on commercial licensing please contact hunterhug.
	Ask for commercial licensing please contact Mail:gdccmcm14@live.com Or QQ:569929309

	2017.7 by hunterhug
*/