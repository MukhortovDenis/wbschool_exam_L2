package main

import "testing"


func Test_shell(t *testing.T){
	slice := []string{
		"cd ..", "pwd", "echo test check", "ps"}
	t.Run("проверка шелла", func(t *testing.T){
		for _, str := range slice{
			err := shell(str)
			if err != nil{
				t.Error(err)
			}
		}
	})
}