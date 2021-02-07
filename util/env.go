package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var (
	globalEnvs        *Envs
	ErrEmptyKey       = fmt.Errorf("env: empty key")
	ErrValueNotNumber = fmt.Errorf("env: the value is not a number")
)

func init() {
	globalEnvs = NewEnvs()
}

func GetEnv(key string) (string, error) {
	return globalEnvs.Get(key)
}

func MustGetEnv(key string) string {
	return globalEnvs.MustGet(key)
}

func GetIntEnv(key string) (int, error) {
	return globalEnvs.GetInt(key)
}

func MustGetIntEnv(key string) int {
	return globalEnvs.MustGetInt(key)
}

func SetEnv(key, val string) error {
	return globalEnvs.Set(key, val)
}

func MustSetEnv(key, val string) {
	globalEnvs.MustSet(key, val)
}

func SetIntEnv(key string, val int) error {
	return globalEnvs.SetInt(key, val)
}

func MustSetIntEnv(key string, val int) {
	globalEnvs.MustSetInt(key, val)
}

func ReadEnvFile(path string) error {
	return globalEnvs.ReadFile(path)
}

type Envs struct {
}

func NewEnvs() *Envs {
	return &Envs{}
}

func (e *Envs) Get(key string) (string, error) {
	if key == "" {
		return "", ErrEmptyKey
	}
	return os.Getenv(key), nil
}

func (e *Envs) MustGet(key string) string {
	val, err := e.Get(key)
	if err != nil {
		panic(err)
	}
	return val
}

func (e *Envs) GetInt(key string) (int, error) {
	val, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		return -1, ErrValueNotNumber
	}
	return val, nil
}

func (e *Envs) MustGetInt(key string) int {
	val, err := e.GetInt(key)
	if err != nil {
		panic(err)
	}
	return val
}

func (e *Envs) Set(key, val string) error {
	if key == "" {
		return ErrEmptyKey
	}
	if err := os.Setenv(key, val); err != nil {
		return fmt.Errorf("env: %w", err)
	}
	return nil
}

func (e *Envs) MustSet(key, val string) {
	err := e.Set(key, val)
	if err != nil {
		panic(err)
	}
}

func (e *Envs) SetInt(key string, val int) error {
	return e.Set(key, strconv.Itoa(val))
}

func (e *Envs) MustSetInt(key string, val int) {
	err := e.SetInt(key, val)
	if err != nil {
		panic(err)
	}
}

func (e *Envs) ReadFile(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("env read file: %w", err)
	}

	for _, line := range strings.Split(string(data), "\n") {
		tokens := strings.SplitN(line, "=", 2)
		if len(tokens) == 2 {
			key := strings.TrimSpace(tokens[0])
			val := strings.TrimSpace(tokens[1])
			if err := e.Set(key, val); err != nil {
				return fmt.Errorf("env read file: %w", err)
			}
		}
	}
	return nil
}
