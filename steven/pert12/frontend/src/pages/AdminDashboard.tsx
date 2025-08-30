import { useState } from "react";
import type { Event } from "../types/event";
import {
  useDeleteEvent,
  useGetEvent,
  usePostEvent,
  useUpdateEvent,
} from "../hooks/event";
import { Notification } from "../components/Notification";

export default function AdminEventDashboard() {
  const { events, loading, error, refetch } = useGetEvent();
  const { doPostEvent, loading: posting, message, errorPost } = usePostEvent();
  const {
    doUpdateEvent,
    loading: updating,
    message: updateMessage,
    errorUpdate,
  } = useUpdateEvent();

  const [showAddModal, setShowAddModal] = useState(false);
  const [showEditModal, setShowEditModal] = useState<null | Event>(null);
  const [newLocation, setNewLocation] = useState("");
  const [showErrorNotification, setShowErrorNotification] = useState(false);
  const [showSuccessNotification, setShowSuccessNotification] = useState(false);
  const [showErrorDeleteNotification, setShowErrorDeleteNotification] =
    useState(false);
  const [showDeleteNotification, setShowDeleteNotification] = useState(false);
  const { messageDelete, errorDelete, doDeleteEvent } = useDeleteEvent();

  const handleAdd = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!newLocation.trim()) return;

    try {
      await doPostEvent(newLocation);
      setShowAddModal(false);
      setShowSuccessNotification(true);
      setNewLocation("");
      refetch();
    } catch {
      setShowErrorNotification(true);
    }
  };

  const handleUpdate = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    if (!showEditModal) return;

    const formData = new FormData(e.currentTarget);

    try {
      const location = formData.get("location") as string;
      const start = formData.get("start") as string;
      const cover = formData.get("cover") as File;
      await doUpdateEvent(showEditModal.id, location, start, cover);
      setShowEditModal(null);
      setShowSuccessNotification(true);
      refetch();
    } catch {
      setShowErrorNotification(true);
    }
  };

  return (
    <div className="p-8 space-y-8">
      {showErrorNotification && (errorPost || errorUpdate) && (
        <Notification
          type="error"
          message={errorPost || errorUpdate || ""}
          onClose={() => setShowErrorNotification(false)}
          duration={2000}
        />
      )}

      {showSuccessNotification && (message || updateMessage) && (
        <Notification
          type="success"
          message={message || updateMessage || ""}
          onClose={() => setShowSuccessNotification(false)}
          duration={2000}
        />
      )}

      <h1 className="text-3xl font-bold">Admin Dashboard - Events</h1>

      <div className="bg-white shadow p-4 rounded">
        <h2 className="text-xl font-semibold mb-4 flex justify-between items-center">
          Events
          <button
            onClick={() => setShowAddModal(true)}
            className="bg-blue-500 text-white px-3 py-1 rounded hover:bg-blue-600"
          >
            Add Event
          </button>
        </h2>

        {loading && <p>Loading events...</p>}
        {error && <p className="text-red-500">{error}</p>}

        {!loading && !error && events && (
          <table className="w-full border">
            <thead>
              <tr>
                <th className="border px-2 py-1">Cover</th>
                <th className="border px-2 py-1">Location</th>
                <th className="border px-2 py-1">Start</th>
                <th className="border px-2 py-1">Actions</th>
              </tr>
            </thead>
            <tbody>
              {Array.isArray(events) &&
                events.map((ev) => (
                  <tr key={ev.id}>
                    <td className="border px-2 py-1">
                      <img
                        src={ev.cover}
                        alt={ev.location}
                        className="w-12 h-12 object-cover rounded"
                      />
                    </td>
                    <td className="border px-2 py-1">{ev.location}</td>
                    <td className="border px-2 py-1">{ev.start}</td>
                    <td className="border px-2 py-1 space-x-2">
                      <button
                        onClick={() => setShowEditModal(ev)}
                        className="bg-yellow-500 text-white px-2 py-1 rounded hover:bg-yellow-600"
                      >
                        Update
                      </button>
                      <button
                        className="bg-red-500 text-white px-2 py-1 rounded hover:bg-red-600"
                        disabled={loading}
                        onClick={async () => {
                          try {
                            await doDeleteEvent(ev.id);
                            setShowDeleteNotification(true); 
                          } catch {
                            setShowErrorDeleteNotification(true);
                          }
                        }}
                      >
                        Delete
                      </button>

                      {messageDelete && showDeleteNotification && (
                        <Notification
                          type="success"
                          message={messageDelete}
                          onClose={() => setShowDeleteNotification(false)}
                          onComplete={refetch}
                          duration={2000}
                        />
                      )}

                      {errorDelete && showErrorDeleteNotification && (
                        <Notification
                          type="error"
                          message={errorDelete}
                          onClose={() => setShowErrorDeleteNotification(false)}
                          onComplete={refetch}
                          duration={2000}
                        />
                      )}
                    </td>
                  </tr>
                ))}
            </tbody>
          </table>
        )}
      </div>

      {showAddModal && (
        <div className="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50">
          <div className="bg-white rounded p-6 w-96">
            <h2 className="text-xl font-semibold mb-4">Add Event</h2>
            <form className="space-y-4" onSubmit={handleAdd}>
              <div>
                <label className="block mb-1">Location</label>
                <input
                  type="text"
                  value={newLocation}
                  onChange={(e) => setNewLocation(e.target.value)}
                  className="w-full border rounded px-3 py-2"
                  placeholder="Enter location"
                />
              </div>
              <div className="flex justify-end space-x-2">
                <button
                  type="button"
                  onClick={() => setShowAddModal(false)}
                  className="px-3 py-1 border rounded"
                >
                  Cancel
                </button>
                <button
                  type="submit"
                  disabled={posting}
                  className="bg-blue-500 text-white px-3 py-1 rounded hover:bg-blue-600 disabled:opacity-50"
                >
                  {posting ? "Saving..." : "Save"}
                </button>
              </div>
            </form>
          </div>
        </div>
      )}

      {showEditModal && (
        <div className="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50">
          <div className="bg-white rounded p-6 w-96">
            <h2 className="text-xl font-semibold mb-4">Update Event</h2>
            <form className="space-y-4" onSubmit={handleUpdate}>
              <div>
                <label className="block mb-1">Location</label>
                <input
                  type="text"
                  name="location"
                  defaultValue={showEditModal.location}
                  className="w-full border rounded px-3 py-2"
                />
              </div>
              <div>
                <label className="block mb-1">Start</label>
                <input
                  type="date"
                  name="start"
                  defaultValue={showEditModal.start?.split("T")[0]}
                  className="w-full border rounded px-3 py-2"
                />
              </div>
              <div>
                <label className="block mb-1">Cover</label>
                <input
                  type="file"
                  name="cover"
                  accept="image/*"
                  className="w-full border rounded px-3 py-2"
                />
              </div>
              <div className="flex justify-end space-x-2">
                <button
                  type="button"
                  onClick={() => setShowEditModal(null)}
                  className="px-3 py-1 border rounded"
                >
                  Cancel
                </button>
                <button
                  type="submit"
                  disabled={updating}
                  className="bg-yellow-500 text-white px-3 py-1 rounded hover:bg-yellow-600 disabled:opacity-50"
                >
                  {updating ? "Updating..." : "Update"}
                </button>
              </div>
            </form>
          </div>
        </div>
      )}
    </div>
  );
}
