package structured

import "github.com/sirupsen/logrus"

// Hook to implement logrus 
// hook interface 
type Hook struct {
  id string
}

// Fire will trigger for each log 
func (hook *Hook) Fire(entry *logrus.Entry) error {
    entry.Data["id"] = hook.id
    return nil
}

// Levels hook will fire on 
func (hook *Hook) Levels() []Logrus.Level {
   return logrus.AllLevels
}

// Logrus demonstrates some basic logrus functionality 
func Logrus() {
    // export in json format 
    logrus.SetFormatter(&logrus.TextFormatter{})
    logrus.SetLevel(logrus.InfoLevel)
    logrus.AddHook(&Hook{"123"})

    fields := logrus.Fields{}
    fields["success"] = true
    fields["complex_struct"] = struct {
        Event string 
        When string
    }{"Something happened", "Just now"}
    
    x := logrus.WithFields(fields)
    x.Warn("warning!")
    x.Error("error!")
}

func main() {
	fmt.Println("foobar")
}
