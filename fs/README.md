# Filesystem Package
A basic filesystem package to overlay the Go standard libraries to add 
missing functionality commonly used. 

### Problem 1: Path validation
Is a given string a valid path:

```
// Empirical Solution

Attempt to create the file and delete it right after.

func IsValid(fp string) bool {
  // Check if file already exists
  if _, err := os.Stat(fp); err == nil {
    return true
  }

  // Attempt to create it
  var d []byte
  if err := ioutil.WriteFile(fp, d, 0644); err == nil {
    os.Remove(fp) // And delete it
    return true
  }

  return false
}
```
