import React, { useState } from "react";
import {
  Container,
  Paper,
  TextField,
  Button,
  Typography,
  Box,
} from "@mui/material";
import { useNavigate, Link } from "react-router-dom";

import { useSignupMutation } from "../../mutations/useSignupMutation";

const Signup = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  const { mutateAsync, isLoading } = useSignupMutation();

  const navigate = useNavigate();

  const handleUsernameChange = (e: React.ChangeEvent<HTMLInputElement>) =>
    setUsername(e.target.value);

  const handlePasswordChange = (e: React.ChangeEvent<HTMLInputElement>) =>
    setPassword(e.target.value);

  const handleSubmit = async (
    e: React.MouseEvent<HTMLElement> | React.FormEvent
  ) => {
    e.preventDefault();
    await mutateAsync({ username, password: btoa(password) });
    navigate("/login");
  };

  return (
    <Container sx={{ padding: 2 }}>
      <Paper elevation={3} sx={{ padding: 2 }}>
        <Typography variant="h3" gutterBottom>
          Signup
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
            <Button
              onClick={handleSubmit}
              variant="contained"
              disabled={isLoading}
            >
              Signup
            </Button>
            <Link to="/login">
              <Button variant="text" disabled={isLoading}>
                Login Instead
              </Button>
            </Link>
          </Box>
        </form>
      </Paper>
    </Container>
  );
};

export default Signup;
