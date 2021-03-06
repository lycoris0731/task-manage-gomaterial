package main

type List struct{}

func (f *List) Help() string {
	return "app List"
}

func (f *List) Run(args []string) int {
	roop := root.GetListTasks()
	for _, v := range roop {
		TaskPrint(root.troot + v)
	}
	return 0
}

func (f *List) Synopsis() string {
	return "display of all tasks"
}
