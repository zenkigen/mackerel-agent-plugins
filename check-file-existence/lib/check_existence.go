package checkexistence

import (
  "fmt"
  "flag"
  "os"

  "github.com/mackerelio/checkers"
  "github.com/mattn/go-pipeline"
)

func run(path string) *checkers.Checker {
  _, err := pipeline.Output(
    []string{"sh", "-c", fmt.Sprintf("ls %s 2>/dev/null", path)},
  )

  if err != nil {
    return checkers.NewChecker(checkers.CRITICAL, fmt.Sprintf("Failed to access to %s by %s", path, err.Error()))
  }

  return checkers.NewChecker(checkers.OK, fmt.Sprintf("Successful access to %s", path))
}

// Do the plugin
func Do() {
  optPath := flag.String("path", "", "Path to file or directory")
  flag.Parse()

  if *optPath == "" {
    flag.PrintDefaults()
    os.Exit(1)
  }

  ckr := run(*optPath)
  ckr.Name = "Existence"
  ckr.Exit()
}
