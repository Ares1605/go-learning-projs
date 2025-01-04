package helper

func GetFreeDiskSpace() int {
  return 100000000
}
func GetTotalDiskSpace() int {
  return 200000000
}

func BytesToMB(bytes int) int {
  return bytes / 1048576
}
