import pytest
from django.contrib.auth import get_user_model
from django.test import TestCase
from django.urls import reverse
from rest_framework.test import APIClient

from .models import Note


@pytest.mark.django_db
class NoteViewsTests(TestCase):
    @classmethod
    def setUpTestData(self):
        self.notes = []
        for i in range(0, 10):
            self.notes.append(
                Note.objects.create(title=f"Title{i}", contents="Some text")
            )

        self.client = APIClient()

    def test_notelistview(self):
        response = self.client.get("/notes", follow=True)
        assert response.status_code == 200
        assert len(response.data) == len(self.notes)

    def test_notecreateview(self):
        data = {"title": "NewTitle", "contents": "Some text"}
        response = self.client.post("/notes", data, follow=True)
        assert response.status_code == 200

    def test_noteupdateview(self):
        data = {"title": "NewTitle", "contents": "Some text"}
        response = self.client.put(f"/notes/{self.notes[0].id}", data, follow=True)
        assert response.status_code == 200

    def test_notedeleteview(self):
        response = self.client.delete(f"/notes/{self.notes[0].id}", follow=True)
        assert response.status_code == 200
        assert dict(response.data)["title"] == self.notes[0].title
        assert dict(response.data)["contents"] == self.notes[0].contents
