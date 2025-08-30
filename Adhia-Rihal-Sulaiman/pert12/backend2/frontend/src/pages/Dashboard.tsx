import { useState } from "react";
import EventCard from "../components/EventCard";
import { useGetEvent } from "../hooks/event";

export default function Dashboard() {
  const images = [
    "/assets/carousel-1.webp",
    "/assets/carousel-2.webp",
    "/assets/carousel-3.webp",
  ];

  const [currentIndex, setCurrentIndex] = useState(0);
  const { events, loading, error } = useGetEvent();

  const prevSlide = () => {
    setCurrentIndex((prev) => (prev === 0 ? images.length - 1 : prev - 1));
  };

  const nextSlide = () => {
    setCurrentIndex((prev) => (prev === images.length - 1 ? 0 : prev + 1));
  };

  return (
    <div className="max-w-7xl mx-auto p-6">
      <h2 className="text-2xl font-semibold mb-4 text-center">
        Selamat Datang di Website Nonton Bareng Merah Putih One For All
      </h2>

      <div className="relative w-full overflow-hidden rounded-lg shadow-lg">
        <img
          src={images[currentIndex]}
          alt={`Slide ${currentIndex + 1}`}
          className="w-full h-96 object-cover"
        />
        <button
          onClick={prevSlide}
          className="absolute left-0 top-1/2 -translate-y-1/2 bg-gray-800 bg-opacity-50 text-white px-4 py-2"
        >
          ❮
        </button>
        <button
          onClick={nextSlide}
          className="absolute right-0 top-1/2 -translate-y-1/2 bg-gray-800 bg-opacity-50 text-white px-4 py-2"
        >
          ❯
        </button>
      </div>

      <h3 className="text-xl font-semibold mt-8 mb-4 text-center">
        Daftar Lokasi Nobar
      </h3>

      {loading && <p>Loading events...</p>}
      {error && <p className="text-red-500">{error}</p>}

      {!loading && !error && (
        <div className="grid gap-6 sm:grid-cols-2 lg:grid-cols-3">
          {Array.isArray(events) && events.length > 0 ? (
            events.map((ev) => (
              <EventCard
                key={ev.id}
                id={ev.id}
                title={ev.location}
                date={
                  ev.start
                    ? new Date(ev.start).toLocaleDateString("id-ID", {
                        year: "numeric",
                        month: "long",
                        day: "numeric",
                      })
                    : "-"
                }
                description={`Nobar di ${ev.location}`}
                image={ev.cover}
              />
            ))
          ) : (
            <p className="col-span-full text-center text-gray-600">
              Belum ada event.
            </p>
          )}
        </div>
      )}
    </div>
  );
}
