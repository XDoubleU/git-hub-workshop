from rest_framework import viewsets

from .models import Note
from .serializer import NoteSerializer


class NoteViewSet(viewsets.ModelViewSet):
    queryset = Note.objects.all().order_by("-title")
    serializer_class = NoteSerializer
