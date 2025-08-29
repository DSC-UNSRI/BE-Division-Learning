import type { EventResponse, Event } from "../types/event";
import api from "./api";

export async function getEvent(): Promise<Event[]> {
  const res = await api.get<EventResponse>("/event");
  return res.data.events;
}

export async function postEvent(
  location: string,
  startpost?: string,
  cover?: File
): Promise<{ message: string }> {
  const formData = new FormData();
  formData.append("location", location);
  if (startpost) formData.append("startpost", startpost);
  if (cover) formData.append("cover", cover);

  const res = await api.post<{ message: string }>("/event", formData, {
    headers: {
      "Content-Type": "multipart/form-data",
    },
  });
  return res.data;
}

export async function updateEvent(
  id: number,
  values: { location?: string; startpost?: string; cover?: File }
): Promise<{ message: string }> {
  const formData = new FormData();

  if (values.location) formData.append("location", values.location);

  if (values.startpost) {
    const isoDate = new Date(values.startpost).toISOString();
    formData.append("startpost", isoDate);
  }

  if (values.cover) formData.append("cover", values.cover);

  const res = await api.put<{ message: string }>(`/event/${id}`, formData, {
    headers: {
      "Content-Type": "multipart/form-data",
      Authorization: `Bearer ${localStorage.getItem("token") || ""}`, // kalau pakai JWT
    },
  });

  return res.data;
}


export async function deleteEvent(id: number): Promise<{ message: string }> {
  const res = await api.delete<{ message: string }>(`/event/${id}`);
  return res.data;
}
