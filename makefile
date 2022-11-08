narcissist: narcissist.go
	go build narcissist.go

narcissist.go: gen_narcissist
	./gen_narcissist > narcissist.go 

gen_narcissist: gen_narcissist.go
	-rm -f narcissist narcissist.go
	go build gen_narcissist.go

almost_narcissist: almost_narcissist.go
	go build almost_narcissist.go

almost_narcissist.go: gen_almost_narcissist
	./gen_almost_narcissist > almost_narcissist.go

gen_almost_narcissist: gen_almost_narcissist.go
	-rm -f gen_almost_narcissist
	go build gen_almost_narcissist.go

self_encrypting: gen_self_encrypting
	./gen_self_encrypting > self_encrypting.b64
	base64 -d  self_encrypting.b64 > self_encrypting.go
	go build self_encrypting.go

gen_self_encrypting: gen_self_encrypting.go
	go build gen_self_encrypting.go


clean:
	-rm -rf gen_narcissist narcissist narcissist.go
	-rm -rf gen_almost_narcissist almost_narcissist.go almost_narcissist
	-rm -rf self_encrypting.b64 self_encrypting.go self_encrypting
	-rm -rf gen_self_encrypting
