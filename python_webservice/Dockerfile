FROM python:3.6-slim
ADD . / 
ENV PYTHONUNBUFFERED=True
RUN apt-get update \ 
    && apt-get install gcc -y \
    && pip install --no-cache-dir -r requirements.txt \
    && rm -rf /var/lib/apt/lists/* 
CMD ["python","main.py"]