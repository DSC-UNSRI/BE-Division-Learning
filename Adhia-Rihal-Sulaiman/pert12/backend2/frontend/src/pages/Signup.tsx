import { useState } from "react";
import { useSignup } from "../hooks/auth";
import { useNavigate } from "react-router-dom";
import { Notification } from "../components/Notification";

export default function Signup() {
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [role] = useState<"user" | "admin">("user");
  const [passwordError, setPasswordError] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");
  const { loading, message, error, doSignup} = useSignup();
  const [showErrorNotification, setShowErrorNotification] = useState(false);
  const [showSuccessNotification, setShowSuccessNotification] = useState(false);
  const navigate = useNavigate();

  const handleSignup = async (e: React.FormEvent) => {
    e.preventDefault();
    if (password !== confirmPassword) {
      setShowErrorNotification(true);
      setPasswordError("Kata sandi tidak sama, mohon ulangi");
      return;
    }
    try {
      await doSignup(name, email, password, role);
      setShowSuccessNotification(true);
    } catch {
      setShowErrorNotification(true);
    }
  };

  return (
    <div className="flex justify-center items-center min-h-screen bg-gray-100">
      <div className="bg-white shadow-md rounded-lg p-8 w-full max-w-md">
        <h2 className="text-2xl font-bold text-center mb-6">Signup</h2>
        <form onSubmit={handleSignup} className="space-y-4">
          <input
            type="text"
            placeholder="Nama Lengkap"
            value={name}
            onChange={(e) => setName(e.target.value)}
            className="w-full px-4 py-2 border rounded focus:outline-none focus:ring-2 focus:ring-red-500"
            required
          />
          <input
            type="email"
            placeholder="Email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            className="w-full px-4 py-2 border rounded focus:outline-none focus:ring-2 focus:ring-red-500"
            required
          />
          <input
            type="password"
            placeholder="Password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            className="w-full px-4 py-2 border rounded focus:outline-none focus:ring-2 focus:ring-red-500"
            required
          />
          <input
            type="password"
            placeholder="Konfirmasi Password"
            value={confirmPassword}
            onChange={(e) => setConfirmPassword(e.target.value)}
            className="w-full px-4 py-2 border rounded focus:outline-none focus:ring-2 focus:ring-red-500"
            required
          />
          <button
            type="submit"
            className="w-full bg-red-600 text-white py-2 rounded hover:bg-red-700"
          >
            {loading ? "Loading..." : "Signup"}
          </button>
        </form>
        {(error || passwordError) && showErrorNotification && (
          <Notification
            type="error"
            message={error || passwordError}
            onClose={() => setShowErrorNotification(false)}
            duration={1000}
          />
        )}

        {message && showSuccessNotification && (
          <Notification
            type="success"
            message={message}
            onClose={() => setShowSuccessNotification(false)}
            onComplete={() => {navigate("/");}}
            duration={3000}
          />
        )}
        <p className="text-center text-gray-600 mt-4">
          Sudah punya akun?{" "}
          <a href="/login" className="text-red-600 hover:underline">
            Login
          </a>
        </p>
      </div>
    </div>
  );
}
