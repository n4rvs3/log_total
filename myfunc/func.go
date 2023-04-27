package myfunc

func BoolSudo(sudo bool, cmd string) string {
	ionice := "ionice -c2 -n7 nice -n19 "
	command := ionice + cmd
	if sudo {
		command := ionice + "sudo" + " " + cmd
		return command
	}
	return command
}
