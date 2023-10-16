from django.db import models


class Note(models.Model):
    title = models.CharField(max_length=255)
    contents = models.CharField(max_length=255)
