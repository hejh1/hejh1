import datacorpus
from django.shortcuts import render, redirect
from django.http import HttpResponse, HttpResponseRedirect
from django.template import Context, context, loader
from django.urls import reverse

from .models import auth, database, storage

def index(request):
	data = database.get()
	data_content = []

	for i, entry in enumerate(data):
		video_id = entry.key()
		video_metadata = entry.val()
		video_reviewed = video_metadata['video_reviewed']
		data_content.append(tuple((video_id, video_reviewed)))
	
	content = {
		"data_content": data_content
	}
	
	return render(request, 'home.html', content)

def get_data(request):
	video_id = request.GET['video_id']
	video_metadata = (database.child(video_id).get()).val()
	video_user_ID = video_metadata['user_ID']
	video_hashtags = video_metadata['video_hashtags']
	video_reviewed = video_metadata['video_reviewed']
	video_text = video_metadata['video_text']
	video_tiktok_url = video_metadata['video_url']
	
	content = {
		"video_id": video_id,
		"video_user_ID": video_user_ID,
		"video_tiktok_url": video_tiktok_url,
		"video_text": video_text,
		"video_hashtags": video_hashtags
	}
	
	return render(request, 'index.html', content)

def save_data(request):
	
	request_data = request.POST
	video_id = request_data['video_id']
	video_comments = []
	for key, value in request_data.items():
		if "comment" in key:
			video_comments.append(value)
	database.child(video_id).child("video_comments").set(video_comments)
	database.child(video_id).child("video_reviewed").set(True)
	return redirect('datacorpus:index')

def delete_data(request):
	
	request_data = request.POST
	video_id = request_data['video_id']
	database.child(video_id).child("video_reviewed").set(True)
	# database.child(video_id).remove()
	# storage.delete('tiktok_data/'+ video_id + '.mp4', user["idToken"])
	return redirect('datacorpus:index')
