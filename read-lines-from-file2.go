func ReadLineFromFile(filepath string) []string {
  b, err:=ioutil.ReadFile(filepath)
  if err!=nil {
    log.Fatal("Error in reading file", err)
  }
  return strings.Split(string(b), "\n")
}
