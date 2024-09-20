import axios from "axios";
import React, { createContext, useContext, useEffect, useState } from "react";

type AuthContextParams = {
  token: string;
  setToken: (token: string) => void;
};

const AuthContext = createContext<AuthContextParams | undefined>(undefined);

const AuthProvider = ({ children }) => {
  const [token, setToken] = useState(localStorage.getItem("token") ?? "");

  useEffect(() => {
    if (token !== "") {
      axios.defaults.headers.common["Authorization"] = `Bearer ${token}`;
      localStorage.setItem("token", token);
      return;
    }

    delete axios.defaults.headers.common["Authorization"];
    localStorage.removeItem("token");
  }, [token]);

  const contextValues = {
    token,
    setToken,
  };

  return (
    <AuthContext.Provider value={contextValues}>
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = () => {
  const context = useContext(AuthContext);
  if (context === undefined) {
    throw new Error("useAuth must be used within an AuthProvider");
  }

  return context;
};

export default AuthProvider;
