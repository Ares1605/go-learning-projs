package system

import (
  "fmt"
  "3-more-packages/system/helper"
)

func GetDiskSpace() string {
  free := helper.BytesToMB(helper.GetFreeDiskSpace())
  total := helper.BytesToMB(helper.GetTotalDiskSpace())
  return fmt.Sprintf("%d/%d", free, total)

}
