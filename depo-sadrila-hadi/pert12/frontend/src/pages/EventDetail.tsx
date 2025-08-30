import { useParams } from "react-router-dom";
import { useState } from "react";

export default function EventDetail() {
  const { id } = useParams();
  const [selectedSeats, setSelectedSeats] = useState<string[]>([]);

  const seatData = [
    { id: "A1", label: "A1", booked: false },
    { id: "A2", label: "A2", booked: false },
    { id: "A3", label: "A3", booked: true },
    { id: "A4", label: "A4", booked: false },
    { id: "B1", label: "B1", booked: false },
    { id: "B2", label: "B2", booked: true },
    { id: "B3", label: "B3", booked: false },
    { id: "B4", label: "B4", booked: false },
    { id: "C1", label: "C1", booked: false },
    { id: "C2", label: "C2", booked: false },
    { id: "C3", label: "C3", booked: false },
    { id: "C4", label: "C4", booked: true },
  ];

  const toggleSeat = (seatId: string) => {
    setSelectedSeats(prev =>
      prev.includes(seatId)
        ? prev.filter(s => s !== seatId)
        : [...prev, seatId]
    );
  };

  return (
    <div className="max-w-4xl mx-auto p-6">
      <h2 className="text-2xl font-semibold mb-4 text-center">
        Event ID: {id}
      </h2>
      <p className="text-gray-700 mb-6 text-center">
        Pilih kursi untuk event ini:
      </p>

      <div className="grid grid-cols-4 sm:grid-cols-8 gap-2 justify-items-center">
        {seatData.map(seat => (
          <button
            key={seat.id}
            onClick={() => toggleSeat(seat.id)}
            disabled={seat.booked}
            className={`w-12 h-12 flex items-center justify-center rounded font-medium transition ${
              seat.booked
                ? "bg-red-500 text-white cursor-not-allowed"
                : selectedSeats.includes(seat.id)
                ? "bg-green-500 text-white"
                : "bg-gray-200 hover:bg-gray-300"
            }`}
          >
            {seat.label}
          </button>
        ))}
      </div>

      {selectedSeats.length > 0 && (
        <div className="mt-6 text-center">
          <p className="mb-2">
            Kursi dipilih:{" "}
            <span className="font-semibold">
              {selectedSeats.join(", ")}
            </span>
          </p>
          <button className="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700">
            Konfirmasi
          </button>
        </div>
      )}
    </div>
  );
}
