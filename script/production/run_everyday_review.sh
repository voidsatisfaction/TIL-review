# It uses absolute path as this script is called by crontab
sudo docker run --rm -v /home/ubuntu/TIL-review:/go/src/til-review-app til-review:1.0.0 go run cmd/everydayReview.go cmd/config.json template/everyday_review.html
