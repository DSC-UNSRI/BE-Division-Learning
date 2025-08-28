import type { EventResponse, Event } from "../types/event";
import api from "./api";

export async function getEvent(): Promise<Event[]> {
  const res = await api.get<EventResponse>("/event");
  return res.data.events;
}

export async function postEvent(
  location: string
): Promise<{ message: string }> {
  const formData = new FormData();
  formData.append("location", location);
  const res = await api.post<{ message: string }>("/event", formData);
  return res.data;
}

export async function updateEvent(
  id: number,
  values: { location?: string; start?: string; cover?: File }
): Promise<{ message: string }> {
  const formData = new FormData();

  if (values.location) formData.append("location", values.location);
  if (values.start) formData.append("start", values.start);
  if (values.cover) formData.append("cover", values.cover);

  const res = await api.patch<{ message: string }>(`/event/${id}`, formData, {
    headers: { "Content-Type": "multipart/form-data" },
  });

  return res.data;
}

export async function deleteEvent(id: number): Promise<{ message: string }> {
  const res = await api.delete<{ message: string }>(`/event/${id}`);
  return res.data;
}
