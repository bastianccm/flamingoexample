.phony: jaeger serve

jaeger:
	docker run --rm -p 16686:16686 -p 14268:14268 -p 9411:9411 jaegertracing/all-in-one:1.13

serve:
	go run . serve
