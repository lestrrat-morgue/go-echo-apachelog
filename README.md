# go-echo-apachelog

Apache-style logger for [echo](https://github.com/labstack/echo)

## Usage

If you just want to see the access log in STDERR:

```go
  e.Use(apachelog.Logger(os.Stderr))
```

If you want to rotate logs and such, you will need to replace the destination,
so you should keep the ApacheLog struct:

```go
  l := &ApacheLog{}
  l.LogFormat = logformat.CombinedLog.Clone()
  l.LogFormat.SetOutput(dst)

  e.Use(l.Wrap)

  // elsewhere in your code...
  l.LogFormat.SetOutput(newLogDestination)
```

