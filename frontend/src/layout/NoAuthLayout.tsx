import React, { useEffect } from "react";
import Box from "@mui/material/Box";
import { useLocalStorage } from "usehooks-ts";
import { useNavigate } from "react-router-dom";

interface NoAuthLayoutProps {
  children: React.ReactNode[] | React.ReactNode;
}

const NoAuthLayout = ({ children }: NoAuthLayoutProps) => {
  const [token, _] = useLocalStorage("token", "");
  const navigate = useNavigate();

  useEffect(() => {
    token && navigate("/");
  }, [token, navigate]);

  return <Box>{children}</Box>;
};

export default NoAuthLayout;
