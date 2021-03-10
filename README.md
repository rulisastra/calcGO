# [calcGO](https://github.com/rulisastra/calcGO)

on going...

- [ ] Flowchart
- [ ] Screenshoot
- [ ] Useful resources and blogs
    - [first youtube](https://www.youtube.com/watch?v=QAwXt-zE7so)

---

### How to use this repo :relaxed:

1. clone this repo
2. run `go run ./calcGo.exe`
---

## flowchart

- when it panicked because divided by zero, do this
```
	// handle panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Ran into an error")
			main()
		}
	}()
```
