docker run -v /home/ubuntu/TIL-review/cmd:/go/src/til-review-app/cmd -v /home/ubuntu/TIL-review/template:/go/src/til-review-app/template -v /home/ubuntu/TIL-review/script:/go/src/til-review-app/script til-review:1.0.0 ./everydayReview cmd/config.json template/everyday_review.html