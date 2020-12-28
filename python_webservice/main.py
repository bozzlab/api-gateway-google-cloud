from datetime import date
from fastapi import FastAPI, Header
from starlette.responses import JSONResponse
from starlette.responses import Response
from starlette.requests import Request

import base64
import time
import json
import uvicorn
import requests
import logging 

import os 

COVID_API_KEY = os.environ['COVID_API_KEY']

app = FastAPI()

@app.get("/info")
async def info():
    return JSONResponse({ 
        "serviceName" : "Covid-19 API",
        "language" : "Python",
        "framework" : "FastAPI"
    } , status_code = 200)

@app.get("/covid-19")
async def get_all():
    url = "https://covid-19-data.p.rapidapi.com/totals"
    headers = {
    'x-rapidapi-key': COVID_API_KEY,
    'x-rapidapi-host': "covid-19-data.p.rapidapi.com"
    }
    response = requests.get(url, headers=headers)
    return response.json()

@app.get("/covid-19/th")
async def get_th():
    current_date = date.today()
    today = current_date.strftime("%Y-%m-%d")
    url = f"https://covid-19-data.p.rapidapi.com/report/country/name?date={today}&name=Thailand"

    headers = {
        'x-rapidapi-key': COVID_API_KEY,
        'x-rapidapi-host': "covid-19-data.p.rapidapi.com"
        }

    response = requests.get(url, headers=headers)
    return JSONResponse(response.json(),status_code = 200)


if __name__ == "__main__":
    uvicorn.run("main:app",host='0.0.0.0',port=8080)