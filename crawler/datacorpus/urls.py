from django.urls import path
from django.contrib.staticfiles.urls import staticfiles_urlpatterns
from django.conf.urls.static import static

from . import views

app_name = "datacorpus"

urlpatterns = [
	path('', views.index, name = "index"),
	path('data/', views.get_data, name = "get_data"),
	path('save/', views.save_data, name = "save_data"),
	path('delete/', views.delete_data, name = "delete_data")
]
