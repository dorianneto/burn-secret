import React from "react";
import { ProtectedRoute } from "./ProtectedRoute";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import { Home, action as homeAction } from "../pages/Home";
import { NewSecret } from "../pages/NewSecret";
import { RevealSecret } from "../pages/RevealSecret";
import { useAuth } from "../contexts/AuthProvider";

export const Routes = () => {
  const { token } = useAuth();

  const protectedRoutes = token
    ? {
        path: "/",
        element: <ProtectedRoute />,
        children: [
          {
            path: "/profile",
            element: <div>Profile Profile</div>,
          },
          {
            path: "/settings",
            element: <div>Settings Profile</div>,
          },
        ],
      }
    : {};

  const router = createBrowserRouter([
    {
      path: "/",
      element: <Home />,
      action: homeAction,
    },
    {
      path: "/secret/new",
      element: <NewSecret />,
    },
    {
      path: "/secret/:id/reveal",
      element: <RevealSecret />,
    },
    // protectedRoutes,
  ]);

  return <RouterProvider router={router} />;
};
