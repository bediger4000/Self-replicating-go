narcissist: narcissist.go
	go build narcissist.go

narcissist.go: gen_narcissist
	./gen_narcissist > narcissist.go 

gen_narcissist: gen_narcissist.go
	-rm -f narcissist narcissist.go
	go build gen_narcissist.go

clean:
	-rm -rf gen_narcissist narcissist narcissist.go
