import { useEffect, useState } from "react";

type NotificationProps = {
  type: "success" | "error" | "warning" | "info";
  message: string;
  onClose: () => void;
  onComplete?: () => void;
  duration?: number;
};

export function Notification({
  type,
  message,
  onClose,
  onComplete,
  duration = 1000,
}: NotificationProps) {
  const colors = {
    success: "bg-green-500",
    error: "bg-red-500",
    warning: "bg-yellow-500 text-black",
    info: "bg-blue-500",
  };

  const [progress, setProgress] = useState(100);

  useEffect(() => {
    const interval = 50;
    const step = 100 / (duration / interval);

    const timer = setInterval(() => {
      setProgress((prev) => {
        if (prev <= 0) {
          clearInterval(timer);
          onClose();
          if (onComplete) onComplete();
          return 0;
        }
        return prev - step;
      });
    }, interval);

    return () => clearInterval(timer);
  }, [duration, onClose, onComplete]);

  return (
    <div className="fixed top-4 left-1/2 -translate-x-1/2 z-50 w-[50%] max-w-md">
      <div
        className={`${colors[type]} text-white px-4 py-2 rounded shadow flex flex-col gap-2`}
      >
        <div className="flex justify-between items-center">
          <span>{message}</span>
          <button onClick={onClose} className="ml-2 font-bold">
            ×
          </button>
        </div>

        <div className="w-full bg-white/30 h-1 rounded">
          <div
            className="bg-white h-1 rounded transition-all duration-100 ease-linear"
            style={{ width: `${progress}%` }}
          />
        </div>
      </div>
    </div>
  );
}
