# dir
ETL Utilities that might be of no use to anybody else :)

### 1. LastModifiedFile: From a glob pattern, get the newest file by modified time.
This is a golang implemention of what we might do in Ruby as:

```ruby
Dir.glob('/path/to/storage/pattern_*.xml').max_by { |file| File.mtime(file) }
```

To use:

```golang
import "github.com/paul-ylz/dir"

func main() {
  // file is a string path to a file
  file := dir.LastModifiedFile("/path/to/storage/pattern_*.xml")

  // do something with file...
}
```
