import React from "react";
import { createRoot } from "react-dom/client";
import { createBrowserRouter, RouterProvider, Link } from "react-router-dom";

const router = createBrowserRouter([
  {
    path: "/",
    element: <div>Hello world! <Link to={"/about"}>about2</Link></div>,
  },
  {
    path: "/about",
    element: <div>About <Link to={"/"}>Home</Link></div>,
  },
]);

createRoot(document.getElementById('app')!).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
);
