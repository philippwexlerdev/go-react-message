import React, { useState } from "react";
import {
  Container,
  Paper,
  TextField,
  Button,
  Typography,
  Box,
} from "@mui/material";
import { useLocalStorage } from "usehooks-ts";
import { Link } from "react-router-dom";

import { useLoginMutation } from "../../mutations/useLoginMutation";

const Login = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  const [_, setToken] = useLocalStorage("token", "");

  const { mutateAsync } = useLoginMutation();

  const handleUsernameChange = (e: React.ChangeEvent<HTMLInputElement>) =>
    setUsername(e.target.value);

  const handlePasswordChange = (e: React.ChangeEvent<HTMLInputElement>) =>
    setPassword(e.target.value);

  const handleSubmit = async (
    e: React.MouseEvent<HTMLElement> | React.FormEvent
  ) => {
    e.preventDefault();
    const loginData = await mutateAsync({ username, password: btoa(password) });
    setToken(
      JSON.stringify({
        token: loginData.data.token,
        username,
        user_id: loginData.data.user_id,
      })
    );
  };

  return (
    <Container sx={{ padding: 2 }}>
      <Paper elevation={3} sx={{ padding: 2 }}>
        <Typography variant="h3" gutterBottom>
          Login
        </Typography>
        <form onSubmit={handleSubmit}>
          <Box
            sx={{
              display: "flex",
              flexDirection: "column",
              gap: 3,
              marginTop: 4,
            }}
          >
            <TextField
              onChange={handleUsernameChange}
              value={username}
              label={"Username"}
            />
            <TextField
              onChange={handlePasswordChange}
              value={password}
              type="password"
              label={"Password"}
            />
          </Box>

          <Box sx={{ gap: 1, display: "flex", paddingTop: 2 }}>
            <Button onClick={handleSubmit} variant="contained">
              Login
            </Button>
            <Link to="/signup">
              <Button variant="text">Signup Instead</Button>
            </Link>
          </Box>
        </form>
      </Paper>
    </Container>
  );
};

export default Login;
