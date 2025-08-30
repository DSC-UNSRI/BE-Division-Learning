import { deleteEvent, getEvent, postEvent, updateEvent } from "../services/event";
import type { Event } from "../types/event";
import { useState, useEffect } from "react";
import { isAxiosError } from "../utils/axios";

export function useGetEvent() {
  const [events, setEvents] = useState<Event[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const fetchEvent = async () => {
    setLoading(true);
    setError(null);

    try {
      const data = await getEvent();
      setEvents(data);
    } catch (err: unknown) {
      if (isAxiosError(err)) {
        setError(err.response?.data?.message || err.message);
      } else {
        setError("Unknown error");
      }
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchEvent();
  }, []);

  return { events, loading, error, refetch: fetchEvent };
}

export function usePostEvent() {
  const [loading, setLoading] = useState(false);
  const [errorPost, setErrorPost] = useState<string | null>(null);
  const [message, setMessage] = useState<string | null>(null);

  const doPostEvent = async (location: string) => {
    setLoading(true);
    setErrorPost(null);
    setMessage(null);

    try {
      const data = await postEvent(location);
      setMessage(data.message);
      return data;
    } catch (err: unknown) {
      if (isAxiosError(err)) {
        setErrorPost(err.response?.data?.message || err.message);
      } else {
        setErrorPost("Unknown error");
      }
      throw err;
    } finally {
      setLoading(false);
    }
  };

  return { loading, errorPost, message, doPostEvent };
}

export function useUpdateEvent() {
  const [loading, setLoading] = useState(false);
  const [errorUpdate, setErrorUpdate] = useState<string | null>(null);
  const [message, setMessage] = useState<string | null>(null);

  const doUpdateEvent = async (
    id: number,
    location: string,
    start: string,
    cover: File | null
  ) => {
    setLoading(true);
    setErrorUpdate(null);
    setMessage(null);

    try {
      const data = await updateEvent(id, { location, start, cover: cover ?? undefined });
      setMessage(data.message);
      return data;
    } catch (err: unknown) {
      if (isAxiosError(err)) {
        setErrorUpdate(err.response?.data?.message || err.message);
      } else {
        setErrorUpdate("Unknown error");
      }
      throw err;
    } finally {
      setLoading(false);
    }
  };

  return { loading, errorUpdate, message, doUpdateEvent };
}

export function useDeleteEvent() {
  const [loading, setLoading] = useState(false);
  const [errorDelete, setErrorDelete] = useState<string | null>(null);
  const [messageDelete, setMessage] = useState<string | null>(null);

  const doDeleteEvent = async (id: number) => {
    setLoading(true);
    setErrorDelete(null);
    setMessage(null);

    try {
      const data = await deleteEvent(id);
      setMessage(data.message);
      return data;
    } catch (err: unknown) {
      if (isAxiosError(err)) {
        setErrorDelete(err.response?.data?.message || err.message);
      } else {
        setErrorDelete("Unknown error");
      }
      throw err;
    } finally {
      setLoading(false);
    }
  };

  return { loading, errorDelete, messageDelete, doDeleteEvent };
}