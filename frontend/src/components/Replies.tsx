import { Divider, Box, Button, TextField } from "@mui/material";
import { useState } from "react";
import { useCreatReplyMutation } from "../mutations/useCreateReplyMutation";
import { useRepliesQuery } from "../queries/useRepliesQuery";
import { Post } from "../types/Post";
import { getUserInfo } from "../util/getUserInfo";
import Reply from "./Reply";

interface RepliesProps {
  post: Post;
}

const Replies = ({ post }: RepliesProps) => {
  const createReplyMutation = useCreatReplyMutation(post.id);
  const [reply, setReply] = useState("");

  const { data: replies } = useRepliesQuery(post.id);

  const handleReplyChange = (e: React.ChangeEvent<HTMLInputElement>) =>
    setReply(e.target.value);

  const handleReply = async () => {
    const user = getUserInfo();

    await createReplyMutation.mutateAsync({
      message: reply,
      post_id: post.id,
      user_id: user.user_id,
    });

    setReply("");
  };

  return (
    <>
      <Box
        sx={{
          marginY: 2,
          display: "flex",
          flexDirection: "column",
          gap: 2,
        }}
      >
        {replies?.map((reply) => (
          <Reply reply={reply} key={reply.id} post_id={post.id} />
        ))}
        {replies?.length === 0 && <h6>No replies!</h6>}
      </Box>
      <Divider />
      <Box
        sx={{
          paddingTop: 2,
          minWidth: 300,
          display: "flex",
          flexDirection: "column",
          gap: 1,
        }}
      >
        <TextField onChange={handleReplyChange} value={reply} label={"Reply"} />
        <Box sx={{ display: "flex", justifyContent: "flex-end" }}>
          <Button variant="contained" onClick={handleReply}>
            Reply
          </Button>
        </Box>
      </Box>
    </>
  );
};

export default Replies;
