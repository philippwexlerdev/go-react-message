import {
  Box,
  Paper,
  IconButton,
  Button,
  Divider,
  Typography,
  TextField,
} from "@mui/material";
import { Edit, Delete } from "@mui/icons-material";
import { Post } from "../types/Post";
import { getUserInfo } from "../util/getUserInfo";
import { useDeletePostMutation } from "../mutations/useDeletePostMutation";
import Replies from "./Replies";
import { useState } from "react";
import { useEditPostMutation } from "../mutations/useEditPostMutation";
import { useUserQuery } from "../queries/useUserQuery";

interface PostProps {
  post: Post;
}

const PostComponent = ({ post }: PostProps) => {
  const userInfo = getUserInfo();
  const [edit, setEdit] = useState(false);
  const [message, setMessage] = useState("");

  const deletePostMutation = useDeletePostMutation();
  const editPostMutation = useEditPostMutation();

  const { data: user } = useUserQuery(post.user_id);

  const handleMessageChange = (e: React.ChangeEvent<HTMLInputElement>) =>
    setMessage(e.target.value);

  const handleEdit = async () => {
    await editPostMutation.mutateAsync({ postId: post.id, message });
    setEdit(false);
  };

  return (
    <Paper sx={{ padding: 2, maxWidth: 350 }}>
      {edit ? (
        <Box
          sx={{
            display: "flex",
            flexDirection: "column",
            gap: 2,
            marginBottom: 2,
          }}
        >
          <TextField
            onChange={handleMessageChange}
            value={message}
            label={"Message"}
          />
          <Box sx={{ display: "flex", justifyContent: "flex-end", gap: 2 }}>
            <Button variant="contained" onClick={handleEdit}>
              Edit
            </Button>
            <Button variant="outlined" onClick={() => setEdit(false)}>
              Cancel
            </Button>
          </Box>
        </Box>
      ) : (
        <Box
          sx={{
            display: "flex",
            gap: 2,
            marginBotttom: 2,
            justifyContent: "space-between",
          }}
        >
          <h2>{post.message}</h2>
          {userInfo.user_id === post.user_id && (
            <>
              <Box>
                <IconButton
                  onClick={() => {
                    setEdit(true);
                    setMessage(post.message);
                  }}
                >
                  <Edit />
                </IconButton>
                <IconButton
                  onClick={() =>
                    deletePostMutation.mutateAsync({ postId: post.id })
                  }
                >
                  <Delete />
                </IconButton>
              </Box>
            </>
          )}
        </Box>
      )}
      <Box sx={{ marginY: 1 }}>
        <Typography variant="body2" align="right" fontStyle="italic">
          - {user?.username}
        </Typography>
      </Box>
      <Divider />
      <Box sx={{ marginTop: 2 }}>
        <Typography variant="h6">Replies</Typography>
        <Replies post={post} />
      </Box>
    </Paper>
  );
};

export default PostComponent;
