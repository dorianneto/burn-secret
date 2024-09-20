import React from "react";
import { useAuth } from "../contexts/AuthProvider";
import { Navigate, Outlet } from "react-router-dom";

export const ProtectedRoute = () => {
  const { token } = useAuth();

  if (!token) {
    return <Navigate to="/" />;
  }

  return <Outlet />;
};
