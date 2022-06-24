import { useState } from "react";
import Box from "@mui/material/Box";
import TextField from "@mui/material/TextField";
import Button from "@mui/material/Button";
import { useCreatePostMutation } from "../mutations/useCreatePostMutation";

const CreatePost = () => {
  const [message, setMessage] = useState("");

  const { mutateAsync } = useCreatePostMutation();

  const handleMessageChange = (e: React.ChangeEvent<HTMLInputElement>) =>
    setMessage(e.target.value);

  const handleSubmit = async (
    e: React.MouseEvent<HTMLElement> | React.FormEvent
  ) => {
    e.preventDefault();

    await mutateAsync({ message });

    setMessage("");
  };

  return (
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
          onChange={handleMessageChange}
          value={message}
          label={"Message"}
        />
      </Box>

      <Box sx={{ gap: 1, display: "flex", paddingTop: 2 }}>
        <Button onClick={handleSubmit} variant="contained">
          Post
        </Button>
      </Box>
    </form>
  );
};

export default CreatePost;
