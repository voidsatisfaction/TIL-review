docker run -v $PWD/cmd:/go/src/til-review-app/cmd -v $PWD/template:/go/src/til-review-app/template -v $PWD/script:/go/src/til-review-app/script til-review:1.0.0 ./everydayReview cmd/config.json template/everyday_review.html.tmpl
