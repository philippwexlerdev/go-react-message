import React, { useEffect, useMemo } from "react";
import Box from "@mui/material/Box";
import Paper from "@mui/material/Paper";
import Typography from "@mui/material/Typography";
import { useLocalStorage } from "usehooks-ts";
import { useNavigate } from "react-router-dom";
import Button from "@mui/material/Button";
import { getUserInfo } from "../util/getUserInfo";

interface AuthLayoutProps {
  children: React.ReactNode[] | React.ReactNode;
}

const AuthLayout = ({ children }: AuthLayoutProps) => {
  const [token, setToken] = useLocalStorage("token", "");
  const navigate = useNavigate();

  useEffect(() => {
    !token && navigate("/login");
  }, [token, navigate]);

  const userInfo = useMemo(() => getUserInfo(), [token]);

  return (
    <Box>
      <Paper
        sx={{
          borderRadius: 0,
          p: 2,
          display: "flex",
          justifyContent: "flex-end",
          alignItems: "center",
          gap: 2,
        }}
      >
        <Typography variant="body1">{userInfo.username}</Typography>
        <Button onClick={() => setToken("")} variant="contained">
          Logout
        </Button>
      </Paper>
      {children}
    </Box>
  );
};

export default AuthLayout;
