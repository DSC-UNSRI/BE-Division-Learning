import { BrowserRouter as Router, Routes, Route } from "react-router-dom";

import Dashboard from "./pages/Dashboard";
import MainLayout from "./layouts/MainLayout";
import EventDetail from "./pages/EventDetail";
import BlankLayout from "./layouts/BlankLayout";
import Login from "./pages/Login";
import Signup from "./pages/Signup";
import AdminDashboard from "./pages/AdminDashboard";
import ProfilePage from "./pages/Profile";

export default function App() {
  return (
    <Router>
      <Routes>
        <Route
          path="/"
          element={
            <MainLayout>
              <Dashboard />
            </MainLayout>
          }
        />
        <Route
          path="/login"
          element={
            <BlankLayout>
              <Login />
            </BlankLayout>
          }
        />
        <Route
          path="/signup"
          element={
            <BlankLayout>
              <Signup />
            </BlankLayout>
          }
        />
        <Route
          path="/event/:id"
          element={
            <MainLayout>
              <EventDetail />
            </MainLayout>
          }
        />
        <Route
          path="/admin/dashboard"
          element={
            <MainLayout>
              <AdminDashboard />
            </MainLayout>
          }
        />
        <Route
          path="/profile"
          element={
            <MainLayout>
              <ProfilePage />
            </MainLayout>
          }
        />
      </Routes>
    </Router>
  );
}
