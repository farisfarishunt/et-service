package logger

import(
    "log"
)

// Very simple logger wrapper :)
func Log(err error) {
    log.Println(err)
}
