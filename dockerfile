FROM golang

##buat folder APP
RUN mkdir /app

##set direktori utama
WORKDIR /app

##copy seluruh folder challenge-2 ke app
ADD ./challenge-2 ./app

##buat executeable
RUN go build -o main .

##jalankan executeable
CMD ["/app/main"]
