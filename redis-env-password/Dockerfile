FROM redis:5-alpine
EXPOSE 6379
#The cmd here cannot be passed as an array because the bash env value will not be interpeted as expected
CMD redis-server --requirepass $REDIS_PASSWORD
