package sys

import "github.com/toolkits/pkg/nux"
// 磁盘
func Df() (map[string]interface{},error) {

	var res = make(map[string]interface{})
	mountPoints, err := nux.ListMountPoint()
	if err != nil{
		return nil,err
	}

	var diskTotal uint64 = 0
	var diskUsed uint64 = 0

	for idx := range mountPoints{
		fsSpec, fsFile, fsVfstype := mountPoints[idx][0], mountPoints[idx][1], mountPoints[idx][2]

		var du *nux.DeviceUsage
		du, err = nux.BuildDeviceUsage(fsSpec, fsFile, fsVfstype)
		if err != nil {
			continue
		}

		diskTotal += du.BlocksAll
		diskUsed += du.BlocksUsed

		/*tags := fmt.Sprintf("mount=%s", du.FsFile)
		fmt.Printf("tags:%s\n",tags)
		fmt.Printf("total:%s\n",du.BlocksAll / 1024)
		fmt.Printf("free:%s\n",du.BlocksFree / 1024)
		fmt.Printf("used:%s\n",du.BlocksUsed / 1024)*/


	}
	res["allTotal"] = diskTotal / 1024
	res["allUsed"] = diskUsed / 1024
	res["ratio"] = int(float64(diskUsed) / float64(diskTotal) * 100)
	return res,nil

}
