import React from "react";
import { createRoot } from "react-dom/client";
import { ChakraProvider } from "@chakra-ui/react";
import AuthProvider from "./contexts/AuthProvider";
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

        <footer>
          <small className="text-gray-500">
            Developed by{" "}
            <a href="https://dorianneto.com" target="_blank">
              Dorian Neto
            </a>{" "}
            - v1.0.0
          </small>
        </footer>
      </div>
    </div>
  </React.StrictMode>
);
