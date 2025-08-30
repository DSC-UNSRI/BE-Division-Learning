import { Link } from "react-router-dom";
import type { EventCardProps } from "../types/event";

export default function EventCard({
  id,
  title,
  date,
  description,
  image,
}: EventCardProps) {
  return (
    <div className="bg-white shadow-md rounded-lg overflow-hidden">
      <img src={image} alt={title} className="w-full h-80 object-cover" />

      <div className="p-4">
        <h3 className="text-lg font-semibold">{title}</h3>
        <p className="text-gray-500 text-sm">{date}</p>
        <p className="mt-2 text-gray-700 text-sm">{description}</p>
        <Link
          to={`/event/${id}`}
          className="mt-4 inline-block bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600"
        >
          Lihat Detail
        </Link>
      </div>
    </div>
  );
}
