import { useState } from "react";
import { Notification } from "../components/Notification";
import { useNavigate } from "react-router-dom";
import { useLogin } from "../hooks/auth";

export default function Login() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const { loading, message, error, doLogin, user } = useLogin();
  const [showErrorNotification, setShowErrorNotification] = useState(false);
  const [showSuccessNotification, setShowSuccessNotification] = useState(false);
  const navigate = useNavigate();

  const handleLogin = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      await doLogin(email, password);
      setShowSuccessNotification(true);
    } catch {
      setShowErrorNotification(true);
    }
  };

  return (
    <div className="flex justify-center items-center min-h-screen bg-gray-100">
      <div className="bg-white shadow-md rounded-lg p-8 w-full max-w-md">
        <h2 className="text-2xl font-bold text-center mb-6">Login</h2>
        <form onSubmit={handleLogin} className="space-y-4">
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
          <button
            type="submit"
            disabled={loading}
            className={`w-full py-2 rounded text-white ${
              loading
                ? "bg-gray-400 cursor-not-allowed"
                : "bg-red-600 hover:bg-red-700"
            }`}
          >
            {loading ? "Loading..." : "Login"}
          </button>
        </form>

        {error && showErrorNotification && (
          <Notification
            type="error"
            message={error}
            onClose={() => setShowErrorNotification(false)}
            duration={1000}
          />
        )}

        {message && showSuccessNotification && (
          <Notification
            type="success"
            message={message}
            onClose={() => setShowSuccessNotification(false)}
            onComplete={() => {
              if (user?.role === "admin") {
                navigate("/admin/dashboard");
              } else {
                navigate("/");
              }
            }}
            duration={1000}
          />
        )}

        <p className="text-center text-gray-600 mt-4">
          Belum punya akun?{" "}
          <a href="/signup" className="text-red-600 hover:underline">
            Signup
          </a>
        </p>
      </div>
    </div>
  );
}
