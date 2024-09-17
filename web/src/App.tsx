import React from "react";
import { createRoot } from "react-dom/client";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import { ChakraProvider } from "@chakra-ui/react";
import { NewSecret } from "./pages/NewSecret";
import { RevealSecret } from "./pages/RevealSecret";
import { action as homeAction, Home } from "./pages/Home";

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
]);

createRoot(document.getElementById("app")!).render(
  <React.StrictMode>
    <div className="bg-slate-900 font-mono text-gray-300">
      <div className="container mx-auto flex flex-col justify-center text-center gap-6 h-screen">
        <ChakraProvider>
          <RouterProvider router={router} />
        </ChakraProvider>
      </div>
    </div>
  </React.StrictMode>
);
