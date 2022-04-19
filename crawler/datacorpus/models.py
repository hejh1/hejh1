from django.db import models
from petthoughts.settings import DATABASE_CONFIG

import pyrebase
import json

firebase = pyrebase.initialize_app(DATABASE_CONFIG)
auth = firebase.auth()
database = firebase.database()
storage = firebase.storage()


