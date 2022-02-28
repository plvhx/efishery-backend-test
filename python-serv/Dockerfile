FROM python:3.9

WORKDIR /usr/src/app

COPY ./python-serv ./serv
COPY ./cache/ ./cache
COPY ./.jwt_pass ./.jwt_pass

RUN cd ./serv && pip install --no-cache-dir -r requirements.txt

EXPOSE 5000/tcp

WORKDIR ./serv

ENTRYPOINT ["python"]

CMD ["app.py"]
