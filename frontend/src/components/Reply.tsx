import { Delete, Edit } from "@mui/icons-material";
import {
  Box,
  Button,
  IconButton,
  Paper,
  TextField,
  Typography,
} from "@mui/material";
import { useMemo, useState } from "react";
import { useDeleteReplyMutation } from "../mutations/useDeleteReplyMutation";
import { useEditReplyMutation } from "../mutations/useEditReplyMutation";
import { useUserQuery } from "../queries/useUserQuery";

import { Reply } from "../types/Reply";
import { getUserInfo } from "../util/getUserInfo";

interface ReplyPropsInterface {
  reply: Reply;
  post_id: string;
}

const ReplyComponent = ({ reply, post_id }: ReplyPropsInterface) => {
  const [edit, setEdit] = useState(false);

  const [message, setMessage] = useState("");

  const { data: replier } = useUserQuery(reply.user_id);

  const editReplyMutation = useEditReplyMutation(post_id);
  const deleteReplyMutation = useDeleteReplyMutation(post_id);

  const user = useMemo(() => getUserInfo(), []);

  const handleMessageChange = (e: React.ChangeEvent<HTMLInputElement>) =>
    setMessage(e.target.value);

  const handleEdit = async () => {
    await editReplyMutation.mutateAsync({ id: reply.id, message });
    setEdit(false);
    setMessage("");
  };

  const handleDelete = () => {
    deleteReplyMutation.mutate({ id: reply.id });
  };

  return (
    <Paper sx={{ padding: 1 }}>
      <Box>
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
              gap: 1,
              justifyContent: "space-between",
              flexWrap: "wrap",
            }}
          >
            <Typography gutterBottom>{reply.message}</Typography>
            {reply.user_id === user.user_id && (
              <>
                <Box
                  sx={{
                    flex: 1,
                    display: "flex",
                    justifyContent: "flex-end",
                    marginBottom: 1,
                  }}
                >
                  <IconButton
                    onClick={() => {
                      setEdit(true);
                      setMessage(reply.message);
                    }}
                  >
                    <Edit />
                  </IconButton>
                  <IconButton onClick={handleDelete}>
                    <Delete />
                  </IconButton>
                </Box>
              </>
            )}
          </Box>
        )}
        {replier && (
          <Typography variant="body2" fontStyle="italic" align="right">
            - {replier?.username}
          </Typography>
        )}
      </Box>
    </Paper>
  );
};

export default ReplyComponent;
