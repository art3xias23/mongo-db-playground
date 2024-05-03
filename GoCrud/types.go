
package main

type DbConfig struct {
	DbPass string `yaml:"dbPass"`
	DbUser string `yaml:"dbUser"`
}

type Animal struct {
	Id   int
	Type string
	Name string
}
