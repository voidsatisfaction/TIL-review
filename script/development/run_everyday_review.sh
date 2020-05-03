docker run --rm --name til-review-run-everyday-review -v $PWD:/go/src/til-review-app til-review:1.0.0 go run cmd/everydayReview.go cmd/config.json template/everyday_review.html.tmpl
