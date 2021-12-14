package files

import ("os"; "bufio")

func ReadFile(filepath string) (*bufio.Scanner, *os.File) {
  file, err := os.Open(filepath)
  if err != nil {
    panic(err)
  }
  return bufio.NewScanner(file), file
}
