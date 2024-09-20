import React from "react";
import { createRoot } from "react-dom/client";
import { ChakraProvider } from "@chakra-ui/react";
import AuthProvider, { useAuth } from "./contexts/AuthProvider";
import { Routes } from "./components/Routes";

createRoot(document.getElementById("app")!).render(
  <React.StrictMode>
    <div className="bg-slate-900 font-mono text-gray-300">
      <div className="container mx-auto flex flex-col justify-center text-center gap-6 h-screen">
        <ChakraProvider>
          <AuthProvider>
            <Routes />
          </AuthProvider>
        </ChakraProvider>
      </div>
    </div>
  </React.StrictMode>
);
