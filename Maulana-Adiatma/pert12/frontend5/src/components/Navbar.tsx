import { useState, useEffect } from "react";
import { useUser } from "../hooks/user";
import { Notification } from "./Notification";
import { useLogout } from "../hooks/auth";

export default function Navbar() {
  const [menuOpen, setMenuOpen] = useState(false);
  const { user, userLoading } = useUser();
  const { handleLogout, logoutLoading, logoutError, logoutSuccess } =
    useLogout();
  const [showError, setShowError] = useState(false);
  const [showSuccess, setShowSuccess] = useState(false);
  const logout = async () => {
    await handleLogout();
    if (logoutSuccess) {
      window.location.href = "/login";
    }
  };

  useEffect(() => {
    if (logoutError) setShowError(true);
  }, [logoutError]);

  useEffect(() => {
    if (logoutSuccess) setShowSuccess(true);
  }, [logoutSuccess]);

  if (userLoading) return null;

  return (
    <>
      {logoutError && showError && (
        <Notification
          type="error"
          message={logoutError.message}
          onClose={() => setShowError(false)}
          duration={1000}
        />
      )}

      {logoutSuccess && showSuccess && (
        <Notification
          type="success"
          message={logoutSuccess}
          onClose={() => setShowSuccess(false)}
          onComplete={() => (window.location.href = "/login")}
          duration={1000}
        />
      )}

      <nav className="bg-red-600 text-white p-4 flex justify-between items-center">
        <h1 className="font-bold text-lg">Nobar Dashboard</h1>
        <ul className="flex gap-4 items-center">
          <li>
            <a href="/" className="hover:underline">
              Dashboard
            </a>
          </li>

          {user ? (
            <>
              <li>
                <a href="/profile" className="hover:underline">
                  Profile
                </a>
              </li>
              <li className="relative">
                <div
                  onClick={() => setMenuOpen(!menuOpen)}
                  className="flex items-center gap-2 bg-red-700 hover:bg-red-800 px-3 py-1 rounded cursor-pointer"
                >
                  <img
                    src={user.profile_picture}
                    alt="User Profile Picture"
                    className="w-8 h-8 rounded-full border-2 border-white"
                  />
                  <span>{user.name}</span>
                </div>

                {menuOpen && (
                  <div className="absolute right-0 mt-2 w-40 bg-white text-black rounded shadow-lg">
                    <button
                      className="w-full text-left px-4 py-2 hover:bg-gray-100"
                      onClick={logout}
                      disabled={logoutLoading}
                    >
                      {logoutLoading ? "Logging out..." : "Logout"}
                    </button>
                  </div>
                )}
              </li>
            </>
          ) : (
            <>
              <li>
                <a href="/login" className="hover:underline">
                  Login
                </a>
              </li>
              <li>
                <a href="/signup" className="hover:underline">
                  Sign Up
                </a>
              </li>
            </>
          )}
        </ul>
      </nav>
    </>
  );
}
